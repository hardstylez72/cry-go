import {Network, Task} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";


export const allNetworks: Network[] = [
  Network.ZKSYNCERA,
  Network.ARBITRUM,
  Network.StarkNet,
  Network.AVALANCHE,
  Network.BinanaceBNB,
  Network.POLIGON,
  Network.OPTIMISM,
  Network.Etherium
]

interface Service {
  link: string
  img: string

  name: string

  op: string
}

interface DescFn {
  (task: Task): string;
}

export interface TaskSpec {
  menu: any
  component: any

  canBeEstimated: boolean
  deprecated: boolean

  descFn: DescFn

  service: Service

  networks: Set<Network>

  airdrops: Set<Airdrop>
  job: TaskJob

  profileType: Set<ProfileType>
}

export enum TaskJob {
  Swap = 'Swap',
  Bridge = 'Bridge',
  Exchange = 'Exchange',
  Other = 'Other',

  NFT = "NFT",

  LP = 'LP'
}

export enum Airdrop {
  LayerZero = "LayerZero",
  ZkSync = "ZkSync",
  StarkNet = "StarkNet",
  Orbiter = "Orbiter",
}

export const Universal = [Airdrop.ZkSync, Airdrop.StarkNet, Airdrop.Orbiter, Airdrop.LayerZero]

