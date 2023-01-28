package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
