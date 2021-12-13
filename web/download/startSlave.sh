#!/bin/bash

chmod u+x slave

#ps -ef | grep slave | awk '{print $2}' | xargs kill -9

cd /zihao/slave

./slave >> slave.log &
