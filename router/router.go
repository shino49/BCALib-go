package router

import (
	"github.com/ytz4178/BCALib-go/app/v1/book"

	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	r.NoMethod(handleNotFound)
	r.NoRoute(handleNotFound)
	r.Use(errHandler())

	r.Static("/static", "./static")
	r.Static("/html", "./html")
	v1 := r.Group("/v1")
	{
		book.AppUrlSet(v1)
	}

	return r
}
