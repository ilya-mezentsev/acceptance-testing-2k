export type TestObject = {name: string, hash: string};

export type TestCommand = {
  name: string,
  hash: string,
  object_name: string,
  method: string,
  base_url: string,
  endpoint: string,
  pass_arguments_in_url: boolean,
  headers: string,
  cookies: string
};

export type TestCommandMeta = {
  headers: KeyValueMapping[],
  cookies: KeyValueMapping[]
};

export type KeyValueMapping = {key: string, value: string};

export type CreateTestCommandResponse = {command_hash: string};
