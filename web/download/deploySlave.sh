#!/bin/bash

ip=$1
slaveId=$2


function installSlave(){
    mkdir -p /zihao/slave

    cd /zihao/slave/

    rm -rf conf slave*
    mkdir conf

    echo "mastIp=$ip" >> conf/zihao.properties
    echo "slaveId=$slaveId" >> conf/zihao.properties

    wget http://$ip:7000/download/slave.tar

    tar -xvf slave.tar  

    chmod u+x slave

    ./slave >> slave.log &
}


# install docker 
function docker_install()
{
	echo "检查Docker......"
	docker -v
    if [ $? -eq  0 ]; then
        echo "检查到Docker已安装!"
    else
    	echo '开始安装docker'
        mkdir /data
        yum install -y yum-utils device-mapper-persistent-data lvm2
        yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
        yum makecache fast
        yum -y install docker-ce
        service docker start
        cat>/etc/docker/daemon.json<<EOF
{"log-driver":"json-file","log-opts":{"max-size":"500m","max-file":"3"}}
EOF
        systemctl daemon-reload
        systemctl restart docker
        groupadd docker
        systemctl restart docker
    fi
    # 创建公用网络==bridge模式
    docker network create java110-net
}



function main(){
    #安装docker
     docker_install
     #安装slave
     installSlave
}
main