package game

import (
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

	stage := Stage{
		Tiles:    tiles,
		Entities: Entities{},
	}

	cargo := Entity{
		&Pos{
			X: 1,
			Y: 1,
		},
		&stage,
		Cargo,
	}
	stage.Entities = append(stage.Entities, cargo)

	goal := Entity{
		&Pos{
			X: 1,
			Y: 2,
		},
		&stage,
		Goal,
	}
	stage.Entities = append(stage.Entities, goal)

	player := Entity{
		&Pos{
			X: 0,
			Y: 0,
		},
		&stage,
		Player,
	}
	stage.Entities = append(stage.Entities, player)

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

	expect := `@..#
.&.#
#_.#
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
	// .&.#
	// #_.#
	// ....

	player := s.Entities.Player()

	// 通常移動
	assert.Equal(t, &Pos{X: 0, Y: 0}, player.Pos)
	player.Right()
	assert.Equal(t, &Pos{X: 1, Y: 0}, player.Pos)
	player.Left()
	assert.Equal(t, &Pos{X: 0, Y: 0}, player.Pos)
	player.Down()
	assert.Equal(t, &Pos{X: 0, Y: 1}, player.Pos)
	player.Up()
	assert.Equal(t, &Pos{X: 0, Y: 0}, player.Pos)

	// 移動不可を検証
	assert.Equal(t, &Pos{X: 0, Y: 0}, player.Pos)
	player.Down()
	player.Down()
	player.Down()
	assert.Equal(t, &Pos{X: 0, Y: 1}, player.Pos) // 移動先が壁タイルの場合
	player.Left()
	assert.Equal(t, &Pos{X: 0, Y: 1}, player.Pos) // 移動先のタイルがない場合
	player.Up()
	player.Up()
	player.Up()
	assert.Equal(t, &Pos{X: 0, Y: 0}, player.Pos) // 移動先のタイルがない場合
	player.Right()
	player.Right()
	player.Right()
	player.Right()
	assert.Equal(t, &Pos{X: 2, Y: 0}, player.Pos) // 移動先のタイルがない場合
}

// 移動したあとに残らないのを検証する
func TestPlayerUnique(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()

	expect := `.@.#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())
}

func TestCollision(t *testing.T) {
	s := InitStage()

	// @..#
	// .&.#
	// #_.#
	// ....

	player := s.Entities.Player()
	player.Right()
	player.Down()
	assert.Equal(t, true, player.isCollision())
}

func TestGetEntityByPos(t *testing.T) {
	s := InitStage()

	_, e := s.Entities.GetEntityByPos(Pos{X: 0, Y: 0})
	assert.Equal(t, Player, e.Kind)
	assert.Equal(t, e.Pos.X, 0)
	assert.Equal(t, e.Pos.Y, 0)

	_, e = s.Entities.GetEntityByPos(Pos{X: 1, Y: 1})
	assert.Equal(t, e.Pos.X, 1)
	assert.Equal(t, e.Pos.Y, 1)
	assert.Equal(t, Cargo, e.Kind)

	ok, e := s.Entities.GetEntityByPos(Pos{X: 0, Y: 1})
	assert.Equal(t, false, ok)
}
