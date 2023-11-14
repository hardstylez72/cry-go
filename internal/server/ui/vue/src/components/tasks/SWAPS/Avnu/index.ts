import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/SWAPS/Avnu/Block.vue";
import Menu from "@/components/tasks/SWAPS/Avnu/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SwapAvnuSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.avnuSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'AvnuSwap',
    link: 'https://app.avnu.fi/',
    img: '/icons/avnu.svg',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet]),
  swapParis: [
    tokenSwapPair(Network.StarkNet, Token.ETH, Token.USDC),
    tokenSwapPair(Network.StarkNet, Token.USDC, Token.ETH),

    tokenSwapPair(Network.StarkNet, Token.ETH, Token.USDT),
    tokenSwapPair(Network.StarkNet, Token.USDT, Token.ETH),

    tokenSwapPair(Network.StarkNet, Token.USDC, Token.USDT),
    tokenSwapPair(Network.StarkNet, Token.USDT, Token.USDC),
  ]
}
