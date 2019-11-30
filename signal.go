package disclosure

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sellernomics/golog"
)

// SignalTermination handles SIGTERM
func (a *Application) SignalTermination(main func()) chan os.Signal {
	notify := make(chan os.Signal)
	signal.Notify(notify, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-notify
		signal.Stop(notify)
		main()
		a.Logger.Warn("Shutting down!", golog.Attributes{"signal": "SIGTERM"})
		os.Exit(0)
	}()

	return notify
}
