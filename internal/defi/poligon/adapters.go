package poligon

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hardstylez72/cry/internal/defi"
	core_dao "github.com/hardstylez72/cry/internal/defi/_bridge/core.dao"
	"github.com/hardstylez72/cry/internal/defi/_bridge/l2pass"
	"github.com/hardstylez72/cry/internal/defi/_bridge/merkly"
	"github.com/hardstylez72/cry/internal/defi/_bridge/stargate"
	stakestg "github.com/hardstylez72/cry/internal/defi/_swap/stakeSTG"
	"github.com/hardstylez72/cry/internal/defi/bozdo"
	"github.com/hardstylez72/cry/internal/defi/generic"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/pkg/errors"
)

func (c *Client) GetBalance(ctx context.Context, req *defi.GetBalanceReq) (*defi.GetBalanceRes, error) {
	return c.defi.GetBalance(ctx, req)
}

func (c *Client) TxViewFn(id string) string {
	return c.defi.TxViewFn(id)
}

func (c *Client) GetNetworkToken() defi.Token {
	return c.defi.GetNetworkToken()
}

func (c *Client) Transfer(ctx context.Context, r *defi.TransferReq) (*defi.TransferRes, error) {
	return c.defi.TransferV2(ctx, r)
}

func (c *Client) GetNetworkId() *big.Int {
	return c.NetworkId
}

func (c *Client) WaitTxComplete(ctx context.Context, tx string) error {
	return c.defi.WaitTxComplete(ctx, tx)
}

func (c *Client) OrbiterBridge(ctx context.Context, req *defi.OrbiterBridgeReq) (*defi.OrbiterBridgeRes, error) {
	return c.defi.OrbiterBridge(ctx, req)
}

func (c *Client) GetPublicKey(pk string, subType v1.ProfileSubType) (string, error) {
	return c.defi.GetPublicKey(pk, subType)
}

func (c *Client) Network() v1.Network {
	return c.defi.Cfg.Network
}

func (c *Client) GetNFTId(ctx context.Context, txHash common.Hash, owner common.Address) (*big.Int, error) {
	return c.defi.GetNFTId(ctx, txHash, owner)
}

func (c *Client) Mint(ctx context.Context, req *defi.SimpleReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	g := generic.NewMinter(c.defi)
	return g.Mint(ctx, req, taskType)
}

func (c *Client) BridgeNft(ctx context.Context, req *defi.BridgeNFTReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_MerklyMintAndBridgeNFT:
		return c.defi.BridgeNftMerkly(ctx, req)
	default:
		return nil, errors.New("not supported")
	}
}

func (c *Client) Bridge(ctx context.Context, req *defi.DefaultBridgeReq, taskType v1.TaskType) (*bozdo.DefaultRes, error) {
	switch taskType {
	case v1.TaskType_CoreDaoBridge:
		return core_dao.NewBridge(c.defi).Bridge(ctx, req)
	case v1.TaskType_StargateBridge:
		return stargate.NewBridge(c.defi).Bridge(ctx, req)
	case v1.TaskType_MerklyRefuel:
		return merkly.NewBridge(c.defi, common.HexToAddress("0x0e1f20075c90ab31fc2dd91e536e6990262cf76d")).Bridge(ctx, req)
	case v1.TaskType_L2PassRefuel:
		return l2pass.NewBridge(c.defi, common.HexToAddress("0x222228060e7efbb1d78bb5d454581910e3922222")).Bridge(ctx, req)
	default:
		return nil, errors.New("bridge unsupported")
	}
}
func (c *Client) WaitForConfirm(ctx context.Context, txId string, taskType v1.TaskType, receiver string) error {
	switch taskType {
	case v1.TaskType_CoreDaoBridge:
		return core_dao.NewBridge(c.defi).WaitForConfirm(ctx, txId, taskType, receiver)
	case v1.TaskType_StargateBridge:
		return stargate.NewBridge(c.defi).WaitForConfirm(ctx, txId, taskType, receiver)
	case v1.TaskType_MerklyRefuel:
		return merkly.NewBridge(c.defi, common.HexToAddress("0x0e1f20075c90ab31fc2dd91e536e6990262cf76d")).WaitForConfirm(ctx, txId)
	case v1.TaskType_L2PassRefuel:
		return l2pass.NewBridge(c.defi, common.HexToAddress("0x222228060e7efbb1d78bb5d454581910e3922222")).WaitForConfirm(ctx, txId)
	default:
		return errors.New("bridge unsupported")
	}
}

func (c *Client) LP(ctx context.Context, req *defi.LPReq, taskType v1.TaskType) (*defi.LPRes, error) {
	switch taskType {
	case v1.TaskType_StakeSTG:
		return stakestg.NewClient(c.defi, common.HexToAddress("0x3AB2DA31bBD886A7eDF68a6b60D3CDe657D3A15D")).LP(ctx, req)
	default:
		return nil, errors.New("unsupported task: " + taskType.String())
	}
}
