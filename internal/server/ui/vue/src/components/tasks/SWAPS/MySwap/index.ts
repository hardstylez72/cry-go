import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MySwap from "@/components/tasks/SWAPS/MySwap/Block.vue";
import MenuMySwap from "@/components/tasks/SWAPS/MySwap/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const MySwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.mySwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: MySwap,
  menu: MenuMySwap,
  service: {
    name: 'MySwap',
    link: 'https://www.myswap.xyz/#/swap',
    img: '/icons/myswap.svg',
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
