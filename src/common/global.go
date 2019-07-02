package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	lGorm "gosports/lib/gorm"
)

var dbPool *lGorm.Pool

func InitDB(configs map[string]*lGorm.Config) (err error) {
	dbPool = lGorm.NewPool()

	for k, v := range configs {
		if err := dbPool.Add(k, v); err != nil {
			return err
		}
	}

	return nil
}

func GetDB(name string) (*gorm.DB, error) {
	return dbPool.Get(name)
}
