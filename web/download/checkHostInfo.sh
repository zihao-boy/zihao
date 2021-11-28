#!/bin/bash

#获取磁盘使用率
data_name=$1
#获取cpu使用率
cpuUsage=`top -n 1 | awk -F '[ %]+' 'NR==3 {print $2}'`

diskUsage=`df -h | grep $data_name | awk -F '[ %]+' '{print $5}'`
#获取内存情况
mem_total=`free -m | awk -F '[ :]+' 'NR==2{print $2}'`
mem_used=`free -m | awk -F '[ :]+' 'NR==3{print $3}'`
mem_free=`${mem_total}-${mem_user}`
#统计内存使用率
mem_used_persent=`awk 'BEGIN{printf "%.0f\n",('$mem_used'/'$mem_total')*100}'`

#获取报警时间
now_time=`date '+%F %T'`
function check(){
     echo "'cpuRate':${cpuUsage}, 'diskRate':${diskUsage},'memRate':${mem_used_persent},'freeMem':${mem_free}}"
}
function main(){
     check
}
main
