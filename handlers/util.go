package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// provides a closure to enable DB interactions
// given a DB and a function that needs (the DB plus whatever its original signature should be)
// this function creates a closure to return a function of the original signature, but with the DB in a closure.
func GormWrapper(db *gorm.DB, f func(db *gorm.DB, c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		f(db, c)
	}
}
