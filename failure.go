package disclosure

// FailOnError panics if there is an error
func (a *Application) FailOnError(err error) {
	if err != nil {
		panic(err)
	}
}
