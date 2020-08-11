package types

import (
	"fmt"
	"strings"
)

type Mapping map[string]string

func (m Mapping) ReduceToRecordable() string {
	var res []string
	for key, value := range m {
		res = append(res, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(res, ";")
}
