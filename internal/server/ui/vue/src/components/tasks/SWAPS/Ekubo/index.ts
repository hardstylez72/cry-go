import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/SWAPS/Ekubo/Block.vue";
import Menu from "@/components/tasks/SWAPS/Ekubo/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SwapEkuboSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.ekuboSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.ekuboSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'Ekubo',
    link: 'https://app.ekubo.org/',
    img: '/icons/ekubo.svg',
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
