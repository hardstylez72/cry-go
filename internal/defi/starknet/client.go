package starknet

import (
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/gateway"
	"github.com/dontpanicdao/caigo/types"
)

type Client struct {
	Provider *gateway.GatewayProvider
}

func NewClient() (*Client, error) {

	provider := gateway.NewProvider(gateway.WithChain(gateway.MAINNET_BASE))
	pk, err := caigo.Curve.GetRandomPrivateKey()
	if err != nil {
		return nil, err
	}
	generatedPk := types.BigToHex(pk)
	println("generated pk: ", generatedPk)

	//x, y, err := caigo.Curve.PrivateToPoint(pk)
	//if err != nil {
	//	return nil, err
	//}

	types.BigToHex(pk)

	//caigo.NewRPCAccount()

	// pk <-> string
	pkInt := types.SNValToBN("0x02faedf3db545d3e9138348efde40ec2a3c3502167fb5ef83e95b978b6ff940b")
	//pkDecoded := types.BigToHex(pkInt)

	x, y, err := caigo.Curve.PrivateToPoint(pkInt)
	if err != nil {
		return nil, err
	}

	println("x: ", types.BigToHex(x))
	println("y: ", types.BigToHex(y))
	//caigo.Curve.i

	// my addr 0x027f52a9401acf44152e469b7e2b2a02cce4a72026451ee4e7b91230b3a7c9f5
	// my pk   0x02faedf3db545d3e9138348efde40ec2a3c3502167fb5ef83e95b978b6ff940b
	//println("secret key: ", pkDecoded)
	return &Client{
		Provider: provider,
	}, nil
}
