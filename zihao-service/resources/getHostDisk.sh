#!/bin/bash

disk_name=`df -h|awk '{if (NR>2){print $1}}'`

size=`df -h|awk '{if (NR>2){print $2}}'`

free_size=`df -h|awk '{if (NR>2){print $4}}'`

dir=`df -h|awk '{if (NR>2){print $6}}'`


echo "$disk_name&&$size&&$free_size&&$dir"
