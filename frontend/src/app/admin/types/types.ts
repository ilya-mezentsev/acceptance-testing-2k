export type TestObject = {name: string, hash: string};

export type TestCommand = {
  name: string,
  hash: string,
  object_name: string,
  method: string,
  base_url: string,
  endpoint: string,
  pass_arguments_in_url: boolean,
  headers: {[k: string]: string},
  cookies: {[k: string]: string}
};

export type KeyValueMapping = {key: string, value: string};
