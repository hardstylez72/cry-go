import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import ProtossSwap from "@/components/tasks/SWAPS/Protoss/Block.vue";
import MenuProtossSwap from "@/components/tasks/SWAPS/Protoss/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const ProtossSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.protosSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.protosSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: ProtossSwap,
  menu: MenuProtossSwap,
  service: {
    name: 'ProtossSwap',
    link: 'https://www.protoss.org/',
    img: '/icons/protossSwap.png',
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
