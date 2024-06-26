import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/BRIDGE/Across/Block.vue";
import Menu from "@/components/tasks/BRIDGE/Across/Menu.vue";

export const AcrossBridgeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.acrossBridgeTask
    return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
  },
  bridge: (task: Task) => {
    const p = task.acrossBridgeTask
    return {
      from: p?.fromToken, to: p?.toToken, fromNetwork: p?.fromNetwork, toNetwork: p?.toNetwork
    }
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
  networks: new Set<Network>([
    Network.ZKSYNCERA,
    Network.Base,
    Network.ARBITRUM,
    Network.OPTIMISM,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
