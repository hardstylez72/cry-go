import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/BridgeAcross/Block.vue";
import Menu from "@/components/tasks/BridgeAcross/Menu.vue";

export const AcrossBridgeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.acrossBridgeTask
    return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'Across',
    link: 'https://across.to/',
    img: '/icons/across.png',
    op: 'bridge',
  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
