import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import JediSwap from "@/components/tasks/SWAPS/Jedi/Block.vue";
import MenuJediSwap from "@/components/tasks/SWAPS/Jedi/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const JediSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.jediSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: JediSwap,
  menu: MenuJediSwap,
  service: {
    name: 'JediSwap',
    link: 'https://app.jediswap.xyz/#/swap',
    img: '/icons/jediswap.png',
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
