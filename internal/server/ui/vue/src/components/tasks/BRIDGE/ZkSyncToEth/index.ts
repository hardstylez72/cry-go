import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Dmail from "@/components/tasks/OTHER/Dmail/Block.vue";
import MenuDmail from "@/components/tasks/OTHER/Dmail/Menu.vue";
import MenuZkSyncOfficialBridgeFromEth from "@/components/tasks/BRIDGE/ZkSyncFromEth/Menu.vue";
import TaskZkSyncOfficialBridgeFromEth from "@/components/tasks/BRIDGE/ZkSyncFromEth/Block.vue";
import MenuZkSyncOfficialBridgeToEth from "@/components/tasks/BRIDGE/ZkSyncToEth/Menu.vue";
import TaskZkSyncOfficialBridgeToEth from "@/components/tasks/BRIDGE/ZkSyncToEth/Block.vue";

export const ZkSyncBridgeToETHSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  menu: MenuZkSyncOfficialBridgeToEth,
  component: TaskZkSyncOfficialBridgeToEth,
  descFn(task) {
    let p = task.zkSyncOfficialBridgeToEthereumTask
    return ` (from ${Network.ZKSYNCERA} to ${Network.Etherium} ETH)`
  },
  service: {
    name: 'zksync to ETH',
    img: '/icons/era.svg',
    link: 'https://portal.zksync.io/bridge/',
    op: 'bridge',
  },
  job: TaskJob.Bridge,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
