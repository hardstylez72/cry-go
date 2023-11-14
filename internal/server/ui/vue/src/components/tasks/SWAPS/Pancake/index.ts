import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskPancakeSwap from "@/components/tasks/SWAPS/Pancake/Block.vue";
import PancakeSwap from "@/components/tasks/SWAPS/Pancake/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const PancakeSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.pancakeSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: TaskPancakeSwap,
  menu: PancakeSwap,
  service: {
    name: 'PancakeSwap',
    link: 'https://pancakeswap.finance/',
    img: '/icons/pancake.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.USDC),
    tokenSwapPair(Network.Base, Token.USDC, Token.ETH),

    tokenSwapPair(Network.Base, Token.USDCBridged, Token.ETH),
    tokenSwapPair(Network.Base, Token.ETH, Token.USDCBridged),

    tokenSwapPair(Network.Base, Token.USDCBridged, Token.USDC),
    tokenSwapPair(Network.Base, Token.USDC, Token.USDCBridged),
  ]
}
