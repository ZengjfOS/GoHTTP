package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, zengjf")
	})
	r.Run(":8888") // r.Run() default listen and serve on 0.0.0.0:8080
}
