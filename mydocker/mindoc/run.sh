#!/bin/bash
export SRC_HOME=../..
# 准备
cp -rf $SRC_HOME/lib zmindoc/
cp -rf $SRC_HOME/static zmindoc/
cp -rf $SRC_HOME/views zmindoc/
cp -rf $SRC_HOME/LICENSE.md zmindoc/
cp -rf $SRC_HOME/README.md zmindoc/
cp -rf $SRC_HOME/favicon.ico zmindoc/
cp -rf $SRC_HOME/mindoc zmindoc/
# 创建
docker build -t wolcen/mindoc .
# 清理
rm -rf zmindoc/lib
rm -rf zmindoc/static
rm -rf zmindoc/views
rm -rf zmindoc/LICENSE.md
rm -rf zmindoc/README.md
rm -rf zmindoc/favicon.ico
rm -rf zmindoc/mindoc

