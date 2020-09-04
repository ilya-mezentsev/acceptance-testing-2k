export type TestObject = {name: string, hash: string};

export interface TestCommandSettings {
  name: string;
  hash: string;
  object_hash: string;
  method: string;
  base_url: string;
  endpoint: string;
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
