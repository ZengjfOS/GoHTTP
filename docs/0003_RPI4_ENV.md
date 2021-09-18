# RPI4 ENV

树莓派直接运行Go代码配置

## 配置环境变量

* 不配置代码下载不了依赖，添加依赖后，也需要配置这个，否则容易无法下载
* export GOPROXY=https://goproxy.io
* go env
* go run main.go
  * 会自动下载依赖库

## 示例

* [Go Web Examples Websockets](https://gowebexamples.com/websockets/)
* [gin websocket 一对一聊天](https://segmentfault.com/a/1190000023581108)

## error

* log
  ```
  [GIN-debug] [WARNING] Now Gin requires Go 1.12+.
  
  [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
  
  [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
   - using env:   export GIN_MODE=release
   - using code:  gin.SetMode(gin.ReleaseMode)
  
  [GIN-debug] GET    /*filepath                --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
  [GIN-debug] HEAD   /*filepath                --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (3 handlers)
  [GIN-debug] POST   /upload                   --> main.main.func1 (3 handlers)
  [GIN-debug] GET    /ws/echo                  --> main.main.func2 (3 handlers)
  panic: '/ws/echo' in new path '/ws/echo' conflicts with existing wildcard '/*filepath' in existing prefix '/*filepath'
  
  goroutine 1 [running]:
  github.com/gin-gonic/gin.(*node).addRoute(0x1856d80, 0x18186d0, 0x8, 0x180c180, 0x3, 0x3)
          /home/pi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/tree.go:240 +0x38c
  github.com/gin-gonic/gin.(*Engine).addRoute(0x1a420d0, 0x506f42, 0x3, 0x18186d0, 0x8, 0x180c180, 0x3, 0x3)
          /home/pi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/gin.go:289 +0xdc
  github.com/gin-gonic/gin.(*RouterGroup).handle(0x1a420d0, 0x506f42, 0x3, 0x50aa01, 0x8, 0x180e2a0, 0x1, 0x1, 0x5a9401, 0x180e2a0)
          /home/pi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/routergroup.go:75 +0x160
  github.com/gin-gonic/gin.(*RouterGroup).GET(0x1a420d0, 0x50aa01, 0x8, 0x180e2a0, 0x1, 0x1, 0x5a94c8, 0x1a420d0)
          /home/pi/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/routergroup.go:103 +0x54
  main.main()
          /home/pi/zengjf/GoHTTP/src/0011_GinWebsocket/main.go:86 +0x154
  exit status 2
  ```
* [Wildcard route conflicts with static files #360](https://github.com/gin-gonic/gin/issues/360)
  * router.StaticFile("/", "static/index.html")

