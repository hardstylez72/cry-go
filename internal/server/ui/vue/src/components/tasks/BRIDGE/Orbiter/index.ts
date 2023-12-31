import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuTaskOrbiterBridge from "@/components/tasks/BRIDGE/Orbiter/Menu.vue";
import TaskOrbiterBridge from "@/components/tasks/BRIDGE/Orbiter/Block.vue";

export const OrbiterBridgeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: MenuTaskOrbiterBridge,
  component: TaskOrbiterBridge,
  descFn(task) {
    let p = task.orbiterBridgeTask
    return ` (from ${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
  },
  bridge: (task: Task) => {
    const p = task.orbiterBridgeTask
    return {
      from: p?.fromToken, to: p?.toToken, fromNetwork: p?.fromNetwork, toNetwork: p?.toNetwork
    }
  },
  service: {
    name: 'Orbiter',
    img: '/icons/orbiter.ico',
    link: 'https://www.orbiter.finance/',
    op: 'bridge',
  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([
    Network.ZKSYNCERA,
    Network.ARBITRUM,
    Network.Base,
    Network.Linea,
    Network.Zora,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.Base, Airdrop.Orbiter, Airdrop.ZkSync, Airdrop.LayerZero, Airdrop.Linea]),
  profileType: new Set([ProfileType.EVM])
}
