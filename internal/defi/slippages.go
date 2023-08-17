package defi

import (
	"math/big"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

type SlippagePercent = string

const (
	SlippagePercent5    SlippagePercent = "5"
	SlippagePercent2    SlippagePercent = "2"
	SlippagePercent1    SlippagePercent = "1"
	SlippagePercent05   SlippagePercent = "0.5"
	SlippagePercent02   SlippagePercent = "0.2"
	SlippagePercent03   SlippagePercent = "0.3"
	SlippagePercent01   SlippagePercent = "0.1"
	SlippagePercent001  SlippagePercent = "0.01"
	SlippagePercentZero SlippagePercent = "0"
)

func Slippage(v *big.Int, slippagePercent SlippagePercent) (*big.Int, error) {
	switch slippagePercent {
	case SlippagePercentZero:
		return v, nil
	case SlippagePercent05:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(995)), nil
	case SlippagePercent01:
		prec := big.NewInt(0).Div(v, big.NewInt(1000))
		return big.NewInt(0).Mul(prec, big.NewInt(999)), nil
	case SlippagePercent001:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9999)), nil
	case SlippagePercent02:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9998)), nil
	case SlippagePercent03:
		prec := big.NewInt(0).Div(v, big.NewInt(10000))
		return big.NewInt(0).Mul(prec, big.NewInt(9997)), nil
	case SlippagePercent1:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(99)), nil
	case SlippagePercent2:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(98)), nil
	case SlippagePercent5:
		prec := big.NewInt(0).Div(v, big.NewInt(100))
		return big.NewInt(0).Mul(prec, big.NewInt(95)), nil
	}
	return nil, errors.New("invalid slippage: " + string(slippagePercent))
}

var SlippageMap = map[v1.TaskType]SlippagePercent{
	v1.TaskType_EzkaliburSwap:  SlippagePercent05,
	v1.TaskType_IzumiSwap:      SlippagePercentZero,
	v1.TaskType_MaverickSwap:   SlippagePercent2,
	v1.TaskType_MuteioSwap:     SlippagePercent01,
	v1.TaskType_SpaceFISwap:    SlippagePercent01,
	v1.TaskType_StargateBridge: SlippagePercent05,
	v1.TaskType_SyncSwap:       SlippagePercent01,
	v1.TaskType_VelocoreSwap:   SlippagePercent05,
	v1.TaskType_VeSyncSwap:     SlippagePercent01,
	v1.TaskType_ZkSwap:         SlippagePercent05,
	v1.TaskType_Swap10k:        SlippagePercent05,
	v1.TaskType_PancakeSwap:    SlippagePercent05,
	v1.TaskType_SithSwap:       SlippagePercent05,
	v1.TaskType_JediSwap:       SlippagePercent05,
}
