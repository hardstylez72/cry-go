import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import swap10k from "@/components/tasks/SWAPS/Swap10k/Block.vue";
import swap10kMenu from "@/components/tasks/SWAPS/Swap10k/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const Swap10kSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.swap10k
    return ''
  },
  swap: (task: Task) => {
    const p = task.swap10k
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: swap10k,
  menu: swap10kMenu,
  service: {
    name: '10kswap',
    link: 'https://10kswap.com/swap',
    img: '/icons/10kswap.png',
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
