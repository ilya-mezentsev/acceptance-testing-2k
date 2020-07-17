package arguments_builder

import "command/http/types"

func Build(data string) types.Arguments {
	if data == "" {
		return emptyArguments{}
	} else {
		return arguments{data}
	}
}
