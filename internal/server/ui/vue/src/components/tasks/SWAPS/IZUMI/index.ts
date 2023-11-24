import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskIzumiSwap from "@/components/tasks/SWAPS/IZUMI/Block.vue";
import MenuTaskIzumiSwap from "@/components/tasks/SWAPS/IZUMI/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const IZUMISwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.izumiSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.izumiSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: TaskIzumiSwap,
  menu: MenuTaskIzumiSwap,
  service: {
    name: 'IZUMI',
    link: 'https://izumi.finance/trade/swap',
    img: '/icons/izumi.svg',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    // ETH

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.IZI),
    tokenSwapPair(Network.ZKSYNCERA, Token.IZI, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.WETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.WETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.WETH, Token.IZI),
    tokenSwapPair(Network.ZKSYNCERA, Token.IZI, Token.WETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDT),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.ETH),

    // Tokens

    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDT),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDT, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDT),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDC),
  ]
}
