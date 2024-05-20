package v1

import (
	"context"
	"math/big"
	"strings"

	paycli "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/lib"
	"github.com/hardstylez72/cry/internal/pay"
	"github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process/task"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/hardstylez72/cry/internal/tg"
	"github.com/pkg/errors"
)

type HelperService struct {
	v1.UnimplementedHelperServiceServer
	settingsService             *settings.Service
	profileRepository           repository.ProfileRepository
	userRepository              repository.UserRepository
	payService                  *pay.Service
	statRepository              repository.StatRepository
	repositoryProcessRepository repository.ProcessRepository
	tgBot                       *tg.Bot
	starkNetClient              *starknet.Client
}

func NewHelperService(
	settingsService *settings.Service,
	profileRepository repository.ProfileRepository,
	userRepository repository.UserRepository,
	payService *pay.Service,
	statRepository repository.StatRepository,
	repositoryProcessRepository repository.ProcessRepository,
	tgBot *tg.Bot,
	starkNetClient *starknet.Client,
) *HelperService {
	return &HelperService{
		settingsService:             settingsService,
		profileRepository:           profileRepository,
		userRepository:              userRepository,
		payService:                  payService,
		statRepository:              statRepository,
		repositoryProcessRepository: repositoryProcessRepository,
		tgBot:                       tgBot,
		starkNetClient:              starkNetClient,
	}
}

func (s *HelperService) EstimateStargateBridgeFee(ctx context.Context, req *v1.EstimateStargateBridgeFeeRequest) (*v1.EstimateStargateBridgeFeeResponse, error) {
	return nil, errors.New("deprecated")
}
func (s *HelperService) ValidatePK(ctx context.Context, req *v1.ValidatePKRequest) (*v1.ValidatePKResponse, error) {

	switch req.Type {
	case v1.ProfileType_EVM:
		w, err := defi.GetEMVPublicKey(req.Pk)
		if err != nil {
			return &v1.ValidatePKResponse{
				Valid:    false,
				WalletId: nil,
			}, nil
		}
		addr := w
		return &v1.ValidatePKResponse{
			Valid:    true,
			WalletId: &addr,
		}, nil
	case v1.ProfileType_StarkNet:
		pub, err := s.starkNetClient.GetPublicKey(req.GetPk(), req.SubType)
		if err != nil {
			return &v1.ValidatePKResponse{
				Valid:    false,
				WalletId: nil,
			}, nil
		}
		addr := pub
		return &v1.ValidatePKResponse{
			Valid:    true,
			WalletId: &addr,
		}, nil
	default:
		return nil, errors.New("invalid pk")
	}
}
func (s *HelperService) ValidateProxy(ctx context.Context, req *v1.ValidateProxyRequest) (*v1.ValidateProxyResponse, error) {

	req.Proxy = strings.TrimSpace(req.Proxy)

	//if req.Proxy == "" {
	//	return &v1.ValidateProxyResponse{
	//		Valid:       true,
	//		CountryName: "disabled",
	//		CountryCode: "",
	//		Ip:          "",
	//	}, nil
	//}

	userAgent := lib.DefaultUserAgent

	p, err := socks5.NewSock5ProxyString(req.Proxy, userAgent)
	if err != nil {
		errMsg := ""
		if errors.Is(err, socks5.ErrParseError) {
			errMsg = "invalid proxy format. want <ip:port:login:password>"
		} else {
			errMsg = "proxy does not responding"
		}

		return &v1.ValidateProxyResponse{
			Valid:        false,
			ErrorMessage: errMsg,
		}, nil
	}

	valid := true
	_, err = p.GetIp(ctx)
	if err != nil {
		valid = false
	}

	return &v1.ValidateProxyResponse{
		Valid:       valid,
		CountryName: "",
		CountryCode: "",
		Ip:          "",
	}, nil
}
func (s *HelperService) CastWEI(ctx context.Context, req *v1.CastWEIRequest) (*v1.CastWEIResponse, error) {

	wei, ok := big.NewInt(0).SetString(req.Wei, 10)
	if !ok {
		return nil, errors.New("invalid wei value: " + req.Wei)
	}

	return &v1.CastWEIResponse{
		Am: defi.AmountUni(wei, req.Network),
	}, nil
}
func (s *HelperService) GetBillingHistory(ctx context.Context, req *v1.GetBillingHistoryReq) (*v1.GetBillingHistoryRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.payService.UserTaskHistory(ctx, &paycli.UserTaskHistoryReq{
		Offset: req.Offset,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	temp := make([]*v1.TaskHistoryRecord, len(res.Records))
	for i, r := range res.Records {
		temp[i] = &v1.TaskHistoryRecord{
			ProcessId: r.ProcessId,
			TaskId:    r.TaskId,
			TaskType:  r.TaskType,
			TaskPrice: r.TaskPrice,
		}
	}

	return &v1.GetBillingHistoryRes{
		Records: temp,
	}, nil
}
func (s *HelperService) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.payService.CreateOrder(ctx, &paycli.CreateOrderReq{
		UserId: userId,
		Am:     req.Am,
		Net:    "ARBI_USDT",
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateOrderRes{
		Id:          res.Id,
		CoinAddrUrl: res.CoinAddrUrl,
		Am:          res.Am,
		ToWallet:    res.ToWallet,
	}, nil
}
func (s *HelperService) GetOrderStatus(ctx context.Context, req *v1.GetOrderStatusReq) (*v1.GetOrderStatusRes, error) {
	res, err := s.payService.CheckOrder(ctx, &paycli.CheckOrderReq{
		Id: req.OrderId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetOrderStatusRes{
		Status: res.Status,
	}, nil
}
func (s *HelperService) GetOrderHistory(ctx context.Context, req *v1.GetOrderHistoryReq) (*v1.GetOrderHistoryRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.payService.GetOrderHistory(ctx, &paycli.GetOrderHistoryReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	tmp := make([]*v1.Order, len(res.Orders))

	for i, r := range res.Orders {
		tmp[i] = &v1.Order{
			Id:          r.Id,
			Net:         r.Net,
			CoinAddrUrl: r.CoinAddrUrl,
			Status:      r.Status,
			CreatedAt:   r.CreatedAt,
			ConfirmedAt: r.ConfirmedAt,
			Am:          r.Am,
			ToWallet:    r.ToWallet,
		}
	}
	return &v1.GetOrderHistoryRes{
		Orders: tmp,
	}, nil
}
func (s *HelperService) TransactionsDailyImpact(ctx context.Context, req *v1.TransactionsDailyImpactReq) (*v1.TransactionsDailyImpactRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	my, err := s.statRepository.GetDailyUserImpact(ctx, userId)
	if err != nil {
		return nil, err
	}

	total, err := s.statRepository.GetDailyTotalImpact(ctx)
	if err != nil {
		return nil, err
	}

	top, err := s.statRepository.GetDailyTopImpact(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.TransactionsDailyImpactRes{
		MyImpact:    *my,
		TotalImpact: *total,
		TopImpact:   *top,
	}, nil
}
func (s *HelperService) SupportMessage(ctx context.Context, req *v1.SupportMessageReq) (*v1.SupportMessageRes, error) {

	var details string
	if req.ProcessId != nil && req.TaskId != nil {
		t, err := s.repositoryProcessRepository.GetProcessTask(ctx, req.GetTaskId())
		if err != nil {
			return nil, err
		}
		tp, err := t.ToPB()
		if err != nil {
			return nil, err
		}
		b, err := task.Marshal(tp)
		if err != nil {
			return nil, err
		}
		details = string(b)
	}

	admin, _, err := s.userRepository.GetOrCreateUser(ctx, &repository.User{Email: config.CFG.AdminEmail})
	if err != nil {
		return nil, err
	}
	chatId, err := s.userRepository.GetUserTelegramChatId(ctx, admin.Id)
	if err != nil {
		return nil, err
	}

	if err := s.tgBot.SupportMessage(*chatId, details, req.GetText()); err != nil {
		return nil, err
	}
	return &v1.SupportMessageRes{}, nil
}

func (s *HelperService) UsePromo(ctx context.Context, req *v1.UsePromoReq) (*v1.UsePromoRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.payService.AddPromo(ctx, &paycli.AddPromoReq{
		UserId: userId,
		Promo:  req.Promo,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UsePromoRes{
		Valid: res.Valid,
		Bonus: res.Bonus,
	}, nil
}
