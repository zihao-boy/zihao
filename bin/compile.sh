#!/bin/bash

cur_pwd=$(pwd)

cd ../

go build .

cd ./slave

go build .

tar -cvf slave.tar slave

cp -r slave.tar ../web/download/

