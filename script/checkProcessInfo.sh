#!/bin/bash

process_name=`ps -ef | grep target_process_name | grep -v grep | awk '{print $1}'`

if [ ! -n "$process_name" ]; then
        process_name=0
fi

echo "{'processName':'$process_name'}"
