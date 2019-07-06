package model

import (
	"fmt"
	"gosports/common/entity"
	"gosports/common/method"
)

type matchModel struct {
	*baseDBModel
}

var MatchModel = &matchModel{
	baseDBModel: createDBModel("main-master"),
}

func (m *matchModel) AddOrUpdate(match *entity.Match) error {
	tempMatch := &entity.Match{ID: match.ID}
	exists, err := m.Get(tempMatch)

	if err != nil {
		return err
	}

	if !exists {
		err = m.Add(match)
		fmt.Printf("Add new match, %+v \n", match)
		return err
	}

	match.CreatedTime = tempMatch.CreatedTime
	tempMatch.UpdatedTime = match.UpdatedTime

	if method.IsTheSame(match, tempMatch) {
		fmt.Printf("nochange don't update, %+v \n", match)
		return nil
	}

	err = m.Update(match, nil)
	fmt.Printf("update match, %+v \n", match)
	return err
}
