#!/bin/bash

container_id=`docker ps -a|awk '{if (NR>1){print $1}}'`

container_name=`docker ps -a|awk '{if (NR>1){print $(NF)}}'`

image=`docker ps -a|awk '{if (NR>1){print $2}}'`

port=`docker ps -a|awk '{if (NR>1){print $(NF-1)}}'`

echo "$container_id&&$container_name&&$image&&$port"
