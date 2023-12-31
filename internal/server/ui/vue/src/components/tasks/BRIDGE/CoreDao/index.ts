import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/BRIDGE/CoreDao/Block.vue";
import Menu from "@/components/tasks/BRIDGE/CoreDao/Menu.vue";

export const CoreDaoBridgeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.coreDaoBridge
    return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
  },
  bridge: (task: Task) => {
    const p = task.coreDaoBridge
    return {
      from: p?.fromToken, to: p?.toToken, fromNetwork: p?.fromNetwork, toNetwork: p?.toNetwork
    }
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'Core DAO',
    link: 'https://bridge.coredao.org/bridge',
    img: '/icons/networks/core.png',
    op: 'bridge',
  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([
    Network.BinanaceBNB,
    Network.Core,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
  profileType: new Set([ProfileType.EVM])
}
