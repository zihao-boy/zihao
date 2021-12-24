#!/bin/bash

# min mem
min_mem=$MIN_MEM

max_mem=$MAX_MEM

active=$ACTIVE

java -jar -Dspring.profiles.active=$active $min_mem $max_mem /root/service-$1.jar