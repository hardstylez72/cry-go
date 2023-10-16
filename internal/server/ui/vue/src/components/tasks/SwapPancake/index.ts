import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskPancakeSwap from "@/components/tasks/SwapPancake/Block.vue";
import PancakeSwap from "@/components/tasks/SwapPancake/Menu.vue";

export const PancakeSwapSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.pancakeSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: TaskPancakeSwap,
  menu: PancakeSwap,
  service: {
    name: 'PancakeSwap',
    link: 'https://pancakeswap.finance/',
    img: '/icons/pancake.ico',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
