package gin_util

import (
	"gogs.sharkgulf.cn/sg/library/utils/qhttp"

	"github.com/gin-gonic/gin"
)

type httpContext struct {
	c *gin.Context
}

func (h *httpContext) ShouldBindJSON(obj interface{}) error {
	return h.c.ShouldBindJSON(obj)
}

func (h *httpContext) AbortWithStatusJSON(code int, jsonObj interface{}) {
	h.c.AbortWithStatusJSON(code, jsonObj)
}

func CheckJsonParam(c *gin.Context, params interface{}) bool {
	return checkJsonParam(httpCtx(c), params)
}

func SendOk(c *gin.Context, data interface{}) {
	sendOk(httpCtx(c), data)
}

func Send(c *gin.Context, code int, data interface{}, stateInfo, customInfo string) {
	send(httpCtx(c), code, data, stateInfo, customInfo)
}

func SendCode(c *gin.Context, code int) {
	send(httpCtx(c), code, nil, "", "")
}

func SendPureFailed(c *gin.Context) {
	sendPureFailed(httpCtx(c))
}

func SendFailedWithErr(c *gin.Context, err error) {
	sendFailedWithErr(httpCtx(c), err)
}

func SendPureOk(c *gin.Context) {
	sendPureOk(httpCtx(c))
}

func SendFile(c *gin.Context, data []byte, filename string) {
	qhttp.Download(data, c.Writer, c.Request, filename)
}
