package defi

import (
	"math/big"

	"github.com/hardstylez72/cry/internal/lib"
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

func slippageSpecial(v *big.Int, s string) (*big.Int, error) {
	f, err := lib.StringToFloat(s)
	if err != nil {
		return nil, errors.Wrap(err, "invalid slippage value: "+s)
	}

	if f > 100 || f < 0.001 {
		return nil, errors.New("slippage must be between [0.001:100]")
	}

	v10 := big.NewInt(0).Mul(v, big.NewInt(10000))
	f10 := big.NewFloat(0).Mul(big.NewFloat(f), big.NewFloat(10000))
	i10, _ := f10.Int(nil)

	slippageValue := big.NewInt(0).Mul(v10, i10)

	slippageValue = big.NewInt(0).Div(slippageValue, big.NewInt(10e9))

	return big.NewInt(0).Sub(v, slippageValue), nil

}

func Slippage(v *big.Int, slippagePercent SlippagePercent) (*big.Int, error) {
	return slippageSpecial(v, slippagePercent)
}

var SlippageMap = map[v1.TaskType]SlippagePercent{
	v1.TaskType_EzkaliburSwap:  SlippagePercent2,
	v1.TaskType_IzumiSwap:      SlippagePercent2,
	v1.TaskType_MaverickSwap:   SlippagePercent2,
	v1.TaskType_MuteioSwap:     SlippagePercent2,
	v1.TaskType_SpaceFISwap:    SlippagePercent2,
	v1.TaskType_StargateBridge: SlippagePercent2,
	v1.TaskType_SyncSwap:       SlippagePercent2,
	v1.TaskType_VelocoreSwap:   SlippagePercent2,
	v1.TaskType_VeSyncSwap:     SlippagePercent2,
	v1.TaskType_ZkSwap:         SlippagePercent2,
	v1.TaskType_Swap10k:        SlippagePercent2,
	v1.TaskType_PancakeSwap:    SlippagePercent2,
	v1.TaskType_SithSwap:       SlippagePercent2,
	v1.TaskType_JediSwap:       SlippagePercent2,
	v1.TaskType_MySwap:         SlippagePercent2,
	v1.TaskType_ProtossSwap:    SlippagePercent2,
	v1.TaskType_OdosSwap:       SlippagePercent05,
	v1.TaskType_AvnuSwap:       SlippagePercent2,
	v1.TaskType_FibrousSwap:    SlippagePercent2,
	v1.TaskType_KyberSwap:      SlippagePercent05,
	v1.TaskType_EkuboSwap:      SlippagePercent05,
}
