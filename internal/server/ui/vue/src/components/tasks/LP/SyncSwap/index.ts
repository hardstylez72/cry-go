import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Dmail from "@/components/tasks/OTHER/Dmail/Block.vue";
import MenuDmail from "@/components/tasks/OTHER/Dmail/Menu.vue";
import MenuTaskSyncSwapLP from "@/components/tasks/LP/SyncSwap/Menu.vue";
import TaskSyncSwapLP from "@/components/tasks/LP/SyncSwap/Block.vue";

export const SyncSwapLPSpec: TaskSpec = {
  deprecated: true,
  canBeEstimated: true,
  descFn(task) {
    let p = task.syncSwapLPTask
    return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p?.a}/${p?.b})`
  },
  menu: MenuTaskSyncSwapLP,
  component: TaskSyncSwapLP,
  service: {
    name: 'SyncSwapLP',
    img: '/icons/syncswap.svg',
    link: 'https://syncswap.xyz/',
    op: 'LP',
  },
  job: TaskJob.LP,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
