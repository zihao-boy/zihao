#!/bin/bash

#
cd /zihao/slave

chmod u+x slave

ps -ef | grep ./slave | grep -v 'grep' | grep -v 'startSlave.sh' | awk '{print $2}' | xargs kill -9

./slave >> slave.log &
