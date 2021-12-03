package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api/v1/users"

func init() {
	Routes = append(Routes,
		Route{
			Method: http.MethodGet,
			Path:   apiPrefix,
			Handle: getUsers,
		}, Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/:userId", apiPrefix),
			Handle: getUser,
		}, Route{
			Method: http.MethodPost,
			Path:   apiPrefix,
			Handle: addUser,
		})
}

func getUsers(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}

func getUser(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}

func addUser(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}
