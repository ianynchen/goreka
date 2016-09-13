package util

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

func TestGetIP(t *testing.T) {

	ip, ok := GetIP()
	fmt.Printf("ip address: %s\n", ip)
	assert.True(t, ok)
}
