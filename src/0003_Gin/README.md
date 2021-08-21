# README

## 安装

* go get -u github.com/gin-gonic/gin
* go get -u golang.org/x/tools

## 使用

* go mod init Gin
* go mod edit -require github.com/gin-gonic/gin@latest
* go mod tidy

## Gin没有提示

* 拷贝`pkg/mod/github.com/gin-gonic/gin@v1.7.4`到`src/github.com/gin-gonic/gin`
* 重启vscode
