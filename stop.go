package disclosure

// Stoper is the interface that wraps the basic Stop method.
//
// The behavior of Stop after the first call is undefined.
// Specific implementations may document their own behavior.
type Stoper interface {
	Stop() error
}

// Stop by checking the error.
func (a *Application) Stop(s Stoper) {
	err := s.Stop()
	if err != nil {
		a.Logger.Error("Could not stop", err)
	}
}
