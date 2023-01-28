package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStageByString(t *testing.T) {
	tileChar := `..#
.#.
...
`
	s := NewStageByString(tileChar)
	assert.Equal(t, tileChar, s.String())
}
