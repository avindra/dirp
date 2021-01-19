package dirp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	x := 1
	y := 1

	assert.Equal(t, x, y, "Bullshit test")
}
