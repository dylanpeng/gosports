package main

import (
	"flag"
	"fmt"
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
	}

	conf := config.GetConfig()
	fmt.Printf("config: %+v \n", conf)

	//break
	breakoff.Breaking()
}
