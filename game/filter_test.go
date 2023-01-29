package game

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCG(t *testing.T) {
	s := NewStagePlane(4)
	s.setCG(3)
}

func TestRandomPos(t *testing.T) {
	s := NewStagePlane(4)
	p := s.randomPos()
	assert.Equal(t, reflect.TypeOf(Pos{}), reflect.TypeOf(p))
}

func TestRandomPoses(t *testing.T) {}

func TestUniq(t *testing.T) {
	dup1 := []Pos{Pos{1, 2}, Pos{1, 2}}
	assert.Equal(t, []Pos{Pos{1, 2}}, uniq(dup1))

	dup2 := []Pos{Pos{1, 1}, Pos{1, 2}, Pos{1, 1}}
	assert.Equal(t, []Pos{Pos{1, 1}, Pos{1, 2}}, uniq(dup2))
}
