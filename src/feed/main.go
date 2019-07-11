package main

import (
	"flag"
	"fmt"
	"gosports/common"
	"gosports/feed/config"
	"gosports/feed/util"
	"gosports/lib/breakoff"
	"runtime"
)

var configFilePath = flag.String("c", "feed.toml", "config file path")

func main() {
	//parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// init config
	var err error
	if err = config.Init(*configFilePath); err != nil {
		fmt.Printf("Init config failed! err: %s \n", err)
		return
	}

	// init logger
	if err = common.InitLogger(config.GetLogConfig()); err != nil {
		fmt.Printf("Init logger failed! err: %s \n", err)
		return
	}

	// init DB
	if err = common.InitDB(config.GetDBConfig()); err != nil {
		fmt.Printf("Init DB failed! err: %s \n", err)
		return
	}

	// init works
	util.InitWorks(config.GetWorkConfig())

	// break
	breakoff.Breaking()

	// stop works
	util.Works.Stop()
}
