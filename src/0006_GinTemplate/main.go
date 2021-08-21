package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 设置HTML中可以被允许调用的处理函数
func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")         // 设置HTML标记符
	router.SetFuncMap(template.FuncMap{ // 设置HTML可使用函数映射
		"formatAsDate": formatAsDate, // 具体的HTML中可使用的函数
	})
	// 加载当前所有的template文件，暂时不知道干啥用的
	router.LoadHTMLFiles("./testdata/index.tmpl", "./testdata/raw.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{
			"name": "zengjf",
		})
	})

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
