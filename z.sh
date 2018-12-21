#!/bin/bash

export TAG=Z3.7
docker build -t wolcen/mindoc:$TAG -t 172.10.60.2/wolcen/mindoc:$TAG .
docker push 172.10.60.2/wolcen/mindoc:$TAG
