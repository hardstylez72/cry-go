import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskIzumiSwap from "@/components/tasks/SWAPS/IZUMI/Block.vue";
import MenuTaskIzumiSwap from "@/components/tasks/SWAPS/IZUMI/Menu.vue";
import TaskSpaceFiSwap from "@/components/tasks/SWAPS/SpaceFI/Block.vue";
import MenuTaskSpaceFiSwap from "@/components/tasks/SWAPS/SpaceFI/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const SpaceFISwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.spaceFiSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  swap: (task: Task) => {
    const p = task.spaceFiSwapTask
    return {from: p?.fromToken, to: p?.toToken, network: p?.network}
  },
  component: TaskSpaceFiSwap,
  menu: MenuTaskSpaceFiSwap,
  service: {
    name: 'SpaceFi',
    link: 'https://swap-zksync.spacefi.io/#/swap',
    img: '/icons/spacefi.png',
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

    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.SPACE),
    tokenSwapPair(Network.ZKSYNCERA, Token.SPACE, Token.ETH),


    // Tokens

  ]
}
