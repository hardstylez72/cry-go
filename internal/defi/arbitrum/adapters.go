package arbitrum

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/across"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/bridge/stargate"
	"github.com/hardstylez72/cry/internal/defi/generic"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func FW1[FUNC func(context.Context, ARG) (RETURNS, error), ARG, RETURNS any](ctx context.Context, f FUNC, t ARG) (res RETURNS, err error) {

	res, err = f(ctx, t)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds for gas") {
			return res, &defi.ErrOutOfGas{
				N:     v1.Network_ARBITRUM,
				Token: v1.Token_ETH,
			}
		} else {
			return res, err
		}
	}
	return res, nil
}

func FW2[FUNC func(context.Context, ARG0, ARG1) (RETURNS, error), ARG0, ARG1, RETURNS any](ctx context.Context, f FUNC, a ARG0, b ARG1) (res RETURNS, err error) {

	res, err = f(ctx, a, b)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds for gas") {
			return res, &defi.ErrOutOfGas{
				N:     v1.Network_ARBITRUM,
				Token: v1.Token_ETH,
			}
		}
	}
	return res, nil
}

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {
	return FW1(ctx, c.defi.GetBalance, req)
}

func (c *Client) TxViewFn(id string) string {
	return c.defi.TxViewFn(id)
}

func (c *Client) GetNetworkToken() defi.Token {
	return c.defi.GetNetworkToken()
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	return FW1(ctx, c.defi.Transfer, r)
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	return c.defi.WaitTxComplete(ctx, tx)
}

func (c *Client) TestNetBridgeSwap(ctx context.Context, req *defi.TestNetBridgeSwapReq) (*defi.TestNetBridgeSwapRes, error) {
	return FW1(ctx, c.defi.TestNetBridgeSwap, req)
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {
	return FW1(ctx, c.defi.OrbiterBridge, req)
}
func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	return FW2(ctx, c.defi.Swapper, req, taskType)
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	return c.defi.GetPublicKey(pk, subType)
}

func (c *Client) Network() v1.Network {
	return c.defi.Cfg.Network
}

func (c *Client) GetNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
	return FW2(ctx, c.defi.GetNFTId, txHash, owner)
}

func (c *Client) Mint(ctx context.Context, req *defi.SimpleReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	g := generic.NewMinter(c.defi)
	return FW2(ctx, g.Mint, req, taskType)
}

func (c *Client) BridgeNft(ctx context.Context, req *defi.BridgeNFTReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_MerklyMintAndBridgeNFT:
		return FW1(ctx, c.defi.BridgeNftMerkly, req)
	default:
		return nil, errors.New("not supported")
	}
}

func (c *Client) Bridge(ctx context.Context, req *defi.DefaultBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_AcrossBridge:
		b := across.NewAcrossBridge(c.defi)
		return FW1(ctx, b.Bridge, req)
	case v1.TaskType_StargateBridge:
		return stargate.NewBridge(c.defi).Bridge(ctx, req)
	default:
		return nil, errors.New("bridge unsupported")
	}
}
func (c *Client) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	switch taskType {
	case v1.TaskType_AcrossBridge:
		b := across.NewAcrossBridge(c.defi)
		return b.WaitForConfirm(ctx, txId, receiver)
	case v1.TaskType_StargateBridge:
		return stargate.NewBridge(c.defi).WaitForConfirm(ctx, txId, taskType, receiver)
	default:
		return errors.New("bridge unsupported")
	}
}
