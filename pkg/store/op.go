package store

import (
	"gorm.io/gorm"
)

func InsertItem(db *gorm.DB, tableName string, item interface{}) error {
	return db.Table(tableName).Create(item).Error
}
