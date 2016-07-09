#!/bin/bash
cd protocal
source ./build.sh
cd ..

cd base
source ./build.sh
echo $GOPATH
cd ..

cd ZoneServer
source ./build.sh
cd ..

cd RecordServer
source ./build.sh
cd ..

