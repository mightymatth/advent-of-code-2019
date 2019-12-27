package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfBlockTiles(t *testing.T) {
	game := NewGame("../input.txt")
	game.Start()

	var blockTileCount int
	for _, tile := range game.Screen.m {
		switch tile {
		case BlockTile:
			blockTileCount++
		default:
			continue
		}
	}

	assert.Equal(t, 452, blockTileCount)
}
