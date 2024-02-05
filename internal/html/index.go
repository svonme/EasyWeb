package html

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"web/pkg/file"
)

type Data struct {
	Title   string `json:"title"`
	Layout  string `json:"layout"`
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

			c.HTML(http.StatusOK, data.Layout, gin.H{
				"title":   data.Title,
				"content": template.HTML(data.Content),
			})
		} else {
			// 返回响应
			c.JSON(http.StatusNotFound, gin.H{
				"message": fileName + " 文件不存在",
			})
		}
	}
}
