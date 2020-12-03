package book

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBookList(c *gin.Context) {
	order := c.DefaultQuery("order", "id")
	limit, errLimit := strconv.Atoi(c.DefaultQuery("limit", "50"))
	filter := c.DefaultQuery("filter", "none")

	if errLimit != nil {
		panic("err limit")
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
