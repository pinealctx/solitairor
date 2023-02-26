package store

import (
	"github.com/pinealctx/solitairor/pkg/sol"
	"gorm.io/gorm"
)

func InsertItem(db *gorm.DB, tableName string, item interface{}) error {
	return db.Table(tableName).Create(item).Error
}

func GetItemByAverageStepAsc(db *gorm.DB, tableName string) ([]sol.Record, error) {
	var rs []sol.Record
	var err = db.Table(tableName).Order("average_step asc").Find(&rs).Error
	if err != nil {
		return nil, err
	}
	return rs, nil
}
