package util

import (
	"gosports/common/consts"
	"gosports/feed/config"
	"gosports/feed/works"
	"gosports/lib/ticker"
	"time"
)

var Works *ticker.WorkSet

func InitWorks(conf *config.WorkConfig) {
	Works = &ticker.WorkSet{}

	Works.AddWork(time.Second*time.Duration(conf.MatchInterval),
		&works.MatchWork{BaseWork: works.NewBaseWork(conf.MatchUrl, nil, consts.HttpGet, nil)})

	Works.AddWork(time.Second*time.Duration(conf.TeamInterval),
		&works.TeamWork{BaseWork: works.NewBaseWork(conf.TeamUrl, nil, consts.HttpGet, nil)})

	Works.Start()

}
