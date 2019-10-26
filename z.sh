#!/bin/bash

export TAG=Z3.18
docker build -t wolcen/mindoc:$TAG  .
docker push wolcen/mindoc:$TAG
