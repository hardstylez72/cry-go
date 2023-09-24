import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/SwapFibrous/Block.vue";
import Menu from "@/components/tasks/SwapFibrous/Menu.vue";

export const SwapFibrousSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.fibrousSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'Fibrous',
    link: 'https://app.fibrous.finance/',
    img: '/icons/fibrous.svg',
    op: 'swap',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet])
}
