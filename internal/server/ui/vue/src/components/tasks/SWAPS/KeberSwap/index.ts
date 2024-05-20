import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import OdosSwap from "@/components/tasks/SWAPS/KeberSwap/Block.vue";
import MenuOdosSwap from "@/components/tasks/SWAPS/KeberSwap/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const KyberSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.kyberSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.kyberSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: OdosSwap,
  menu: MenuOdosSwap,
  service: {
    name: 'KyberSwap',
    link: 'https://kyberswap.com/swap/zksync/eth-to-usdc',
    img: '/icons/kyberSwap.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [

    // ETH
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.ETH),

    // Tokens

    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),

  ]
}
