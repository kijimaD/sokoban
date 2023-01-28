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

func TestStageStrToArray(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		input := `...
.#.
..#
`
		expect := []string{
			"...", ".#.", "..#",
		}

		err, result := stageStrToArray(input)
		assert.Equal(t, err, nil)
		assert.Equal(t, expect, result)
	})

	t.Run("fail", func(t *testing.T) {
		input := `..
...
..
`
		err, _ := stageStrToArray(input)
		assert.Equal(t, StageInvalidError, err)
	})
}
