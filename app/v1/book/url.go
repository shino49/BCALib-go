package book

import (
	"github.com/gin-gonic/gin"
)

func AppUrlSet(rg *gin.RouterGroup) {
	rg.GET("/book/:id", getBookList)
	rg.GET("/book/:id/info", getBookInfo)
	rg.GET("/book/:id/chapter", getBookChapterList)
	rg.GET("/book/:id/allcontent", getBookContent)
}
