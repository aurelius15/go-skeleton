package route

import (
	"fmt"
	"net/http"

	"github.com/aurelius15/go-skeleton/internal/repository"
	"github.com/gin-gonic/gin"
)

const APIPrefix = "/api/v1/users"

func init() {
	Routes = append(Routes,
		Route{
			Method: http.MethodGet,
			Path:   fmt.Sprintf("%s/:userId", APIPrefix),
			Handle: getUser,
		}, Route{
			Method: http.MethodPost,
			Path:   APIPrefix,
			Handle: addUser,
		})
}

func getUser(c *gin.Context) {
	u, err := repository.UserRepository().GetUserByID(c, c.Param("userId"))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, u)
}

func addUser(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!\n")
}
