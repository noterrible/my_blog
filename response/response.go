package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
	})
}
func OKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

//	func Failed(c *gin.Context, code int64, msg string) {
//		c.JSON(http.StatusOK, Response{
//			Code: code,
//			Msg:  msg,
//		})
//	}
//
//	func FailedWithCode(c *gin.Context, code int64) {
//		if _, ok := ErrMap[code]; !ok {
//			c.JSON(http.StatusOK, Response{
//				Code: code,
//				Msg:  "未知错误",
//			})
//		} else {
//			c.JSON(http.StatusOK, Response{
//				Code: code,
//				Msg:  ErrMap[code],
//			})
//		}
//	}
func FailedWithMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: -1,
		Msg:  msg,
	})
}
