package task

import (
	"fmt"
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func ErrGasIsOverMax(max, estimated string) error {
	return errors.New(fmt.Sprintf("gas (%s USD) is higher than max (%s USD)", estimated, max))
}

var (
	ErrUserHasNoBalance              = errors.New("User has not enough balance. visit billing page please")
	ErrProfileHasNoConnectedOkexAddr = errors.New("Profile is not connected to okex deposit wallet")
	ErrTransactionIsNotReady         = errors.New("transaction is not ready")
	ErrAccountIsZero                 = errors.New("profile has Insufficient balance")
	ErrSwapRateTooBig                = errors.New("swap rate too big")
	ErrProfileHasInsufficientBalance = func(token v1.Token, want, have *big.Int) error {
		return errors.Wrap(ErrAccountIsZero, fmt.Sprintf("want: %s have: %s of %s", want.String(), have.String(), token.String()))
	}
	ErrSwapRateBiggerThenAllowed = func(limit, max float64) error {
		return errors.Wrap(ErrSwapRateTooBig, fmt.Sprintf("max: %f%% actual: %f%%", max, limit))
	}
)
