## 后端文档说明


#### 1.0 如何安装

git clone https://github.com/zihao-boy/zihao.git

国内建议开启代理
export GO111MODULE=on
export GOPROXY=https://goproxy.io

window 

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io

git build .

./zihao > zihao.log &

github 太慢走自己的代理

git config --global http.proxy http://127.0.0.1:1080

git config --global https.proxy http://127.0.0.1:1080

git config --global --unset http.proxy

git config --global --unset https.proxy

#### 主机需要实现内容

df -h


## 梓豪平台

## 系统截图

![image](doc/1.png)

![image](doc/2.png)

![image](doc/3.png)

![image](doc/4.png)

![image](doc/5.png)

![image](doc/6.png)
