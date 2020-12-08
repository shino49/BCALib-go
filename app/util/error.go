package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Success     = NewError(http.StatusOK, 0, "success")
	ServerError = NewError(http.StatusInternalServerError, 1000, "系统异常，请稍后重试!")
	NotFound    = NewError(http.StatusNotFound, 100, http.StatusText(http.StatusNotFound))
)

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

func (e *Error) Error() string {
	return e.Msg
}
func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 10000, message)
}

func NewError(statusCode, code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
	}
}

func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}
