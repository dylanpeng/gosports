package main

import (
	"fmt"
	"gosports/gateway/config"
)

func main() {
	err := config.Init("/Users/chenpeng/working/gosports/src/conf/gateway.toml")
	if err != nil{
		fmt.Printf("Init config failed! err: %s", err)
	}

	conf := config.GetConfig()
	fmt.Printf("config: %+v", conf)
}
