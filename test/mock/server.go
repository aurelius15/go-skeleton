package mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockedConfig struct {
	mock.Mock
}

func (m *MockedConfig) Config(fieldName string) (string, error) {
	args := m.Called(fieldName)
	return args.String(0), args.Error(1)
}

func (m *MockedConfig) Command() string {
	args := m.Called()
	return args.String(0)
}

type MockedWebEngine struct {
	mock.Mock
}

func (m *MockedWebEngine) Run(adr ...string) error {
	args := m.Called(adr)
	return args.Error(0)
}

func (m *MockedWebEngine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = m.Called(w, r)
}
