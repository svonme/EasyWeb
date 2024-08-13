package html

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"web/pkg/config"
	"web/pkg/file"
)

type Data struct {
	Title   string `json:"title"`
	Layout  string `json:"layout"`
	Assets  string `json:"assets"`
	Content string
}

func getConfig(text string) *Data {
	lines := strings.Split(text, "\n")
	firstLine := lines[0]

	start := strings.IndexRune(firstLine, '{')

	end := strings.LastIndex(firstLine, "}")

	if start >= 0 && end >= 0 {
		data := &Data{}

		value := firstLine[start : end+1]
		err := json.Unmarshal([]byte(value), &data)
		if err == nil {
			str := strings.Join(lines[1:], "\n")
			data.Content = str
			return data
		}
	}

	return &Data{
		Title:   "",
		Layout:  "layout/index.html",
		Content: text,
	}
}

func Template(templateName string) gin.HandlerFunc {
	conf := config.Load()
	return func(c *gin.Context) {
		fileName := c.Request.URL.Path
		if fileName == "/" || fileName == "index" {
			fileName = "/index.html"
		}
		fileName = file.DecorateHtml(fileName)
		filePath := filepath.Join(templateName, fileName)
		status := file.IsFileExist(filePath)

		fmt.Printf("url =%s\n", filePath)

		var text string
		if status && file.IsHtml(filePath) {
			text = file.GetContent(filePath)
		}
		if len(text) > 0 {
			data := getConfig(text)

			if len(data.Title) < 1 {
				data.Title = conf.Title
			}

			if len(data.Assets) < 1 {
				data.Assets = conf.Assets
			}

			// 定义模板数据
			htmlData := gin.H{
				"assets": data.Assets,
				"title":  data.Title,
			}
			// 创建新的模板对象并解析模板字符串
			tmpl, err := template.New("").Parse(data.Content)
			if err != nil {
				log.Fatalf("Failed to parse template: %v", err)
			}
			// 使用缓冲区来捕获渲染结果
			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, htmlData); err != nil {
				log.Fatalf("Failed to execute template: %v", err)
			}

			res := gin.H{
				"assets":  data.Assets,
				"title":   data.Title,
				"content": template.HTML(buf.String()),
			}
			c.HTML(http.StatusOK, data.Layout, res)
		} else {
			// 返回响应
			c.JSON(http.StatusNotFound, gin.H{
				"message": fileName + " 文件不存在",
			})
		}
	}
}
