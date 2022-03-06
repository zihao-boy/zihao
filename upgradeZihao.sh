#!/bin/bash

cd ..

rm -rf zihao_release*

wget https://homecommunity.oss-cn-beijing.aliyuncs.com/zihao_release.tar

tar -xvf zihao_release.tar

rm -rf ./master/web

rm -rf ./master/zihao

rm -rf ./slave/slave*

cp -r ./zihao_release/web ./master/

cp -r ./zihao_release/zihao ./master/

cp -r ./zihao_release/web/download/slave.tar ./slave/

chmod u+x ./master/zihao

tar -xvf ./slave/slave.tar -C ./slave/

chmod u+x ./slave/slave

sh ./slave/startSlave.sh &

sleep 3s

sh ./master/restart_zihao.sh &