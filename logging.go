package disclosure

import (
	newrelic "github.com/newrelic/go-agent"
	"gitlab.com/lyticaa-public/golog"
)

type shim struct {
	logger golog.Logger
}

func (s *shim) Error(message string, c map[string]interface{}) {
	s.logger.Warn(message, c)
}

func (s *shim) Warn(message string, c map[string]interface{}) {
	s.logger.Warn(message, c)
}

func (s *shim) Info(message string, c map[string]interface{}) {
	s.logger.Info(message, c)
}

func (s *shim) Debug(message string, c map[string]interface{}) {
	s.logger.Debug(message, c)
}

func (s *shim) DebugEnabled() bool {
	return false
}

func createLogger(logger golog.Logger) newrelic.Logger {
	return &shim{logger: logger}
}
