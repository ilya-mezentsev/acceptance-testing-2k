export type TestObject = {name: string, hash: string};

export interface TestCommandSettings {
  name: string;
  hash: string;
  object_hash: string;
  method: string;
  base_url: string;
  endpoint: string;
  timeout: number;
  pass_arguments_in_url: boolean;
}

export interface TestCommandRecord extends TestCommandSettings, TestCommandMeta {}

export interface TestCommandMeta {
  headers: KeyValueMapping[];
  cookies: KeyValueMapping[];
}

export type KeyValueMapping = {
  key: string,
  value: string,
  hash: string,
  object_hash: string
};

export type CreateTestCommandResponse = {command_hash: string};

export type CreateTestsFile = {filename: string};

export type TransactionError = {
  code: string,
  description: string,
  transactionText?: string,
  testCaseText?: string,
};

export type TestsReport = {
  passedCount: number,
  failedCount: number,
  errors?: TransactionError[],
};

export type RunTestsResult = {
  report: TestsReport,
  applicationError: TransactionError,
};

export type ServiceError = {
  code: string,
  description: string,
};

export interface GeneralUpdateRequest {
  command_hashes: string[];
}

export interface GeneralBaseURLsUpdateRequest extends GeneralUpdateRequest {
  base_url: string;
}

export interface GeneralTimeoutUpdateRequest extends GeneralUpdateRequest {
  timeout: number;
}

export interface GeneralHeadersCreateRequest extends GeneralUpdateRequest {
  headers: KeyValueMapping[];
}

export interface GeneralCookiesCreateRequest extends GeneralUpdateRequest {
  cookies: KeyValueMapping[];
}
