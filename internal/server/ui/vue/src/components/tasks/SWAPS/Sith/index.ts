import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import SithSwap from "@/components/tasks/SWAPS/Sith/Block.vue";
import MenuSithSwap from "@/components/tasks/SWAPS/Sith/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SithSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.sithSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.sithSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: SithSwap,
  menu: MenuSithSwap,
  service: {
    name: 'SythSwap',
    link: 'https://sithswap.com/',
    img: '/icons/sithswap.png',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet]),
  swapParis: [
    tokenSwapPair(Network.StarkNet, Token.ETH, Token.USDC),
    tokenSwapPair(Network.StarkNet, Token.USDC, Token.ETH),
  ]
}
