package cmd

import (
	"context"
	"fmt"

	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/internal/repository"
)

type NumberOfUsersCmd struct {
	config config.Configure
}

func (c *NumberOfUsersCmd) BindConfig(i config.Configure) {
	c.config = i
}

func (c *NumberOfUsersCmd) Execute() {
	n, err := repository.UserRepository().NumberOfUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Number of records: %d \n", n)
}

func init() {
	CommandCollection[config.NumberOfUsersCmd] = &NumberOfUsersCmd{}
}
