#!/bin/bash

# restart zihao

# kill zihao process
ps -ef | grep ./zihao | grep -v 'grep' | awk '{print $2}' | xargs kill -9

chmod u+x zihao

# start zihao
cd /zihao/master

./zihao > zihao.log &



