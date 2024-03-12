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

export class StatService {
  /**
   *
   */
  statServiceZkSyncStat(
    params: {
      /**  */
      body: ZkSyncStatReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ZkSyncStatRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/stat/profile/zksync';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface Interaction {
  /**  */
  to: string;

  /**  */
  toUrl: string;

  /**  */
  serviceName?: string;

  /**  */
  serviceUrl?: string;

  /**  */
  amounts: InteractionAmount[];

  /**  */
  txs: string;

  /**  */
  amountUsd: number;
}

export interface InteractionAmount {
  /**  */
  token?: Token;

  /**  */
  tokenUrl: string;

  /**  */
  amountWei: string;

  /**  */
  amountUsd?: string;
}

export interface Stat {
  /**  */
  activeDays?: string;

  /**  */
  activeMonths?: string;

  /**  */
  lastActivity?: Date;

  /**  */
  interactions?: Interaction[];

  /**  */
  totalUsdAmount?: number;

  /**  */
  txCount?: string;

  /**  */
  uniqAddress?: string;
}

export interface ZkSyncStatReq {
  /**  */
  profileId: string;
}

export interface ZkSyncStatRes {
  /**  */
  stat: Stat;
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
  'CFX' = 'CFX',
  'FUSE' = 'FUSE',
  'AGLD' = 'AGLD',
  'KLAY' = 'KLAY',
  'CELO' = 'CELO',
  'SMR' = 'SMR',
  'JEWEL' = 'JEWEL',
  'STRK' = 'STRK',
  'FTM' = 'FTM'
}
