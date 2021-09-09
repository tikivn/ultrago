package u_hash

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func HashSHA1(key string) string {
	hashFunc := sha1.New()
	io.WriteString(hashFunc, key)
	return fmt.Sprintf("%x", hashFunc.Sum(nil))
}
