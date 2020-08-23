export interface Fetcher {
  get(endpoint: string): Promise<ServerResponse>;
  post(endpoint: string, data: any): Promise<ServerResponse>;
  patch(endpoint: string, data: any): Promise<ServerResponse>;
  delete(endpoint: string): Promise<ServerResponse>;
}

export type ServerResponse = DefaultResponse | ErrorResponse | Response<any>;

export type Response<T> = {
  status: 'ok' | 'error'
  data: T
};

export type DefaultResponse = Response<null>;

export type ErrorResponse = Response<{
  code: string
  description: string
}>;
