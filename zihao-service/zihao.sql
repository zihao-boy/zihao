
create table u_user(
                       user_id varchar(64) not null primary key comment '用户ID',
                       username varchar(64) not null comment '用户名',
                       real_name varchar(64) not null comment '真实名称',
                       passwd varchar(128) not null comment '密码',
                       sex int not null default 1 comment '1 男 2 女',
                       phone varchar(11) not null comment '手机号',
                       email varchar(64) comment '邮箱',
                       state varchar(12) not null default '100201' comment '用户状态 100201 在用',
                       create_time timestamp not null default current_timestamp comment '创建时间',
                       status_cd varchar(2) not null default '0' comment '数据有效'
)