package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ....
// ....
// ..@.
// ....

func InitStage() Stage {
	tiles := map[Pos]Tile{
		Pos{X: 0, Y: 0}: Tile{Kind: 1},
		Pos{X: 0, Y: 1}: Tile{Kind: 0},
		Pos{X: 1, Y: 0}: Tile{Kind: 1},
		Pos{X: 1, Y: 1}: Tile{Kind: 0},
	}
	player := Entity{
		&Pos{
			X: 0,
			Y: 0,
		},
	}
	entities := map[Pos]Entity{}

	stage := Stage{
		Tiles:    tiles,
		Player:   player,
		Entities: entities,
	}
	return stage
}

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

func TestPlayerMove(t *testing.T) {
	s := InitStage()

	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
	s.Player.Right()
	assert.Equal(t, &Pos{X: 1, Y: 0}, s.Player.Pos)
	s.Player.Left()
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
	s.Player.Down()
	assert.Equal(t, &Pos{X: 0, Y: 1}, s.Player.Pos)
	s.Player.Up()
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
}
