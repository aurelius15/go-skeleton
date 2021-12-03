package server

import (
	"fmt"

	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/server/middleware"
	"github.com/aurelius15/go-skeleton/internal/server/route"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Up(port string) {
	router := gin.New()

	router.Use(middleware.AccessLogging)

	for _, r := range route.Routes {
		router.Handle(r.Method, r.Path, r.Handle)
	}

	log.Default().Fatal("Error during running server", zap.Error(
		router.Run(getAddr(port)),
	))
}

func getAddr(port string) string {
	return fmt.Sprintf(":%s", port)
}
