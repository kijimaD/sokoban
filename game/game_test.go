package game

import (
	"testing"
)

// ....
// ....
// ..@.
// ....

func TestInit(t *testing.T) {
	s := InitStage()
	s.Show()
	s.ToSlice()
}
