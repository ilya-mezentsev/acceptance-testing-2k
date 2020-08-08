package request_decoder

import (
	"encoding/json"
	"io"
)

func Decode(request io.ReadCloser, dest interface{}) error {
	return json.NewDecoder(request).Decode(dest)
}
