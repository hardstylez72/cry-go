import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskMaverickSwap from "@/components/tasks/SWAPS/Maverick/Block.vue";
import MenuTaskMaverickSwap from "@/components/tasks/SWAPS/Maverick/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const MaverickSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.maverickSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.maverickSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: TaskMaverickSwap,
  menu: MenuTaskMaverickSwap,
  service: {
    name: 'Maverick',
    img: '/icons/maverick.ico',
    link: 'https://app.mav.xyz/?chain=324',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.LayerZero]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    // ETH
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.LUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.LUSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.MAV),
    tokenSwapPair(Network.ZKSYNCERA, Token.MAV, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.ETH),

    //Tokens
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDC),

  ]
}
