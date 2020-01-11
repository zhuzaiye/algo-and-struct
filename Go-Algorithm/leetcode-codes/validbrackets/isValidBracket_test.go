package validbrackets

import (
	"fmt"
	"testing"
)

func TestIsValidBracket(t *testing.T) {
	inputString := "{[[{}]{}]}"
	rst := IsValidBracket(inputString)
	fmt.Println(rst)
}
