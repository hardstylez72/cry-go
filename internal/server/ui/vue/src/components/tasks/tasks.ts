import {Network, Task, TaskType} from "@/generated/flow";

import {ProfileType} from "@/generated/profile";
import {AcrossBridgeSpec} from "@/components/tasks/BRIDGE/Across";
import {Airdrop, allNetworks, TaskJob, TaskSpec, Universal} from "@/components/tasks/utils";
import {StargateBridgeSpec} from "@/components/tasks/BRIDGE/Stargate";
import {TraderJoeSpec} from "@/components/tasks/SWAPS/TraderJoe";
import {ExchangeDeposit, ExchangeSwapSpec, ExchangeWithdrawSpec} from "@/components/tasks/EXCHANGE";
import {SwapAvnuSpec} from "@/components/tasks/SWAPS/Avnu";
import {SwapFibrousSpec} from "@/components/tasks/SWAPS/Fibrous";
import {LPZkLendLPSpec} from "@/components/tasks/LP/LPZkLend";
import {PancakeSwapSpec} from "@/components/tasks/SWAPS/Pancake";
import {SwapOdosSpec} from "@/components/tasks/SWAPS/Odos";
import {OrbiterBridgeSpec} from "@/components/tasks/BRIDGE/Orbiter";
import {SwapWoofiSpec} from "@/components/tasks/SWAPS/Woofi";
import {LPAaveSpec} from "@/components/tasks/LP/LPAave";
import {NFTMintFunSpec} from "@/components/tasks/NFT/MintFun";
import {NFTMerklyMintSpec, NFTMerklySpec} from "@/components/tasks/NFT/Merkly";
import {NFTZeriusMintSpec} from "@/components/tasks/NFT/Zeruis";
import {IZUMISwapSpec} from "@/components/tasks/SWAPS/IZUMI";
import {SyncSwapSpec} from "@/components/tasks/SWAPS/SyncSwap";
import {VelocoreSwapSpec} from "@/components/tasks/SWAPS/Velocore";
import {MaverickSwapSpec} from "@/components/tasks/SWAPS/Maverick";
import {SpaceFISwapSpec} from "@/components/tasks/SWAPS/SpaceFI";
import {KyberSwapSpec} from "@/components/tasks/SWAPS/KeberSwap";
import {MySwapSpec} from "@/components/tasks/SWAPS/MySwap";
import {Swap10kSpec} from "@/components/tasks/SWAPS/Swap10k";
import {SithSwapSpec} from "@/components/tasks/SWAPS/Sith";
import {ProtossSwapSpec} from "@/components/tasks/SWAPS/Protoss";
import {JediSwapSpec} from "@/components/tasks/SWAPS/Jedi";
import {ZkSwapSpec} from "@/components/tasks/SWAPS/ZkSwap";
import {MuteIOSwapSpec} from "@/components/tasks/SWAPS/muteio";
import {EzkaliburSwapSpec} from "@/components/tasks/SWAPS/ezkalibur";
import {SnapshotSpec} from "@/components/tasks/OTHER/Snapshot";
import {DmailSpec} from "@/components/tasks/OTHER/Dmail";
import {VeSyncSpec} from "@/components/tasks/SWAPS/VeSync";
import {DeploySpec} from "@/components/tasks/OTHER/StarknetDeployAccount";
import {SyncSwapLPSpec} from "@/components/tasks/LP/SyncSwap";
import {StarknetIdMintSpec} from "@/components/tasks/NFT/StarknetId";
import {WETHSwapSpec} from "@/components/tasks/SWAPS/WETH";
import {BridgeStarknetSpec} from "@/components/tasks/BRIDGE/StarkNet";
import {TestNetSpec} from "@/components/tasks/BRIDGE/TestNet";
import {ZkSyncBridgeFromETHSpec} from "@/components/tasks/BRIDGE/ZkSyncFromEth";
import {ZkSyncBridgeToETHSpec} from "@/components/tasks/BRIDGE/ZkSyncToEth";
import {DelaySpec} from "@/components/tasks/OTHER/Delay";
import {SwapEkuboSpec} from "@/components/tasks/SWAPS/Ekubo";
import {LPNostraLPSpec} from "@/components/tasks/LP/Nostra";
import {CoreDaoBridgeSpec} from "@/components/tasks/BRIDGE/CoreDao";
import {StakeSTGSpec} from "@/components/tasks/LP/StakeSTG";
import {merklyRefuelSpec} from "@/components/tasks/BRIDGE/MerklyRefuel";
import {l2PassRefuelSpec} from "@/components/tasks/BRIDGE/L2Pass";
import {StarknetClaim} from "@/components/tasks/OTHER/StarknetClaim";
import {LPEraLendLPSpec} from "@/components/tasks/LP/EraLend";


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
    Delay: DelaySpec,
    StargateBridge: StargateBridgeSpec,
    WithdrawExchange: ExchangeWithdrawSpec,
    OkexDeposit: ExchangeDeposit,
    TestNetBridgeSwap: TestNetSpec,
    SnapshotVote: SnapshotSpec,
    SyncSwap: SyncSwapSpec,
    ZkSyncOfficialBridgeToEthereum: ZkSyncBridgeToETHSpec,
    Mock: deprecated,
    OkexBinance: deprecated,
    Swap1inch: deprecated,
    OrbiterBridge: OrbiterBridgeSpec,
    ZkSyncOfficialBridgeFromEthereum: ZkSyncBridgeFromETHSpec,
    WETH: WETHSwapSpec,
    MuteioSwap: MuteIOSwapSpec,
    SyncSwapLP: SyncSwapLPSpec,
    MaverickSwap: MaverickSwapSpec,
    SpaceFISwap: SpaceFISwapSpec,
    VelocoreSwap: VelocoreSwapSpec,
    IzumiSwap: IZUMISwapSpec,
    VeSyncSwap: VeSyncSpec,
    EzkaliburSwap: EzkaliburSwapSpec,
    ZkSwap: ZkSwapSpec,
    TraderJoeSwap: TraderJoeSpec,
    MerklyMintAndBridgeNFT: NFTMerklySpec,
    DeployStarkNetAccount: DeploySpec,
    Swap10k: Swap10kSpec,
    PancakeSwap: PancakeSwapSpec,
    SithSwap: SithSwapSpec,
    JediSwap: JediSwapSpec,
    MySwap: MySwapSpec,
    ProtossSwap: ProtossSwapSpec,
    StarkNetBridge: BridgeStarknetSpec,
    Dmail: DmailSpec,
    StarkNetIdMint: StarknetIdMintSpec,
    OdosSwap: SwapOdosSpec,
    AcrossBridge: AcrossBridgeSpec,
    AvnuSwap: SwapAvnuSpec,
    FibrousSwap: SwapFibrousSpec,
    ExchangeSwap: ExchangeSwapSpec,
    ZkLendLP: LPZkLendLPSpec,
    WoofiSwap: SwapWoofiSpec,
    AaveLP: LPAaveSpec,
    MintFun: NFTMintFunSpec,
    MintMerkly: NFTMerklyMintSpec,
    MintZerius: NFTZeriusMintSpec,
    KyberSwap: KyberSwapSpec,
    EkuboSwap: SwapEkuboSpec,
    NostraLP: LPNostraLPSpec,
    CoreDaoBridge: CoreDaoBridgeSpec,
    StakeSTG: StakeSTGSpec,
    MerklyRefuel: merklyRefuelSpec,
    L2PassRefuel: l2PassRefuelSpec,
    StarknetClaim: StarknetClaim,
    EraLend: LPEraLendLPSpec,
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


