package data

import "github.com/jinzhu/gorm"

type CommentTable struct {
	gorm.Model
	Comment string `gorm:"unique_index:unique_comments"`
	URL     string `gorm:"unique_index:unique_comments"`
}

func (CommentTable) TableName() string {
	return "comments"
}
