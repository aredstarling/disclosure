package disclosure

// Trace a specific name
func (a *Application) Trace(name string, main func(tracePoint *TracePoint) error) error {
	transaction := a.newrelic.StartTransaction(name, nil, nil)
	defer a.End(transaction)

	tracePoint := a.createTracePoint(transaction)

	return main(tracePoint)
}
