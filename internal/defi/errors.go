package defi

import (
	"fmt"
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func ErrTokenNotSupportedFn(t v1.Token) error {
	return errors.Wrap(ErrTokenNotSupported, "token: "+string(t))
}

var (
	ErrTokenNotSupported = errors.New("token is not supported in this network")
)

var ErrTxStatusFailed = errors.New("transaction failed")
var ErrTxNotFound = errors.New("transaction is not found")

type ErrOutOfGas struct {
	N     v1.Network
	Token v1.Token
	Want  *big.Int
	Have  *big.Int
}

func (e ErrOutOfGas) Error() string {
	return fmt.Sprintf("нехватает газа %s на %s", e.Token.String(), e.N.String())
}
