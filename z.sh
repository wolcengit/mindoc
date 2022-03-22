#!/bin/sh

# build mindoc-base
# docker build -f Dockerfile-mindoc-base -t mindoc-base .

export MINDOC_VER="2.0.1"
#docker build --progress plain --build-arg TAG=$MINDOC_VER --tag mindoc:$MINDOC_VER .

export TAG=Z4.01
sudo docker tag mindoc:$MINDOC_VER registry.cn-hangzhou.aliyuncs.com/wolcen/mindoc:$TAG
sudo docker push registry.cn-hangzhou.aliyuncs.com/wolcen/mindoc:$TAG
