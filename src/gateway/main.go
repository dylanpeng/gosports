package main

import (
	"flag"
	"fmt"
	"gosports/gateway/config"
	"os"
	"os/signal"
	"syscall"
)

var configFilePath = flag.String("c", "gateway.toml", "config file path")

func main() {
	//parse flag
	flag.Parse()

	err := config.Init(*configFilePath)
	if err != nil{
		fmt.Printf("Init config failed! err: %s \n", err)
	}

	conf := config.GetConfig()
	fmt.Printf("config: %+v \n", conf)

	waitForSignal()
}

func waitForSignal(){
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
