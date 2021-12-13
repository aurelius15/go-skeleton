package cmd

import (
	"testing"

	"github.com/aurelius15/go-skeleton/internal/config"
	"github.com/aurelius15/go-skeleton/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (s *TestSuite) TestServerCmd_Init() {
	cmd := CommandCollection[config.ServerCmd]

	assert.NotNil(s.T(), cmd)
}

func (s *TestSuite) TestServerCmd_BindConfig() {
	cmd := CommandCollection[config.ServerCmd]
	cnf := new(mock.MockedConfig)

	cmd.BindConfig(cnf)
}

func (s *TestSuite) TestServerCmd_Execute() {
	cnf := new(mock.MockedConfig)
	webEngine := new(mock.MockedWebEngine)

	cmd := ServerCmd{
		webEngine: webEngine,
		config:    cnf,
	}

	cnf.On("Config", "Port").Return(":9000", nil)
	webEngine.On("Run", []string{":9000"}).Return(nil)

	cmd.Execute()

	cnf.AssertExpectations(s.T())
	webEngine.AssertExpectations(s.T())
}

func (s *TestSuite) TestServerCmd_Execute_Without_Configs() {
	cmd := ServerCmd{}

	assert.PanicsWithValuef(s.T(), "before Execute trigger BindConfig", func() {
		cmd.Execute()
	}, "it should throw panic with message")
}

func TestServerCmd(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
