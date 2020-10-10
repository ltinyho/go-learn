package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ltinyho/go-learn/pkg/gin_util"
	"gogs.sharkgulf.cn/sg/library/utils/qutil"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"runtime"
)

func route(r *gin.Engine) {
	r.Use(gin_util.LogReq(log))
	r.GET("/ping", ping)
	r.POST("/upload", upload)
	r.POST("/save", save)
	r.GET("/test", func(c *gin.Context) {
		fileData, err := ioutil.ReadFile("test.img")
		if err != nil {
			log.Error(err)
			gin_util.SendPureFailed(c)
			return
		}
		fmt.Println(len(fileData))
		fileData = nil
		gin_util.SendPureOk(c)
	})
	r.GET("/bak", func(c *gin.Context) {
		f, err := os.Open("test.img")
		if err != nil {
			log.Error(err)
			gin_util.SendPureFailed(c)
			return
		}
		err = _saveFileByReader(f, "test.bak.img")
		if err != nil {
			log.Error(err)
			gin_util.SendPureFailed(c)
			return
		}
		gin_util.SendPureOk(c)
	})
}

func ping(c *gin.Context) {
	runtime.NumGoroutine()
	gin_util.SendOk(c, "pong")
	return
}

func test(c *gin.Context) {
	f, err := os.Open("test.img")
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	var buf bytes.Buffer
	buf.ReadFrom(f)
	f.Close()
	log.Debugf("name%s", f.Name())
	runtime.GC()
	gin_util.SendPureOk(c)
	return
}
func save(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	err = saveFile(f)
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	gin_util.SendPureOk(c)
	return
}

func upload(c *gin.Context) {
	ff, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	data, err := ff.Open()
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	defer data.Close()
	err = _saveFileByReader(data, ff.Filename)
	if err != nil {
		log.Error(err)
		gin_util.SendPureFailed(c)
		return
	}
	gin_util.SendPureOk(c)
	return
}

func _saveFileByReader(file io.Reader, name string) (err error) {
	of, err := os.Create(name)
	if err != nil {
		return
	}
	_, err = io.Copy(of, file)
	return
}

func _saveFileByBytes(data []byte, name string) (err error) {
	of, err := os.Create(name)
	if err != nil {
		return
	}
	_, err = of.Write(data)
	return
}
func saveFile(f *multipart.FileHeader) (err error) {
	body, err := qutil.PostFile("http://localhost:8080/upload", f)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	return
}
