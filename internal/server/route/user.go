package route

import (
	"fmt"
	"net/http"

	"github.com/aurelius15/go-skeleton/internal/entity"
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
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func addUser(c *gin.Context) {
	u, err := repository.UserRepository().SaveUser(c, &entity.User{
		FirstName: "Test",
		LastName:  "Test LastName",
		Address:   nil,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}
