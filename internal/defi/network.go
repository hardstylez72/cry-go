package defi

import (
	"context"
	"math/big"
)

func (c *EtheriumClient) GetNetworkId(ctx context.Context) (*big.Int, error) {
	id, err := c.Cli.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	c.Cfg.NetworkId = id
	return id, nil
}
func (c *EtheriumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.Cli.SuggestGasPrice(ctx)
}

func (c *EtheriumClient) GetPublicKey(pk string) (string, error) {
	t, err := newWalletTransactor(pk)
	if err != nil {
		return "", err
	}
	return t.WalletAddrHR, nil
}
