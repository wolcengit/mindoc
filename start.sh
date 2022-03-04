#!/bin/bash
set -eux

cp --no-clobber -rf /mindoc/conf.template/*  /mindoc/conf/
# 如果配置文件不存在就复制
cp --no-clobber /mindoc/conf.template/app.conf.example /mindoc/conf/app.conf

# 数据库等初始化
/mindoc/mindoc_linux_amd64 install

# 运行
/mindoc/mindoc_linux_amd64