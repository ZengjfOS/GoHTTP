# Environment

环境变量说明

## 环境变量说明

* GOROOT: go的安装路径
* GOPATH: 作为编译后二进制文件存放目的地和import包时的搜索路径(其实也是你的工作路径)

```js
{
	// Use IntelliSense to learn about possible attributes.
	// Hover to view descriptions of existing attributes.
	// For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
	"version": "0.2.0",
	"configurations": [
		{
			"name": "Launch Package",
			"type": "go",
			"request": "launch",
			"mode": "auto",
			"program": "${fileDirname}",
			"env": {
				"GOROOT": "D:\\zengjf\\software\\golang",
				"GOPATH": "D:\\zengjf\\github\\GoHTTP",
			},
		}
	]
}
```

## mod

* cd src
* go mod init GoHTTP

## 镜像站

* 打开cmd
* go env -w GOPROXY=https://goproxy.cn,direct

## error

* error 1
  ```
  Starting: D:\zengjf\github\GoHTTP\bin\dlv-dap.exe dap --listen=127.0.0.1:56764
  DAP server listening at: 127.0.0.1:56764
  Build Error: go build -o d:\zengjf\github\GoHTTP\src\__debug_bin.exe -gcflags all=-N -l d:\zengjf\github\GoHTTP\src
  cannot find package "d:\\zengjf\\github\\GoHTTP\\src" in any of:
  	D:\zengjf\software\golang\src\d:\zengjf\github\GoHTTP\src (from $GOROOT)
  	D:\zengjf\github\GoHTTP\src\d:\zengjf\github\GoHTTP\src (from $GOPATH) (exit status 1)
  ```
* go env -w GO111MODULE=on

## Run

* F5
  ```
  Starting: D:\zengjf\github\GoHTTP\bin\dlv-dap.exe dap --listen=127.0.0.1:61306
  DAP server listening at: 127.0.0.1:61306
  Hello, World
  Process 17600 has exited with status 0
  Detaching
  dlv dap (14828) exited with code: 0
  ```
