package task

import (
	"fmt"
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func ErrGasIsOverMax(max, estimated string) error {
	return errors.Wrap(ErrHighGas, fmt.Sprintf("Текущий газ (%s USD) выше заданного ограничения (%s USD)", estimated, max))
}

var (
	ErrHighGas                       = errors.New("Высокий газ")
	ErrUserHasNoBalance              = errors.New("Недостаточный баланс. Перейдите в раздел биллинг")
	ErrProfileHasNoConnectedOkexAddr = errors.New("Профиль не соединен с okex суб-аккаунтом")
	ErrTransactionIsNotReady         = errors.New("transaction is not ready")
	ErrAccountIsZero                 = errors.New("profile has Insufficient balance")
	ErrSwapRateTooBig                = errors.New("Недопустимое расхождение с курсом Binance")
	ErrProfileHasInsufficientBalance = func(token v1.Token, want, have *big.Int) error {
		return errors.Wrap(ErrAccountIsZero, fmt.Sprintf("want: %s have: %s of %s", want.String(), have.String(), token.String()))
	}
	ErrSwapRateBiggerThenAllowed = func(limit, max float64) error {
		return errors.Wrap(ErrSwapRateTooBig, fmt.Sprintf("Ограничение: %f%% Текущее: %f%%", max, limit))
	}
)
