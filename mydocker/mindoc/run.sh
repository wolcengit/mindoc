#!/bin/bash
TAG="$1"
export SRC_HOME=../..
# 准备
mkdir zmindoc
mkdir -p zmindoc/conf
mkdir -p zmindoc/uploads
cp -rf $SRC_HOME/conf/app.conf.example zmindoc/conf/
cp -rf $SRC_HOME/lib zmindoc/
cp -rf $SRC_HOME/static zmindoc/
cp -rf $SRC_HOME/views zmindoc/
cp -rf $SRC_HOME/LICENSE.md zmindoc/
cp -rf $SRC_HOME/README.md zmindoc/
cp -rf $SRC_HOME/favicon.ico zmindoc/
cp -rf $SRC_HOME/mindoc_linux_amd64 zmindoc/
cp -rf $SRC_HOME/start.sh zmindoc/
# 创建
docker build -t wolcen/mindoc:$TAG .
# 清理
rm -rf zmindoc

docker push wolcen/mindoc:$TAG

