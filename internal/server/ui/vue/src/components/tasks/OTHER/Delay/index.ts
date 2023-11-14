import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, allNetworks, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import MenuSnapshotTask from "@/components/tasks/OTHER/Snapshot/Menu.vue";
import TaskSnapshotVote from "@/components/tasks/OTHER/Snapshot/Block.vue";
import MenuDelayTask from "@/components/tasks/OTHER/Delay/Menu.vue";
import TaskDelay from "@/components/tasks/OTHER/Delay/Block.vue";
import {humanDuration} from "@/components/helper";

export const DelaySpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: false,
  menu: MenuDelayTask,
  component: TaskDelay,
  descFn(task: Task): string {
    let p = task.delayTask
    return ` (${p?.random ? `${humanDuration(p?.minRandom)}:${humanDuration(p?.maxRandom)}` : humanDuration(p?.duration)})`
  },
  service: {
    name: 'Delay',
    link: '',
    img: '',
    op: '',
  },
  job: TaskJob.Other,
  networks: new Set<Network>(allNetworks),
  airdrops: new Set<Airdrop>(Universal),
  profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
}
