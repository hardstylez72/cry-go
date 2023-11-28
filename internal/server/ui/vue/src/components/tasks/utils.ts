import {Network, Task, Token} from "@/generated/flow";
import {ProfileType} from "@/generated/profile";
import {BridgePair, SwapPair} from "@/components/helper";


export const allNetworks: Network[] = [
  Network.ZKSYNCERA,
  Network.ARBITRUM,
  Network.StarkNet,
  Network.AVALANCHE,
  Network.BinanaceBNB,
  Network.POLIGON,
  Network.OPTIMISM,
  Network.Etherium,
  Network.Base,
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

  swap?: (task: Task) => { from: Token, to: Token, network: Network }
  nft?: (task: Task) => { network: Network, toNetwork?: Network }
  simple?: (task: Task) => { network: Network }
  service: Service

  networks: Set<Network>

  airdrops: Set<Airdrop>
  job: TaskJob

  profileType: Set<ProfileType>

  swapParis?: SwapPair[]
  bridgePairs?: BridgePair[]
}

export enum TaskJob {
  Swap = 'Swap',
  Bridge = 'Bridge',
  Exchange = 'Exchange',
  Other = 'Other',

  NFT = "NFT",

  LP = 'LP',

  Simple = "Simple"
}

export enum Airdrop {
  LayerZero = "LayerZero",
  ZkSync = "ZkSync",
  StarkNet = "StarkNet",
  Orbiter = "Orbiter",
  Base = "Base",
  Linea = "Linea",
}

export const Universal = [Airdrop.ZkSync, Airdrop.StarkNet, Airdrop.Orbiter, Airdrop.LayerZero, Airdrop.Base]

