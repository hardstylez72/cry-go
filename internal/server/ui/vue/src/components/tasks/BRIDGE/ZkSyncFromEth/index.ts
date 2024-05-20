import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import MenuZkSyncOfficialBridgeFromEth from "@/components/tasks/BRIDGE/ZkSyncFromEth/Menu.vue";
import TaskZkSyncOfficialBridgeFromEth from "@/components/tasks/BRIDGE/ZkSyncFromEth/Block.vue";

export const ZkSyncBridgeFromETHSpec: TaskSpec = {
    deprecated: true,
    canBeEstimated: true,
    menu: MenuZkSyncOfficialBridgeFromEth,
    component: TaskZkSyncOfficialBridgeFromEth,
    descFn(task) {
        let p = task.zkSyncOfficialBridgeFromEthereumTask
        return ` (from ${Network.Etherium} to ${Network.ZKSYNCERA} ETH)`
    },
    bridge: (task: Task) => {
        return {
            from: Token.ETH, to: Token.ETH, fromNetwork: Network.Etherium, toNetwork: Network.ZKSYNCERA
        }
    },
    service: {
        name: 'zksync from ETH',
        img: '/icons/era.svg',
        link: 'https://portal.zksync.io/bridge/',
        op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
}
