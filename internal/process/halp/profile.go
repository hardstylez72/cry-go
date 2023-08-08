package halp

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/zksyncera"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type Halp struct {
	profileRepository repository.ProfileRepository
	settingsService   *settings.Service
}

func NewHalp(profileRepository repository.ProfileRepository,
	settingsService *settings.Service) *Halp {
	return &Halp{
		profileRepository: profileRepository,
		settingsService:   settingsService,
	}
}

type Profile struct {
	h           *Halp
	UserAgent   string
	UserId      string
	Id          string
	DB          *repository.Profile
	ProxyString string
	WalletAddr  common.Address
	WalletPK    string
	Wallet      *defi.WalletTransactor
}

func (h *Halp) Profile(ctx context.Context, profileId string) (*Profile, error) {

	profile, err := h.profileRepository.GetProfile(ctx, profileId)
	if err != nil {
		return nil, err
	}

	proxyString := ""
	if profile.Proxy.Valid {
		proxyString = profile.Proxy.String
	}

	tx, err := defi.NewWalletTransactor(string(profile.MmskPk))
	if err != nil {
		return nil, err
	}
	return &Profile{
		UserAgent:   profile.UserAgent,
		DB:          profile,
		ProxyString: proxyString,
		WalletAddr:  tx.WalletAddr,
		WalletPK:    tx.PK,
		Wallet:      tx,
		Id:          profileId,
		UserId:      profile.UserId,
		h:           h,
	}, nil
}

type Settings struct {
	Source *v1.NetworkSettings
	p      *Profile
}

func (p *Profile) GetNetworkSettings(ctx context.Context, network v1.Network) (*Settings, error) {
	ns, err := p.h.settingsService.GetSettings(ctx, p.UserId, network)
	if err != nil {
		return nil, err
	}

	return &Settings{Source: ns, p: p}, nil
}

func (p *Settings) BaseConfig() *uniclient.BaseClientConfig {

	return &uniclient.BaseClientConfig{
		ProxyString:     p.p.ProxyString,
		RPCEndpoint:     p.Source.RpcEndpoint,
		UserAgentHeader: p.p.DB.UserAgent,
	}
}

func (p *Settings) EraWallet() (*zksyncera.WalletTransactor, error) {

	id, ok := big.NewInt(0).SetString(p.Source.Id, 10)
	if !ok {
		return nil, errors.New("invalid network id: " + p.Source.Id)
	}

	wallet, err := zksyncera.NewWalletTransactor(string(p.p.DB.MmskPk), id)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (p *Profile) EthWallet() (*defi.WalletTransactor, error) {

	wallet, err := defi.NewWalletTransactor(string(p.DB.MmskPk))
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (p *Settings) GetWalletAddr() (*common.Address, error) {
	var walletAddr common.Address
	if p.Source.Network == v1.Network_ZKSYNCERA || p.Source.GetNetwork() == v1.Network_ZKSYNCLITE {
		w, err := p.EraWallet()
		if err != nil {
			return nil, err
		}
		walletAddr = w.WalletAddr
	} else {
		w, err := p.p.EthWallet()
		if err != nil {
			return nil, err
		}
		walletAddr = w.WalletAddr
	}
	return &walletAddr, nil
}
