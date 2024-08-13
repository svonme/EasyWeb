package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	// 模板地址
	Template string `json:"template"`
	// 页面地址
	Views string `json:"views"`
	// 静态文件地址
	Static string `json:"static"`
	// 服务端口
	Port string `json:"port"`
	// 静态文件域名
	Assets string `json:"assets"`
	// 网站默认标题
	Title string `json:"title"`
}

var _config *Config

func getConfigDir() string {
	// 设置配置文件地址
	configSrc := "config.json"
	if len(os.Args) >= 2 && len(os.Args[1]) > 0 {
		// 获取程序形参中传递的 config.json 文件地址
		configSrc = os.Args[1]
	}
	return configSrc
}

func Load() *Config {
	if _config != nil {
		return _config
	}
	dir := getConfigDir()
	file, err := os.Open(dir)
	if err != nil {
		return nil
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(byteValue, &_config)
	if err != nil {
		return nil
	}

	return _config
}
