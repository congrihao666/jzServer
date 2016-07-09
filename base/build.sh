#!/bin/bash

CURDIR=`pwd`
export GOPATH=$GOPATH:$CURDIR
go build ./src/jznet
go install ./src/jznet
go build ./src/jzcode
go install ./src/jzcode

echo $GOPATH
echo 'build base ok' 

