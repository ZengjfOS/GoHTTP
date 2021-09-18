package main

import (
	"bufio"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./public/echart.tmpl")

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.StaticFile("/", "./public/index.html")
	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

		// Bind file
		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		// Save uploaded file
		file := bindFile.File
		dst := "tmp/" + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		fmt.Printf("File %s uploaded successfully with fields name=%s and email=%s.\n", file.Filename, bindFile.Name, bindFile.Email)

		data, err := os.Open(dst)
		if err != nil {
			c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, bindFile.Name, bindFile.Email))
		} else {

			scanner := bufio.NewScanner(data)
			scanner.Split(bufio.ScanLines)
			var dataArray []float32
			var line string

			for scanner.Scan() {
				line = scanner.Text()
				value, err := strconv.ParseFloat(line, 32)
				if err != nil {
					value = 0
				}
				dataArray = append(dataArray, float32(value)*100)
			}

			/*
				c.HTML(http.StatusOK, "echart.tmpl", map[string][]float32{
					"data": {820, 932, 901, 934, 1290, 1330, 1320},
				})
			*/

			c.HTML(http.StatusOK, "echart.tmpl", map[string][]float32{
				"data": dataArray,
			})
		}
		defer data.Close()

	})

	router.GET("/ws/echo", func(c *gin.Context) {
		conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	router.Run(":8080")
}
