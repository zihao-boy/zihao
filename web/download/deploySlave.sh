#!/bin/bash

ip=$1
slaveId=$2


function installSlave(){
    mkdir -p /zihao/slave

    cd /zihao/slave/

    ps -ef | grep slave | grep -v grep | awk '{print $2}' | xargs kill -9

    rm -rf conf slave*
    mkdir conf

    echo "mastIp=$ip" >> conf/zihao.properties
    echo "slaveId=$slaveId" >> conf/zihao.properties

    wget http://$ip/download/slave.tar

    tar -xvf slave.tar  
    #download start slave shell
    wget http://$ip/download/startSlave.sh
    chmod u+x startSlave.sh
    ./startSlave.sh

    # power on execute start slave shell
    chmod +x /etc/rc.d/rc.local
    echo '/zihao/slave/startSlave.sh' >> /etc/rc.d/rc.local
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
        systemctl enable docker
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