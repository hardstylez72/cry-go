import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Dmail from "@/components/tasks/OTHER/StarknetClaim/Block.vue";
import MenuDmail from "@/components/tasks/OTHER/StarknetClaim/Menu.vue";

export const StarknetClaim: TaskSpec = {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
        const p = task.starknetClaim
        return ''
    },
    simple(task) {
        const p = task.starknetClaim
        return {network: p.network}
    },
    component: Dmail,
    menu: MenuDmail,
    service: {
        name: 'ClaimStarknet',
        link: 'https://starkgate.starknet.io/',
        img: '/icons/starknet.svg',
        op: 'claims',
    },
    job: TaskJob.Other,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
}
