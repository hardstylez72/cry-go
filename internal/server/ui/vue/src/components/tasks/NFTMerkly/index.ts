import {Network} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import TaskMerklyNFT from "@/components/tasks/NFTMerkly/Block.vue";
import MenuTaskMerklyNFT from "@/components/tasks/NFTMerkly/Menu.vue";

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
    name: 'Merkly',
    link: 'https://minter.merkly.com/',
    img: '/icons/merkly.webp',
    op: 'mint/brdge NFT',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([Network.ZKSYNCERA]),
  airdrops: new Set<Airdrop>([Airdrop.LayerZero, Airdrop.ZkSync]),
  profileType: new Set([ProfileType.EVM])
}
