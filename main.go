package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raidancampbell/browser-metrics/conf"
	"github.com/raidancampbell/browser-metrics/data"
	"github.com/raidancampbell/browser-metrics/handlers"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
)

var db = Database{}

type Database struct {
	DB *gorm.DB
}

var cfg *conf.Conf

func init() {
	cfg = conf.Initialize()
	logrus.Info("initialization complete")
	var err error
	logrus.Info("connecting to database...")
	db.DB, err = gorm.Open("sqlite3", cfg.DatasourceLocation)
	if err != nil {
		panic(err)
	}

	db.DB.AutoMigrate(&data.HistoryTable{})
	db.DB.AutoMigrate(&data.CommentTable{})
	logrus.Info("database connection complete")
}

func main() {
	r := gin.Default()
	r.Handle(http.MethodGet, "/debug/pprof/*method", gin.WrapF(http.DefaultServeMux.ServeHTTP))
	r.POST(fmt.Sprintf("/api/v1/visit/*%s", handlers.URLParameterHolder), handlers.GormWrapper(db.DB, handlers.HandleURL))

	logrus.Errorf("HTTP Server stopped with reason '%w'", r.Run(fmt.Sprintf("0.0.0.0:%d", cfg.ListenPort)))
}
