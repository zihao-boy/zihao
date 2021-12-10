#!/bin/bash

# min mem
min_mem=$MIN_MEM

max_mem=$MAX_MEM

java -jar $min_mem $max_mem /root/service-mall-$1.jar