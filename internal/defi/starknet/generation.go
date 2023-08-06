package starknet

import (
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/types"
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
