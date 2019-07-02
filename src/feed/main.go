package main

import (
	"flag"
	"fmt"
	"gosports/common"
	"gosports/common/entity"
	"gosports/common/model"
	"gosports/feed/config"
	"gosports/feed/works"
	"gosports/lib/breakoff"
	"gosports/lib/ticker"
	"time"
)

var configFilePath = flag.String("c", "feed.toml", "config file path")

func main() {
	flag.Parse()

	err := config.Init(*configFilePath)

	if err != nil{
		fmt.Printf("Init config failed! err: %s \n", err)
		return
	}

	//init DB
	if err = common.InitDB(config.GetDBConfig()); err != nil{
		fmt.Printf("Init DB failed! err: %s \n", err)
		return
	}

	match := &entity.Match{
		ID:1,
		LeagueID:2,
		LeagueName:"联赛",
		MatchDate:time.Now(),
		HomeTeamID:3,
		HomeTeamName:"主队",
		AwayTeamID:4,
		AwayTeamName:"客队",
		HalfTimeHomeScore:5,
		HalfTimeAwayScore:6,
		HomeScore:7,
		AwayScore:8,
		Round:9,
		CreatedTime:time.Now(),
		UpdatedTime:time.Now(),
		IsDelete:true,
	}

	err = model.MatchModel.Add(match)
	if err != nil{
		fmt.Printf("DB add match failed! err: %s \n", err)
		return
	}

	searchMatch := &entity.Match{
		ID:1,
	}

	_, err = model.MatchModel.Get(searchMatch)

	if err != nil{
		fmt.Printf("DB get match failed! err: %s \n", err)
		return
	}

	fmt.Printf("DB get match success! err: %+v \n", searchMatch)

	worker := &ticker.WorkSet{}
	worker.AddWork(time.Second*2, &works.MatchWork{})
	worker.AddWork(time.Second*3, &works.TeamWork{})
	worker.Start()

	time.Sleep(time.Second * 30)

	worker.Stop()

	breakoff.Breaking()
}
