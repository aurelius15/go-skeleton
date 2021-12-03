package server

import (
	"fmt"

	"github.com/aurelius15/go-skeleton/internal/log"
	"github.com/aurelius15/go-skeleton/internal/server/middleware"
	"github.com/aurelius15/go-skeleton/internal/server/route"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	port   string
	Router *gin.Engine
}

func CreateServer(port string) *Server {
	router := gin.New()

	router.Use(middleware.AccessLogging)

	for _, r := range route.Routes {
		router.Handle(r.Method, r.Path, r.Handle)
	}

	return &Server{
		port:   getAddr(port),
		Router: router,
	}
}

func (s *Server) Up() {
	log.Default().Fatal("Error during running server", zap.Error(
		s.Router.Run(s.port),
	))
}

func getAddr(port string) string {
	return fmt.Sprintf(":%s", port)
}
