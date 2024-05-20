import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {Airdrop, TaskJob, TaskSpec} from "@/components/tasks/utils";
import Block from "@/components/tasks/NFT/MintFun/Block.vue";
import Menu from "@/components/tasks/NFT/MintFun/Menu.vue";

export const NFTMintFunSpec: TaskSpec = {
  deprecated: false,
  canBeEstimated: true,
  descFn(task) {
    const p = task.mintFunTask
    return ` (${p?.network})`
  },
  nft(task: Task) {
    const p = task.mintFunTask
    return {
      network: p.network,
    }
  },
  component: Block,
  menu: Menu,
  service: {
    name: 'MintFun',
    link: 'https://mint.fun/',
    img: '/icons/mintFun.png',
    op: 'mint NFT',
  },
  job: TaskJob.NFT,
  networks: new Set<Network>([Network.Base]),
  airdrops: new Set<Airdrop>([Airdrop.Base]),
  profileType: new Set([ProfileType.EVM])
}
