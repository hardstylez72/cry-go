package process

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/exchange"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/halp"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type AutoRefuel struct {
	Profile   *halp.Profile
	Err       defi.ErrOutOfGas
	TaskId    string
	ProcessId string
}

func (d *Dispatcher) AutoRefuel(ctx context.Context, a AutoRefuel) error {

	networkSettings, err := d.settingsService.GetSettings(ctx, a.Profile.UserId, a.Err.N)
	if err != nil {
		return err
	}

	ar := networkSettings.GetAutoRefuel()
	if !ar.GetEnabled() {
		return errors.New("auto-refuel disabled")
	}

	w, err := d.runner.withdrawerRepository.GetWithdrawer(ctx, ar.GetWithdrawerId(), a.Profile.UserId)
	if err != nil {
		return err
	}

	cli, err := uniclient.NewExchangeWithdrawer(w, a.Profile.DB)
	if err != nil {
		return err
	}

	nativeToken, ok := bozdo.NativeTokenMap[a.Err.N]
	if !ok {
		return errors.New("native token is not found for " + a.Err.N.String())
	}

	if nativeToken.String() != a.Err.Token.String() {
		return errors.New("must be native token")
	}

	am := lib.FloatToString(lib.RandFloatRange(ar.GetMin(), ar.GetMax()))
	wr, err := cli.Withdraw(ctx, &exchange.WithdrawRequest{
		ToAddress: a.Profile.Addr,
		Amount:    am,
		Network:   CastNetworkToOkex(a.Err.N),
		Token:     CastTokenToOkex(a.Err.Token),
	})
	if err != nil {
		return err
	}

	err = d.r.UpdateProcessTaskStatus(ctx, v1.ProcessStatus_StatusRunning.String(), a.TaskId, a.ProcessId)
	if err != nil {
		return errors.Wrap(err, "UpdateProcessTaskStatus")
	}

	_ = d.runner.processRepository.RecordStatusChanged(ctx, a.TaskId, v1.ProcessStatus_StatusError, v1.ProcessStatus_StatusRunning, "авто-пополнение на "+am+" "+CastTokenToOkex(a.Err.Token))

	over := make(chan struct{})
	defer close(over)
	finish := make(chan struct{})
	defer close(finish)
	pt, _ := d.pTable.Get(a.ProcessId)

	go func() {
		for {
			select {
			case _ = <-finish:
				over <- struct{}{}
				return
			case signal := <-pt.signals:
				switch signal {
				case SignalStop:
					over <- struct{}{}
					return
				}
			}
		}
	}()

	go func() {
		if _, err := cli.WaitConfirm(ctx, wr.WithdrawId); err != nil {
		}
		finish <- struct{}{}
	}()

	<-over

	return nil
}

func CastTokenToOkex(t v1.Token) string {
	switch t {
	case v1.Token_ETH:
		return "ETH"
	case v1.Token_AVAX:
		return "AVAX"
	case v1.Token_BNB:
		return "BNB"
	case v1.Token_MATIC:
		return "MATIC"
	case v1.Token_CORE:
		return "CORE"
	default:
		return ""
	}
}

func CastNetworkToOkex(t v1.Network) string {
	switch t {
	case v1.Network_StarkNet:
		return "ETH-Starknet"
	case v1.Network_Linea:
		return "ETH-Linea"
	case v1.Network_ZKSYNCERA:
		return "ETH-zkSync Era"
	case v1.Network_POLIGON:
		return "MATIC-Polygon"
	case v1.Network_BinanaceBNB:
		return "BNB-BSC"
	case v1.Network_ARBITRUM:
		return "ETH-Arbitrum One"
	case v1.Network_Base:
		return "ETH-Base"
	case v1.Network_OPTIMISM:
		return "ETH-Optimism"
	case v1.Network_AVALANCHE:
		return "AVAX-Avalanche C-Chain"
	case v1.Network_Etherium:
		return "ETH-ERC20"
	case v1.Network_Core:
		return "CORE-CORE"
	default:
		return ""
	}
}
