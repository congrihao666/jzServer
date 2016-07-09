#!/bin/bash
CURDIR=`pwd`
export GOPATH=$GOPATH:$CURDIR
go build ./src/RecordServer
go install ./src/RecordServer
