package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFile("/", "./public/index.html")
	router.Static("/js", "public/js")
	router.Static("/css", "public/css")

	router.Run(":8080")
}
