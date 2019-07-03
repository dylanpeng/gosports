package works

import (
	"fmt"
	"time"
)

type TeamWork struct {
	*BaseWork
}

func (m *TeamWork) DoWork() {
	_, err := m.Request()

	if err != nil {
		fmt.Printf("TeamWork request failed: %s \n", err)
	} else {
		fmt.Printf("TeamWork request success: %s \n", time.Now())
	}
}
