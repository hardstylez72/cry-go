package arbitrum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	"github.com/hardstylez72/cry/internal/defi/across"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/nft/merkly"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {
	return c.defi.GetBalance(ctx, req)
}

func (c *Client) TxViewFn(id string) string {
	return c.defi.TxViewFn(id)
}

func (c *Client) StargateBridgeSwap(ctx context.Context, req *defi.DefaultBridgeReq) (*bozdo.DefaultRes, error) {
	return c.defi.StargateBridgeSwap(ctx, req)
}
func (c *Client) GetStargateBridgeFee(ctx context.Context, req *defi.GetStargateBridgeFeeReq) (*defi.GetStargateBridgeFeeRes, error) {
	return c.defi.GetStargateBridgeFee(ctx, req)
}
func (c *Client) GetNetworkToken() defi.Token {
	return c.defi.GetNetworkToken()
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	return c.defi.Transfer(ctx, r)
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	return c.defi.WaitTxComplete(ctx, tx)
}

func (c *Client) TestNetBridgeSwap(ctx context.Context, req *defi.TestNetBridgeSwapReq) (*defi.TestNetBridgeSwapRes, error) {
	return c.defi.TestNetBridgeSwap(ctx, req)
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {
	return c.defi.OrbiterBridge(ctx, req)
}
func (c *Client) Swap(ctx context.Context, req *defi.DefaultSwapReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	return c.defi.Swapper(ctx, req, taskType)
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	return c.defi.GetPublicKey(pk, subType)
}

func (c *Client) Network() v1.Network {
	return c.defi.Cfg.Network
}

func (c *Client) GetMerklyNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
	return c.defi.GetMerklyNFTId(ctx, txHash, owner)
}

func (c *Client) MerklyMintNft(ctx context.Context, req *merkly.MintNFTReq) (*bozdo.DefaultRes, error) {
	return c.defi.MerklyMintNft(ctx, req)
}

func (c *Client) MerklyBridgeNft(ctx context.Context, req *merkly.BridgeNFTReq) (*bozdo.DefaultRes, error) {
	return c.defi.MerklyBridgeNft(ctx, req)
}

func (c *Client) Bridge(ctx context.Context, req *defi.DefaultBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_AcrossBridge:
		b := across.NewAcrossBridge(c.defi)
		return b.Bridge(ctx, req)
	default:
		return nil, errors.New("bridge unsupported")
	}
}
func (c *Client) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	switch taskType {
	case v1.TaskType_AcrossBridge:
		b := across.NewAcrossBridge(c.defi)
		return b.WaitForConfirm(ctx, txId, receiver)
	default:
		return errors.New("bridge unsupported")
	}
}
