package works

import (
	"fmt"
	"time"
)

type TeamWork struct {
}

func (m *TeamWork) DoWork() {
	fmt.Printf("TeamWork do time: %s \n", time.Now())
}
