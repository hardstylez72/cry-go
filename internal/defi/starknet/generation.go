package starknet

import (
	"context"

	"github.com/hardstylez72/cry/internal/defi/starknet/halper"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	starknetgo "github.com/hardstylez72/cry/starknet.go"
	"github.com/hardstylez72/cry/starknet.go/types"
	"github.com/pkg/errors"
)

type Account struct {
	Seed string
	Pub  string
	PK   string
}

func (c *Client) DegenerateAccount(ctx context.Context, profileType v1.ProfileSubType, count int64) ([]Account, error) {
	accs, err := c.halper.Generate(ctx, &halper.GenerateReq{
		SubType: profileType.String(), Count: int(count),
	})
	if err != nil {
		return nil, err
	}

	out := make([]Account, len(*accs))
	for i, a := range *accs {
		out[i] = Account{
			Seed: a.Seed,
			Pub:  a.Pub,
			PK:   a.Pk,
		}
	}
	return out, nil
}

//func DegenerateAccount() (*Account, error) {
//	pk, err := starknetgo.Curve.GetRandomPrivateKey()
//	if err != nil {
//		return nil, err
//	}
//
//	x, _, err := starknetgo.Curve.PrivateToPoint(pk)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Account{
//		PublicKey:  types.BigToHex(x),
//		PrivateKey: types.BigToHex(pk),
//	}, nil
//}

func GetPublicKeyHash(pk string) (string, error) {

	pkb := types.HexToBN(pk)
	if pkb == nil {
		return "", errors.New("can not cast to big.Int")
	}

	x, _, err := starknetgo.Curve.PrivateToPoint(pkb)
	if err != nil {
		return "", err
	}

	return types.BigToHex(x), nil
}
