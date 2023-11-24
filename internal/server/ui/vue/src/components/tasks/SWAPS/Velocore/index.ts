import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskVelocoreSwap from "@/components/tasks/SWAPS/Velocore/Block.vue";
import MenuTaskVelocoreSwap from "@/components/tasks/SWAPS/Velocore/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const VelocoreSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.velocoreSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.velocoreSwapTask
    return {from: p.fromToken, to: p.toToken, network: p.network}
  },
  component: TaskVelocoreSwap,
  menu: MenuTaskVelocoreSwap,
  service: {
    name: 'Velocore',
    link: 'https://app.velocore.xyz/swap',
    img: '/icons/velocore.png',
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

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.ETH),

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.VC),
    tokenSwapPair(Network.ZKSYNCERA, Token.VC, Token.ETH),

    // Tokens
    tokenSwapPair(Network.ZKSYNCERA, Token.USDp, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.USDp),

    tokenSwapPair(Network.ZKSYNCERA, Token.USDp, Token.BUSD),
    tokenSwapPair(Network.ZKSYNCERA, Token.BUSD, Token.USDp),

  ]
}
