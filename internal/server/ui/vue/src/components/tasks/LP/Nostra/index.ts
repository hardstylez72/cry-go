import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Menu from "@/components/tasks/LP/Nostra/Menu.vue";
import Block from "@/components/tasks/LP/Nostra/block.vue";

export const LPNostraLPSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.nostraLPTask
    return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p?.tokens.map(t => t.token).join(' - ')})`
  },
  service: {
    name: 'Nostra',
    img: '/icons/nostra.svg',
    link: 'https://app.nostra.finance/',
    op: 'lp',
  },
  networks: new Set<Network>([
    Network.StarkNet,
  ]),
  job: TaskJob.LP,
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet])
}
