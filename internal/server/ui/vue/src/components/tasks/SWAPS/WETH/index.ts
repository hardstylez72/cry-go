import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuTaskWethSwap from "@/components/tasks/SWAPS/WETH/Menu.vue";
import TaskWethSwap from "@/components/tasks/SWAPS/WETH/Block.vue";

export const WETHSwapSpec: TaskSpec = {
  deprecated: true,
  menu: MenuTaskWethSwap,
  component: TaskWethSwap,
  canBeEstimated: true,
  descFn(task) {
    return task.wETHTask?.wrap ? ' (wrap)' : ' (unwrap)'
  },
  service: {
    link: '',
    img: '',
    name: 'WETH <> ETH',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
