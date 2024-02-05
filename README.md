# 操作说明

[Releases](https://github.com/svonme/EasyWeb/releases)


## 修改配置文件

编辑 config.json 文件中的配置参数, 如果没有则自行创建

```
{
  "template": "模板文件目录地址",
  "views": "网页文件目录地址",
  "static": "静态文件目录地址",
  "Port": "服务端口"
}

```

### static

除 html 文件外的资源文件, 比如 css, js 等文件

### template [使用教程](https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/gin%E6%B8%B2%E6%9F%93/html%E6%A8%A1%E6%9D%BF%E6%B8%B2%E6%9F%93.html)

配置公共文件, 比如配置统一的头部, 统一的尾部等内容。

##### 目录结构
    common
        head.html
    layout
        index.html

- define 定义模板名称
- template 导入公共文件

```html
{{ define "common/head.html" }}
<title>{{ .title }}</title>
{{ end }}
```

```html
{{ define "layout/index.html" }}
<!DOCTYPE html>
<html lang="en">
<head>
    {{ template "common/head.html" . }}
</head>
<body>
    {{ .content }}
</body>
</html>
{{ end }}
```

### views

所有 html 文件

##### 案例

html 文件中的第一行为配置信息

- title 网页标题
- layout 指定 template 中的模板名称

```html
<!--{ "title": "XXX", "layout": "layout/index.html" }-->
<div>
    <p>hello world</p>
</div>
```