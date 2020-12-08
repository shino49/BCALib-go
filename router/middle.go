package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ytz4178/BCALib-go/app/util"
)

func errHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *util.Error
				if e, ok := err.(*util.Error); ok {
					Err = e
				} else if e, ok := err.(error); gin.Mode() == gin.DebugMode && ok {
					Err = util.OtherError(e.Error())
				} else {
					Err = util.ServerError
				}
				log.Println(Err.Code, Err.Msg)
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}

func handleNotFound(c *gin.Context) {
	err := util.NotFound
	c.JSON(err.StatusCode, err)
	return
}
