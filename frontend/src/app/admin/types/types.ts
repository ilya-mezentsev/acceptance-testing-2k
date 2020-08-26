export type Object = {name: string, hash: string};

export type Command = {
  name: string,
  hash: string,
  object_name: string,
  method: string,
  base_url: string,
  endpoint: string,
  pass_arguments_in_url: boolean,
  headers: Map<string, string>,
  cookies: Map<string, string>
};
