package config

import (
	"github.com/aurelius15/go-skeleton/internal/reflection"
)

const NumberOfUsersCmd = "number-users"

type NumberOfUsersConfig struct{}

func (c *NumberOfUsersConfig) Config(fieldName string) (s string) {
	s, _ = reflection.StringFieldByName(c, fieldName)
	return
}

func (c *NumberOfUsersConfig) Command() string {
	return NumberOfUsersCmd
}
