package layerzero

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

const (
	TypeFuncSwap uint8 = 1
)

const LayerZeroBoostPercent = 5

var LayerZeroChainMap = map[v1.Network]uint16{
	v1.Network_Etherium:     101,
	v1.Network_BinanaceBNB:  102,
	v1.Network_AVALANCHE:    106,
	v1.Network_POLIGON:      109,
	v1.Network_ARBITRUM:     110,
	v1.Network_OPTIMISM:     111,
	v1.Network_GOERLIETH:    154,
	v1.Network_ZKSYNCERA:    165,
	v1.Network_Meter:        176,
	v1.Network_Tenet:        173,
	v1.Network_Canto:        159,
	v1.Network_ArbitrumNova: 175,
	v1.Network_PolygonZKEVM: 158,
	v1.Network_Fantom:       112,
	v1.Network_Base:         184,
	v1.Network_opBNB:        202,
	v1.Network_Core:         153,
	v1.Network_Conflux:      212,
	v1.Network_Fuse:         138,
	v1.Network_Loot:         197,
	v1.Network_Klaytn:       150,

	//SBNameMetis:         151,
}

func MakeLayerZeroAdapterParams(a uint16, b *big.Int) ([]byte, error) {

	structThing, err := abi.NewType("tuple", "lzAdapterParams", []abi.ArgumentMarshaling{
		{Name: "a", Type: "uint16"},
		{Name: "b", Type: "uint256"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "abi.NewType")
	}

	record := struct {
		A uint16
		B *big.Int
	}{
		A: a,
		B: b,
	}

	args := abi.Arguments{
		{Type: structThing, Name: "lzAdapterParams"},
	}
	adapterParams, err := args.Pack(&record)
	if err != nil {
		return nil, errors.Wrap(err, "args.Pack(&record)")
	}

	// 34 байтов
	// https://github.com/LayerZero-Labs/solidity-examples/blob/main/contracts/lzApp/LzApp.sol#L64
	adapterParams = adapterParams[30:]

	return adapterParams, nil
}

type GetStargateBridgeFeeReq struct {
	ToChain v1.Network
	WA      common.Address
}

type GetStargateBridgeFeeRes struct {
	Fee1 *big.Int
	Fee2 *big.Int
}

type AdapterParamsV2 struct {
	Version      uint16
	GasAmount    *big.Int
	NativeForDst *big.Int
	AddressOnDst common.Address
}

func MakeLayerZeroAdapterParamsV2(p AdapterParamsV2) ([]byte, error) {

	vesion := fmt.Sprintf("%0.4X", p.Version)
	gasAmount := fmt.Sprintf("%0.64X", p.GasAmount)
	NativeForDst := fmt.Sprintf("%0.64X", p.NativeForDst)
	addr := p.AddressOnDst.String()[2:]

	out := vesion + gasAmount + NativeForDst + addr

	return hex.DecodeString(out)
}
