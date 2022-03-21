#!/bin/bash
export CONTAINER_NAME="fs-1"
docker kill ${CONTAINER_NAME}
docker rm ${CONTAINER_NAME}
docker run --name ${CONTAINER_NAME} -v `pwd`/${CONTAINER_NAME}:/data -w /data -d fileserver:1.0 fileserver