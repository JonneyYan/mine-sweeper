package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"infinitemining.com/storage"
)

const API_PORT = 3030

func Run() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	s := Service{}
	gin.SetMode("debug")

	s.Router = gin.Default()
	s.Router.Use(CORSMiddleware())

	s.Storage = storage.NewMemStorage(5)

	v1 := s.Router.Group("/api/v1")

	v1.GET("/load", s.LoadGame)
	v1.POST("/start", s.StartNewGame)
	v1.GET("/history", s.GetGameHistory)
	v1.GET("/conf", s.GetConfig)

	game := v1.Group("/game")
	game.POST("/clear", s.ClearMine)
	game.POST("/clearAdjacent", s.ClearAdjacent)
	game.POST("/flagging", s.ChangeFlag)
	game.GET("/scanMinefield", s.ScanMinefield)

	addr := fmt.Sprintf("localhost:%d", API_PORT)

	log.Info().Int("port", API_PORT).Msg("Listening to port")
	log.Fatal().Err(http.ListenAndServe(addr, s.Router)).Msg("API server failed to start")
}

// This enables us interact with the React Frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
