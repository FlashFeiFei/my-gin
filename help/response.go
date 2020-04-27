package help

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//400
func Gin400NotFoundResponse(c *gin.Context, errmsg error, data interface{}) {
	h := gin.H{"errmsg": errmsg.Error()}
	if data != nil {
		h["data"] = data
	}
	c.JSON(http.StatusBadRequest, h)
}

//200返回.请求正常
func Gin200SuccessResponse(c *gin.Context, msg string, data interface{}) {
	Gin200Response(c, 0, msg, data)
}

//200返回
func Gin200Response(c *gin.Context, code int, msg string, data interface{}) {
	h := gin.H{"code": code, "msg": msg}
	if data != nil {
		h["data"] = data
	}
	c.JSON(http.StatusOK, h)
}

//200返回,业务异常,错误的返回
func Gin200ErrorResponse(c *gin.Context, code int, msg string, data interface{}) {
	Gin200Response(c, code, msg, data)
}

//500服务器错误
func Gin500ErrorResponse(c *gin.Context, errmsg error, data interface{}) {
	h := gin.H{"msg": errmsg.Error()}
	if data != nil {
		h["data"] = data
	}
	c.JSON(http.StatusInternalServerError, h)
}
