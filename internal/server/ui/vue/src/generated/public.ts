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

export class PublicService {
  /**
   *
   */
  publicServiceSwapStat(
    params: {
      /**  */
      body: SwapStatReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SwapStatRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/public/swap-stat';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface SwapStatItem {
  /**  */
  type: TaskType;

  /**  */
  rateRatio: number;

  /**  */
  slippage: number;

  /**  */
  sum: number;
}

export interface SwapStatReq {}

export interface SwapStatRes {
  /**  */
  zkSyncETHUSDC?: SwapStatItem[];

  /**  */
  zkSyncUSDCETH?: SwapStatItem[];

  /**  */
  starknetETHUSDC?: SwapStatItem[];

  /**  */
  starknetUSDCETH?: SwapStatItem[];

  /**  */
  updated: Date;
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
  'ZkLandLP' = 'ZkLandLP'
}
