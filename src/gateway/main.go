package main

import (
	"flag"
	"fmt"
	"gosports/common"
	"gosports/gateway/config"
	"gosports/lib/breakoff"
	"runtime"
)

var configFilePath = flag.String("c", "gateway.toml", "config file path")

func main() {
	//parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	//init config
	err := config.Init(*configFilePath)
	if err != nil{
		fmt.Printf("Init config failed! err: %s \n", err)
		return
	}

	//init logger
	if err = common.InitLogger(config.GetLogConfig()); err != nil{
		fmt.Printf("Init logger failed! err: %s \n", err)
		return
	}

	//break
	breakoff.Breaking()
}
