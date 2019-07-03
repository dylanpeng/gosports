package works

import (
	"fmt"
	"gosports/common/consts"
	"gosports/feed/config"
	"time"
)

type MatchWork struct {
	*BaseWork
}

func (m *MatchWork) DoWork() {
	now, matchUrl, dayRange := time.Now(), config.GetWorkConfig().MatchUrl, config.GetWorkConfig().MatchDayRange

	for i := dayRange * -1; i <= dayRange; i++{
		if i < 0{
			m.Client.Url = fmt.Sprintf(matchUrl, consts.PathMatchFinished, now.AddDate(0, 0, i).Format("20060102"))
		} else {
			m.Client.Url = fmt.Sprintf(matchUrl, consts.PathMatchNotStart, now.AddDate(0, 0, i).Format("20060102"))
		}

		_, err := m.Request()
		if err != nil {
			fmt.Printf("MatchWork request failed: %s \n", err)
		} else {
			fmt.Printf("MatchWork request success: %s \n", m.Client.Url)
		}
	}
}
