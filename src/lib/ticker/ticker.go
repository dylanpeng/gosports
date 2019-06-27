package ticker

import "time"

type Ticker struct{
	Interval time.Duration
	Function func()
}

func NewTicker(d time.Duration, f func()) *Ticker{
	return &Ticker{Interval:d, Function:f}
}

func (t *Ticker) Start(){
	ticker := time.NewTicker(t.Interval)
	go func(){
		for _ = range ticker.C{
			t.Function()
		}
	}()
}
