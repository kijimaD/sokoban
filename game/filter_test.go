package game

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutCG(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := NewStagePlane(4)
		s.putCG(1)
		assert.Equal(t, 1, len(s.Entities))

		countTile := 0
		for _, v := range s.Tiles {
			if v.Kind == Goal {
				countTile += 1
			}
		}
		assert.Equal(t, 1, countTile)

		c := strings.Count(s.String(), PassChar)
		assert.Equal(t, 1, c)
	})

	t.Run("2", func(t *testing.T) {
		s := NewStagePlane(4)
		s.putCG(2)
		assert.Equal(t, 2, len(s.Entities))
		countTile := 0
		for _, v := range s.Tiles {
			if v.Kind == Goal {
				countTile += 1
			}
		}
		assert.Equal(t, 2, countTile)

		c := strings.Count(s.String(), PassChar)
		assert.Equal(t, 2, c)
	})
}

func TestPutWall(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := NewStagePlane(4)
		s.putWall(2)

		c := strings.Count(s.String(), WallChar)
		assert.Equal(t, 2, c)
	})

	t.Run("not overwrite Goal tile", func(t *testing.T) {
		s := NewStagePlane(4)
		s.putCG(2)
		s.putWall(2)

		cg := strings.Count(s.String(), PassChar)
		assert.Equal(t, 2, cg)
		wc := strings.Count(s.String(), WallChar)
		assert.Equal(t, 2, wc)
	})
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
