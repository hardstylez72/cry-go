import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskMerklyNFT from "@/components/tasks/NFT/NFTMerkly/Block.vue";
import MenuTaskMerklyNFT from "@/components/tasks/NFT/NFTMerkly/Menu.vue";

import MintBlock from "@/components/tasks/NFT/NFTMerkly/mintBlock.vue";
import MintMenu from "@/components/tasks/NFT/NFTMerkly/mintMenu.vue";

export const NFTMerklySpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.merklyMintAndBridgeNFTTask
    return ` (from ${p?.fromNetwork} to ${p?.toNetwork})`
  },
  component: TaskMerklyNFT,
  menu: MenuTaskMerklyNFT,
  service: {
    name: 'Merkly mint/bridge',
    link: 'https://minter.merkly.com/',
    img: '/icons/merkly.webp',
    op: 'mint/brdge NFT',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([
    Network.ZKSYNCERA,
    Network.ARBITRUM,
    Network.BinanaceBNB,
    Network.POLIGON,
    Network.Base,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.LayerZero, Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}

export const NFTMerklyMintSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.mintMerklyTask
    return ` (${p?.network})`
  },
  component: MintBlock,
  menu: MintMenu,
  service: {
    name: 'Merkly mint',
    link: 'https://minter.merkly.com/',
    img: '/icons/merkly.webp',
    op: 'mint NFT',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([
    Network.ZKSYNCERA,
    Network.ARBITRUM,
    Network.BinanaceBNB,
    Network.POLIGON,
    Network.Base,
  ]),
  airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
