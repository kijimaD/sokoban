package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	p := Pos{X: 1, Y: 1}
	s := Stage{}
	k := Player

	e := NewEntity(p, &s, k)

	assert.Equal(t, p, *e.Pos)
	assert.Equal(t, s, *e.Stage)
	assert.Equal(t, k, e.Kind)
}

func TestCollisionEntities(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.Right()
	player.down() // 強制的に動かして重ねる
	another := player.collisionEntities()
	assert.Equal(t, 1, len(another))
	assert.NotEqual(t, Player, another[0].Kind)

	// 重なってないとき
	player.Right()
	another = player.collisionEntities()
	assert.Equal(t, 0, len(another))
}

// プレイヤーは上に表示される
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

	cargo := NewEntity(Pos{X: 1, Y: 1}, s, Cargo)
	s.Entities = append(s.Entities, cargo)

	assert.Equal(t, true, cargo.isCollision())
}

func TestPull(t *testing.T) {
	s := InitStage()

	player := s.Entities.Player()
	player.moveRelative(1, 2)
	player.PullDown()
	expect := `...#
...#
#✓.#
.@..
`
	assert.Equal(t, expect, s.String())

	player.moveRelative(0, -2)
	player.Down() // 荷物を押す
	player.moveRelative(1, 1)

	expect = `...#
...#
#_.#
.&@.
`
	assert.Equal(t, expect, s.String())

	player.PullRight()
	expect = `...#
...#
#_.#
..&@
`
	assert.Equal(t, expect, s.String())

	player.moveRelative(-2, 0)
	expect = `...#
...#
#_.#
.@&.
`
	assert.Equal(t, expect, s.String())

	player.PullLeft()
	expect = `...#
...#
#_.#
@&..
`
	assert.Equal(t, expect, s.String())

	player.moveRelative(1, -1)
	expect = `...#
...#
#@.#
.&..
`
	assert.Equal(t, expect, s.String())
	player.PullUp()
	expect = `...#
.@.#
#✓.#
....
`
	assert.Equal(t, expect, s.String())
}

func TestRandomWalk(t *testing.T) {
	s := InitBigStage()

	player := s.Entities.Player()

	for i := 0; i < 10; i++ {
		player.randomWalk()
	}
}
