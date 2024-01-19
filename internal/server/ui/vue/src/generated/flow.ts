/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */

/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */
export interface IRequestOptions {
  headers?: any;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

export interface IRequestPromise<T = any> extends Promise<IRequestResponse<T>> {}

export interface IRequestResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: any;
  config: any;
  request?: any;
}

export interface IRequestInstance {
  (config: any): IRequestPromise;
  (url: string, config?: any): IRequestPromise;
  request<T = any>(config: any): IRequestPromise<T>;
}

export interface IRequestConfig {
  method?: any;
  headers?: any;
  url?: any;
  data?: any;
  params?: any;
}

// Add options interface
export interface ServiceOptions {
  axios?: IRequestInstance;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

// Add default options
export const serviceOptions: ServiceOptions = {};

// Instance selector
export function axios(configs: IRequestConfig, resolve: (p: any) => void, reject: (p: any) => void): Promise<any> {
  if (serviceOptions.axios) {
    return serviceOptions.axios
      .request(configs)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  } else {
    throw new Error('please inject yourself instance like axios  ');
  }
}

export function getConfigs(method: string, contentType: string, url: string, options: any): IRequestConfig {
  const configs: IRequestConfig = {
    loading: serviceOptions.loading,
    showError: serviceOptions.loading,
    ...options,
    method,
    url
  };
  configs.headers = {
    ...options.headers,
    'Content-Type': contentType
  };
  return configs;
}

export const basePath = '';

export interface IList<T> extends Array<T> {}
export interface List<T> extends Array<T> {}
export interface IDictionary<TValue> {
  [key: string]: TValue;
}
export interface Dictionary<TValue> extends IDictionary<TValue> {}

export interface IListResult<T> {
  items?: T[];
}

export class ListResultDto<T> implements IListResult<T> {
  items?: T[];
}

export interface IPagedResult<T> extends IListResult<T> {
  totalCount?: number;
  items?: T[];
}

export class PagedResultDto<T = any> implements IPagedResult<T> {
  totalCount?: number;
  items?: T[];
}

// customer definition
// empty

export class FlowService {
  /**
   *
   */
  flowServiceOnlyRandomFlowPreview(
    params: {
      /**  */
      body: OnlyRandomFlowPreviewReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<OnlyRandomFlowPreviewRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow-random/only-preview';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceOnlyRandomFlowFromTokens(
    params: {
      /**  */
      body: OnlyRandomFlowPreviewReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<OnlyRandomFlowFromTokensRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow-random/only-preview-tokens';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceSharedFlow(
    params: {
      /**  */
      body: SharedFlowReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SharedFlowRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow-shared';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceUseSharedFlow(
    params: {
      /**  */
      body: UseSharedFlowReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UseSharedFlowRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow-shared/use';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceCopyFlow(
    params: {
      /**  */
      body: CopyFlowReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CopyFlowRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/copy';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceCreateFlow(
    params: {
      /**  */
      body: CreateFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceDeleteFlow(
    params: {
      /**  */
      body: DeleteFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceGetFlow(
    params: {
      /**  */
      body: GetFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceHideFlow(
    params: {
      /**  */
      body: HideFlowReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<HideFlowRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/hide';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceListFlow(
    params: {
      /**  */
      body: ListFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceFlowPreview(
    params: {
      /**  */
      body: FlowPreviewReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<FlowPreviewRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/preview';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceShareFlow(
    params: {
      /**  */
      body: ShareFlowReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ShareFlowRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/publish';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceSharedFlows(
    params: {
      /**  */
      body: SharedFlowsReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SharedFlowsRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/shared/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceUpdateFlow(
    params: {
      /**  */
      body: UpdateFlowRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateFlowResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/flow/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceCreateFlowV2(
    params: {
      /**  */
      body: CreateFlowV2Req;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateFlowV2Res> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v2/flow/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceGetFlowV2(
    params: {
      /**  */
      body: GetFlowV2Req;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetFlowV2Res> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v2/flow/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  flowServiceUpdateFlowV2(
    params: {
      /**  */
      body: UpdateFlowV2Request;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateFlowV2Response> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v2/flow/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface AmUni {
  /**  */
  gwei?: string;

  /**  */
  eth?: string;

  /**  */
  usd?: string;

  /**  */
  network?: Network;

  /**  */
  wei?: string;
}

export interface Amount {
  /**  */
  sendAll?: boolean;

  /**  */
  sendPercent?: number;

  /**  */
  sendAmount?: number;

  /**  */
  sendValue?: string;

  /**  */
  percRange?: PercentRange;

  /**  */
  send?: AmUni;

  /**  */
  balance?: AmUni;

  /**  */
  gasEstimated?: AmUni;
}

export interface CopyFlowReq {
  /**  */
  id: string;
}

export interface CopyFlowRes {
  /**  */
  id: string;
}

export interface CreateFlowRequest {
  /**  */
  label: string;

  /**  */
  tasks: Task[];

  /**  */
  randomTasks: Task[];
}

export interface CreateFlowResponse {
  /**  */
  flow: flow_Flow;
}

export interface CreateFlowV2Req {
  /**  */
  label: string;

  /**  */
  blocks: FlowBlock[];
}

export interface CreateFlowV2Res {
  /**  */
  id: string;
}

export interface DefaultBridge {
  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  amount: Amount;

  /**  */
  approveTx?: TaskTx;

  /**  */
  tx?: TaskTx;

  /**  */
  received: boolean;
}

export interface DefaultLP {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  a: Token;

  /**  */
  b: Token;

  /**  */
  tx?: TaskTx;

  /**  */
  add: boolean;

  /**  */
  approveA?: TaskTx;

  /**  */
  approveB?: TaskTx;

  /**  */
  tokens: LPToken[];
}

export interface DefaultSwap {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  tx?: TaskTx;

  /**  */
  approveTx?: TaskTx;
}

export interface DelayTask {
  /**  */
  duration: string;

  /**  */
  waitFor?: Date;

  /**  */
  random: boolean;

  /**  */
  minRandom?: string;

  /**  */
  maxRandom?: string;

  /**  */
  randomDuration?: string;
}

export interface DeleteFlowRequest {
  /**  */
  id: string;
}

export interface DeleteFlowResponse {}

export interface DeployStarkNetAccountTask {
  /**  */
  network: Network;

  /**  */
  tx: TaskTx;
}

export interface ExchangeSwapTask {
  /**  */
  amount: Amount;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  withdrawerId: string;

  /**  */
  tradeId?: string;

  /**  */
  pair?: string;

  /**  */
  before?: boolean;

  /**  */
  after?: boolean;

  /**  */
  swapCompleted?: boolean;
}

export interface FlowBase {
  /**  */
  token?: Token;
}

export interface FlowBlock {
  /**  */
  weight: string;

  /**  */
  man?: FlowBlockMan;

  /**  */
  rand?: FlowBlockRand;
}

export interface FlowBlockMan {
  /**  */
  tasks: Task[];

  /**  */
  randomTasks: Task[];
}

export interface FlowBlockRand {
  /**  */
  startToken?: Token;

  /**  */
  finishToken?: Token;

  /**  */
  startNetwork?: Network;

  /**  */
  tasks: RandomTask[];

  /**  */
  taskCount?: string;

  /**  */
  minDelay: string;

  /**  */
  maxDelay: string;

  /**  */
  nativeTokenMinPercent?: string;

  /**  */
  nativeTokenMaxPercent?: string;

  /**  */
  nonnativeTokenMinPercent?: string;

  /**  */
  nonnativeTokenMaxPercent?: string;
}

export interface FlowListItem {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  nextId?: string;

  /**  */
  createdAt: Date;

  /**  */
  deletedAt?: Date;

  /**  */
  version: string;
}

export interface FlowPreviewReq {
  /**  */
  label: string;

  /**  */
  blocks: FlowBlock[];
}

export interface FlowPreviewRes {
  /**  */
  flow: UniqueFlow[];

  /**  */
  uniquePercent: number;
}

export interface FlowShared {
  /**  */
  id: string;

  /**  */
  description: string;

  /**  */
  label: string;

  /**  */
  parentId: string;

  /**  */
  createdAt: Date;

  /**  */
  deletedAt?: Date;

  /**  */
  tasks: Task[];

  /**  */
  creatorId: string;
}

export interface GetFlowRequest {
  /**  */
  id: string;
}

export interface GetFlowResponse {
  /**  */
  flow: flow_Flow;
}

export interface GetFlowV2Req {
  /**  */
  id: string;
}

export interface GetFlowV2Res {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  blocks: FlowBlock[];
}

export interface HideFlowReq {
  /**  */
  id: string;
}

export interface HideFlowRes {}

export interface LPToken {
  /**  */
  token: Token;

  /**  */
  approveTx?: TaskTx;
}

export interface LiquidityBridgeTask {
  /**  */
  amount: Amount;

  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  token: Token;

  /**  */
  tx?: TaskTx;

  /**  */
  approveTx?: TaskTx;
}

export interface ListFlowRequest {}

export interface ListFlowResponse {
  /**  */
  flows: FlowListItem[];
}

export interface MerklyMintAndBridgeNFTTask {
  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  mintTx: TaskTx;

  /**  */
  bridgeTx: TaskTx;

  /**  */
  nftId?: string;

  /**  */
  fee?: string;
}

export interface MockTask {}

export interface OkexBinanaceTask {
  /**  */
  okexWithdrawerId: string;

  /**  */
  okexToken: string;

  /**  */
  okexNetwork: string;

  /**  */
  binanaceDepositAddr: string;

  /**  */
  txId?: string;

  /**  */
  withdrawTxId?: string;
}

export interface OkexDepositTask {
  /**  */
  network: Network;

  /**  */
  okexAccName?: string;

  /**  */
  okexAddr?: string;

  /**  */
  token: Token;

  /**  */
  address?: string;

  /**  */
  txId?: string;

  /**  */
  txComplete?: boolean;

  /**  */
  subMainTransfer?: boolean;

  /**  */
  amount?: Amount;

  /**  */
  tx?: TaskTx;

  /**  */
  approveTx?: TaskTx;
}

export interface OnlyRandomFlowFromTokensRes {
  /**  */
  tokens: Token[];
}

export interface OnlyRandomFlowPreviewReq {
  /**  */
  startToken?: Token;

  /**  */
  finishToken?: Token;

  /**  */
  startNetwork?: Network;

  /**  */
  tasks?: RandomTask[];

  /**  */
  taskCount?: string;

  /**  */
  ignoreStartToken?: boolean;

  /**  */
  ignoreFinishToken?: boolean;

  /**  */
  minDelay: string;

  /**  */
  maxDelay: string;
}

export interface OnlyRandomFlowPreviewRes {
  /**  */
  flow: UniqueFlow[];

  /**  */
  uniquePercent: number;

  /**  */
  tokens: TokenArr[];
}

export interface OrbiterBridgeTask {
  /**  */
  amount: Amount;

  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  orbiterReceiverAddr?: string;

  /**  */
  orbiterSenderAddr?: string;

  /**  */
  swapCompleted?: boolean;

  /**  */
  tx?: TaskTx;
}

export interface PercentRange {
  /**  */
  min?: string;

  /**  */
  max?: string;
}

export interface RPsimple {
  /**  */
  network: Network;
}

export interface RPswap {
  /**  */
  items: RPswapItem[];
}

export interface RPswapItem {
  /**  */
  network: Network;

  /**  */
  from: Token;

  /**  */
  to: Token;

  /**  */
  am?: Amount;
}

export interface RandomTask {
  /**  */
  optional: boolean;

  /**  */
  taskType: TaskType;

  /**  */
  swap?: RPswap;

  /**  */
  simple?: RPsimple;
}

export interface ShareFlowReq {
  /**  */
  id: string;

  /**  */
  description: string;
}

export interface ShareFlowRes {
  /**  */
  id: string;
}

export interface SharedFlowReq {
  /**  */
  id: string;
}

export interface SharedFlowRes {
  /**  */
  flow: FlowShared;
}

export interface SharedFlowsReq {}

export interface SharedFlowsRes {
  /**  */
  items: FlowShared[];
}

export interface SimpleTask {
  /**  */
  network: Network;

  /**  */
  tx?: TaskTx;

  /**  */
  approveTx?: TaskTx;
}

export interface SnapshotVoteTask {
  /**  */
  space: string;

  /**  */
  proposal: object;
}

export interface StargateBridgeTask {
  /**  */
  fromNetwork: Network;

  /**  */
  toNetwork: Network;

  /**  */
  fromToken: Token;

  /**  */
  toToken: Token;

  /**  */
  fee?: string;

  /**  */
  txId?: string;

  /**  */
  layerZeroStatus?: string;

  /**  */
  lzscanUrl?: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  amount: Amount;

  /**  */
  tx?: TaskTx;
}

export interface Swap1inchTask {
  /**  */
  network: Network;

  /**  */
  fromTokenName: string;

  /**  */
  fromTokenCode: string;

  /**  */
  fromTokenAddr: string;

  /**  */
  toTokenName: string;

  /**  */
  toTokenCode: string;

  /**  */
  toTokenAddr: string;

  /**  */
  sendAll?: boolean;

  /**  */
  sendPercent?: string;

  /**  */
  sendAmount?: number;

  /**  */
  txId?: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  swapCompleted?: boolean;
}

export interface Task {
  /**  */
  weight: string;

  /**  */
  taskType: TaskType;

  /**  */
  description: string;

  /**  */
  stargateBridgeTask?: StargateBridgeTask;

  /**  */
  mockTask?: MockTask;

  /**  */
  delayTask?: DelayTask;

  /**  */
  withdrawExchangeTask?: WithdrawExchangeTask;

  /**  */
  okexDepositTask?: OkexDepositTask;

  /**  */
  testNetBridgeSwapTask?: TestNetBridgeSwapTask;

  /**  */
  snapshotVoteTask?: SnapshotVoteTask;

  /**  */
  okexBinanaceTask?: OkexBinanaceTask;

  /**  */
  swap1inchTask?: Swap1inchTask;

  /**  */
  syncSwapTask?: DefaultSwap;

  /**  */
  zkSyncOfficialBridgeToEthereumTask?: ZkSyncOfficialBridgeToEthereumTask;

  /**  */
  orbiterBridgeTask?: OrbiterBridgeTask;

  /**  */
  zkSyncOfficialBridgeFromEthereumTask?: ZkSyncOfficialBridgeFromEthereumTask;

  /**  */
  wETHTask?: WETHTask;

  /**  */
  muteioSwapTask?: DefaultSwap;

  /**  */
  syncSwapLPTask?: DefaultLP;

  /**  */
  maverickSwapTask?: DefaultSwap;

  /**  */
  spaceFiSwapTask?: DefaultSwap;

  /**  */
  velocoreSwapTask?: DefaultSwap;

  /**  */
  izumiSwapTask?: DefaultSwap;

  /**  */
  veSyncSwapTask?: DefaultSwap;

  /**  */
  ezkaliburSwapTask?: DefaultSwap;

  /**  */
  zkSwapTask?: DefaultSwap;

  /**  */
  traderJoeSwapTask?: DefaultSwap;

  /**  */
  merklyMintAndBridgeNFTTask?: MerklyMintAndBridgeNFTTask;

  /**  */
  deployStarkNetAccountTask?: DeployStarkNetAccountTask;

  /**  */
  swap10k?: DefaultSwap;

  /**  */
  pancakeSwapTask?: DefaultSwap;

  /**  */
  sithSwapTask?: DefaultSwap;

  /**  */
  jediSwapTask?: DefaultSwap;

  /**  */
  mySwapTask?: DefaultSwap;

  /**  */
  protosSwapTask?: DefaultSwap;

  /**  */
  starkNetBridgeTask?: LiquidityBridgeTask;

  /**  */
  dmailTask?: SimpleTask;

  /**  */
  starkNetIdMintTask?: SimpleTask;

  /**  */
  odosSwapTask?: DefaultSwap;

  /**  */
  acrossBridgeTask?: DefaultBridge;

  /**  */
  avnuSwapTask?: DefaultSwap;

  /**  */
  fibrousSwapTask?: DefaultSwap;

  /**  */
  exchangeSwapTask?: ExchangeSwapTask;

  /**  */
  zkLendLPTask?: DefaultLP;

  /**  */
  woofiSwapTask?: DefaultSwap;

  /**  */
  aaveLPTask?: DefaultLP;

  /**  */
  mintFunTask?: SimpleTask;

  /**  */
  mintMerklyTask?: SimpleTask;

  /**  */
  mintZeriusTask?: SimpleTask;

  /**  */
  kyberSwapTask?: DefaultSwap;

  /**  */
  ekuboSwapTask?: DefaultSwap;

  /**  */
  nostraLPTask?: DefaultLP;

  /**  */
  wethSwapTask?: DefaultSwap;

  /**  */
  coreDaoBridge?: DefaultBridge;

  /**  */
  stargateBridge?: DefaultBridge;

  /**  */
  stakeSTG?: DefaultLP;

  /**  */
  merklyRefuel?: DefaultBridge;
}

export interface TaskTx {
  /**  */
  txCompleted: boolean;

  /**  */
  txId: string;

  /**  */
  retryCount: string;

  /**  */
  url?: string;

  /**  */
  network?: Network;

  /**  */
  code?: string;

  /**  */
  gasEstimated?: AmUni;

  /**  */
  gasResult?: AmUni;

  /**  */
  gasLimit?: AmUni;

  /**  */
  multiplier?: number;

  /**  */
  details: TxDetail[];

  /**  */
  ts?: string;

  /**  */
  completeTs?: string;
}

export interface TestNetBridgeSwapTask {
  /**  */
  network: Network;

  /**  */
  minAmount: string;

  /**  */
  maxAmount: string;

  /**  */
  amount?: string;

  /**  */
  tx?: TaskTx;
}

export interface TokenArr {
  /**  */
  from: Token;

  /**  */
  to: Token[];
}

export interface TxDetail {
  /**  */
  key: string;

  /**  */
  value: string;
}

export interface UniqueFlow {
  /**  */
  tasks: Task[];
}

export interface UpdateFlowRequest {
  /**  */
  flow: flow_Flow;
}

export interface UpdateFlowResponse {
  /**  */
  flow: flow_Flow;
}

export interface UpdateFlowV2Request {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  blocks: FlowBlock[];
}

export interface UpdateFlowV2Response {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  blocks: FlowBlock[];
}

export interface UseSharedFlowReq {
  /**  */
  id: string;
}

export interface UseSharedFlowRes {
  /**  */
  id: string;
}

export interface WETHTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  wrap: boolean;

  /**  */
  tx?: TaskTx;
}

export interface WithdrawExchangeTask {
  /**  */
  withdrawerId: string;

  /**  */
  network: string;

  /**  */
  amountMin: string;

  /**  */
  amountMax: string;

  /**  */
  token: string;

  /**  */
  withdrawOrderId: string;

  /**  */
  amount?: string;

  /**  */
  txId?: string;

  /**  */
  withdrawAddr?: string;

  /**  */
  useExternalAddr?: boolean;

  /**  */
  sendAllCoins?: boolean;

  /**  */
  sendToRelatedProfile?: boolean;
}

export interface ZkSyncOfficialBridgeFromEthereumTask {
  /**  */
  amount: Amount;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  tx?: TaskTx;
}

export interface ZkSyncOfficialBridgeToEthereumTask {
  /**  */
  amount: Amount;

  /**  */
  network: Network;

  /**  */
  amountKind: string;

  /**  */
  txCompleted?: boolean;

  /**  */
  txId?: string;

  /**  */
  tx?: TaskTx;
}

export interface flow_Flow {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  tasks: Task[];

  /**  */
  nextId?: string;

  /**  */
  createdAt: Date;

  /**  */
  deletedAt?: Date;

  /**  */
  randomTasks: Task[];

  /**  */
  base?: FlowBase;
}

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE',
  'GOERLIETH' = 'GOERLIETH',
  'ZKSYNCERA' = 'ZKSYNCERA',
  'ZKSYNCERATESTNET' = 'ZKSYNCERATESTNET',
  'ZKSYNCLITE' = 'ZKSYNCLITE',
  'StarkNet' = 'StarkNet',
  'Meter' = 'Meter',
  'Tenet' = 'Tenet',
  'Canto' = 'Canto',
  'ArbitrumNova' = 'ArbitrumNova',
  'PolygonZKEVM' = 'PolygonZKEVM',
  'Fantom' = 'Fantom',
  'Base' = 'Base',
  'opBNB' = 'opBNB',
  'Linea' = 'Linea',
  'Zora' = 'Zora',
  'Core' = 'Core',
  'Conflux' = 'Conflux'
}

export enum TaskType {
  'StargateBridge' = 'StargateBridge',
  'Mock' = 'Mock',
  'Delay' = 'Delay',
  'WithdrawExchange' = 'WithdrawExchange',
  'OkexDeposit' = 'OkexDeposit',
  'TestNetBridgeSwap' = 'TestNetBridgeSwap',
  'SnapshotVote' = 'SnapshotVote',
  'OkexBinance' = 'OkexBinance',
  'Swap1inch' = 'Swap1inch',
  'SyncSwap' = 'SyncSwap',
  'ZkSyncOfficialBridgeToEthereum' = 'ZkSyncOfficialBridgeToEthereum',
  'OrbiterBridge' = 'OrbiterBridge',
  'ZkSyncOfficialBridgeFromEthereum' = 'ZkSyncOfficialBridgeFromEthereum',
  'WETH' = 'WETH',
  'MuteioSwap' = 'MuteioSwap',
  'SyncSwapLP' = 'SyncSwapLP',
  'MaverickSwap' = 'MaverickSwap',
  'SpaceFISwap' = 'SpaceFISwap',
  'VelocoreSwap' = 'VelocoreSwap',
  'IzumiSwap' = 'IzumiSwap',
  'VeSyncSwap' = 'VeSyncSwap',
  'EzkaliburSwap' = 'EzkaliburSwap',
  'ZkSwap' = 'ZkSwap',
  'TraderJoeSwap' = 'TraderJoeSwap',
  'MerklyMintAndBridgeNFT' = 'MerklyMintAndBridgeNFT',
  'DeployStarkNetAccount' = 'DeployStarkNetAccount',
  'Swap10k' = 'Swap10k',
  'PancakeSwap' = 'PancakeSwap',
  'SithSwap' = 'SithSwap',
  'JediSwap' = 'JediSwap',
  'MySwap' = 'MySwap',
  'ProtossSwap' = 'ProtossSwap',
  'StarkNetBridge' = 'StarkNetBridge',
  'Dmail' = 'Dmail',
  'StarkNetIdMint' = 'StarkNetIdMint',
  'OdosSwap' = 'OdosSwap',
  'AcrossBridge' = 'AcrossBridge',
  'AvnuSwap' = 'AvnuSwap',
  'FibrousSwap' = 'FibrousSwap',
  'ExchangeSwap' = 'ExchangeSwap',
  'ZkLendLP' = 'ZkLendLP',
  'WoofiSwap' = 'WoofiSwap',
  'AaveLP' = 'AaveLP',
  'MintFun' = 'MintFun',
  'MintMerkly' = 'MintMerkly',
  'MintZerius' = 'MintZerius',
  'KyberSwap' = 'KyberSwap',
  'EkuboSwap' = 'EkuboSwap',
  'NostraLP' = 'NostraLP',
  'CoreDaoBridge' = 'CoreDaoBridge',
  'StakeSTG' = 'StakeSTG',
  'MerklyRefuel' = 'MerklyRefuel'
}

export enum Token {
  'USDT' = 'USDT',
  'ETH' = 'ETH',
  'USDC' = 'USDC',
  'STG' = 'STG',
  'BNB' = 'BNB',
  'MATIC' = 'MATIC',
  'AVAX' = 'AVAX',
  'veSTG' = 'veSTG',
  'WETH' = 'WETH',
  'LUSD' = 'LUSD',
  'LSD' = 'LSD',
  'MUTE' = 'MUTE',
  'MAV' = 'MAV',
  'SPACE' = 'SPACE',
  'VC' = 'VC',
  'IZI' = 'IZI',
  'USDCBridged' = 'USDCBridged',
  'BUSD' = 'BUSD',
  'USDp' = 'USDp',
  'CORE' = 'CORE',
  'CFX' = 'CFX'
}
