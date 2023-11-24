import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskMerklyNFT from "@/components/tasks/NFT/Merkly/Block.vue";
import MenuTaskMerklyNFT from "@/components/tasks/NFT/Merkly/Menu.vue";

import MintBlock from "@/components/tasks/NFT/Merkly/mintBlock.vue";
import MintMenu from "@/components/tasks/NFT/Merkly/mintMenu.vue";

export const NFTMerklySpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.merklyMintAndBridgeNFTTask
    return ` (from ${p?.fromNetwork} to ${p?.toNetwork})`
  },
  nft(task: Task) {
    const p = task.merklyMintAndBridgeNFTTask
    return {
      network: p.fromNetwork,
      toNetwork: p.toNetwork
    }
  },
  component: TaskMerklyNFT,
  menu: MenuTaskMerklyNFT,
  service: {
    name: 'Merkly',
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
  nft(task: Task) {
    const p = task.mintMerklyTask
    return {
      network: p.network,
    }
  },
  component: MintBlock,
  menu: MintMenu,
  service: {
    name: 'Merkly',
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
