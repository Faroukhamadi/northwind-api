package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	router *gin.Engine
}

// func (s *server) run() {
// }

func main() {
	// var s *server
	s := server{router: gin.Default()}
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "cool",
		})
	})

	s.router.Run()
}

func run() error {
	return fmt.Errorf("yikes")
}
