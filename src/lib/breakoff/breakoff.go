package breakoff

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Breaking() {
	// waitting for exit signal
	exit := make(chan os.Signal)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	fmt.Printf("stop by exit signal '%s'", sign)

}
