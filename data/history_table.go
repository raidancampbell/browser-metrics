package data

import "github.com/jinzhu/gorm"

type HistoryTable struct {
	gorm.Model
	URL string
	Title string
}

func (HistoryTable) TableName() string {
	return "history"
}
