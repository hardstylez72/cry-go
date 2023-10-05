import {Network, Task, TaskType} from "@/generated/flow";
import {flow_Flow} from "@/generated/process";


import TaskStargateBridge from "@/components/tasks/BridgeStargate/Block.vue";
import TaskDelay from "@/components/tasks/block/Delay.vue";
import TaskExchangeWithdraw from "@/components/tasks/Exchange/Withdraw.vue";
import TaskOkexDeposit from "@/components/tasks/Exchange/Deposit.vue";
import TaskTestNetBridgeSwap from "@/components/tasks/block/TaskTestNetBridgeSwap.vue";
import TaskSnapshotVote from "@/components/tasks/block/TaskSnapshotVote.vue";
import TaskSyncSwap from "@/components/tasks/block/SyncSwap.vue";
import MenuTaskSyncSwap from "@/components/tasks/menu/MenuSyncSwap.vue";
import MenuDelayTask from "@/components/tasks/menu/DelayMenu.vue";
import MenuExchangeWithdraw from "@/components/tasks/Exchange/WithdrawMenu.vue";
import MenuOkexDeposit from "@/components/tasks/Exchange/DepositMenu.vue";
import MenuSnapshotTask from "@/components/tasks/menu/Snapshot.vue";
import MenuTaskStargateBridge from "@/components/tasks/BridgeStargate/Menu.vue";
import MenuTaskTestNetBridge from "@/components/tasks/menu/MenuTaskTestNetBridge.vue";
import TaskZkSyncOfficialBridgeToEth from "@/components/tasks/block/TaskZkSyncOfficialBridgeToEth.vue";
import TaskOrbiterBridge from "@/components/tasks/block/TaskOrbiterBridge.vue";
import MenuTaskOrbiterBridge from "@/components/tasks/menu/MenuTaskOrbiterBridge.vue";
import MenuZkSyncOfficialBridgeToEth from "@/components/tasks/menu/MenuZkSyncOfficialBridgeToEth.vue";
import TaskZkSyncOfficialBridgeFromEth from "@/components/tasks/block/TaskZkSyncOfficialBridgeFromEth.vue";
import MenuZkSyncOfficialBridgeFromEth from "@/components/tasks/menu/MenuZkSyncOfficialBridgeFromEth.vue";
import TaskWethSwap from "@/components/tasks/block/TaskWethSwap.vue";
import MenuTaskWethSwap from "@/components/tasks/menu/MenuTaskWethSwap.vue";
import TaskMuteIoSwap from "@/components/tasks/block/MuteIoSwap.vue";
import MenuTaskMuteioSwap from "@/components/tasks/menu/MenuMuteioSwap.vue";
import {humanDuration} from "@/components/helper";
import TaskSyncSwapLP from "@/components/tasks/block/TaskSyncSwapLP.vue";
import MenuTaskSyncSwapLP from "@/components/tasks/menu/MenuTaskSyncSwapLP.vue";
import TaskMaverickSwap from "@/components/tasks/block/MaverickSwap.vue";
import MenuTaskMaverickSwap from "@/components/tasks/menu/MenuMaverickSwap.vue";
import TaskSpaceFiSwap from "@/components/tasks/block/SpaceFiSwap.vue";
import MenuTaskSpaceFiSwap from "@/components/tasks/menu/MenuSpaceFiSwap.vue";
import TaskVelocoreSwap from "@/components/tasks/block/VelocoreSwap.vue";
import MenuTaskVelocoreSwap from "@/components/tasks/menu/MenuVelocoreSwap.vue";
import TaskIzumiSwap from "@/components/tasks/block/IzumiSwap.vue";
import MenuTaskIzumiSwap from "@/components/tasks/menu/MenuIzumiSwap.vue";
import TaskVeSyncSwap from "@/components/tasks/block/VeSyncSwap.vue";
import MenuTaskVeSyncSwap from "@/components/tasks/menu/MenuVeSyncSwap.vue";
import TaskEzkaliburSwap from "@/components/tasks/block/EzkaliburSwap.vue";
import MenuTaskEzkaliburSwap from "@/components/tasks/menu/MenuEzkaliburSwap.vue";
import TaskZkSwap from "@/components/tasks/block/ZkSwap.vue";
import MenuTaskZkSwap from "@/components/tasks/menu/MenuZkSwap.vue";
import TaskTraderJoeSwap from "@/components/tasks/TraderJoe/Block.vue";
import MenuTaskTraderJoeSwap from "@/components/tasks/TraderJoe/Menu.vue";
import MenuTaskMerklyNFT from "@/components/tasks/NFTMerkly/Menu.vue";
import TaskMerklyNFT from "@/components/tasks/NFTMerkly/Block.vue";
import TaskDeployStarkNetAccount from "@/components/tasks/block/DeployStarkNetAccount.vue";
import DeployStarkNetAccount from "@/components/tasks/menu/DeployStarkNetAccount.vue";
import {ProfileType} from "@/generated/profile";
import swap10k from "@/components/tasks/block/swap10k.vue";
import swap10kMenu from "@/components/tasks/menu/MenuSwap10k.vue";
import TaskPancakeSwap from "@/components/tasks/block/PancakeSwap.vue";
import PancakeSwap from "@/components/tasks/menu/MenuPancakeSwap.vue";
import SithSwap from "@/components/tasks/block/SithSwap.vue";
import MenuSithSwap from "@/components/tasks/menu/MenuSithSwap.vue";
import MenuJediSwap from "@/components/tasks/menu/MenuJediSwap.vue";
import JediSwap from "@/components/tasks/block/JediSwap.vue";
import MySwap from "@/components/tasks/block/MySwap.vue";
import MenuMySwap from "@/components/tasks/menu/MenuMySwap.vue";
import MenuProtossSwap from "@/components/tasks/menu/MenuProtossSwap.vue";
import ProtossSwap from "@/components/tasks/block/ProtossSwap.vue";
import StrarkNetBridge from "@/components/tasks/block/StrarkNetBridge.vue";
import MenuStarkNetBridge from "@/components/tasks/menu/MenuStarkNetBridge.vue";
import Dmail from "@/components/tasks/block/Dmail.vue";
import MenuDmail from "@/components/tasks/menu/MenuDmail.vue";
import StarkNetId from "@/components/tasks/block/StarkNetId.vue";
import MenuStarkNetId from "@/components/tasks/menu/MenuStarkNetId.vue";
import MenuOdosSwap from "@/components/tasks/menu/MenuOdosSwap.vue";
import OdosSwap from "@/components/tasks/block/OdosSwap.vue";
import {AcrossBridgeSpec} from "@/components/tasks/BridgeAcross";
import {Airdrop, allNetworks, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import {StargateBridgeSpec} from "@/components/tasks/BridgeStargate";
import {TraderJoeSpec} from "@/components/tasks/TraderJoe";
import {ExchangeDeposit, ExchangeSwapSpec, ExchangeWithdrawSpec} from "@/components/tasks/Exchange";
import {SwapAvnuSpec} from "@/components/tasks/SwapAvnu";
import {SwapFibrousSpec} from "@/components/tasks/SwapFibrous";
import {LPZkLendLPSpec} from "@/components/tasks/LPZkLend";


export interface TaskArg {
  task?: Task
  component: any
  weight: number

  taskType: TaskType
}


const deprecated: TaskSpec = {
  deprecated: true,
  canBeEstimated: false,
  menu: null,
  component: null,
  descFn(task) {
    return ''
  },

  service: {
    name: '',
    img: '',
    link: '',
    op: '',

  },
  job: TaskJob.Other,
  networks: new Set<Network>(allNetworks),
  airdrops: new Set<Airdrop>(),
  profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
}

export const taskProps: Record<TaskType, TaskSpec> = {
  Delay: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuDelayTask,
    component: TaskDelay,
    descFn(task: Task): string {
      let p = task.delayTask
      return ` (${p?.random ? `${humanDuration(p?.minRandom)}:${humanDuration(p?.maxRandom)}` : humanDuration(p?.duration)})`
    },
    service: {
      name: 'Delay',
      link: '',
      img: '',
      op: '',
    },
    job: TaskJob.Other,
    networks: new Set<Network>(allNetworks),
    airdrops: new Set<Airdrop>(Universal),
    profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
  },
  StargateBridge: StargateBridgeSpec,
  WithdrawExchange: ExchangeWithdrawSpec,
  OkexDeposit: ExchangeDeposit,
  TestNetBridgeSwap: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuTaskTestNetBridge,
    component: TaskTestNetBridgeSwap,
    descFn(task) {
      return `[${task.testNetBridgeSwapTask?.minAmount}:${task.testNetBridgeSwapTask?.maxAmount}]`
    },
    service: {
      name: 'TestNet-bridge',
      img: '/icons/testnetbridge.svg',
      link: 'https://testnetbridge.com/',
      op: 'bridge',

    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.ARBITRUM]),
    airdrops: new Set<Airdrop>([Airdrop.LayerZero]),
    profileType: new Set([ProfileType.EVM])
  },
  SnapshotVote: {
    deprecated: false,
    canBeEstimated: false,
    menu: MenuSnapshotTask,
    component: TaskSnapshotVote,
    descFn(task) {
      let p = task.snapshotVoteTask
      return ` (space: ${p?.space})`
    },
    service: {
      name: 'Snapshot.org',
      img: '/icons/snapshot.png',
      link: 'https://snapshot.org/#/',
      op: 'vote',
    },
    job: TaskJob.Other,
    networks: new Set<Network>([Network.ARBITRUM]),
    airdrops: new Set<Airdrop>(Universal),
    profileType: new Set([ProfileType.EVM])
  },
  SyncSwap: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuTaskSyncSwap,
    component: TaskSyncSwap,
    descFn(task) {
      let p = task.syncSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    service: {
      name: 'SyncSwap',
      img: '/icons/syncswap.svg',
      link: 'https://syncswap.xyz/',
      op: 'swap',
    },
    networks: new Set<Network>([Network.ZKSYNCERA]),
    job: TaskJob.Swap,
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  ZkSyncOfficialBridgeToEthereum: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuZkSyncOfficialBridgeToEth,
    component: TaskZkSyncOfficialBridgeToEth,
    descFn(task) {
      let p = task.zkSyncOfficialBridgeToEthereumTask
      return ` (from ${Network.ZKSYNCERA} to ${Network.Etherium} ETH)`
    },
    service: {
      name: 'zksync to ETH',
      img: '/icons/era.svg',
      link: 'https://portal.zksync.io/bridge/',
      op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  Mock: deprecated,
  OkexBinance: deprecated,
  Swap1inch: deprecated,
  OrbiterBridge: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuTaskOrbiterBridge,
    component: TaskOrbiterBridge,
    descFn(task) {
      let p = task.orbiterBridgeTask
      return ` (from ${p?.fromNetwork} ${p?.fromToken} to ${p?.toNetwork} ${p?.toToken})`
    },
    service: {
      name: 'Orbiter',
      img: '/icons/orbiter.ico',
      link: 'https://www.orbiter.finance/',
      op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>(Universal),
    profileType: new Set([ProfileType.EVM])
  },
  ZkSyncOfficialBridgeFromEthereum: {
    deprecated: false,
    canBeEstimated: true,
    menu: MenuZkSyncOfficialBridgeFromEth,
    component: TaskZkSyncOfficialBridgeFromEth,
    descFn(task) {
      let p = task.zkSyncOfficialBridgeFromEthereumTask
      return ` (from ${Network.Etherium} to ${Network.ZKSYNCERA} ETH)`
    },
    service: {
      name: 'zksync from ETH',
      img: '/icons/era.svg',
      link: 'https://portal.zksync.io/bridge/',
      op: 'bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  WETH: {
    deprecated: false,
    menu: MenuTaskWethSwap,
    component: TaskWethSwap,
    canBeEstimated: true,
    descFn(task) {
      return task.wETHTask?.wrap ? ' (wrap)' : ' (unwrap)'
    },
    service: {
      link: '',
      img: '',
      name: 'WETH <> ETH',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  MuteioSwap: {
    deprecated: false,
    component: TaskMuteIoSwap,
    menu: MenuTaskMuteioSwap,
    canBeEstimated: true,
    descFn(task) {
      const p = task.muteioSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    service: {
      name: 'Mute.io',
      img: '/icons/muteio.png',
      link: 'https://app.mute.io/swap',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  SyncSwapLP: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      let p = task.syncSwapLPTask
      return ` (${p?.network} ${p?.add ? 'deposit' : "withdraw"} pool: ${p?.a}/${p?.b})`
    },
    menu: MenuTaskSyncSwapLP,
    component: TaskSyncSwapLP,
    service: {
      name: 'SyncSwap',
      img: '/icons/syncswap.svg',
      link: 'https://syncswap.xyz/',
      op: 'LP',
    },
    job: TaskJob.LP,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  MaverickSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.maverickSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskMaverickSwap,
    menu: MenuTaskMaverickSwap,
    service: {
      name: 'Maverick',
      img: '/icons/maverick.ico',
      link: 'https://app.mav.xyz/?chain=324',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.LayerZero]),
    profileType: new Set([ProfileType.EVM])
  },
  SpaceFISwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.spaceFiSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskSpaceFiSwap,
    menu: MenuTaskSpaceFiSwap,
    service: {
      name: 'SpaceFi',
      link: 'https://swap-zksync.spacefi.io/#/swap',
      img: '/icons/spacefi.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  VelocoreSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.velocoreSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskVelocoreSwap,
    menu: MenuTaskVelocoreSwap,
    service: {
      name: 'Velocore',
      link: 'https://app.velocore.xyz/swap',
      img: '/icons/velocore.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  IzumiSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.izumiSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskIzumiSwap,
    menu: MenuTaskIzumiSwap,
    service: {
      name: 'IZUMI',
      link: 'https://izumi.finance/trade/swap',
      img: '/icons/izumi.svg',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  VeSyncSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.veSyncSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskVeSyncSwap,
    menu: MenuTaskVeSyncSwap,
    service: {
      name: 'VeSync',
      link: 'https://app.vesync.finance/swap',
      img: '/icons/vesync.svg',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  EzkaliburSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.ezkaliburSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskEzkaliburSwap,
    menu: MenuTaskEzkaliburSwap,
    service: {
      name: 'Ezkalibur',
      link: 'https://dapp.ezkalibur.com/',
      img: '/icons/ezkalibur.ico',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  ZkSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.zkSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskZkSwap,
    menu: MenuTaskZkSwap,
    service: {
      name: 'ZkSwap',
      link: 'https://zkswap.finance/swap',
      img: '/icons/zkswap.ico',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  TraderJoeSwap: TraderJoeSpec,
  MerklyMintAndBridgeNFT: {
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
  },
  DeployStarkNetAccount: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.deployStarkNetAccountTask
      return ''
    },
    component: TaskDeployStarkNetAccount,
    menu: DeployStarkNetAccount,
    service: {
      name: 'StarkNet acc deploy',
      link: 'https://starkgate.starknet.io/',
      img: '/icons/starknet.svg',
      op: 'deploy account',
    },
    job: TaskJob.Other,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  Swap10k: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.swap10k
      return ''
    },
    component: swap10k,
    menu: swap10kMenu,
    service: {
      name: '10kswap',
      link: 'https://10kswap.com/swap',
      img: '/icons/10kswap.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  PancakeSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.pancakeSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: TaskPancakeSwap,
    menu: PancakeSwap,
    service: {
      name: 'PancakeSwap',
      link: 'https://pancakeswap.finance/',
      img: '/icons/pancake.ico',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  SithSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.sithSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: SithSwap,
    menu: MenuSithSwap,
    service: {
      name: 'SythSwap',
      link: 'https://sithswap.com/',
      img: '/icons/sithswap.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  JediSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.jediSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: JediSwap,
    menu: MenuJediSwap,
    service: {
      name: 'JediSwap',
      link: 'https://app.jediswap.xyz/#/swap',
      img: '/icons/jediswap.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  MySwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.mySwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: MySwap,
    menu: MenuMySwap,
    service: {
      name: 'MySwap',
      link: 'https://www.myswap.xyz/#/swap',
      img: '/icons/myswap.svg',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  ProtossSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.protosSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: ProtossSwap,
    menu: MenuProtossSwap,
    service: {
      name: 'ProtossSwap',
      link: 'https://www.protoss.org/',
      img: '/icons/protossSwap.png',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  StarkNetBridge: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.starkNetBridgeTask
      return ''
    },
    component: StrarkNetBridge,
    menu: MenuStarkNetBridge,
    service: {
      name: 'StarkNet liquidity bridge',
      link: 'https://starkgate.starknet.io/',
      img: '/icons/starknet.svg',
      op: 'liquidity bridge',
    },
    job: TaskJob.Bridge,
    networks: new Set<Network>([Network.Etherium, Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.EVM])
  },
  Dmail: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.dmailTask
      return ''
    },
    component: Dmail,
    menu: MenuDmail,
    service: {
      name: 'Dmail',
      link: 'https://dmail.ai/',
      img: '/icons/dmail.ico',
      op: 'email send',
    },
    job: TaskJob.Other,
    networks: new Set<Network>([Network.ZKSYNCERA, Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync, Airdrop.StarkNet]),
    profileType: new Set([ProfileType.EVM, ProfileType.StarkNet])
  },
  StarkNetIdMint: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.starkNetIdMintTask
      return ''
    },
    component: StarkNetId,
    menu: MenuStarkNetId,
    service: {
      name: 'StarkNetId',
      link: 'https://app.starknet.id/',
      img: '/icons/starknetid.svg',
      op: 'mint uniq id',
    },
    job: TaskJob.NFT,
    networks: new Set<Network>([Network.StarkNet]),
    airdrops: new Set<Airdrop>([Airdrop.StarkNet]),
    profileType: new Set([ProfileType.StarkNet])
  },
  OdosSwap: {
    deprecated: false,
    canBeEstimated: true,
    descFn(task) {
      const p = task.odosSwapTask
      return ` (${p?.network} from ${p?.fromToken} to ${p?.toToken})`
    },
    component: OdosSwap,
    menu: MenuOdosSwap,
    service: {
      name: 'Odos',
      link: 'https://app.odos.xyz',
      img: '/icons/odos.ico',
      op: 'swap',
    },
    job: TaskJob.Swap,
    networks: new Set<Network>([Network.ZKSYNCERA]),
    airdrops: new Set<Airdrop>([Airdrop.ZkSync]),
    profileType: new Set([ProfileType.EVM])
  },
  AcrossBridge: AcrossBridgeSpec,
  AvnuSwap: SwapAvnuSpec,
  FibrousSwap: SwapFibrousSpec,
  ExchangeSwap: ExchangeSwapSpec,
  ZkLendLP: LPZkLendLPSpec,
}

export const getFlow = (flow: { tasks: Task[] }): string[] => {
  const result: string[] = []
  if (!flow || !flow.tasks) {
    return result
  }

  flow.tasks.forEach((task) => {

    if (!task || !task.taskType) {
      return
    }

    const out = task.taskType + " " + taskProps[task.taskType].descFn(task)
    result.push(out)
  })
  return result
}

export const taskComponentMap = new Map<TaskType, any>()
export const menuTaskComponentMap = new Map<TaskType, any>()
export const estimatedTaskMap = new Map<TaskType, boolean>()
export const taskTypes: TaskType[] = []

for (let name of Object.getOwnPropertyNames(taskProps)) {
  const n = name as TaskType

  const v = taskProps[n]
  if (v) {
    if (v.deprecated) {
      continue
    }

    taskComponentMap.set(n, v.component)
    menuTaskComponentMap.set(n, v.menu)
    taskTypes.push(n)
    estimatedTaskMap.set(n, v.canBeEstimated)
  }
}


