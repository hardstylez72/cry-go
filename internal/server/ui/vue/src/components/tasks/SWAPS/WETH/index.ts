import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuTaskWethSwap from "@/components/tasks/SWAPS/WETH/Menu.vue";
import TaskWethSwap from "@/components/tasks/SWAPS/WETH/Block.vue";
import {tokenSwapPair} from "@/components/helper";

export const WETHSwapSpec: TaskSpec = {
  deprecated: false,
  menu: MenuTaskWethSwap,
  component: TaskWethSwap,
  canBeEstimated: true,
  descFn(task) {
    const p = task.wethSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.wethSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  service: {
    link: '',
    img: '',
    name: 'WETH',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.WETH),
    tokenSwapPair(Network.ZKSYNCERA, Token.WETH, Token.ETH),

    tokenSwapPair(Network.Base, Token.ETH, Token.WETH),
    tokenSwapPair(Network.Base, Token.WETH, Token.ETH),
  ]
}
