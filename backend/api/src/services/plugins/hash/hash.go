package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GetHashWithTimeAsKey(s string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%s%d", s, time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))
}
