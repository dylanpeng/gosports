package entity

import (
	"fmt"
	"time"
)

type Match struct {
	ID                int64     `gorm:"column:id;primary_key" json:"id"`
	LeagueID          int64     `gorm:"column:league_id" json:"league_id"`
	LeagueName        string    `gorm:"column:league_name" json:"league_name"`
	MatchDate         time.Time `gorm:"column:match_date" json:"match_date"`
	HomeTeamID        int64     `gorm:"column:home_team_id" json:"home_team_id"`
	HomeTeamName      string    `gorm:"column:home_team_name" json:"home_team_name"`
	AwayTeamID        int64    `gorm:"column:away_team_id" json:"away_team_id"`
	AwayTeamName      string    `gorm:"column:away_team_name" json:"away_team_name"`
	HalfTimeHomeScore int    `gorm:"column:half_time_home_score" json:"half_time_home_score"`
	HalfTimeAwayScore int    `gorm:"column:half_time_away_score" json:"half_time_away_score"`
	HomeScore         int    `gorm:"column:home_score" json:"home_score"`
	AwayScore         int    `gorm:"column:away_score" json:"away_score"`
	Round             int    `gorm:"column:round" json:"round"`
	CreatedTime       time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime       time.Time `gorm:"column:updated_time" json:"updated_time"`
	IsDelete          bool      `gorm:"column:is_delete" json:"is_delete"`
}

func (e *Match) TableName() string {
	return "t_match"
}

func (e *Match) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.ID}
}

func (e *Match) IsSetPrimary() bool {
	return e.ID > 0
}

func (e *Match) String() string {
	return fmt.Sprintf("%+v", *e)
}
