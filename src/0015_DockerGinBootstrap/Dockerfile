FROM golang:1.17
MAINTAINER zengjf42@163.com

RUN apt-get -y update

#环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

#指定操作目录
WORKDIR /go/src
#复制根目录内所有源码文件到操作目录下
COPY . .

RUN go mod download

#暴露的端口
EXPOSE 8080

CMD go run main.go --port 8080
