package disclosure

import (
	"github.com/newrelic/go-agent"
)

// TracePoint is used for finer control of tracing
type TracePoint struct {
	application *Application
	transaction newrelic.Transaction
}

func (a *Application) createTracePoint(transaction newrelic.Transaction) *TracePoint {
	return &TracePoint{transaction: transaction, application: a}
}

// Trace a specific name
func (t *TracePoint) Trace(name string, main func()) {
	segment := newrelic.StartSegment(t.transaction, name)
	defer t.application.End(segment)

	main()
}

// AddMetrics to the trace point.
func (t *TracePoint) AddMetrics(metrics map[string]interface{}) {
	for key, value := range metrics {
		if err := t.transaction.AddAttribute(key, value); err != nil {
			t.application.Logger.Error("Could not add attribute", err)
		}
	}
}
