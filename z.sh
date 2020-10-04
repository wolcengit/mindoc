#!/bin/bash

export TAG=Z3.22
sudo docker build -t wolcen/mindoc:$TAG  .
sudo docker push wolcen/mindoc:$TAG
