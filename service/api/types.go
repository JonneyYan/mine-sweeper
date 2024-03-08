package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"infinitemining.com/minesweeper"
	"infinitemining.com/storage"
)

type Service struct {
	Router  *gin.Engine
	Storage storage.Storage

	CurrentGame *minesweeper.MinesweeperGame
}

type ApiBaseResponse struct {
	Status  int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiBaseResponse{
		Status:  0,
		Message: "success",
		Data:    data,
	})
}
func ReturnGameStatus(c *gin.Context, game *minesweeper.MinesweeperGame) {
	grid, _ := game.ScanMinefield()
	c.JSON(http.StatusOK, ApiBaseResponse{
		Status:  0,
		Message: "success",
		Data: minesweeper.MinesweeperGameAPI{
			MinesweeperGame: *game,
			Grid:            grid,
		},
	})
}
func Error(c *gin.Context, errMsg string) {
	log.Error().
		Str("url", c.Request.URL.Path).
		Msg(errMsg)
	c.JSON(http.StatusOK, ApiBaseResponse{
		Status:  3,
		Message: errMsg,
		Data:    nil,
	})
}
