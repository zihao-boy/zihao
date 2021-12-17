# 梓豪平台是一个docker 容器化管理平台 包括主机管理，手工制作镜像，镜像管理，ci/cd 构建和应用管理的paas平台。欢迎安装使用，也欢迎贡献！


## 1.0 手动安装(install)

git clone https://github.com/zihao-boy/zihao.git

https://github.com.cnpmjs.org/zihao-boy/zihao.git

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

## 2.0 Linux 一键安装(auto install)

yum install -y wget && wget https://homecommunity.oss-cn-beijing.aliyuncs.com/install.sh && sh install.sh

yum install -y wget && wget https://homecommunity.oss-cn-beijing.aliyuncs.com/install.sh -O zihaoinstall.sh && sh zihaoinstall.sh

访问地址 http://ip:7000

账号为 zihao 
密码为 123456

## 3.0 演示地址(demo)

http://zihao.homecommunity.cn/

## 4.0 视频教程（video）

https://www.bilibili.com/video/BV1BM4y1w7Ti?spm_id_from=333.999.0.0

## 5.0 系统截图

![image](doc/1.png)

![image](doc/2.png)

![image](doc/3.png)

![image](doc/4.png)

![image](doc/5.png)

![image](doc/6.png)
