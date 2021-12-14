package route

import (
	"fmt"
	"net/http"

	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/gin-gonic/gin"
)

const Internal = "/internal/"

type statusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func init() {
	Routes = append(Routes,
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/health", Internal),
			Handle: health,
		},
	)
}

func health(c *gin.Context) {
	_, err := storage.Instance().Ping(c).Result()
	if err == nil {
		c.String(http.StatusOK, "ok")
		return
	}

	c.JSON(http.StatusServiceUnavailable, statusMessage{
		Status:  "error",
		Message: "redis instance is not reachable",
	})
}
