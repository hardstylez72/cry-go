import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/TraderJoe/Block.vue";
import Menu from "@/components/tasks/TraderJoe/Menu.vue";

export const TraderJoeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.traderJoeSwapTask
    return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'TraderJoe',
    link: 'https://traderjoexyz.com/arbitrum/trade',
    img: '/icons/traderjoe.png',
    op: 'trade',
  },
  job: TaskJob.Swap,
  networks: new Set<Network>([Network.ARBITRUM]),
  airdrops: new Set<Airdrop>(),
  profileType: new Set([ProfileType.EVM])
}
