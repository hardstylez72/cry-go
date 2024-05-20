import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Menu from "@/components/tasks/LP/LPZkLend/Menu.vue";
import Block from "@/components/tasks/LP/LPZkLend/block.vue";

export const LPZkLendLPSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.zkLendLPTask
    return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p.tokens.map(t => t.token).join(' - ')})`
  },
  service: {
    name: 'ZkLend',
    img: '/icons/zklend.ico',
    link: 'https://app.zklend.com/markets',
    op: 'lp',
  },
  networks: new Set<Network>([
    Network.StarkNet,
  ]),
  job: TaskJob.LP,
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.StarkNet])
}
