package disclosure

import "io"

// Close by checking the error.
func (a *Application) Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		a.Logger.Error("Could not close", err)
	}
}
