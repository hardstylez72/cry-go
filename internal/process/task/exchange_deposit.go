package task

import (
	"context"
	"strings"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/exchange/okex"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type OkexDepositTask struct {
	cancel func()
}

func (t *OkexDepositTask) Stop() error {
	t.cancel()
	return nil
}

func (t *OkexDepositTask) Type() v1.TaskType {
	return v1.TaskType_OkexDeposit
}

func (t *OkexDepositTask) Run(ctx context.Context, a *Input) (*v1.ProcessTask, error) {

	taskContext, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	t.cancel = cancel

	task := a.Task
	l, ok := a.Task.Task.Task.(*v1.Task_OkexDepositTask)
	if !ok {
		return nil, errors.New("panic.a.Task.Task.Task.(*v1.Task_OkexDepositTask) call an ambulance!")
	}

	p := l.OkexDepositTask

	profile, err := a.Halper.Profile(ctx, a.ProfileId)
	if err != nil {
		return nil, err
	}

	s, err := profile.GetNetworkSettings(ctx, p.Network)
	if err != nil {
		return nil, err
	}
	client, err := uniclient.NewTransfer(p.Network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	switch a.Task.Status {
	case v1.ProcessStatus_StatusDone, v1.ProcessStatus_StatusError:
		return a.Task, nil
	case v1.ProcessStatus_StatusRetry, v1.ProcessStatus_StatusReady:
		task.Status = v1.ProcessStatus_StatusRunning
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	//if p.Network == v1.Network_StarkNet {
	//
	//	starkClient, err := NewStarkNetClient(profile)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if p.GetApproveTx().GetTxId() == "" {
	//
	//		txId, err := StarkNetApprove(taskContext, p.Token, starkClient, profile, t.Type())
	//		if err != nil {
	//			return nil, err
	//		}
	//
	//		if txId != nil {
	//			p.ApproveTx = NewStarkNetApproveTx(*txId)
	//			if err := a.AddTx2(ctx, p.ApproveTx); err != nil {
	//				return nil, err
	//			}
	//			if err := a.UpdateTask(ctx, task); err != nil {
	//				return nil, err
	//			}
	//		}
	//	}
	//
	//	if err := WaitTxComplete(taskContext, p.ApproveTx, task, client, a); err != nil {
	//		return nil, err
	//	}
	//}

	if p.GetTx().GetTxId() == "" {

		addr, err := GetOkexDepositAddr(taskContext, profile, p, a.WithdrawerRepository)
		if err != nil {
			return nil, err
		}

		p.Address = addr
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}

		b, err := client.GetBalance(taskContext, &defi.GetBalanceReq{
			WalletAddress: profile.Addr,
			Token:         p.Token,
		})
		if err != nil {
			return nil, err
		}

		am, err := defi.ResolveAmount(p.Amount, b.WEI)
		if err != nil {
			return nil, err
		}

		estimate, err := EstimateOkexDepositCost(taskContext, profile, p, a.WithdrawerRepository)
		if err != nil {
			return nil, err
		}
		gas, err := GasManager(estimate, s.Source, p.Network, t.Type())
		if err != nil {
			return nil, err
		}

		if p.Token == client.GetNetworkToken() {
			am = ResolveNetworkTokenAmount(b.WEI, &gas.TotalGas, am)
		}

		res, err := client.Transfer(taskContext, &defi.TransferReq{
			Pk:           profile.WalletPK,
			ToAddr:       *addr,
			Token:        p.Token,
			Amount:       am,
			Gas:          gas,
			PSubType:     profile.SubType,
			EstimateOnly: false,
		})
		if err != nil {
			return nil, err
		}

		p.Tx = NewTx(res.Tx, gas)
		if err := a.AddTx2(ctx, p.Tx); err != nil {
			return nil, err
		}

		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	if err := WaitTxComplete(taskContext, p.Tx, task, client, a); err != nil {
		return nil, err
	}

	// перевод с sub на main okex аккаунт
	if p.GetTx().GetTxCompleted() && p.SubMainTransfer == nil {
		if err := t.OkexSubMainTransfer(taskContext, a); err != nil {
			if errors.Is(err, ErrZeroBalance) {
				return task, nil
			}
			return nil, err
		}
		tmp := true
		p.SubMainTransfer = &tmp
		task.Status = v1.ProcessStatus_StatusDone
		if err := a.UpdateTask(ctx, task); err != nil {
			return nil, err
		}
	}

	return task, nil
}

var ErrZeroBalance = errors.New("balance is zero")

func (t *OkexDepositTask) OkexSubMainTransfer(ctx context.Context, a *Input) error {
	l, ok := a.Task.Task.Task.(*v1.Task_OkexDepositTask)
	if !ok {
		return errors.New("panic.a.Task.Task.Task.(*v1.Task_OkexDepositTask) call an ambulance!")
	}

	p := l.OkexDepositTask

	profile, err := a.ProfileRepository.GetProfile(ctx, a.ProfileId)
	if err != nil {
		return err
	}

	attached, err := a.WithdrawerRepository.GetAttachedAddr(ctx, profile.Id, a.UserId, profile.SubType)
	if err != nil {
		if errors.Is(err, pg.EntityNotFound) {
			return errors.Wrap(ErrProfileHasNoConnectedOkexAddr, err.Error())
		}
		return err
	}

	sub, err := a.WithdrawerRepository.GetWithdrawer(ctx, attached.WithdrawerId, a.UserId)
	if err != nil {
		return err
	}
	subb, err := GetOkexCli(ctx, a.WithdrawerRepository, attached.WithdrawerId, a.UserId, profile.UserAgent)
	if err != nil {
		return err
	}

	token := p.Token.String()
	if p.Token == v1.Token_USDCBridged {
		token = "USDC"
	}

	b, err := subb.GetBalance(ctx, token)
	if err != nil {
		return err
	}
	if b == 0 {
		return errors.Wrap(ErrZeroBalance, "")
	}
	main, err := GetOkexCli(ctx, a.WithdrawerRepository, sub.PrevId.String, a.UserId, profile.UserAgent)
	if err != nil {
		return err
	}

	if err := main.SubsToMain(ctx, &okex.SubsToMainReq{
		Ccy: token,
	}); err != nil {
		return err
	}

	return nil
}

func GetOkexDepositAddr(ctx context.Context, profile *halp.Profile, p *v1.OkexDepositTask, withdrawerRepository repository.WithdrawerRepository) (*string, error) {

	attached, err := withdrawerRepository.GetAttachedAddr(ctx, profile.Id, profile.UserId, profile.SubType.String())
	if err != nil {
		return nil, errors.Wrap(err, "к профилю не привязан адресс депозита с okex")
	}

	addr := attached.Addr

	okexCli, err := GetOkexCli(ctx, withdrawerRepository, attached.WithdrawerId, profile.UserId, profile.UserAgent)
	if err != nil {
		return nil, err
	}

	token := p.Token.String()
	if p.Token == v1.Token_USDCBridged {
		token = "USDC"
	}

	addrMap, err := okexCli.GetDepositAddress(ctx, token)
	if err != nil {
		return nil, err
	}

	_, exist := addrMap[attached.Addr]
	if !exist {
		return nil, errors.New("okex deposit address in settings is invalid. update please")
	}
	return &addr, nil
}

func GetOkexCli(ctx context.Context, rep repository.WithdrawerRepository, withdrawerId, userId, userAgentHeader string) (*okex.Service, error) {

	wd, err := rep.GetWithdrawer(ctx, withdrawerId, userId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if wd.Proxy.Valid {
		proxyString = wd.Proxy.String
	}

	proxy, err := socks5.NewSock5ProxyString(proxyString, userAgentHeader)
	if err != nil {
		return nil, err
	}

	sub := strings.Split(string(wd.ApiKey), ":")

	return okex.NewService(ctx, &okex.Config{
		ApiKey:     sub[0],
		SecretKey:  string(wd.SecretKey),
		PassPhrase: sub[1],
		HttpClient: proxy.Cli,
	})
}

func EstimateOkexDepositCost(ctx context.Context, profile *halp.Profile, p *v1.OkexDepositTask, withdrawerRepository repository.WithdrawerRepository) (*v1.EstimationTx, error) {

	s, err := profile.GetNetworkSettings(ctx, p.Network)

	walletAddr := profile.Addr

	swapper, err := uniclient.NewTransfer(p.Network, s.BaseConfig())
	if err != nil {
		return nil, err
	}

	b, err := swapper.GetBalance(ctx, &defi.GetBalanceReq{WalletAddress: walletAddr, Token: p.Token})
	if err != nil {
		return nil, err
	}

	if b.Float == 0 {
		return nil, errors.New("not enough balance in " + p.Token.String())
	}

	am, err := defi.ResolveAmount(p.Amount, b.WEI)
	if err != nil {
		return nil, err
	}

	addr, err := GetOkexDepositAddr(ctx, profile, p, withdrawerRepository)
	if err != nil {
		return nil, err
	}

	swap, err := swapper.Transfer(ctx, &defi.TransferReq{
		Pk:           profile.WalletPK,
		ToAddr:       *addr,
		Token:        p.Token,
		Amount:       bozdo.Percent(am, 80),
		EstimateOnly: true,
		PSubType:     profile.SubType,
	})
	if err != nil {
		return nil, err
	}

	if p.Token == swapper.GetNetworkToken() {
		am = ResolveNetworkTokenAmount(b.WEI, swap.ECost.TotalGasWei, am)
	}

	return GasStation(swap.ECost, p.Network), nil
}
