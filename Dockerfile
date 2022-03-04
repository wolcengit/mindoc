FROM amd64/golang:1.13 AS build

ARG TAG=2.0.1

# 编译-环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=1
ENV GOARCH=amd64
ENV GOOS=linux

# 工作目录
ADD . /go/src/github.com/mindoc-org/mindoc
WORKDIR /go/src/github.com/mindoc-org/mindoc

# 编译
RUN go env
RUN go mod tidy -v
RUN go build -o mindoc_linux_amd64 -ldflags "-w -s -X 'main.VERSION=$TAG' -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`'"
RUN cp conf/app.conf.example conf/app.conf
# 清理不需要的文件
RUN rm appveyor.yml docker-compose.yml Dockerfile .travis.yml .gitattributes .gitignore go.mod go.sum main.go README.md simsun.ttc start.sh conf/*.go
RUN rm -rf cache commands controllers converter .git .github graphics mail models routers utils

# 测试编译的mindoc是否ok
RUN ./mindoc_linux_amd64 version

# 必要的文件复制
ADD simsun.ttc /usr/share/fonts/win/
ADD start.sh /go/src/github.com/mindoc-org/mindoc


# Ubuntu 20.04
FROM mindoc-base

COPY --from=build /go/src/github.com/mindoc-org/mindoc /mindoc
RUN cp -rf /mindoc/conf /mindoc/conf.template && rm -f /mindoc/conf.template/app.conf
WORKDIR /mindoc

EXPOSE 8181/tcp

ENV ZONEINFO=/mindoc/lib/time/zoneinfo.zip
RUN chmod +x /mindoc/start.sh

ENTRYPOINT ["/bin/bash", "/mindoc/start.sh"]

# export MINDOC_VER="2.0.1"
# docker build --progress plain --build-arg TAG=$MINDOC_VER --tag mindoc:$MINDOC_VER .
#
# mkdir -p  ~/mindoc/{conf,database,uploads}
# docker run -d -p 8181:8181 -v ~/mindoc/conf:/mindoc/conf -v ~/mindoc/database:/mindoc/database -v ~/mindoc/uploads:/mindoc/uploads --name mindoc -e MINDOC_ENABLE_EXPORT=true mindoc:$MINDOC_VER
