package route

import (
	"github.com/gin-gonic/gin"
)

var Routes []Route

type Route struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}
