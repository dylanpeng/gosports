package ticker

import (
	"fmt"
	"time"
)

type Ticker struct {
	Interval    time.Duration
	Do          func()
	ticker      *time.Ticker
	stopChannel chan bool
}

func NewTicker(d time.Duration, f func()) *Ticker {
	return &Ticker{Interval: d, Do: f, ticker: time.NewTicker(d), stopChannel: make(chan bool)}
}

func (t *Ticker) Start() {
	go func() {
	ForBegin:
		for {
			select {
			case <-t.ticker.C:
				t.Do()
			case <-t.stopChannel:
				break ForBegin
			}
		}

		fmt.Printf("exit ticker range \n")
	}()
}

func (t *Ticker) Stop() {
	t.stopChannel <- true
	t.ticker.Stop()

	fmt.Printf("stop ticker \n")
}
