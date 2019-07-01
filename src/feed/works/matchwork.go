package works

import (
	"fmt"
	"time"
)

type MatchWork struct {
}

func (m *MatchWork) DoWork() {
	fmt.Printf("MatchWork do time: %s \n", time.Now())
}
