package halp

import (
	"context"
	"math/big"

	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet"
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
	starknetClient    *starknet.Client
}

func NewHalp(
	profileRepository repository.ProfileRepository,
	settingsService *settings.Service,
	starknetClient *starknet.Client) *Halp {
	return &Halp{
		profileRepository: profileRepository,
		settingsService:   settingsService,
		starknetClient:    starknetClient,
	}
}

type Profile struct {
	h           *Halp
	UserAgent   string
	UserId      string
	Id          string
	DB          *repository.Profile
	ProxyString string
	WalletPK    string
	Addr        string
	Type        v1.ProfileType
	SubType     v1.ProfileSubType
	Num         int
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

	p, err := profile.ToPB(h.starknetClient)
	if err != nil {
		return nil, err
	}
	return &Profile{
		UserAgent:   profile.UserAgent,
		DB:          profile,
		ProxyString: proxyString,
		WalletPK:    string(profile.MmskPk),
		Id:          profileId,
		UserId:      profile.UserId,
		h:           h,
		Addr:        p.MmskId,
		Type:        p.Type,
		SubType:     p.SubType,
		Num:         int(p.Num),
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

func (p *Settings) GetWalletAddr() (*string, error) {
	var walletAddr string
	if p.Source.Network == v1.Network_ZKSYNCERA || p.Source.GetNetwork() == v1.Network_ZKSYNCLITE {
		w, err := p.EraWallet()
		if err != nil {
			return nil, err
		}
		walletAddr = w.WalletAddr.String()
	} else {
		w, err := defi.GetEMVPublicKey(string(p.p.DB.MmskPk))
		if err != nil {
			return nil, err
		}
		walletAddr = w
	}
	return &walletAddr, nil
}
