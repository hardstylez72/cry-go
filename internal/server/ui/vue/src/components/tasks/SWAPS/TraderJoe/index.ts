import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/SWAPS/TraderJoe/Block.vue";
import Menu from "@/components/tasks/SWAPS/TraderJoe/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const TraderJoeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.traderJoeSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'TraderJoe',
    link: 'https://traderjoexyz.com/arbitrum/trade',
    img: '/icons/traderjoe.png',
    op: 'trade',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ARBITRUM]),
  airdrops: new Set<Airdrop>(),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.ARBITRUM, Token.STG, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ARBITRUM, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ARBITRUM, Token.USDT, Token.USDC),
    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.USDT),

    tokenSwapPair(Network.ARBITRUM, Token.USDC, Token.USDCBridged),
    tokenSwapPair(Network.ARBITRUM, Token.USDCBridged, Token.USDC),

    tokenSwapPair(Network.ARBITRUM, Token.ETH, Token.USDCBridged),
    tokenSwapPair(Network.ARBITRUM, Token.USDCBridged, Token.ETH),
  ]
}
