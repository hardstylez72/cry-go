package starknet

import (
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
	"github.com/pkg/errors"
)

type Account struct {
	PublicKey  string
	PrivateKey string
}

func DegenerateAccount() (*Account, error) {
	pk, err := caigo.Curve.GetRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	x, _, err := caigo.Curve.PrivateToPoint(pk)
	if err != nil {
		return nil, err
	}

	return &Account{
		PublicKey:  types.BigToHex(x),
		PrivateKey: types.BigToHex(pk),
	}, nil
}

func GetPublicKeyHash(pk string) (string, error) {

	pkb := types.HexToBN(pk)
	if pkb == nil {
		return "", errors.New("can not cast to big.Int")
	}

	x, _, err := caigo.Curve.PrivateToPoint(pkb)
	if err != nil {
		return "", err
	}

	return types.BigToHex(x), nil
}
