package router

import (
	"github.com/ytz4178/BCALib-go/app/v1/book"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		book.AppUrlSet(v1)
	}

	return r
}
