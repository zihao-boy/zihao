# 如何安装go

## 1.0 下载go

```shell

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
```

## 2.0 下载代码

```shell
git clone https://github.com/zihao-boy/zihao.git

https://github.com.cnpmjs.org/zihao-boy/zihao.git

github 太慢走自己的代理

git config --global http.proxy http://127.0.0.1:1080

git config --global https.proxy http://127.0.0.1:1080

git config --global --unset http.proxy

git config --global --unset https.proxy
```


## 3.0 编译并启动
```shell
git build .
./zihao > zihao.log &
```

## 4.0 访问系统

http://ip:7000 

账号密码：梓豪/123456