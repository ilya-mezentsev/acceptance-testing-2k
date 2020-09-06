export interface Fetcher {
  get(endpoint: string): Promise<ServerResponse>;
  post(endpoint: string, data: any): Promise<ServerResponse>;
  patch(endpoint: string, data: any): Promise<ServerResponse>;
  delete(endpoint: string): Promise<ServerResponse>;
}

export interface FileSender {
  sendFile<T>(endpoint: string, file: FormData): Promise<ErrorResponse | Response<T>>;
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

export type UpdatePayload = {
  hash: string,
  field_name: string,
  new_value: any
};
