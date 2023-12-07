#!/bin/bash

CUR_DIR=$(pwd)
PROTOC_DIR="${CUR_DIR}/protoc"
GANESHA_SRC_DIR="${CUR_DIR}/server"

# install the protocol buffer compiler
cp ${PROTOC_DIR}/bin/protoc /usr/local/bin
cp -fr ${PROTOC_DIR}/include/google /usr/local/include
# install the Go protocol buffers plugin
cp ${PROTOC_DIR}/bin/protoc-gen-go /usr/local/bin
cp ${PROTOC_DIR}/bin/protoc-gen-go-grpc /usr/local/bin

protoc -I=${GANESHA_SRC_DIR} --go_out=${GANESHA_SRC_DIR}/protocol \
	--go-grpc_out=${GANESHA_SRC_DIR}/protocol \
	${GANESHA_SRC_DIR}/protocol/ganesha.proto
if [ $? != 0 ];then
    echo "compile protocol buffer failed"
    exit 1
fi

go_ver=`go version`
echo "${go_ver}"
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go build -o ./server/server ./server
if [ $? != 0 ];then
    echo "build ganesha rpc server failed"
    exit 1
fi
echo "build ganesha rpc server success"

go build -o ./test/ganesh_rpc_test ./test
if [ $? != 0 ];then
    echo "build ganesha rpc server test failed"
fi
echo "build ganesha rpc server test success"
