package game

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCG(t *testing.T) {
	s := NewStagePlane(4)
	s.setCG(1)
	assert.Equal(t, 1, len(s.Entities))

	// FIXME: バグってる
	// s.setCG(2)
	// assert.Equal(t, 2, len(s.Entities))
}

func TestRandomPos(t *testing.T) {
	s := NewStagePlane(4)
	for i := 0; i < 10; i++ {
		p := s.randomPos()
		assert.Equal(t, reflect.TypeOf(Pos{}), reflect.TypeOf(p))
		assert.Equal(t, true, 0 <= p.X && p.X <= 3)
		assert.Equal(t, true, 0 <= p.Y && p.Y <= 3)
	}
}

func TestRandomPoses(t *testing.T) {
	s := NewStagePlane(4)
	poses := s.randomPoses(3)
	assert.Equal(t, 3, len(poses))
	assert.Equal(t, 3, len(uniq(poses)))

	posesMax := s.randomPoses(16)
	assert.Equal(t, 16, len(posesMax))
	assert.Equal(t, 16, len(uniq(posesMax)))

}

func TestUniq(t *testing.T) {
	dup1 := []Pos{Pos{1, 2}, Pos{1, 2}}
	assert.Equal(t, []Pos{Pos{1, 2}}, uniq(dup1))

	dup2 := []Pos{Pos{1, 1}, Pos{1, 2}, Pos{1, 1}}
	assert.Equal(t, []Pos{Pos{1, 1}, Pos{1, 2}}, uniq(dup2))
}
