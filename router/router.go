package router

import (
	"github.com/ytz4178/BCALib-go/app/book"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/book/list", book.GetBookList)
	}

	return r
}
