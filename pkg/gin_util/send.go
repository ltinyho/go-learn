package gin_util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpContext interface {
	ShouldBindJSON(obj interface{}) error
	AbortWithStatusJSON(code int, jsonObj interface{})
}

func httpCtx(c *gin.Context) HttpContext {
	return &httpContext{c: c}
}

func checkJsonParam(c HttpContext, params interface{}) bool {
	err := c.ShouldBindJSON(params)
	if err != nil {
		ret := gin.H{
			"code": 02,
			"data": "",
			"msg":  "无效参数",
			"info": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusOK, ret)
		return false
	}
	return true
}

func sendOk(c HttpContext, data interface{}) {
	send(c, 0, data, "", "")
}

func send(c HttpContext, code int, data interface{}, stateInfo, CustomInfo string) {
	info := stateInfo
	ret := gin.H{
		"code": code,
		"data": data,
		"msg":  info,
		"info": CustomInfo,
	}
	c.AbortWithStatusJSON(http.StatusOK, ret)
}

func sendPureFailed(c HttpContext) {
	send(c, 01, nil, "", "")
}

func sendFailedWithErr(c HttpContext, err error) {
	send(c, 01, err.Error(), "", "")
}

func sendPureOk(c HttpContext) {
	sendOk(c, nil)
}
