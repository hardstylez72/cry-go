import {Network, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskVeSyncSwap from "@/components/tasks/SWAPS/VeSync/Block.vue";
import MenuTaskVeSyncSwap from "@/components/tasks/SWAPS/VeSync/Menu.vue";
import {tokenSwapPair} from "@/components/helper";

export const VeSyncSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.veSyncSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: TaskVeSyncSwap,
  menu: MenuTaskVeSyncSwap,
  service: {
    name: 'VeSync',
    link: 'https://app.vesync.finance/swap',
    img: '/icons/vesync.svg',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM]),
  swapParis: [
    tokenSwapPair(Network.ZKSYNCERA, Token.ETH, Token.USDC),
    tokenSwapPair(Network.ZKSYNCERA, Token.USDC, Token.ETH),
  ]
}
