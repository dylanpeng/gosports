package main

import (
	"fmt"
	"gosports/lib/breakoff"
	localticker "gosports/lib/ticker"
	"time"
)

func main(){
	ticker := localticker.NewTicker(time.Second * 5, doWrite)
	ticker.Start()

	//ticker := time.NewTicker(time.Second * 5)
	//
	//go func(){
	//	for _ = range ticker.C{
	//		fmt.Printf("ticked at %v \n", time.Now())
	//	}
	//}()

	breakoff.Breaking()
}

func doWrite(){
	fmt.Printf("ticker: %v \n", time.Now())
}
