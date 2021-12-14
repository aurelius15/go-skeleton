package route

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/aurelius15/go-skeleton/internal/storage"
	"github.com/gin-gonic/gin"
)

const Internal = "/internal/"

const DebugPrefix = "/internal/debug/pprof"

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

		// pprof endpoints
		Route{
			Method: http.MethodGet,
			Path:   DebugPrefix,
			Handle: gin.WrapF(pprof.Index),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/cmdline", DebugPrefix),
			Handle: gin.WrapF(pprof.Cmdline),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/profile", DebugPrefix),
			Handle: gin.WrapF(pprof.Profile),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/symbol", DebugPrefix),
			Handle: gin.WrapF(pprof.Symbol),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/trace", DebugPrefix),
			Handle: gin.WrapF(pprof.Trace),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/allocs", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("allocs")),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/block", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("block")),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/goroutine", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("goroutine")),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/heap", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("heap")),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/mutex", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("mutex")),
		},
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/threadcreate", DebugPrefix),
			Handle: gin.WrapH(pprof.Handler("threadcreate")),
		},
		// pprof endpoints END
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
