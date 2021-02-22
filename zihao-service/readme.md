## 后端文档说明

参考文档： https://gitee.com/yhm_my/go-iris/tree/master

webshell  https://github.com/shibingli/webconsole

#### 1.0 创建数据库和用户

```sql

./mysql -h127.0.0.1  -e "create user 'zihao'@'%' identified by 'zihao@12345678';"
./mysql -h127.0.0.1  -e "flush privileges;"
./mysql -h127.0.0.1  -e "CREATE DATABASE zihao ;"
./mysql -h127.0.0.1  -e "grant all privileges on zihao.* to 'zihao'@'%' ;"
```