package cfgman

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

func TestSearchFile(t *testing.T) {

	if file, ok := findFile([]string{".", "/Users/ianynchen", "/Users/ianynchen/programming/code-study/nifty"}, "README.md"); ok {
		assert.Equal(t, file, "/Users/ianynchen/programming/code-study/nifty/README.md")
	} else {
		fmt.Printf("error searching for file")
	}
}
