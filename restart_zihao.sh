#!/bin/bash

# restart zihao

# kill zihao process
ps -ef | grep zihao | awk '{print $2}' | xargs kill -9

chmod u+x zihao

# start zihao
./zihao > zihao.log &



