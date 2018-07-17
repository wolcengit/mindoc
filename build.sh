#!/bin/bash

MODE="$1"
export MINDOC_VER="0.10.1-Z6"

if [ "$MODE"x = "push"x ]
then
  docker push wolcen/mindoc:$MINDOC_VER
  echo "push over!"
  exit 0
fi

rm -rf mindoc_linux_amd64
#dep ensure -v
go build -v -work -x -o mindoc_linux_amd64 -ldflags="-w -X github.com/lifei6671/mindoc/conf.VERSION=$MINDOC_VER -X 'github.com/lifei6671/mindoc/conf.BUILD_TIME=`date`' -X 'conf.GO_VERSION=`go version`'"
# 准备
rm -rf zmindoc
mkdir zmindoc
mkdir -p zmindoc/conf
mkdir -p zmindoc/uploads
cp -rf conf/app.conf.example zmindoc/conf/
cp -rf lib zmindoc/
cp -rf static zmindoc/
cp -rf views zmindoc/
cp -rf LICENSE.md zmindoc/
cp -rf README.md zmindoc/
cp -rf favicon.ico zmindoc/
cp -rf mindoc_linux_amd64 zmindoc/
cp -rf start.sh zmindoc/
# 创建
docker build -t wolcen/mindoc:$MINDOC_VER .
# 清理
rm -rf zmindoc
echo "build over!"


