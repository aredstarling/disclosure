package disclosure

import (
	"errors"
	"fmt"

	"gitlab.com/lyticaa-public/golog"
)

// Monitor the application
func (a *Application) Monitor(main func() error) {
	if err := a.monitor(main); err != nil {
		a.Notify(a.Name, err)
		a.Logger.Fatal("Received an error", golog.Attributes{"error": err})
	}
}

func (a *Application) monitor(main func() error) (err error) {
	defer func() {
		switch recoveredError := recover().(type) {
		case nil:
			return
		case error:
			err = recoveredError
		default:
			err = errors.New(fmt.Sprint(recoveredError))
		}
	}()

	return main()
}
