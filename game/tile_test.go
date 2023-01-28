package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTileString(t *testing.T) {
	assert.Equal(t, WallChar, Tile{Kind: 0}.String())
	assert.Equal(t, FloorChar, Tile{Kind: 1}.String())
}
