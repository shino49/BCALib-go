package book

import (
	"github.com/gin-gonic/gin"
)

func AppUrlSet(rg *gin.RouterGroup) {
	rg.GET("/book/list", getBookList)
}
