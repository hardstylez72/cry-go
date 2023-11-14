package v1

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"sync"
	"time"

	"github.com/corpix/uarand"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/starknet"
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/user"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/socks5"
	"github.com/hardstylez72/cry/internal/uniclient"
	"github.com/pkg/errors"
)

type ProfileService struct {
	v1.UnimplementedProfileServiceServer
	repository      repository.ProfileRepository
	settingsService *settings.Service
	starkNetClient  *starknet.Client
}

func NewProfileService(repository repository.ProfileRepository, settingsService *settings.Service, starkNetClient *starknet.Client) *ProfileService {

	return &ProfileService{
		repository:      repository,
		settingsService: settingsService,
		starkNetClient:  starkNetClient,
	}
}

func (s *ProfileService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (*v1.UpdateProfileResponse, error) {

	if err := NotWellcome(ctx, user.GroupWorker); err != nil {
		return nil, err
	}

	p := &repository.Profile{
		Id:        req.ProfileId,
		Label:     req.Label,
		Proxy:     sql.NullString{},
		Meta:      sql.NullString{},
		UserAgent: req.UserAgent,
	}

	if req.Meta != nil {
		p.Meta.Valid = true
		p.Meta.String = req.GetMeta()
	}

	if req.Proxy != nil {
		p.Proxy.Valid = true
		p.Proxy.String = req.GetProxy()
	}

	if err := s.repository.UpdateProfile(ctx, p); err != nil {
		return nil, err
	}
	return &v1.UpdateProfileResponse{}, nil
}
func (s *ProfileService) ValidateLabel(ctx context.Context, req *v1.ValidateLabelRequest) (*v1.ValidateLabelResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	valid, err := s.repository.ValidateLabel(ctx, &repository.ValidateLabelReq{
		ProfileId: req.ProfileId,
		Label:     req.Label,
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.ValidateLabelResponse{
		Valid: *valid,
	}, nil
}
func (s *ProfileService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.GetProfileResponse, error) {
	p, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	pr, err := p.ToPB(s.starkNetClient)
	if err != nil {
		return nil, err
	}
	return &v1.GetProfileResponse{
		Profile: pr,
	}, nil
}
func (s *ProfileService) CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error) {

	if err := NotWellcome(ctx, user.GroupWorker); err != nil {
		return nil, err
	}

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	a := &repository.Profile{
		Id:        uuid.New().String(),
		Label:     req.Label,
		Proxy:     sql.NullString{},
		MmskPk:    []byte(req.MmskPk),
		Meta:      sql.NullString{},
		UserId:    userId,
		CreatedAt: time.Now(),
		UserAgent: uarand.NewWithCustomList(lib.UserAgents).GetRandom(),
		Type:      req.Type.String(),
		SubType:   req.SubType.String(),
	}

	switch req.Type {
	case v1.ProfileType_EVM:
		w, err := defi.GetEMVPublicKey(req.MmskPk)
		if err != nil {
			return nil, errors.New("invalid pk")
		}
		a.MmskId = []byte(w)
	case v1.ProfileType_StarkNet:
		publicKey, err := starknet.GetPublicKeyHash(req.MmskPk)
		if err != nil {
			return nil, err
		}
		a.MmskId = []byte(publicKey)
	}

	if req.GetProxy() != "" {
		p, err := socks5.NewSock5ProxyString(req.GetProxy(), "")
		if err != nil {
			return nil, errors.New("invalid proxy")
		}
		_, err = p.GetIpStat(ctx)
		if err != nil {
			return nil, errors.New("proxy stat is not available")
		}
	}

	if req.Meta != nil {
		a.Meta.Valid = true
		a.Meta.String = *req.Meta
	}

	if req.Proxy != nil {
		a.Proxy.Valid = true
		a.Proxy.String = *req.Proxy
	}

	err = s.repository.CreateProfile(ctx, a)
	if err != nil {
		if errors.Is(err, pg.EntityAlreadyExist) {
			return nil, errors.New("pk already exist")
		}
		return nil, err
	}

	res, err := s.repository.GetProfile(ctx, a.Id)
	if err != nil {
		return nil, err
	}
	pr, err := res.ToPB(s.starkNetClient)
	if err != nil {
		return nil, err
	}

	return &v1.CreateProfileResponse{
		Profile: pr,
	}, nil
}
func (s *ProfileService) ListProfile(ctx context.Context, req *v1.ListProfileRequest) (*v1.ListProfileResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}
	res, err := s.repository.ListProfiles(ctx, userId, req.Type.String(), req.Offset)
	if err != nil {
		return nil, err
	}

	profiles := make([]*v1.Profile, 0)

	for i := range res {
		p := res[i]
		pp, err := p.ToPB(s.starkNetClient)
		if err != nil {
			return nil, err
		}

		profiles = append(profiles, pp)
	}

	return &v1.ListProfileResponse{
		Profiles: profiles,
	}, nil
}
func (s *ProfileService) DeleteProfile(ctx context.Context, req *v1.DeleteProfileRequest) (*v1.DeleteProfileResponse, error) {

	if err := NotWellcome(ctx, user.GroupWorker); err != nil {
		return nil, err
	}
	_, _ = s.repository.DeleteProfile(ctx, req)
	return &v1.DeleteProfileResponse{}, nil
}
func (s *ProfileService) SearchProfilesNotConnectedToOkexDeposit(ctx context.Context, req *v1.SearchProfilesNotConnectedToOkexDepositRequest) (*v1.SearchProfilesNotConnectedToOkexDepositResponse, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	profiles, err := s.repository.SearchNotConnectedOkexDepositProfile(ctx, userId, req.SubType)
	if err != nil {
		return nil, err
	}

	out := make([]*v1.Profile, 0)

	for _, p := range profiles {
		o, err := p.ToPB(s.starkNetClient)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}
	return &v1.SearchProfilesNotConnectedToOkexDepositResponse{Profiles: out}, nil
}
func (s *ProfileService) SearchProfile(ctx context.Context, req *v1.SearchProfileRequest) (*v1.SearchProfileResponse, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	profiles, err := s.repository.SearchProfile(ctx, req.Pattern, userId, req.GetType().String())
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Profile, 0)

	for _, p := range profiles {
		o, err := p.ToPB(s.starkNetClient)
		if err != nil {
			return nil, err
		}
		out = append(out, o)
	}

	return &v1.SearchProfileResponse{Profiles: out}, nil
}
func (s *ProfileService) GetBalance(ctx context.Context, req *v1.GetBalanceRequest) (*v1.GetBalanceResponse, error) {

	tokens := []v1.Token{
		v1.Token_USDT,
		v1.Token_USDC,
		v1.Token_USDCBridged,
		v1.Token_ETH,
		v1.Token_STG,
		v1.Token_WETH,
		v1.Token_LSD,
		v1.Token_LUSD,
		v1.Token_MUTE,
		v1.Token_MAV,
		v1.Token_SPACE,
		v1.Token_VC,
		v1.Token_IZI,
		v1.Token_BUSD,
		v1.Token_USDp,
	}

	var err error
	wg := sync.WaitGroup{}

	profile, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	p, err := profile.ToPB(s.starkNetClient)
	if err != nil {
		return nil, err
	}

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	stgs, err := s.settingsService.GetSettings(ctx, userId, req.Network)
	if err != nil {
		return nil, err
	}

	cli, err := uniclient.NewBaseClient(req.Network, &uniclient.BaseClientConfig{
		RPCEndpoint: stgs.RpcEndpoint,
		ProxyString: p.GetProxy(),
	})
	if err != nil {
		return nil, err
	}

	tokens = append(tokens, cli.GetNetworkToken())

	m := make(map[v1.Token]bool)
	for _, token := range tokens {
		m[token] = true
	}

	wg.Add(len(m))
	balance := make([]*v1.Balance, len(m))

	for i := range tokens {
		_, ok := m[tokens[i]]
		if ok {
			delete(m, tokens[i])
		} else {
			continue
		}

		go func(balancer defi.Networker, i int) {
			defer wg.Done()

			pubKey, err := balancer.GetPublicKey(string(profile.MmskPk), p.SubType)
			if err != nil {
				return
			}

			b, err := balancer.GetBalance(ctx, &defi.GetBalanceReq{
				WalletAddress: pubKey,
				Token:         tokens[i],
			})
			if err != nil {
				msg := err.Error()
				balance[i] = &v1.Balance{Token: tokens[i], Error: &msg}
			} else {
				balance[i] = &v1.Balance{Token: tokens[i], Amount: b.HumanReadable, Wei: b.WEI.String()}
			}

		}(cli, i)
	}

	wg.Wait()

	tmp := make([]*v1.Balance, 0)
	for _, b := range balance {

		if b.Error != nil {
			continue
		}

		if b == nil {
			continue
		}
		if b.Wei == "0" {
			continue
		}
		tmp = append(tmp, b)
	}

	return &v1.GetBalanceResponse{
		Balance: tmp,
	}, nil
}

//	func (s *ProfileService) ExportProfiles(ctx context.Context, req *v1.ExportProfilesReq) (*v1.ExportProfilesRes, error) {
//		userId, err := user.ResolveUserId(ctx)
//		if err != nil {
//			return nil, err
//		}
//
//		profiles, err := s.repository.ExportProfiles(ctx, userId)
//		if err != nil {
//			return nil, err
//		}
//
//		buf := bytes.NewBuffer([]byte{})
//		buf.Grow(100000)
//		w := csv.NewWriter(buf)
//
//		w.UseCRLF = true
//		w.Comma = ';'
//		for i, p := range profiles {
//			if i%10 == 0 {
//				if err := w.Write([]string{}); err != nil {
//					return nil, err
//				}
//			}
//			if err := w.Write([]string{string(p.MmskPk), p.Proxy.String, p.Label}); err != nil {
//				return nil, err
//			}
//		}
//		w.Flush()
//		return &v1.ExportProfilesRes{
//			Text: buf.String(),
//		}, nil
//	}
func (s *ProfileService) GenerateProfiles(ctx context.Context, req *v1.GenerateProfilesReq) (*v1.GenerateProfilesRes, error) {

	if err := NotWellcome(ctx, user.GroupWorker); err != nil {
		return nil, err
	}

	pb := bytes.NewBuffer([]byte{})
	pb.Grow(100000)

	buf := bytes.NewBuffer([]byte{})
	buf.Grow(100000)
	w := csv.NewWriter(buf)

	w.UseCRLF = true
	w.Comma = ';'

	switch req.Type {
	case v1.ProfileType_EVM:
		for i := 0; i < int(req.Count); i++ {
			key, err := crypto.GenerateKey()
			if err != nil {
				return nil, err
			}
			privateKey := hex.EncodeToString(key.D.Bytes())

			pk, err := crypto.HexToECDSA(privateKey)
			if err != nil {
				return nil, errors.Wrap(err, "crypto.HexToECDSA(privateKey)")
			}

			publicKey := pk.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			walletAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

			if err := w.Write([]string{privateKey, "", ""}); err != nil {
				return nil, err
			}

			if _, err := pb.Write([]byte(
				" pk: " + privateKey +
					" pub: " + walletAddr.String() +
					"\n")); err != nil {
				return nil, err
			}
		}

	case v1.ProfileType_StarkNet:

		acc, err := s.starkNetClient.DegenerateAccount(ctx, req.SubType, req.GetCount())
		if err != nil {
			return nil, err
		}

		for _, a := range acc {
			if err := w.Write([]string{a.PK, "", ""}); err != nil {
				return nil, err
			}

			if _, err = pb.Write([]byte(" pub: " + a.Pub + " pk: " + a.PK + " seed: " + a.Seed + "\n")); err != nil {
				return nil, err
			}
		}
	}

	w.Flush()

	return &v1.GenerateProfilesRes{
		Text:    buf.String(),
		Preview: pb.String(),
	}, nil

}

func (s *ProfileService) StarkNetAccountDeployed(ctx context.Context, req *v1.StarkNetAccountDeployedReq) (*v1.StarkNetAccountDeployedRes, error) {
	p, err := s.repository.GetProfile(ctx, req.ProfileId)
	if err != nil {
		return nil, err
	}

	profile, err := p.ToPB(s.starkNetClient)
	if err != nil {
		return nil, err
	}
	deployed, err := s.starkNetClient.IsAccountDeployed(ctx, string(p.MmskPk), profile.SubType)
	if err != nil {
		return nil, err
	}

	return &v1.StarkNetAccountDeployedRes{Deployed: *deployed}, nil
}

func (s *ProfileService) GetProfileRelations(ctx context.Context, req *v1.GetProfileRelationsReq) (*v1.GetProfileRelationsRes, error) {
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	r := &repository.GetProfileRelationsReq{
		P1Type:    req.GetP1Type(),
		P2Type:    req.P2Type,
		P1SubType: req.P1SubType,
		P2SubType: req.P2SubType,
		UserId:    userId,
	}

	res, err := s.repository.GetProfileRelations(ctx, r)
	if err != nil {
		return nil, err
	}

	tmp := []*v1.ProfileRelation{}

	for _, r := range res {
		t := &v1.ProfileRelation{
			Type:    v1.ProfileType(v1.ProfileType_value[r.Type]),
			SubType: v1.ProfileSubType(v1.ProfileSubType_value[r.SubType]),
			Id:      r.Id,
			Num:     r.Num,
		}

		if r.Related != nil {
			t.Rel = &v1.UnrelatedProfile{
				Type:    v1.ProfileType(v1.ProfileType_value[r.Related.Type]),
				SubType: v1.ProfileSubType(v1.ProfileSubType_value[r.Related.SubType]),
				Id:      r.Related.Id,
				Num:     r.Related.Num,
			}
		}

		tmp = append(tmp, t)
	}

	return &v1.GetProfileRelationsRes{Items: tmp}, nil

}
func (s *ProfileService) DeleteProfileRelation(ctx context.Context, req *v1.DeleteProfileRelationReq) (*v1.DeleteProfileRelationRes, error) {

	//if err := NotWellcome(ctx, user.GroupWorker); err != nil {
	//	return nil, err
	//}
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	err = s.repository.DeleteProfileRelation(ctx, &repository.ProfileRelation{
		P1Id:      req.P1Id,
		P1SubType: req.P1SubType.String(),
		P1Type:    req.P1Type.String(),
		P2Id:      req.P2Id,
		P2SubType: req.P2SubType.String(),
		P2Type:    req.P2Type.String(),
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteProfileRelationRes{}, nil
}
func (s *ProfileService) AddProfileRelation(ctx context.Context, req *v1.AddProfileRelationReq) (*v1.AddProfileRelationRes, error) {

	//if err := NotWellcome(ctx, user.GroupWorker); err != nil {
	//	return nil, err
	//}
	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	err = s.repository.AddProfileRelation(ctx, &repository.ProfileRelation{
		P1Id:      req.P1Id,
		P1SubType: req.P1SubType.String(),
		P1Type:    req.P1Type.String(),
		P2Id:      req.P2Id,
		P2SubType: req.P2SubType.String(),
		P2Type:    req.P2Type.String(),
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddProfileRelationRes{}, nil
}
func (s *ProfileService) SearchProfileRelation(ctx context.Context, req *v1.SearchProfileRelationReq) (*v1.SearchProfileRelationRes, error) {

	userId, err := user.ResolveUserId(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.repository.SearchProfilesNotRelated(ctx, &repository.SearchProfilesNotRelated{
		PType:       req.P1Type.String(),
		PSubType:    req.P1SubType.String(),
		NotPType:    req.P2Type.String(),
		NotPSubType: req.P2SubType.String(),
		UserId:      userId,
	})
	if err != nil {
		return nil, err
	}

	tmp := make([]*v1.UnrelatedProfile, len(res))
	for i := range tmp {
		t := res[i]
		tmp[i] = &v1.UnrelatedProfile{
			Id:      t.Id,
			SubType: v1.ProfileSubType(v1.ProfileSubType_value[t.SubType]),
			Type:    v1.ProfileType(v1.ProfileType_value[t.Type]),
			Num:     t.Num,
		}
	}

	return &v1.SearchProfileRelationRes{Items: tmp}, nil
}
