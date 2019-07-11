package model

import (
	"gosports/common/consts"
	"gosports/common/entity"
	"gosports/common/method"
)

type matchModel struct {
	*baseDBModel
}

var MatchModel = &matchModel{
	baseDBModel: createDBModel("main-master"),
}

func (m *matchModel) AddOrUpdate(match *entity.Match) (string, error) {
	tempMatch := &entity.Match{ID: match.ID}
	exists, err := m.Get(tempMatch)

	if err != nil {
		return "", err
	}

	if !exists {
		err = m.Add(match)
		return consts.MatchAdd, err
	}

	match.CreatedTime = tempMatch.CreatedTime
	tempMatch.UpdatedTime = match.UpdatedTime

	if method.IsTheSame(match, tempMatch) {
		return consts.MatchNoNeedUpdate, nil
	}

	err = m.Update(match, nil)
	return consts.MatchUpdate, err
}
