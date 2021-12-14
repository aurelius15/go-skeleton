package middleware

import "github.com/gin-gonic/gin"

var Middlewares = []Middleware{
	{
		Handle: Correlation,
	},
	{
		Handle: AccessLogging,
	},
}

type Middleware struct {
	Handle gin.HandlerFunc
}
