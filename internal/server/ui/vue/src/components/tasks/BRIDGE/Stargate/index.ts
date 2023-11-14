import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Menu from "@/components/tasks/BRIDGE/Stargate/Menu.vue";
import Block from "@/components/tasks/BRIDGE/Stargate/Block.vue";

export const StargateBridgeSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: Menu,
  component: Block,
  descFn(task) {
    let p = task.stargateBridgeTask
    return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
  },
  service: {
    name: 'Stargate',
    img: '/icons/stg.svg',
    link: 'https://stargate.finance/',
    op: 'bridge',
  },
  networks: new Set<Network>([
    Network.BinanaceBNB,
    Network.POLIGON,
    Network.AVALANCHE,
    Network.OPTIMISM,
    Network.ARBITRUM,
    Network.ZKSYNCERA]),
  job: TaskJob.Bridge,
  airdrops: new Set<Airdrop>([Airdrop.LayerZero, Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
