# image wolcen/calibre
FROM ubuntu:17.10
LABEL author="wolcen@msn.com"
# 安装calibre
RUN apt-get update -y \
  && apt-get install -y calibre \
  && apt-get clean \
  && rm -rf /var/cache/apt/* /var/lib/apt/lists/*
