package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func IsHtml(name string) bool {
	reg := regexp.MustCompile(`(?i)\.html$`)
	// 判断文件路径是否匹配正则表达式
	return reg.MatchString(name)
}

func IsFileExist(src string) bool {
	// 使用Stat函数获取文件信息
	_, err := os.Stat(src)
	// 判断文件是否存在
	if err == nil {
		return true
	}
	return false
}

func GetContent(src string) string {
	// 读取文件内容
	content, err := ioutil.ReadFile(src)
	if err != nil {
		return ""
	}
	return string(content)
}

func DecorateHtml(name string) string {
	ext := filepath.Ext(name)
	// 如果文件存在格式，则直接返回
	if len(ext) > 0 {
		return name
	}
	// 默认为 html 文件
	return DecorateHtml(name + ".html")
}
