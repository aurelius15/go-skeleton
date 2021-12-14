package server

import (
	"net/http"

	"github.com/aurelius15/go-skeleton/internal/server/middleware"
	"github.com/aurelius15/go-skeleton/internal/server/route"
	"github.com/gin-gonic/gin"
)

type WebEngine interface {
	Run(adr ...string) error
	http.Handler
}

func NewServer() WebEngine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	for _, m := range middleware.Middlewares {
		router.Use(m.Handle)
	}

	for _, r := range route.Routes {
		router.Handle(r.Method, r.Path, r.Handle)
	}

	return router
}
