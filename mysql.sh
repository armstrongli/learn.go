#!/bin/bash
set -x
export DATA_PATH="/Users/jianqli/go/src/learn.go/mysql"
docker run \
  --name learn.go \
  -p 3306:3306 \
  -v ${DATA_PATH}:/var/lib/mysql \
  -e MYSQL_ROOT_PASSWORD=learngo \
  -d mysql:5.7.37