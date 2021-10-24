# README

可视化zip数据包

## 官方示例

https://github.com/gin-gonic/examples

## 文件上传

* 参考：https://github.com/gin-gonic/examples/tree/master/file-binding

## file test log

* zip
  ```
  Content Type: application/zip
  [GIN] 2021/08/21 - 16:49:37 | 200 |      2.4308ms |       127.0.0.1 | POST     "/upload"
  ```
* txt
  ```
  Content Type: application/octet-stream
  [GIN] 2021/08/21 - 16:49:46 | 200 |      1.6817ms |       127.0.0.1 | POST     "/upload"
  ```
* png
  ```
  Content Type: image/png
  [GIN] 2021/08/21 - 16:51:44 | 200 |      1.6176ms |       127.0.0.1 | POST     "/upload"
  ```
