import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import StrarkNetBridge from "@/components/tasks/BRIDGE/StarkNet/Block.vue";
import MenuStarkNetBridge from "@/components/tasks/BRIDGE/StarkNet/Menu.vue";

export const BridgeStarknetSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.starkNetBridgeTask
    return ''
  },
  component: StrarkNetBridge,
  menu: MenuStarkNetBridge,
  service: {
    name: 'StarkNet liquidity bridge',
    link: 'https://starkgate.starknet.io/',
    img: '/icons/starknet.svg',
    op: 'liquidity bridge',
  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([Network.Etherium, Network.StarkNet]),
  airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
  profileType: new Set([ProfileType.EVM])
}
