package works

import (
	"fmt"
	"time"
)

type MatchWork struct {
	*BaseWork
}

func (m *MatchWork) DoWork() {
	_, err := m.Request()

	if err != nil {
		fmt.Printf("MatchWork request failed: %s \n", err)
	} else {
		fmt.Printf("MatchWork request success: %s \n", time.Now())
	}
}
