import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Menu from "@/components/tasks/LP/StakeSTG/Menu.vue";
import Block from "@/components/tasks/LP/StakeSTG/block.vue";

export const StakeSTGSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.stakeSTG
    return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p.tokens.map(t => t.token).join(' - ')})`
  },
  service: {
    name: 'Stake STG',
    img: '/icons/stg.svg',
    link: 'https://stargate.finance/',
    op: 'lp',
  },
  networks: new Set<Network>([
    Network.POLIGON,
  ]),
  job: TaskJob.LP,
  airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
  profileType: new Set([ProfileType.EVM])
}
