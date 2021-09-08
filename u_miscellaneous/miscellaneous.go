package miscellaneous

import (
	"encoding/json"
	"fmt"
)

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
