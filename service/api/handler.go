package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	ms "infinitemining.com/minesweeper"
)

type StartGameReq struct {
	Difficulty      int    `json:"difficulty"`
	Username        string `json:"username" binding:"required"`
	InitialPosition [2]int `json:"initialPosition"`
}

func (s *Service) LoadGame(c *gin.Context) {
	game, err := s.Storage.LoadLatest()

	if err != nil {
		Error(c, "please create new game")
		return
	}
	gamePtr, ok := (game).(ms.MinesweeperGame)
	if !ok {
		fmt.Println("Conversion failed")
		Error(c, "Conversion failed")
		return
	}
	s.CurrentGame = &gamePtr
	ReturnGameStatus(c, s.CurrentGame)
}

func (s *Service) StartNewGame(c *gin.Context) {
	params := StartGameReq{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Error().Err(err).Send()
		Error(c, "参数错误")
		return
	}

	game, _ := ms.NewGame(
		params.Username,
		params.InitialPosition[0],
		params.InitialPosition[1],
		params.Difficulty,
	)

	game.ClearMine(params.InitialPosition[0], params.InitialPosition[1])

	s.CurrentGame = game
	ReturnGameStatus(c, game)
}
func (s *Service) GetConfig(c *gin.Context) {
	Ok(c, ms.DefaultGameConfig)
}
func (s *Service) GetGameHistory(c *gin.Context) {
	games := s.Storage.LoadAll()
	// var result []ms.MinesweeperGame

	// for _, g := range games {
	// 	gt, _ := g.(ms.MinesweeperGame)
	// 	if gt.Status != ms.MinesweeperStatusLive {
	// 		result = append(result, gt)
	// 	}
	// }
	Ok(c, games)
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (s *Service) ClearMine(c *gin.Context) {
	params := Point{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Error().Err(err).Send()
		Error(c, "参数错误")
		return
	}

	game := s.CurrentGame
	game.ClearMine(params.X, params.Y)

	if game.Status == ms.MinesweeperStatusGameOver {
		s.Storage.Save(game)
	}

	ReturnGameStatus(c, game)
}

func (s *Service) ClearAdjacent(c *gin.Context) {
	params := Point{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		Error(c, "参数错误")
		return
	}

	game := s.CurrentGame
	game.ClearAdjacent(params.X, params.Y)

	if game.Status == ms.MinesweeperStatusGameOver {
		s.Storage.Save(game)
	}

	ReturnGameStatus(c, game)
}

func (s *Service) ChangeFlag(c *gin.Context) {
	params := Point{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		Error(c, "参数错误")
		return
	}

	game := s.CurrentGame
	game.Flagging(params.X, params.Y)

	ReturnGameStatus(c, game)
}

func (s *Service) ScanMinefield(c *gin.Context) {
	game := s.CurrentGame

	ReturnGameStatus(c, game)
}
