#!/bin/bash
CURDIR=`pwd`
export GOPATH=$GOPATH:$CURDIR
go build ./src/ZoneServer
go install ./src/ZoneServer

