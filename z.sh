#!/bin/bash

export TAG=Z3.22
#sudo docker build -t wolcen/mindoc:$TAG  .
sudo docker push wolcen/mindoc:$TAG
sudo docker tag wolcen/mindoc:$TAG registry.cn-hangzhou.aliyuncs.com/wolcen/mindoc:$TAG
sudo docker push registry.cn-hangzhou.aliyuncs.com/wolcen/mindoc:$TAG
