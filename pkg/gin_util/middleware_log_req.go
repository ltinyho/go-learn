package gin_util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type resWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r resWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogReq(log *logrus.Entry) gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &resWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = w
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		requestId := GetRequestId(c)
		contentType := c.Request.Header.Get("Content-Type")
		uid := getuid(c)
		_log := log.WithField("requestId", requestId)
		if !strings.Contains(contentType, "multipart/form-data") {
			_log.Infof("url=%s uid=%s req=%s", c.Request.URL, uid, bodyBytes)
		} else {
			_log.Infof("url=%s uid=%s", c.Request.URL, uid)
		}
		_ = c.Request.Body.Close() //  must close
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		_log.Debugf("url=%s uid=%s res=%s", c.Request.URL, uid, w.body.Bytes())
	}
}
