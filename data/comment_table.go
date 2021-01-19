package data

import "github.com/jinzhu/gorm"

type CommentTable struct {
	gorm.Model
	Comment string `gorm:"unique_index"`
	URL     string
}

func (CommentTable) TableName() string {
	return "comments"
}
