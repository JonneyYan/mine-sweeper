package minesweeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartGame(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyPrimary)
	grid := game.Minefield.GetGrid()

	assert.EqualValues(t, DefaultGameConfig.Difficulties[DifficultyPrimary][0], len(grid))
	assert.EqualValues(t, DefaultGameConfig.Difficulties[DifficultyPrimary][1], len(grid[1]))
}

func TestDifficult(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyMedium)
	grid := game.Minefield.GetGrid()

	assert.EqualValues(t, DefaultGameConfig.Difficulties[DifficultyMedium][0], len(grid))
	assert.EqualValues(t, DefaultGameConfig.Difficulties[DifficultyMedium][1], len(grid[1]))
}

func TestClearMine(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyPrimary)

	mine := GridCell{Value: GridCellValueMine}
	g0 := GridCell{Value: 0}
	g1 := GridCell{Value: 1}
	g2 := GridCell{Value: 2}

	// set mock map
	game.Minefield = &Minefield{
		Rows:         5,
		Cols:         5,
		WinThreshold: 25 - 3,
		grid: [][]GridCell{
			{mine, mine, g1, g0, g0},
			{g2, g2, g1, g0, g0},
			{g0, g0, g0, g0, g0},
			{g0, g1, g1, g1, g0},
			{g0, g1, mine, g1, g0},
		},
	}

	// find mine, game over
	game.ClearMine(0, 0)
	grid, _ := game.ScanMinefield()

	assert.Equal(t, GridCellValueMine, grid[0][0].Value)
	assert.Equal(t, [2]int{0, 0}, game.Steps[0])
	assert.Equal(t, MinesweeperStatusGameOver, game.Status)

	grid, err := game.ScanMinefield()
	assert.NoError(t, err)
	assert.Equal(t, GridCellValueMine, grid[0][1].Value)

}

func TestWin(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyPrimary)

	mine := GridCell{Value: GridCellValueMine}
	g0 := GridCell{Value: 0}
	g1 := GridCell{Value: 1}
	g2 := GridCell{Value: 2}

	// set mock map
	game.Minefield = &Minefield{
		Rows:         3,
		Cols:         3,
		WinThreshold: 9 - 3,
		grid: [][]GridCell{
			{mine, mine, g1},
			{g2, g2, g1},
			{g0, g1, mine},
		},
	}

	game.ClearMine(0, 2)
	game.ClearMine(1, 0)
	game.ClearMine(1, 1)
	game.ClearMine(1, 2)
	game.ClearMine(2, 0)
	assert.Equal(t, MinesweeperStatusLive, game.Status)
	game.ClearMine(2, 1)

}

func TestFlagging(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyPrimary)

	grid, _ := game.ScanMinefield()
	assert.False(t, grid[0][2].Flagging)

	game.Flagging(0, 2)
	grid, _ = game.ScanMinefield()
	assert.True(t, grid[0][2].Flagging)

	game.Flagging(0, 2)
	grid, _ = game.ScanMinefield()
	assert.False(t, grid[0][2].Flagging)
}

func TestClearAndFlagging(t *testing.T) {
	game, _ := NewGame("User", 3, 3, DifficultyPrimary)

	mine := GridCell{Value: GridCellValueMine}
	g0 := GridCell{Value: 0}
	g1 := GridCell{Value: 1}
	g2 := GridCell{Value: 2}

	// set mock map
	game.Minefield = &Minefield{
		Rows:         5,
		Cols:         5,
		WinThreshold: 25 - 3,
		grid: [][]GridCell{
			{mine, mine, g1, g0, g0},
			{g2, g2, g1, g0, g0},
			{g0, g0, g0, g0, g0},
			{g0, g1, g1, g1, g0},
			{g0, g1, mine, g1, g0},
		},
	}

	game.Flagging(2, 0)

	game.ClearMine(2, 2)
	game.ClearAdjacent(2, 1)
	grid, _ := game.ScanMinefield()

	assert.True(t, grid[1][0].Cleared)
	assert.True(t, grid[1][1].Cleared)
	assert.True(t, grid[1][2].Cleared)

	assert.False(t, grid[2][0].Cleared)
	assert.True(t, grid[2][1].Cleared)
	assert.True(t, grid[2][2].Cleared)

	assert.True(t, grid[3][0].Cleared)
	assert.True(t, grid[3][1].Cleared)
	assert.True(t, grid[3][2].Cleared)
}
