package encrypt

import (
	"fmt"
	"testing"
)

func TestGenerateRandomStringWithLength(t *testing.T) {
	text1 := GenerateRandomStringWithLength(64)
	fmt.Println(text1)
}
