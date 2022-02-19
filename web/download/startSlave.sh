#!/bin/bash

chmod u+x slave

ps -ef | grep ./slave | grep -v 'grep' | awk '{print $2}' | xargs kill -9

cd /zihao/slave

./slave >> slave.log &
