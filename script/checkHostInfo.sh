#!/bin/bash

mem_total=`free -m | awk '/Mem/ {print $2}'`

mem_used=`free -m | awk '/Mem/ {print $3}'`

disk_total=`df -m $1 | grep $1 | awk '{print $2}'`

if [ ! -n "$disk_total" ]; then
        disk_total=0
fi

disk_used=`df -m $1 | grep $1 | awk '{print $3}'`

if [ ! -n "$disk_used" ]; then
        disk_used=0
fi

cpu_us=`top -b -n 1 | grep Cpu | awk '{print $2}'`

echo "{'cpuRate':$cpu_us,'memTotal':$mem_total,'memUsed':$mem_used,'diskTotal':$disk_total,'diskUsed':$disk_used}"
