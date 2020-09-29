package gin_util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const MiddleAppKey = "app"
const MiddleUidKey = "ruserid"    // uid
const MiddleAdminKey = "daccount" // 管理员用户名

type SingleId struct {
	Id int `json:"id" binding:"required"`
}
type PageParams struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func GetUidAndReqId(c *gin.Context) map[string]string {
	headers := map[string]string{
		MiddleUidKey: getuid(c),
		"requestid":  GetRequestId(c)}
	return headers
}

func GetRequestId(c *gin.Context) string {
	return c.Request.Header.Get("requestid")
}
func GetUid(c *gin.Context) int {
	ruserid, _ := strconv.Atoi(getuid(c)) //网关层保证ruserid有效
	return ruserid
}
func GetApp(c *gin.Context) string {
	return c.Request.Header.Get(MiddleAppKey)
}
func GetAdminAccount(c *gin.Context) string {
	return c.Request.Header.Get(MiddleAdminKey)
}

func getuid(c *gin.Context) string {
	return c.Request.Header.Get(MiddleUidKey)
}

func QueryInt(c *gin.Context, key string, defaultVal int) int {
	val, ok := c.GetQuery(key)
	if ok {
		res, _ := strconv.Atoi(val)
		return res
	}
	return defaultVal
}
func QueryStr(c *gin.Context, key string, defaultVal string) string {
	val, ok := c.GetQuery(key)
	if ok {
		return val
	}
	return defaultVal
}
func ParamInt(c *gin.Context, key string, defaultVal int) int {
	val := c.Param(key)
	if val == "" {
		res, _ := strconv.Atoi(val)
		return res
	}
	return defaultVal
}
func ParamStr(c *gin.Context, key string, defaultVal string) string {
	val := c.Param(key)
	if val == "" {
		return val
	}
	return defaultVal
}
