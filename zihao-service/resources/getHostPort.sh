#!/bin/bash

protocol=`netstat -tunlp|awk '{if (NR>2){print $1}}'`

port=`netstat -tunlp|awk '{if (NR>2){print $4}}'`

program_name=`netstat -tunlp|awk '{if (NR>2){print $(NF)}}'`

echo "$protocol&&$port&&$program_name"
