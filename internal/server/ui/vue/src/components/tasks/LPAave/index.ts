import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Menu from "@/components/tasks/LPAave/Menu.vue";
import Block from "@/components/tasks/LPAave/block.vue";

export const LPAaveSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.aaveLPTask
    return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p.tokens.map(t => t.token).join(' - ')})`
  },
  service: {
    name: 'Aave',
    img: '/icons/aave.svg',
    link: 'https://app.aave.com/',
    op: 'lp',
  },
  networks: new Set<Network>([
    Network.Base,
  ]),
  job: TaskJob.LP,
  airdrops: new Set<Airdrop>([Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
