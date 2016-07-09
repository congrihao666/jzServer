#!/bin/bash
#protoc --go_out=./src/protocal *.proto
CURDIR=`pwd`
export GOPATH=$GOPATH:$CURDIR

protoc --go_out=./src/SceneServerMsg/ SceneServerMsg.proto
protoc --go_out=./src/ZoneServerMsg/ ZoneServerMsg.proto
go build ./src/SceneServerMsg
go install ./src/SceneServerMsg
go build ./src/ZoneServerMsg
go install ./src/ZoneServerMsg


#go build ./src/protocal
#go install ./src/protocal
#echo 'go build protocal ok'

