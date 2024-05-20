import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/BRIDGE/L2Pass/Block.vue";
import Menu from "@/components/tasks/BRIDGE/L2Pass/Menu.vue";

export const l2PassRefuelSpec: TaskSpec = {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
        const p = task.l2passRefuel
        return ` (${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
    },
    bridge: (task: Task) => {
        const p = task.l2passRefuel
        return {
            from: p?.fromToken, to: p?.toToken, fromNetwork: p?.fromNetwork, toNetwork: p?.toNetwork
        }
    },
    component: Block,
    menu: Menu,
    service: {
        name: 'L2Pass refuel',
        link: 'https://l2pass.com/refuel',
        img: '/icons/l2pass.svg',
        op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([
        Network.POLIGON,
    ]),
    airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
    profileType: new Set([ProfileType.EVM])
}
