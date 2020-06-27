package data

import "github.com/jinzhu/gorm"

type HistoryTable struct {
	gorm.Model
	URL string
}

func (HistoryTable) TableName() string {
	return "history"
}
