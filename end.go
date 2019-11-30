package disclosure

// Ender is the interface that wraps the basic End method.
//
// The behavior of End after the first call is undefined.
// Specific implementations may document their own behavior.
type Ender interface {
	End() error
}

// End by checking the error.
func (a *Application) End(e Ender) {
	err := e.End()
	if err != nil {
		a.Logger.Error("Could not end", err)
	}
}
