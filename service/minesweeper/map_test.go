package minesweeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMap(t *testing.T) {
	var mines int = 20
	sweeperMap := GenerateRandomMinefield(10, 10, mines, 2, 2)

	sweeperMap.Print()

	grid := sweeperMap.GetGrid()
	var expect int = 0
	for _, r := range grid {
		for _, c := range r {
			if c.Value == GridCellValueMine {
				expect++
			}
		}
	}

	assert.Equal(t, mines, expect, "mines should be equal")
	assert.NotEqual(t, grid[2][2], '*', "the starting position should have no mines")
}
