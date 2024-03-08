package minesweeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrParams    = errors.New("params error")
	ErrOverflow  = errors.New("overflow")
	ErrGameOver  = errors.New("game over")
	ErrPrivilege = errors.New("game not finish")
	ErrFlagged   = errors.New("have flagged")
	ErrEmpty     = errors.New("no game")
)

const (
	GridCellValuePlaceholder = '❔'
	GridCellValueInitial     = '🎯'
	GridCellValueMine        = '💣'
)

type GridCell struct {
	Value    rune `json:"value"`
	Cleared  bool `json:"cleared"`
	Flagging bool `json:"flagging"`
}
type GridCellJson struct {
	Value    string `json:"value"`
	Cleared  bool   `json:"cleared"`
	Flagging bool   `json:"flagging"`
}

func (c GridCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(GridCellJson{
		Value:    string(c.Value),
		Cleared:  c.Cleared,
		Flagging: c.Flagging,
	})

}

type Grid = [][]GridCell

type Minefield struct {
	Rows         int
	Cols         int
	WinThreshold int

	grid Grid
}

func GenerateRandomMinefield(rows, cols, mines, xInti, yInti int) *Minefield {
	grid := make(Grid, rows)
	for i := range grid {
		grid[i] = make([]GridCell, cols)
		for j := range grid[i] {
			if i == int(xInti) && j == int(yInti) {
				grid[i][j] = GridCell{Value: GridCellValueInitial}
			} else {
				grid[i][j] = GridCell{Value: GridCellValuePlaceholder}
			}
		}
	}

	m := &Minefield{
		Rows:         rows,
		Cols:         cols,
		WinThreshold: (rows * cols) - mines,
		grid:         grid,
	}

	m.placeMine(mines)
	m.placeCounter()
	return m
}

func (m *Minefield) GetGrid() Grid {
	return m.grid
}
func (m *Minefield) Print() {
	for _, row := range m.grid {
		for _, cell := range row {
			if int(cell.Value) < 9 {
				fmt.Printf("[%d-%v-%v]", cell.Value, cell.Cleared, cell.Flagging)
			} else {
				fmt.Printf("[%c-%v-%v]", cell.Value, cell.Cleared, cell.Flagging)
			}

		}
		fmt.Println()
	}
}

func (m *Minefield) countAdjacentMines(row, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < int(m.Rows) && c >= 0 && c < int(m.Cols) && m.grid[r][c].Value == GridCellValueMine {
				count++
			}
		}
	}
	return count
}

func (m *Minefield) CountAdjacentFlags(row, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < int(m.Rows) && c >= 0 && c < int(m.Cols) && m.grid[r][c].Flagging {
				count++
			}
		}
	}
	return count
}

func (m *Minefield) placeCounter() {
	for i := range m.grid {
		for j := range m.grid[i] {
			if m.grid[i][j].Value != GridCellValueMine {
				m.grid[i][j].Value = rune(m.countAdjacentMines(i, j) + '0')
			}
		}
	}
}

func (m *Minefield) placeMine(mines int) {
	for i := 0; i < int(mines); i++ {
		m.randomPlace()
	}

}
func (m *Minefield) randomPlace() {
	rand.NewSource(time.Now().UnixNano())
	x := rand.Intn(int(m.Rows))
	y := rand.Intn(int(m.Cols))

	if m.grid[x][y].Value == GridCellValuePlaceholder {
		m.grid[x][y].Value = GridCellValueMine
	} else {
		m.randomPlace()
	}
}
