package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	ok, _ := s.Entities.GetEntitiesByPos(Pos{X: 0, Y: 1})
	assert.Equal(t, false, ok)
}
