package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ....
// ....
// ..@.
// ....

func TestInit(t *testing.T) {
	s := InitStage()
	s.String()
	s.ToSlice()
}

func TestTileString(t *testing.T) {
	assert.Equal(t, WALL, Tile{Kind: 0}.String())
	assert.Equal(t, FLOOR, Tile{Kind: 1}.String())
}

func TestStageString(t *testing.T) {
	s := InitStage()
	expect := `.##
.##
###
`
	assert.Equal(t, expect, s.String())
}

func TestToSlice(t *testing.T) {
	s := InitStage()

	expect := [][]Tile{
		{Tile{Kind: 1}, Tile{Kind: 0}},
		{Tile{Kind: 1}, Tile{Kind: 0}},
	}

	assert.Equal(t, expect, s.ToSlice())
}
