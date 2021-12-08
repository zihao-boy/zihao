#!/bin/bash

cur_pwd=$(pwd)

# build master
go build .

# build slave
cd ./slave

go build .

tar -cvf slave.tar slave

cp -r slave.tar ../web/download/

cd $cur_pwd

cd ..
# if not exits please new
if [ ! -d "zihao_release" ]; then
    mkdir zihao_release
fi

cd $cur_pwd

cp -r web ../zihao_release

cp -r zihao ../zihao_release/

if [ ! -d "../zihao_release/conf" ]; then
    cp -r conf ../zihao_release/
fi

if [ ! -d "../zihao_release/db" ]; then
    cp -r db ../zihao_release/
fi


