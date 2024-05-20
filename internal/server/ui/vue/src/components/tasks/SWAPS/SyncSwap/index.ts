import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuTaskSyncSwap from "@/components/tasks/SWAPS/SyncSwap/Menu.vue";
import TaskSyncSwap from "@/components/tasks/SWAPS/SyncSwap/Block.vue";
import {tokenSwapPair} from "@/components/helper";

export const SyncSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: MenuTaskSyncSwap,
  component: TaskSyncSwap,
  descFn(task) {
    let p = task.syncSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.syncSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  service: {
    name: 'SyncSwap',
    img: '/icons/syncswap.svg',
    link: 'https://syncswap.xyz/',
    op: 'swap',
  },
  networks: new Set<Network>([Network.ZKSYNCERA]),
  job: TaskJob.Swap,
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [

    // ETH
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.LUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.LUSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.LSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.LSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.MUTE),
    tokenSwapPair(Network.ZKSYNCERA, Token.MUTE, Token.ETH),

    //tokens
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDp, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDp),
  ]
}
