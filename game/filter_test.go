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

// func TestRandomPoses(t *testing.T) {}
