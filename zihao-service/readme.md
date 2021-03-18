## 后端文档说明


#### 1.0 创建数据库和用户

```sql

./mysql -h127.0.0.1  -e "create user 'zihao'@'%' identified by 'zihao@12345678';"
./mysql -h127.0.0.1  -e "flush privileges;"
./mysql -h127.0.0.1  -e "CREATE DATABASE zihao ;"
./mysql -h127.0.0.1  -e "grant all privileges on zihao.* to 'zihao'@'%' ;"
```

#### 主机需要实现内容

df -h

