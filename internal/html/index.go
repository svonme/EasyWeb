package html

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
	"web/pkg/file"
)

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
			content := template.HTML(text)
			c.HTML(http.StatusOK, "layout.html", gin.H{
				"content": content,
			})
		} else {
			// 返回响应
			c.JSON(http.StatusNotFound, gin.H{
				"message": fileName + " 文件不存在",
			})
		}
	}
}
