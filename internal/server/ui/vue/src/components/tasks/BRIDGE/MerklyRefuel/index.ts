import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/BRIDGE/MerklyRefuel/Block.vue";
import Menu from "@/components/tasks/BRIDGE/MerklyRefuel/Menu.vue";

export const merklyRefuelSpec: TaskSpec = {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
        const p = task.merklyRefuel
        return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
    },
    bridge: (task: Task) => {
        const p = task.merklyRefuel
        return {
            from: p?.fromToken, to: p?.toToken, fromNetwork: p?.fromNetwork, toNetwork: p?.toNetwork
        }
    },
    component: Block,
    menu: Menu,
    service: {
        name: 'Merkly refuel',
        link: 'https://minter.merkly.com/gas',
        img: '/icons/merkly.webp',
        op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([
        Network.ARBITRUM,
        Network.POLIGON,
        Network.BinanaceBNB,
    ]),
    airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
    profileType: new Set([ProfileType.EVM])
}
