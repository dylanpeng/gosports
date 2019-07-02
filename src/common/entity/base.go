package entity

type IEntity interface {
	PrimaryPairs() []interface{}
	IsSetPrimary() bool
	String() string
}
