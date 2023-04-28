package u_miscellaneous

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"regexp"
)

var removeSpecialCharsRegex = regexp.MustCompile(`[^a-zA-Z0-9()@:%_\+.~#?&//=\- ]+`)

func CountCharacterInStr(input string, character string) int {
	count := 0
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) == character {
			count++
		}
	}
	return count
}

func Copy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("unable to unmarshal into dst: %s", err)
	}
	return nil
}

func Contains[T comparable](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

// Abbreviate accept not correct in some cases
func Abbreviate(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return fmt.Sprintf("%s...", str[0:maxLength])
}

func UUID2UInt(uid string) (uint64, error) {
	h := md5.New()
	h.Write([]byte(uid))
	hexStr := hex.EncodeToString(h.Sum(nil))

	bi := big.NewInt(0)
	bi, success := bi.SetString(hexStr, 36)
	if !success {
		return 0, fmt.Errorf("cast %s to uint64 failed", uid)
	}
	return bi.Uint64(), nil
}

func RemoveSpecialChars(str string) string {
	return removeSpecialCharsRegex.ReplaceAllString(str, "")
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
