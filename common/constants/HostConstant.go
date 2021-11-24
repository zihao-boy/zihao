package constants

const (
	Check_host_resource_shell = "#!/bin/bash\n\nmem_total=`free -m | awk '/Mem/ {print $2}'`\n\nmem_used=`free -m | awk '/Mem/ {print $3}'`\n\ndisk_total=`df -m $1 | grep $1 | awk '{print $2}'`\n\nif [ ! -n \"$disk_total\" ]; then\n        disk_total=0\nfi\n\ndisk_used=`df -m $1 | grep $1 | awk '{print $3}'`\n\nif [ ! -n \"$disk_used\" ]; then\n        disk_used=0\nfi\n\ncpu_us=`top -b -n 1 | grep Cpu | awk '{print $2}'`\n\necho \"{'cpuRate':$cpu_us,'memTotal':$mem_total,'memUsed':$mem_used,'diskTotal':$disk_total,'diskUsed':$disk_used}\"\n"

	Get_host_port_shell = "#!/bin/bash\n\nprotocol=`netstat -tunlp|awk '{if (NR>2){print $1}}'`\n\nport=`netstat -tunlp|awk '{if (NR>2){print $4}}'`\n\nprogram_name=`netstat -tunlp|awk '{if (NR>2){print $(NF)}}'`\n\necho \"$protocol&&$port&&$program_name\""

	Get_host_disk_shell = "#!/bin/bash\n\ndisk_name=`df -h|awk '{if (NR>2){print $1}}'`\n\nsize=`df -h|awk '{if (NR>2){print $2}}'`\n\nfree_size=`df -h|awk '{if (NR>2){print $4}}'`\n\ndir=`df -h|awk '{if (NR>2){print $6}}'`\n\n\necho \"$disk_name&&$size&&$free_size&&$dir\""

	DownLoad_slave_shell = ""
)
