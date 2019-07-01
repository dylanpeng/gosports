package ticker

import "time"

type WorkSet struct {
	works []*Ticker
}

func (b *WorkSet) Start() {
	if b.works == nil || len(b.works) == 0 {
		return
	}

	for _, work := range b.works {
		work.Start()
	}
}

func (b *WorkSet) Stop() {
	if b.works == nil || len(b.works) == 0 {
		return
	}

	for _, work := range b.works {
		work.Stop()
	}
}

func (b *WorkSet) AddWork(d time.Duration, work IWork) {
	ti := NewTicker(d, work)
	b.works = append(b.works, ti)
}
