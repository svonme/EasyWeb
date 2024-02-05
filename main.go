package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"web/internal/html"
	"web/pkg/config"
	"web/pkg/file"
)

var Mode string // 用于存储环境变量的值

func main() {
	conf := config.Load()

	if conf == nil {
		fmt.Printf("请正确配置 config.json 文件\n")
		return
	}

	if Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.LoadHTMLGlob(filepath.Join(conf.Template, "*"))

	template := html.Template(conf.Views)

	// 拦截 html 文件, 直接返回
	r.Use(func(c *gin.Context) {
		fileName := c.Request.URL.Path
		if file.IsHtml(fileName) {
			template(c)
			c.Abort()
		} else {
			c.Next()
		}
	})

	// 配置静态文件服务
	if len(conf.Static) > 0 {
		_, fileName := filepath.Split(conf.Static)
		if len(fileName) > 0 {
			name := "/" + fileName
			r.Static(name, conf.Static)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		fileName := c.Request.URL.Path
		ext := filepath.Ext(fileName)
		// 如果存在文件后缀，则默认为异常
		if len(ext) > 0 {
			// 返回响应
			c.JSON(http.StatusNotFound, gin.H{
				"message": fileName + " 文件不存在",
			})
		} else {
			template(c)
		}
		c.Abort()
	})
	if len(conf.Port) > 0 {
		fmt.Printf("正常启动, 请访问 http://127.0.0.1:%s/\n", conf.Port)
		_ = r.Run(fmt.Sprintf(":%s", conf.Port))
	} else {
		fmt.Printf("正常启动, 请访问 http://127.0.0.1:8080/\n")
		_ = r.Run(":8080")
	}

}
