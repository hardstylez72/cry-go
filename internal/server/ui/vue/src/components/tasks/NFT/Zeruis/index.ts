import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";

import MintBlock from "@/components/tasks/NFT/Zeruis/mintBlock.vue";
import MintMenu from "@/components/tasks/NFT/Zeruis/mintMenu.vue";


export const NFTZeriusMintSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.mintZeriusTask
    return ` (${p?.network})`
  },
  nft(task: Task) {
    const p = task.mintZeriusTask
    return {
      network: p.network,
    }
  },
  component: MintBlock,
  menu: MintMenu,
  service: {
    name: 'Zerius',
    link: 'https://zerius.io/',
    img: '/icons/zerius.svg',
    op: 'mint NFT',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([
    // Network.ZKSYNCERA,
    // Network.ARBITRUM,
    // Network.BinanaceBNB,
    // Network.POLIGON,
    Network.Base,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
