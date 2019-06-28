package main

import (
	"fmt"
	"gosports/lib/breakoff"
	localTicker "gosports/lib/ticker"
	"time"
)

func main() {
	ticker := localTicker.NewTicker(time.Second, doWrite)
	ticker.Start()

	time.Sleep(time.Second * 5)
	ticker.Stop()

	fmt.Printf("out stop")


	//ticker := time.NewTicker(time.Second * 5)
	//
	//go func(){
	//	for _ = range ticker.C{
	//		fmt.Printf("ticked at %v \n", time.Now())
	//	}
	//}()

	breakoff.Breaking()
}

func doWrite() {
	fmt.Printf("ticker: %v \n", time.Now())
}
