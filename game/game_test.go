package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func InitStage() Stage {
	// @..#
	// .&.#
	// #_.#
	// ....

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

	// FIXME: entity初期化の順番に依存する
	player := Entity{
		&Pos{
			X: 0,
			Y: 0,
		},
		&stage,
		Player,
	}
	stage.Entities = append(stage.Entities, player)

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

func TestPlayer(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	assert.Equal(t, Player, player.Kind)
}

func TestNotPlayer(t *testing.T) {
	s := InitStage()

	notPlayer := s.Entities.NotPlayer()
	assert.Equal(t, Cargo, notPlayer.Kind)
}

func TestCollisionEntity(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.Right()
	player.Down()
	player.Down()
	player.Left()
	_, another := player.collisionEntity()
	assert.NotEqual(t, Player, another.Kind)

	// 重なってないとき
	player.Right()
	ok, _ := player.collisionEntity()
	assert.Equal(t, false, ok)
}

func TestPlayerMove(t *testing.T) {
	s := InitStage()

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

// 移動したあとにプレイヤーが残らないのを検証する
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

	player.Left()
	expect = `@..#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())
}

// 他entityと重なったとき、プレイヤーは上に表示される
func TestPlayerOver(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.Right()
	player.Down()
	player.Left()

	expect := `...#
.@.#
#_.#
....
`
	assert.Equal(t, expect, s.String())
}

func TestCollision(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.Right()
	player.Down()
	player.Down()
	player.Left()
	assert.Equal(t, true, player.isCollision())
}

// 位置で正しくentityを探せていることを確認する
func TestGetEntitiesByPos(t *testing.T) {
	s := InitStage()

	_, es := s.Entities.GetEntitiesByPos(Pos{X: 0, Y: 0})
	assert.Equal(t, Player, es[0].Kind)
	assert.Equal(t, es[0].Pos.X, 0)
	assert.Equal(t, es[0].Pos.Y, 0)

	_, es = s.Entities.GetEntitiesByPos(Pos{X: 1, Y: 1})
	assert.Equal(t, es[0].Pos.X, 1)
	assert.Equal(t, es[0].Pos.Y, 1)
	assert.Equal(t, Cargo, es[0].Kind)

	_, es = s.Entities.GetEntitiesByPos(Pos{X: 1, Y: 2})
	assert.Equal(t, es[0].Pos.X, 1)
	assert.Equal(t, es[0].Pos.Y, 2)
	assert.Equal(t, Goal, es[0].Kind)

	ok, _ := s.Entities.GetEntitiesByPos(Pos{X: 0, Y: 1})
	assert.Equal(t, false, ok)
}

func TestPush(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.Down()

	expect := `...#
.@.#
#&.#
....
`
	assert.Equal(t, expect, s.String())

	// ヌルポ
	// 	player.Down()
	// 	expect = `...#
	// ...#
	// #@.#
	// .&..
	// `
	// 	assert.Equal(t, expect, s.String())
}
