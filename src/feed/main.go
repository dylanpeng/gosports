package main

import (
	"gosports/feed/works"
	"gosports/lib/breakoff"
	"gosports/lib/ticker"
	"time"
)

func main() {
	worker := &ticker.WorkSet{}
	worker.AddWork(time.Second*2, &works.MatchWork{})
	worker.AddWork(time.Second*3, &works.TeamWork{})
	worker.Start()

	time.Sleep(time.Second * 30)

	worker.Stop()

	breakoff.Breaking()
}
