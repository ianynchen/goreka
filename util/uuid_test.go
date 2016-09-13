package util

import (
	"testing"

	"github.com/issue9/assert"
)

func TestUniqueness(t *testing.T) {

	id1 := GetUuid()
	id2 := GetUuid()
	id3 := GetUuid()

	assert.True(t, id1 != id2)
	assert.True(t, id2 != id3)
	assert.True(t, id3 != id1)
}
