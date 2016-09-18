package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInstanceJson(t *testing.T) {

	request := NewInstance()
	if content, err := json.Marshal(request); err == nil {

		fmt.Printf("Content is %s\n", string(content))
	}
}
