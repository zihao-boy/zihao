#!/bin/bash

# install 7z
wget https://homecommunity.oss-cn-beijing.aliyuncs.com/hc/p7zip_9.20.1.tar

tar -xvf p7zip_9.20.1.tar

cd p7zip_9.20.1

make && make install

if [ ! -d "/home/data" ];then
  mkdir -p /home/data
fi

if [ ! -d "/home/data/web" ];then
  cd /home/data
  # download web
  wget https://homecommunity.oss-cn-beijing.aliyuncs.com/hc/web.7z
  7za x web.7z -r -o./
fi

if [ ! -d "/home/data/nginx" ];then
  cd /home/data
  # download nginx
  wget https://homecommunity.oss-cn-beijing.aliyuncs.com/hc/nginx.7z
  7za x nginx.7z -r -o./
fi

if [ ! -d "/home/data/mysql" ];then
  cd /home/data
  # download nginx
  wget https://homecommunity.oss-cn-beijing.aliyuncs.com/hc/mysql.7z
  7za x mysql.7z -r -o./
fi
