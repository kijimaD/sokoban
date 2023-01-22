package game

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitStage() Stage {
	tiles := map[Pos]Tile{
		Pos{X: 0, Y: 0}: Tile{Kind: 1},
		Pos{X: 1, Y: 0}: Tile{Kind: 1},
		Pos{X: 2, Y: 0}: Tile{Kind: 1},
		Pos{X: 3, Y: 0}: Tile{Kind: 0},

		Pos{X: 0, Y: 1}: Tile{Kind: 1},
		Pos{X: 1, Y: 1}: Tile{Kind: 1},
		Pos{X: 2, Y: 1}: Tile{Kind: 1},
		Pos{X: 3, Y: 1}: Tile{Kind: 0},

		Pos{X: 0, Y: 2}: Tile{Kind: 0},
		Pos{X: 1, Y: 2}: Tile{Kind: 1},
		Pos{X: 2, Y: 2}: Tile{Kind: 1},
		Pos{X: 3, Y: 2}: Tile{Kind: 0},

		Pos{X: 0, Y: 3}: Tile{Kind: 1},
		Pos{X: 1, Y: 3}: Tile{Kind: 1},
		Pos{X: 2, Y: 3}: Tile{Kind: 1},
		Pos{X: 3, Y: 3}: Tile{Kind: 1},
	}
	player := Entity{
		&Pos{
			X: 0,
			Y: 0,
		},
		nil,
	}
	entities := map[Pos]Entity{}

	stage := Stage{
		Tiles:    tiles,
		Player:   player,
		Entities: entities,
	}
	stage.Player.Stage = &stage
	return stage
}

func TestInit(t *testing.T) {
	s := InitStage()
	s.ToSlice()
}

func TestTileString(t *testing.T) {
	assert.Equal(t, WallChar, Tile{Kind: 0}.String())
	assert.Equal(t, FloorChar, Tile{Kind: 1}.String())
}

func TestStageString(t *testing.T) {
	s := InitStage()
	fmt.Println(s)

	expect := `@..#
...#
#..#
....
`
	assert.Equal(t, expect, s.String())
}

func TestToSlice(t *testing.T) {
	s := InitStage()

	expect := [][]Tile{
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 0}, Tile{Kind: 1}},
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}},
		{Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}, Tile{Kind: 1}},
		{Tile{Kind: 0}, Tile{Kind: 0}, Tile{Kind: 0}, Tile{Kind: 1}},
	}

	assert.Equal(t, expect, s.ToSlice())
}

func TestPlayerMove(t *testing.T) {
	s := InitStage()
	// @..#
	// ...#
	// #..#
	// ....

	// 通常移動
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
	s.Player.Right()
	assert.Equal(t, &Pos{X: 1, Y: 0}, s.Player.Pos)
	s.Player.Left()
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
	s.Player.Down()
	assert.Equal(t, &Pos{X: 0, Y: 1}, s.Player.Pos)
	s.Player.Up()
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)

	// 移動不可を検証
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos)
	s.Player.Down()
	s.Player.Down()
	s.Player.Down()
	assert.Equal(t, &Pos{X: 0, Y: 1}, s.Player.Pos) // 移動先が壁タイルの場合
	s.Player.Left()
	assert.Equal(t, &Pos{X: 0, Y: 1}, s.Player.Pos) // 移動先のタイルがない場合
	s.Player.Up()
	s.Player.Up()
	s.Player.Up()
	assert.Equal(t, &Pos{X: 0, Y: 0}, s.Player.Pos) // 移動先のタイルがない場合
	s.Player.Right()
	s.Player.Right()
	s.Player.Right()
	s.Player.Right()
	assert.Equal(t, &Pos{X: 2, Y: 0}, s.Player.Pos) // 移動先のタイルがない場合
}
