import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";

import MintBlock from "@/components/tasks/NFT/NFTZeruis/mintBlock.vue";
import MintMenu from "@/components/tasks/NFT/NFTZeruis/mintMenu.vue";


export const NFTZeriusMintSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.mintZeriusTask
    return ` (${p?.network})`
  },
  component: MintBlock,
  menu: MintMenu,
  service: {
    name: 'Zerius mint',
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
