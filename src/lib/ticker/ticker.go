package ticker

import (
	"fmt"
	"time"
)

type IWork interface {
	DoWork()
}

type Ticker struct {
	Interval    time.Duration
	Work        IWork
	ticker      *time.Ticker
	stopChannel chan bool
}

func NewTicker(d time.Duration, f IWork) *Ticker {
	return &Ticker{Interval: d, Work: f, ticker: time.NewTicker(d), stopChannel: make(chan bool)}
}

func (t *Ticker) Start() {
	go func() {
		t.Work.DoWork()

	ForBegin:
		for {
			select {
			case <-t.ticker.C:
				t.Work.DoWork()
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
