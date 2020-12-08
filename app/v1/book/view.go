package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getBookList(c *gin.Context) {
	if c.Param("id") == "list" {
		order := c.DefaultQuery("order", "id")
		limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "50"))
		filter := c.DefaultQuery("filter", "none")

		if errLimit != nil {
			panic(errLimit)
			// return
		}
		if limit >= 200 {
			limit = 200
		}

		c.JSON(200, gin.H{
			"order":  order,
			"limit":  limit,
			"filter": filter,
		})
	}
}

func getBookInfo(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"id":           id,
		"title":        "test1",
		"author":       "admin",
		"word_num":     0,
		"cover_pic":    nil,
		"chapter_list": []string{"a", "b", "c", "d", "e"},
	})
}

func getBookChapterList(c *gin.Context) {
	// id := c.Param("id")
	c.JSON(200, []string{"a", "b", "c", "d", "e"})
}

func getBookContent(c *gin.Context) {
	// id := c.Param("id")
	c.Redirect(http.StatusMovedPermanently, "/static/book/黄金罗盘.txt")
}
