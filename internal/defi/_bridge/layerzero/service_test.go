package layerzero

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/hardstylez72/cry/internal/tests"
	"github.com/stretchr/testify/assert"
)

func TestXuy(t *testing.T) {

	//data := "00020000000000000000000000000000000000000000000000000000000000030d400000000000000000000000000000000000000000000000000de0b6b3a76400004a6e7c137a6691d55693ca3bc7e5c698d9d43815"
	//
	//p := []byte(data)
	//
	//structThing, err := abi.NewType("tuple", "lzAdapterParams", []abi.ArgumentMarshaling{
	//	{Name: "version", Type: "uint16"},
	//	{Name: "gasAmount", Type: "uint256"},
	//	{Name: "nativeForDst", Type: "uint256"},
	//	{Name: "addressOnDst", Type: "address"},
	//})
	//assert.NoError(t, err)
	//
	//args := abi.Arguments{
	//	{Type: structThing, Name: "lzAdapterParams"},
	//}
	//adapterParams, err := args.Unpack(p)
	//assert.NoError(t, err)
	//
	//assert.NotNil(t, adapterParams)

	t.Run("anus", func(t *testing.T) {
		b, err := MakeLayerZeroAdapterParamsV2(AdapterParamsV2{
			Version:      2,
			GasAmount:    big.NewInt(200000),
			NativeForDst: big.NewInt(1000000000000000000),
			AddressOnDst: tests.GetConfig().Wallet,
		})
		assert.NoError(t, err)
		boba := hex.EncodeToString(b)

		assert.Equal(t, boba, "00020000000000000000000000000000000000000000000000000000000000030d400000000000000000000000000000000000000000000000000de0b6b3a76400004a6e7c137a6691d55693ca3bc7e5c698d9d43815")

		// 0x0002
		//0000000000000000000000000000000000000000000000000000000000030d4000000000000000000000000000000000000000000000000000b1a2bc2ec50000
		//4a6e7c137a6691d55693ca3bc7e5c698d9d43815
	})

}
