package model

type matchModel struct{
	*baseDBModel
}

var MatchModel = &matchModel{
	baseDBModel: createDBModel("main-master"),
}