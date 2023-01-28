package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	p := Pos{X: 1, Y: 1}
	s := Stage{}
	k := Player

	e := NewEntity(&p, &s, k)

	assert.Equal(t, p, *e.Pos)
	assert.Equal(t, s, *e.Stage)
	assert.Equal(t, k, e.Kind)
}
