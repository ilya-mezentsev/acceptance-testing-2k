package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5WithTimeAsKey(s string) string {
	return Md5(fmt.Sprintf("%s%d", s, rand.Intn(int(time.Now().UnixNano()))))
}
