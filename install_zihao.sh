#!/bin/bash

pwd_cur=$(pwd)

local_ip=127.0.0.1:7000
local_host=5f3761ed-008a-475b-ad37-8cc35c88402c

# install zihao
chmod u+x restart_zihao.sh

sh restart_zihao.sh

chmod +x /etc/rc.d/rc.local
echo $pwd_cur'/restart_zihao.sh' >> /etc/rc.d/rc.local

sleep 10s

sh ./web/download/deploySlave.sh $local_ip $local_host

chmod +x /etc/rc.d/rc.local
echo '/zihao/slave/startSlave.sh' >> /etc/rc.d/rc.local