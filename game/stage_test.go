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

func TestNewStagePlane(t *testing.T) {
	s := NewStagePlane(5)
	expect := `.....
.....
.....
.....
.....
`
	assert.Equal(t, expect, s.String())

}

func TestStageString(t *testing.T) {
	s := InitStage()

	expect := `@..#
.&.#
#_.#
....
`
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

func TestStageToSlice(t *testing.T) {
	s := InitStage()

	// FIXME: 合ってない
	expect := [][]Tile{
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 0}, Tile{Kind: 1}},
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 2}, Tile{Kind: 1}},
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}},
		{Tile{Kind: 0}, Tile{Kind: 0}, Tile{Kind: 0}, Tile{Kind: 1}},
	}

	assert.Equal(t, expect, s.ToSlice())
}

func TestIsFinish(t *testing.T) {
	s := InitStage()
	player := s.Entities.Player()

	player.Right()
	assert.Equal(t, false, s.IsFinish())

	player.Down()
	assert.Equal(t, true, s.IsFinish())

	player.Down()
	assert.Equal(t, false, s.IsFinish())
}

func TestReset(t *testing.T) {
	s := InitStage()

	var poses []Pos

	for _, e := range s.Entities {
		poses = append(poses, *e.Pos)
	}

	player := s.Entities.Player()
	player.Right()

	s.ResetPos(poses)
	expect := `@..#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())

	player = s.Entities.Player()
	player.Right()
	expect = `.@.#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())

	s.ResetPos(poses)
	expect = `@..#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())
}
