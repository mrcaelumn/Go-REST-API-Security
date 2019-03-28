#!/usr/bin/env bash
export GO_REST_API_SECURITY=go-rest-api-security

# remove container
[ "$(docker ps -a | grep $GO_REST_API_SECURITY)" ] && docker rm -f $GO_REST_API_SECURITY

#build go project
make build_binary

#dockerize app
docker build -t $GO_REST_API_SECURITY:develop .

#run docker image
docker run -d \
    --name $GO_REST_API_SECURITY \
    -p 8088:8088 \
    --restart always \
    -e GORESTSECURITY_SOCKET=tcp://0.0.0.0:8080 \
    $GO_REST_API_SECURITY:develop