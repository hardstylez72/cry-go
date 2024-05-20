import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import MenuSnapshotTask from "@/components/tasks/OTHER/Snapshot/Menu.vue";
import TaskSnapshotVote from "@/components/tasks/OTHER/Snapshot/Block.vue";

export const SnapshotSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: false,
  menu: MenuSnapshotTask,
  component: TaskSnapshotVote,
  descFn(task) {
    let p = task.snapshotVoteTask
    return ` (space: ${p?.space})`
  },
  service: {
    name: 'Snapshot.org',
    img: '/icons/snapshot.png',
    link: 'https://snapshot.org/#/',
    op: 'vote',
  },
  job: TaskJob.Other,
  networks: new Set<Network>([Network.ARBITRUM]),
  airdrops: new Set<Airdrop>(Universal),
  profileType: new Set([ProfileType.EVM])
}
