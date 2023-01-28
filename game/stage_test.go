package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStageByString(t *testing.T) {
	tileChar := `..#
..#
_..
`
	entityChar := `@~~
~&~
~~~
`
	expect := `@.#
.&#
_..
`
	s := NewStageByString(tileChar, entityChar)
	assert.Equal(t, expect, s.String())
}
