# 指定源于一个基础镜像
FROM registry.cn-beijing.aliyuncs.com/sxd/ubuntu-java8:1.0
# 维护者/拥有者
MAINTAINER xxx <xxx@xx.com>
# 从宿主机上传文件 ，这里上传一个脚本，
# 点击选择业务包上传
# ADD bin/start_api.sh /root/

# 容器内执行相应指令
RUN chmod u+x /root/start_jar.sh
# 运行命令
# CMD <command>   or CMD [<command>]
# 整个Dockerfile 中只能有一个,多个会被覆盖的
CMD ["/root/start_jar.sh", "eureka"]