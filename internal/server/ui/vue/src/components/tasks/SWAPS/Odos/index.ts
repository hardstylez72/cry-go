import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import OdosSwap from "@/components/tasks/SWAPS/Odos/Block.vue";
import MenuOdosSwap from "@/components/tasks/SWAPS/Odos/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SwapOdosSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.odosSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: OdosSwap,
  menu: MenuOdosSwap,
  service: {
    name: 'Odos',
    link: 'https://app.odos.xyz',
    img: '/icons/odos.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [

    tokenSwapPair(Network.Base, Token.ETH, Token.WETH),
    tokenSwapPair(Network.Base, Token.WETH, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.USDC),
    tokenSwapPair(Network.Base, Token.USDC, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.USDCBridged),
    tokenSwapPair(Network.Base, Token.USDCBridged, Token.ETH),

    tokenSwapPair(Network.Base, Token.USDC, Token.USDCBridged),
    tokenSwapPair(Network.Base, Token.USDCBridged, Token.USDC),

    // ETH
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.ETH),

    // Tokens
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.BUSD),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),

    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.BUSD),
  ]
}