# 如何安装go

下载go

wget https://golang.google.cn/dl/go1.17.3.linux-amd64.tar.gz

 tar zxvf go1.17.3.linux-amd64.tar.gz -C /usr/local

 vim /etc/profile

 文件最后输入如下信息

 export GOROOT=/usr/local/go

export GOPATH=/usr/local/go/go_path

export GOBIN=$GOPATH/bin

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

保存

source /etc/profile

export GO111MODULE=on
export GOPROXY=https://goproxy.io
