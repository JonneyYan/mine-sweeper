package minesweeper

import (
	"time"
)

const (
	MinesweeperStatusLive = iota
	MinesweeperStatusGameOver
	MinesweeperStatusWin
)
const (
	DifficultyPrimary = iota // 0
	DifficultyMedium
	DifficultyExpert
)

// ref：https://zh.wikipedia.org/wiki/%E8%B8%A9%E5%9C%B0%E9%9B%B7#Windows_%E7%89%88%E6%9C%AC
var DefaultGameConfig = GameConfig{
	Difficulties: map[int][3]int{
		DifficultyPrimary: {8, 8, 10},
		DifficultyMedium:  {16, 16, 40},
		DifficultyExpert:  {16, 30, 99},
	},
}

type MinesweeperGame struct {
	Username  string     `json:"username"`
	ID        int64      `json:"id"`
	Minefield *Minefield `json:"-"`

	Difficulty int       `json:"difficulty"`
	Status     int       `json:"status"`
	Cleared    int       `json:"cleared"`
	Steps      [][2]int  `json:"steps"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

type MinesweeperGameAPI struct {
	MinesweeperGame
	Grid [][]GridCell `json:"grid"`
}
type GameConfig struct {
	Difficulties map[int][3]int `json:"difficulties"`
}

func NewGame(username string, x, y, difficulty int) (*MinesweeperGame, error) {
	diff := DefaultGameConfig.Difficulties[difficulty]

	if diff == [3]int{} {
		return nil, ErrParams
	}

	game := &MinesweeperGame{
		Username:   username,
		Difficulty: difficulty,
		ID:         time.Now().UnixNano(),
		Minefield:  GenerateRandomMinefield(diff[0], diff[1], diff[2], x, y),
		Cleared:    0,
		Status:     MinesweeperStatusLive,
		Steps:      [][2]int{},
		StartTime:  time.Now(),
	}

	return game, nil
}
func (game *MinesweeperGame) ClearMine(x, y int) {
	if game.Minefield.Rows < x || game.Minefield.Cols < y {
		return
	}

	if game.Status != MinesweeperStatusLive {
		return
	}

	grid := game.Minefield.GetGrid()

	if grid[x][y].Flagging || grid[x][y].Cleared {
		return
	}

	game.Cleared++
	grid[x][y].Cleared = true
	game.Steps = append(game.Steps, [2]int{x, y})

	// trigger mine, game over
	if grid[x][y].Value == GridCellValueMine {
		game.Status = MinesweeperStatusGameOver
		game.EndTime = time.Now()
	}

	if game.Cleared == game.Minefield.WinThreshold {
		game.Status = MinesweeperStatusWin
		game.EndTime = time.Now()
	}
}

func (game *MinesweeperGame) ClearAdjacent(x, y int) {
	if game.Minefield.CountAdjacentFlags(x, y) == 0 {
		return
	}
	grid := game.Minefield.GetGrid()

loop:
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i >= 0 && i < game.Minefield.Rows && j >= 0 && j < game.Minefield.Cols && !grid[i][j].Cleared && !grid[i][j].Flagging {
				game.ClearMine(i, j)
				if game.Status == MinesweeperStatusGameOver {
					break loop
				}
			}
		}
	}

}
func (game *MinesweeperGame) Flagging(x, y int) {
	if game.Minefield.Rows < x || game.Minefield.Cols < y {
		return
	}

	if game.Status != MinesweeperStatusLive {
		return
	}
	grid := game.Minefield.GetGrid()

	if grid[x][y].Cleared {
		return
	}

	grid[x][y].Flagging = !grid[x][y].Flagging
}

func (game *MinesweeperGame) ScanMinefield() (Grid, error) {
	grid := game.Minefield.GetGrid()
	res := make(Grid, game.Minefield.Rows)

	for i := 0; i < int(game.Minefield.Rows); i++ {
		res[i] = make([]GridCell, game.Minefield.Cols)
		for j := 0; j < int(game.Minefield.Cols); j++ {
			res[i][j] = grid[i][j]
			if !res[i][j].Cleared && game.Status == MinesweeperStatusLive {
				res[i][j].Value = GridCellValuePlaceholder
			}
		}
	}

	return res, nil
}
