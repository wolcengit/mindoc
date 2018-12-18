#FROM golang:1.10.3-alpine3.7 AS build

#RUN apk add --update && apk add --no-cache --upgrade gcc g++ make curl git

FROM wolcen/golang-compile:1.10.3-alpine3.7 AS build

#RUN mkdir -p /go/src/github.com/lifei6671/ && cd /go/src/github.com/lifei6671/ && git clone https://github.com/wolcengit/mindoc.git && cd mindoc

ADD . /go/src/github.com/lifei6671/mindoc
WORKDIR /go/src/github.com/lifei6671/mindoc

RUN	CGO_ENABLE=1 go build -v -a -o mindoc_linux_amd64 -ldflags="-w -s -X main.VERSION=$TAG -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"

#FROM registry.cn-hangzhou.aliyuncs.com/mindoc/mindoc:v0.13
FROM wolcen/mindoc:base

COPY --from=build /go/src/github.com/lifei6671/mindoc/conf/app.conf.example /mindoc/conf/app.conf.example
COPY --from=build /go/src/github.com/lifei6671/mindoc/static /mindoc/static
COPY --from=build /go/src/github.com/lifei6671/mindoc/views /mindoc/views
COPY --from=build /go/src/github.com/lifei6671/mindoc/mindoc_linux_amd64 /mindoc/mindoc_linux_amd64

WORKDIR /mindoc

RUN chmod +x start.sh

CMD ["./start.sh"]