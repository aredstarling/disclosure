package disclosure

import (
	"github.com/newrelic/go-agent"
)

// Notify an error has occured
func (a *Application) Notify(name string, err error) {
	transaction := a.newrelic.StartTransaction(name, nil, nil)
	defer a.End(transaction)

	a.notify(transaction, err)
}

func (a *Application) notify(transaction newrelic.Transaction, err error) {
	if err := transaction.NoticeError(err); err != nil {
		a.Logger.Error("Could not notify newrelic", err)
	}

	a.sentry.CaptureErrorAndWait(err, nil)
}
