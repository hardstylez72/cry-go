import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Dmail from "@/components/tasks/OTHER/Dmail/Block.vue";
import MenuDmail from "@/components/tasks/OTHER/Dmail/Menu.vue";

export const DmailSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.dmailTask
    return ''
  },
  simple(task) {
    const p = task.dmailTask
    return {network: p.network}
  },
  component: Dmail,
  menu: MenuDmail,
  service: {
    name: 'Dmail',
    link: 'https://dmail.ai/',
    img: '/icons/dmail.ico',
    op: 'email send',
  },
  job: TaskJob.Simple,
  networks: new Set<Network>([Network.ZKSYNCERA, Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.StarkNet]),
  profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
}
