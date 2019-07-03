package main

import (
	"flag"
	"fmt"
	"gosports/common"
	"gosports/feed/config"
	"gosports/feed/util"
	"gosports/lib/breakoff"
)

var configFilePath = flag.String("c", "feed.toml", "config file path")

func main() {
	flag.Parse()

	err := config.Init(*configFilePath)

	if err != nil {
		fmt.Printf("Init config failed! err: %s \n", err)
		return
	}

	//init DB
	if err = common.InitDB(config.GetDBConfig()); err != nil {
		fmt.Printf("Init DB failed! err: %s \n", err)
		return
	}

	util.InitWorks(config.GetWorkConfig())

	breakoff.Breaking()

	util.Works.Stop()
}
