package works

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gosports/common/consts"
	"gosports/common/entity"
	"gosports/common/model"
	"gosports/feed/config"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MatchWork struct {
	*BaseWork
}

func (m *MatchWork) DoWork() {
	now, matchUrl, dayRange := time.Now(), config.GetWorkConfig().MatchUrl, config.GetWorkConfig().MatchDayRange

	for i := dayRange * -1; i <= dayRange; i++ {
		requestDate := now.AddDate(0, 0, i)

		if i < 0 {
			m.Client.Url = fmt.Sprintf(matchUrl, consts.PathMatchFinished, requestDate.Format("20060102"))
		} else {
			m.Client.Url = fmt.Sprintf(matchUrl, consts.PathMatchNotStart, requestDate.Format("20060102"))
		}

		resp, err := m.Request()
		if err != nil {
			fmt.Printf("MatchWork request failed: %s \n", err)
			continue
		}

		matches, err := m.GetMatches(resp, requestDate)
		for _, match := range matches {
			err = model.MatchModel.AddOrUpdate(match)

			if err != nil {
				fmt.Printf("Add match failed. match: %+v \n", err)
			}
		}

	}
}

func (m *MatchWork) GetMatches(body []byte, now time.Time) (matches []*entity.Match, err error) {
	matches = make([]*entity.Match, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))

	if err != nil {
		fmt.Printf("goquery failed. err: %s \n", err)
		return
	}

	_ = doc.Find("li.list-item").Each(func(i int, selection *goquery.Selection) {
		match := &entity.Match{}
		id, exists := selection.Attr("data-id")

		if exists {
			var innerErr error
			match.ID, innerErr = strconv.ParseInt(id, 10, 64)

			if innerErr != nil {
				fmt.Printf("convert id failed. err: %s \n", innerErr)
				return
			}
		}

		leagueNode := selection.Find("a.event-name").First()
		leagueUrl, exists := leagueNode.Attr("href")

		if exists {
			regStr := regexp.MustCompile(consts.RegexMatchLeagueId).FindStringSubmatch(leagueUrl)
			if len(regStr) > 1 {
				match.LeagueID, _ = strconv.ParseInt(regStr[1], 10, 64)
			}
		}

		match.LeagueName = leagueNode.Find("span").First().Text()

		match.Round, _ = strconv.Atoi(strings.TrimSpace(selection.Find("span.lab-round").First().Text()))

		timeStr := strings.Split(strings.TrimSpace(selection.Find("span.lab-time").First().Text()), ":")
		if len(timeStr) == 2 {
			nowHour, _ := strconv.Atoi(timeStr[0])
			nowMin, _ := strconv.Atoi(timeStr[1])
			match.MatchDate = time.Date(now.Year(), now.Month(), now.Day(), nowHour, nowMin, 0, 0, time.Local)
		}

		if strings.TrimSpace(selection.Find("span.lab-status").First().Text()) == "完场" {
			match.MatchStatus = consts.End
		}

		homeNode := selection.Find("span.lab-team-home").First().Find("a").First()
		homeHref, exists := homeNode.Attr("href")

		if exists {
			regStr := regexp.MustCompile(consts.RegexMatchTeamId).FindStringSubmatch(homeHref)
			match.HomeTeamID, _ = strconv.ParseInt(regStr[1], 10, 64)
		}

		match.HomeTeamName = homeNode.Text()

		scoreArr := strings.Split(strings.TrimSpace(selection.Find("span.score").First().Find("b").First().Text()), "-")
		if len(scoreArr) > 1 {
			match.HomeScore, _ = strconv.Atoi(scoreArr[0])
			match.AwayScore, _ = strconv.Atoi(scoreArr[1])
		}

		awayNode := selection.Find("span.lab-team-away").First().Find("a").First()
		awayHref, exists := awayNode.Attr("href")

		if exists {
			regStr := regexp.MustCompile(consts.RegexMatchTeamId).FindStringSubmatch(awayHref)
			match.AwayTeamID, _ = strconv.ParseInt(regStr[1], 10, 64)
		}

		match.AwayTeamName = awayNode.Text()

		halfScoreArr := strings.Split(strings.TrimSpace(strings.TrimSpace(selection.Find("span.lab-half").First().Text())), "-")
		if len(halfScoreArr) > 1 {
			match.HalfTimeHomeScore, _ = strconv.Atoi(halfScoreArr[0])
			match.HalfTimeAwayScore, _ = strconv.Atoi(halfScoreArr[1])
		}

		matchResult := strings.TrimSpace(selection.Find("span.lab-bet-odds").First().Find("span").First().Text())
		switch matchResult {
		case "胜":
			match.MatchResult = consts.MatchResultWin
		case "平":
			match.MatchResult = consts.MatchResultDraw
		case "负":
			match.MatchResult = consts.MatchResultLoss
		default:
			match.MatchResult = consts.MatchResultNormal
		}

		match.CreatedTime = time.Now()
		match.UpdatedTime = time.Now()
		matches = append(matches, match)
	})
	return
}
