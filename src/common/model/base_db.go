package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gosports/common"
	"gosports/common/entity"
	"strings"
)

type baseDBModel struct {
	writeInstance string
}

func createDBModel(writeInstance string) *baseDBModel{
	return &baseDBModel{writeInstance:writeInstance}
}

func (b *baseDBModel) getDB() (*gorm.DB, error){
	return common.GetDB(b.writeInstance)
}

func (b *baseDBModel) Add(e entity.IEntity) error{
	db, err := b.getDB()

	if err != nil{
		return err
	}

	return db.Create(e).Error
}

func (b *baseDBModel) Get(e entity.IEntity) (exist bool, err error){
	if !e.IsSetPrimary(){
		err = errors.New("primary attribute is empty")
		return
	}

	db, err := b.getDB()

	if err != nil{
		return false, err
	}

	err = db.First(e).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (b *baseDBModel) Update(e entity.IEntity, props map[string]interface{}) error{
	if !e.IsSetPrimary(){
		return errors.New("primary attribute is empty")
	}

	db, err := b.getDB()

	if err != nil{
		return err
	}

	if props == nil{
		return db.Save(e).Error
	} else{
		return db.Model(e).Update(props).Error
	}
}

func (b *baseDBModel) Remove(e entity.IEntity) error{
	if !e.IsSetPrimary(){
		return errors.New("primary attribute is empty")
	}

	db, err := b.getDB()

	if err != nil {
		return err
	}

	return db.Delete(e).Error
}

func (b *baseDBModel) getInsertStatement(fields []string) (list, statement string) {
	list = "(" + strings.Join(fields, ",") + ")"
	statement = "(" + strings.Repeat("?,", len(fields)-1) + "?)"
	return
}