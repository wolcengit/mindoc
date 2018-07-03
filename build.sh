#!/bin/bash
MODE="$1"
export MINDOC_VER="0.10.1-Z1"
rm -rf mindoc
#dep ensure -v
go build -v -work -x -o mindoc -ldflags="-w -X github.com/lifei6671/mindoc/conf.VERSION=$MINDOC_VER -X 'github.com/lifei6671/mindoc/conf.BUILD_TIME=`date`' -X 'conf.GO_VERSION=`go version`'"
if [ "$MODE"x = "push"x ]
then
  cd mydocker/mindoc
  sh run.sh $MINDOC_VER
fi
echo "build over!"

