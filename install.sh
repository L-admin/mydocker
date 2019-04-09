#!/bin/bash

# 辅助安装cli和logrus包

# todo: 检查环境变量GOROOT,GOPATH是否存在

WORKDIR=${GOPATH}/src

cd ${WORKDIR}

# todo: 失败退出
go get github.com/urfave/cli
echo "cli install success."

# todo: 失败退出
mkdir -p golang.org/x && cd golang.org/x                                  
git clone https://github.com/golang/sys.git
cd sys/unix
go install
go get github.com/Sirupsen/logrus
echo "logrus install success."
