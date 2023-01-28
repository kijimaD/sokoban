package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCollisionEntities(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.Right()
	player.Down()
	player.Down()
	player.Left()
	another := player.collisionEntities()
	assert.Equal(t, 1, len(another))
	assert.NotEqual(t, Player, another[0].Kind)

	// 重なってないとき
	player.Right()
	another = player.collisionEntities()
	assert.Equal(t, 0, len(another))
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
	player.Down()
	player.Left()

	expect := `...#
.&.#
#@.#
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

func TestCollisionCargo(t *testing.T) {
	s := InitStage()

	cargo := Entity{
		&Pos{
			X: 1,
			Y: 1,
		},
		s,
		Cargo,
	}
	s.Entities = append(s.Entities, cargo)

	assert.Equal(t, true, cargo.isCollision())
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

	// 押せる
	player := s.Entities.Player()
	player.Right()
	player.Down()

	expect := `...#
.@.#
#✓.#
....
`
	assert.Equal(t, expect, s.String())

	// ゴールの上の荷物を押せる
	player.Down()
	expect = `...#
...#
#@.#
.&..
`
	assert.Equal(t, expect, s.String())

	// 押す先が壁だと移動できない
	player.Down()
	expect = `...#
...#
#@.#
.&..
`
	assert.Equal(t, expect, s.String())

}

// goalが他のentityと同じ座標にあるときは表示が変わる
func TestDisplayPriority(t *testing.T) {
	s := InitStage()
	player := s.Entities.Player()
	player.Right()
	expect := `.@.#
.&.#
#_.#
....
`
	assert.Equal(t, expect, s.String())

	// cargoと重なるとき✓
	player.Down()
	expect = `...#
.@.#
#✓.#
....
`
	assert.Equal(t, expect, s.String())

	// playerと重なるとき@
	player.Down()
	expect = `...#
...#
#@.#
.&..
`
	assert.Equal(t, expect, s.String())
}

// 2つは押せない
func TestPushDouble(t *testing.T) {
	s := InitStage()
	cargo := Entity{
		&Pos{
			X: 2,
			Y: 1,
		},
		s,
		Cargo,
	}
	s.Entities = append(s.Entities, cargo)
	player := s.Entities.Player()
	player.Down()
	expect := `...#
@&&#
#_.#
....
`
	assert.Equal(t, expect, s.String())
	player.Right()
	expect = `...#
@&&#
#_.#
....
`
	assert.Equal(t, expect, s.String())
}

func TestPushDoubleNoWall(t *testing.T) {
	s := InitStage()
	cargo := Entity{
		&Pos{
			X: 1,
			Y: 2,
		},
		s,
		Cargo,
	}
	s.Entities = append(s.Entities, cargo)
	player := s.Entities.Player()
	player.Right()
	expect := `.@.#
.&.#
#✓.#
....
`
	assert.Equal(t, expect, s.String())

	player.Down()
	expect = `.@.#
.&.#
#✓.#
....
`
	assert.Equal(t, expect, s.String())
}

func TestIsFinish(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()

	player.Right()
	assert.Equal(t, false, s.Entities.isFinish())

	player.Down()
	assert.Equal(t, true, s.Entities.isFinish())

	player.Down()
	assert.Equal(t, false, s.Entities.isFinish())
}
