import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskZkSwap from "@/components/tasks/SWAPS/ZkSwap/Block.vue";
import MenuTaskZkSwap from "@/components/tasks/SWAPS/ZkSwap/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const ZkSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.zkSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.zkSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: TaskZkSwap,
  menu: MenuTaskZkSwap,
  service: {
    name: 'ZkSwap',
    link: 'https://zkswap.finance/swap',
    img: '/icons/zkswap.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),
  ]
}
