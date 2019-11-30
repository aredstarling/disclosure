package disclosure

import (
	"os"
	"time"

	"gitlab.com/sellernomics/golog"
	raven "github.com/getsentry/raven-go"
	_ "github.com/heroku/x/hmetrics/onload" // Allows us to remove in one place if we don't need it.
	newrelic "github.com/newrelic/go-agent"
)

// Application defines a way to seperate what we want dislosed.
type Application struct {
	Logger   golog.Logger
	Name     string
	newrelic newrelic.Application
	sentry   *raven.Client
}

// CreateApplication to seperate what we want dislose
func CreateApplication(name string, logger golog.Logger) (*Application, error) {
	sentry, err := raven.New(os.Getenv("SENTRY_DSN"))
	if err != nil {
		return nil, err
	}

	license := os.Getenv("NEW_RELIC_LICENSE_KEY")
	config := newrelic.NewConfig(name, license)
	config.Enabled = len(license) > 0
	config.Logger = createLogger(logger)

	app, err := newrelic.NewApplication(config)
	if err != nil {
		return nil, err
	}

	if err := app.WaitForConnection(time.Second * 60); err != nil {
		return nil, err
	}

	return &Application{Name: name, Logger: logger, sentry: sentry, newrelic: app}, nil
}
