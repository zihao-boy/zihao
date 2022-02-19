#!/bin/bash

export GO111MODULE=on
export GOPROXY=https://goproxy.io

sh compile.sh

rm -rf /zihao/master/web

rm -rf /zihao/master/zihao

rm -rf /zihao/slave/slave*

cp -r ../zihao_release/web /zihao/master/

cp -r ../zihao_release/zihao /zihao/master/

cp -r ../zihao_release/web/download/slave.tar /zihao/slave/

chmod u+x /zihao/master/zihao

tar -xvf /zihao/slave/slave.tar -C /zihao/slave/

chmod u+x /zihao/slave/slave

sh /zihao/slave/startSlave.sh &

sh /zihao/master/restart_zihao.sh &