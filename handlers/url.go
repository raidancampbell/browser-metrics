package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/raidancampbell/browser-metrics/data"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

const URLParameterHolder = "url"

func HandleURL(db *gorm.DB, c *gin.Context) {

	db.Save(&data.HistoryTable{
		URL: strings.TrimPrefix(c.Request.RequestURI, "/api/v1/visit/"),
	})

	body, err := html.Parse(c.Request.Body)
	if err != nil {
		logrus.Errorf("failed to extract HTML from payload '%s'", err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, comment := range walkForComments(body) {
		logrus.Debugf("saving comment '%s'", comment.Data)
		db.Save(&data.CommentTable{
			Comment: comment.Data,
			URL:     strings.TrimPrefix(c.Param(URLParameterHolder), "/"),
		})
	}

}

func walkForComments(doc *html.Node) []*html.Node {
	if doc.Type == html.CommentNode {
		return []*html.Node{doc}
	}
	var acc []*html.Node
	for child := doc.FirstChild; child != nil; child = child.NextSibling {
		acc = append(acc, walkForComments(child)...)
	}
	return acc
}
