#!/bin/bash

set -e

app=nas-spider

echo "==================开始构建${app}至bin目录=================="
GOOS=linux GOARCH=386 go build -o bin/${app}

echo ""==================构建完成"=================="

echo ""==================开始构建镜像"=================="
docker build --platform=linux/386 -t ${app}:latest .
echo ""==================构建完成"=================="

echo ""==================开始推送镜像"=================="
# 
docker login --username=小熊的专职浇水工 registry.cn-beijing.aliyuncs.com

docker tag ${app}:latest registry.cn-beijing.aliyuncs.com/levicy/${app}:latest

docker push registry.cn-beijing.aliyuncs.com/levicy/${app}:latest