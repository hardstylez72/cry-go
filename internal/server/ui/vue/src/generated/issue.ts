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

export class IssueService {
  /**
   *
   */
  issueServiceAddComment(
    params: {
      /**  */
      body: AddCommentReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<AddCommentRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/comment/add';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceDeleteComment(
    params: {
      /**  */
      body: DeleteCommentReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteCommentRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/comment/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceCreateIssue(
    params: {
      /**  */
      body: CreateIssueReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateIssueRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/create';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceDeleteIssue(
    params: {
      /**  */
      body: DeleteIssueReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteIssueRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceIssue(
    params: {
      /**  */
      body: IssueReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<IssueRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceIssues(
    params: {
      /**  */
      body: IssuesReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<IssuesRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceIssueUpdateStatus(
    params: {
      /**  */
      body: IssueUpdateStatusReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<IssueUpdateStatusRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/status/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  issueServiceIssuesUser(
    params: {
      /**  */
      body: IssuesReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<IssuesRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/issue/user-list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface AddCommentReq {
  /**  */
  issueId?: string;

  /**  */
  text: string;
}

export interface AddCommentRes {}

export interface CreateIssueReq {
  /**  */
  title: string;

  /**  */
  description: string;

  /**  */
  processId?: string;

  /**  */
  taskid?: string;
}

export interface CreateIssueRes {}

export interface DeleteCommentReq {
  /**  */
  commentId: string;
}

export interface DeleteCommentRes {}

export interface DeleteIssueReq {
  /**  */
  issueId: string;
}

export interface DeleteIssueRes {}

export interface Issue {
  /**  */
  id: string;

  /**  */
  createdAt: Date;

  /**  */
  finishedAt?: Date;

  /**  */
  createdBy: string;

  /**  */
  title: string;

  /**  */
  description: string;

  /**  */
  status: IssueStatus;

  /**  */
  my: boolean;
}

export interface IssueComment {
  /**  */
  id: string;

  /**  */
  issueId: string;

  /**  */
  creatorId: string;

  /**  */
  description: string;

  /**  */
  createdAt: Date;
}

export interface IssueReq {
  /**  */
  id: string;
}

export interface IssueRes {
  /**  */
  issue?: Issue;

  /**  */
  comments: IssueComment[];
}

export interface IssueUpdateStatusReq {
  /**  */
  issueId?: string;

  /**  */
  status: IssueStatus;
}

export interface IssueUpdateStatusRes {}

export interface IssuesReq {
  /**  */
  limit: string;

  /**  */
  offset: string;
}

export interface IssuesRes {
  /**  */
  items: Issue[];
}

export enum IssueStatus {
  'Created' = 'Created',
  'Processing' = 'Processing',
  'Done' = 'Done',
  'Deleted' = 'Deleted',
  'Rejected' = 'Rejected'
}
