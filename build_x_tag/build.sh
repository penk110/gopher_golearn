#!/bin/sh

COMMIT_ID=`git log |head -n 1| awk '{print $2;}'`
AUTHOR=`git log |head -n 3| grep Author| awk '{print $2;}'`
BRANCH_NAME=`git branch | awk '/\*/ { print $2; }'`
SERVICE_INFO="$COMMIT_ID,$AUTHOR,$BRANCH_NAME"
echo $SERVICE_INFO

# 通过上面的方式获取当前分支、版本等信息，在构建时通过构建参数注入到程序中
go build -ldflags "-X main.SERVICE_INFO=$SERVICE_INFO" -o build_x_tag