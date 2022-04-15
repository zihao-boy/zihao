
CREATE TABLE `app_service` (
  `as_id` varchar(64) NOT NULL ,
  `as_name` varchar(128) NOT NULL ,
  `as_type` varchar(12) NOT NULL,
  `tenant_id` varchar(64) NOT NULL ,
  `as_desc` varchar(512) DEFAULT NULL ,
  `state` varchar(12) NOT NULL DEFAULT '10012' ,
  `as_count` int(11) NOT NULL DEFAULT '1' ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`as_id`)
);

alter table app_service add column ver_id varchar(64) ;

-- ----------------------------
-- Records of app_service
-- ----------------------------
INSERT INTO `app_service` VALUES ('808eb406-db6f-49d5-8d79-e5de6d601169', 'api服务', '003', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测试', '10012', '2', '2021-03-23 22:39:50', '1');
INSERT INTO `app_service` VALUES ('d1cf8ffb-3eaf-479d-9b7f-d22283c92d6a', 'api服务', '003', '512c369e-0642-41e5-9ea6-5fe737984ae6', 'api服务', '10012', '1', '2021-03-23 23:08:10', '0');

-- ----------------------------
-- Table structure for app_var
-- ----------------------------
DROP TABLE IF EXISTS `app_var`;
CREATE TABLE `app_var` (
  `av_id` varchar(64) NOT NULL ,
  `avg_id` varchar(64) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL,
  `var_name` varchar(128) NOT NULL ,
  `var_type` varchar(12) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  `var_spec` varchar(64) NOT NULL ,
  PRIMARY KEY (`av_id`)
);

-- ----------------------------
-- Records of app_var
-- ----------------------------
INSERT INTO `app_var` VALUES ('', '6eb0bbf1-47c5-461f-98f2-da8a006d4de4', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测试', '1001', '2021-03-24 22:48:54', '1', '');
INSERT INTO `app_var` VALUES ('0aa03297-7f2c-4c8f-b589-98e3026d1984', '9dc3295e-54e3-4259-ac28-c29bc03523e1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测试', '1001', '2021-03-24 22:50:37', '1', '');
INSERT INTO `app_var` VALUES ('72f96e23-3108-4136-a90e-f4281db50045', '9dc3295e-54e3-4259-ac28-c29bc03523e1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测', '1001', '2021-03-24 23:00:57', '0', '1001');
INSERT INTO `app_var` VALUES ('a3fb251c-0572-41e6-95a5-85c8bee7eb95', '9dc3295e-54e3-4259-ac28-c29bc03523e1', '512c369e-0642-41e5-9ea6-5fe737984ae6', 'Java内存', '1001', '2021-03-24 23:00:25', '1', '512');

-- ----------------------------
-- Table structure for app_var_group
-- ----------------------------
DROP TABLE IF EXISTS `app_var_group`;
CREATE TABLE `app_var_group` (
  `avg_id` varchar(64) NOT NULL ,
  `avg_name` varchar(128) NOT NULL ,
  `avg_type` varchar(12) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `avg_desc` varchar(512) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`avg_id`)
);

-- ----------------------------
-- Records of app_var_group
-- ----------------------------
INSERT INTO `app_var_group` VALUES ('27552e4a-4ba5-4836-9d6b-bdcd2418f2da', '测试组1', '001', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测试组', '2021-03-23 23:52:08', '1');
INSERT INTO `app_var_group` VALUES ('9dc3295e-54e3-4259-ac28-c29bc03523e1', '测试组1', '001', '512c369e-0642-41e5-9ea6-5fe737984ae6', '测试组1', '2021-03-23 23:53:12', '0');

-- ----------------------------
-- Table structure for app_version
-- ----------------------------
DROP TABLE IF EXISTS `app_version`;
CREATE TABLE `app_version` (
  `av_id` varchar(64) NOT NULL,
  `name` varchar(128) NOT NULL ,
  `remark` varchar(512) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`av_id`)
) ;

-- ----------------------------
-- Records of app_version
-- ----------------------------

-- ----------------------------
-- Table structure for app_version_attr
-- ----------------------------
DROP TABLE IF EXISTS `app_version_attr`;
CREATE TABLE `app_version_attr` (
  `attr_id` varchar(64) NOT NULL ,
  `av_id` varchar(64) NOT NULL,
  `version` varchar(32) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`attr_id`)
) ;

-- ----------------------------
-- Records of app_version_attr
-- ----------------------------

-- ----------------------------
-- Table structure for app_version_job
-- ----------------------------
DROP TABLE IF EXISTS `app_version_job`;
CREATE TABLE `app_version_job` (
       `job_id` varchar(64) NOT NULL,
       `job_name` varchar(128) NOT NULL,
       git_url varchar(512) not null,
       git_username varchar(128) ,
       git_passwd varchar(128) ,
       work_dir varchar(512) not null,
       `job_shell` longtext NOT NULL,
       `tenant_id` varchar(64) NOT NULL ,
       `state` varchar(12) NOT NULL DEFAULT '10012' ,
       `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
       `status_cd` varchar(2) NOT NULL DEFAULT '0',
       `job_time` datetime NOT NULL ,
       PRIMARY KEY (`job_id`)
);

-- ----------------------------
-- Records of app_version_job
-- ----------------------------
INSERT INTO `app_version_job` VALUES ('27f9cf18-b0f7-4fa9-8b39-3d26f9b6f06f', 'test', 'sdf2', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-27 15:32:14', '2021-03-27 15:32:14', '10012', '2021-03-27 15:32:14', '1');
INSERT INTO `app_version_job` VALUES ('2fc748ec-c257-437a-aa06-debc476dc5d8', '测试', 'mv a.txt b.txt1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-27 23:26:16', '2021-03-27 23:26:16', '1001', '2021-03-27 23:26:16', '1');
INSERT INTO `app_version_job` VALUES ('4045be60-fa1a-4e58-8fde-ed00d23023b8', '测试装填', '水电费', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-27 16:05:51', '2021-03-27 16:05:51', '10012', '2021-03-27 16:05:51', '1');
INSERT INTO `app_version_job` VALUES ('dbf958c1-676b-497b-bf44-f686ee1019a4', 'service-api', 'git clone https://gitee.com/wuxw7/MicroCommunity.git\ncd MicroCommunity\nmvn clean install', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-28 00:13:16', '2021-03-28 00:13:16', '2002', '2021-03-28 00:13:16', '0');
INSERT INTO `app_version_job` VALUES ('e5782206-2ebc-4bbf-be75-547763506eb7', '测试', 'mv a.txt b.txt1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-27 16:07:32', '2021-03-27 16:07:32', '1001', '2021-03-27 16:07:32', '1');

-- ----------------------------
-- Table structure for app_version_job_detail
-- ----------------------------
DROP TABLE IF EXISTS `app_version_job_detail`;
CREATE TABLE `app_version_job_detail` (
  `detail_id` varchar(64) NOT NULL ,
  `job_id` varchar(64) NOT NULL ,
  `log_path` varchar(256) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `state` varchar(12) NOT NULL DEFAULT '10012' ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`detail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_version_job_detail
-- ----------------------------
INSERT INTO `app_version_job_detail` VALUES ('03b12a15-6dc6-42b1-ae16-2b175183c2af', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:00:24', '0');
INSERT INTO `app_version_job_detail` VALUES ('1cee60a9-2f97-4549-82d4-f24e56adf490', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:50:00', '0');
INSERT INTO `app_version_job_detail` VALUES ('2bbf1078-80c1-472f-bf16-b384c4e3da61', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 22:28:03', '0');
INSERT INTO `app_version_job_detail` VALUES ('3c8aadca-967e-48af-8b22-01ff0075ef30', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 22:55:08', '0');
INSERT INTO `app_version_job_detail` VALUES ('571e7cd7-82fa-40a9-b346-d3030ac9f5de', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:26:55', '0');
INSERT INTO `app_version_job_detail` VALUES ('70093a83-78cf-453a-82f5-6b99dc8b9e2d', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:04:43', '0');
INSERT INTO `app_version_job_detail` VALUES ('757d83ae-e8b7-4c0c-947d-b9aa84abf341', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 21:45:28', '0');
INSERT INTO `app_version_job_detail` VALUES ('8a19fe81-79fc-412d-bf51-3a5244d825a5', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 22:47:25', '0');
INSERT INTO `app_version_job_detail` VALUES ('94b82462-6f17-4b1b-8a4c-3508838df6f8', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 21:27:43', '0');
INSERT INTO `app_version_job_detail` VALUES ('970b5dad-c30d-43b6-a69b-5da3b2007499', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:35:07', '0');
INSERT INTO `app_version_job_detail` VALUES ('c4ca9c16-4a30-42ab-85ce-dba306f05f02', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:30:11', '0');
INSERT INTO `app_version_job_detail` VALUES ('ca24303c-aeed-4b62-9561-ca2bec925108', 'dbf958c1-676b-497b-bf44-f686ee1019a4', '/Users/wuxuewen/zihao/dbf958c1-676b-497b-bf44-f686ee1019a4/dbf958c1-676b-497b-bf44-f686ee1019a4.log', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001', '2021-03-28 23:25:53', '0');

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `host_id` varchar(64) NOT NULL,
  `group_id` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `ip` varchar(20) NOT NULL,
  `username` varchar(64) NOT NULL,
  `passwd` varchar(64) NOT NULL,
  `cpu` int(11) NOT NULL,
  `mem` decimal(10,2) NOT NULL,
  `disk` int(11) NOT NULL,
  `tenant_id` varchar(64) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`host_id`)
);

alter table host add column heartbeat_time datetime ;

-- ----------------------------
-- Records of host
-- ----------------------------
INSERT INTO `host` VALUES ('0638e278-484c-4e49-b59f-9ed73ab418ed', 'fd96a59e-f04b-49d7-bd6a-0c9d284940e2', 'HC测试机', '106.52.221.206:9004', 'root', 'wuxw2015', '4', '16.00', '100', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-03 23:37:03', '0');
INSERT INTO `host` VALUES ('0a7575df-2c9a-4bf3-9596-13446fbdb56a', '测试组', '测试主机', '192.168.1.100', 'root', 'wuxw2015', '4', '8.00', '100', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-02-23 09:48:12', '1');
INSERT INTO `host` VALUES ('55dfa429-5f96-44f9-ace8-88870b6c3741', 'fd96a59e-f04b-49d7-bd6a-0c9d284940e2', 'HC测试主机1', '106.52.221.206:9007', 'root', 'wuxw2015', '6', '16.00', '100', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-02-23 10:10:54', '0');
INSERT INTO `host` VALUES ('aaefbf33-ad23-4f0a-9854-0f960dba397e', '02f35061-cab9-44a2-a44a-4de3ba8c39e8', '家里测试', '192.168.1.100', 'root', 'wuxw2015', '4', '16.00', '100', '2d4451e6-83ea-4028-99d4-6f5a8e1d4864', '2021-02-23 22:24:40', '0');
INSERT INTO `host` VALUES ('ccf0a121-60b5-4564-866e-6f6aa5f90fa3', 'fd96a59e-f04b-49d7-bd6a-0c9d284940e2', '192.168.1.106', '192.168.1.106:22', 'root', 'wuxw2015', '4', '16.00', '100', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-03-18 01:15:44', '0');

-- ----------------------------
-- Table structure for host_group
-- ----------------------------
DROP TABLE IF EXISTS `host_group`;
CREATE TABLE `host_group` (
  `group_id` varchar(64) NOT NULL,
  `name` varchar(64) NOT NULL,
  `description` varchar(200) DEFAULT NULL,
  `tenant_id` varchar(64) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`group_id`)
);

-- ----------------------------
-- Records of host_group
-- ----------------------------
INSERT INTO `host_group` VALUES ('', '生产组', '生产组', '', '2021-02-22 23:13:22', '0');
INSERT INTO `host_group` VALUES ('02f35061-cab9-44a2-a44a-4de3ba8c39e8', 'A组', 'A组', '2d4451e6-83ea-4028-99d4-6f5a8e1d4864', '2021-02-23 22:24:05', '0');
INSERT INTO `host_group` VALUES ('1', '测试组', '测试组', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-02-20 21:51:16', '0');
INSERT INTO `host_group` VALUES ('c282f3ed-769a-4efc-969a-77fe4c346097', '删除组', '删除组', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-02-22 23:41:09', '1');
INSERT INTO `host_group` VALUES ('fd96a59e-f04b-49d7-bd6a-0c9d284940e2', '生产组', '生产组', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2021-02-22 23:16:16', '0');

-- ----------------------------
-- Table structure for mapping
-- ----------------------------
DROP TABLE IF EXISTS `mapping`;
CREATE TABLE `mapping` (
  `id` INTEGER  PRIMARY KEY AUTOINCREMENT NOT NULL ,
  `domain` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL ,
  `zkeys` varchar(100) NOT NULL ,
  `value` varchar(1000) NOT NULL ,
  `remark` varchar(200) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' 
);

-- ----------------------------
-- Records of mapping
-- ----------------------------
INSERT INTO `mapping` VALUES ('1', 'DOMAIN.COMMON', '环境标识', 'ZIHAO_ENV', 'DEV', '环境标识', '2021-02-23 20:27:43', '0');
INSERT INTO `mapping` VALUES ('2', 'DOMAIN.COMMON', 'xx', 'xx', 'xx', 'xx', '2021-02-23 20:31:28', '1');

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `m_id` varchar(64) NOT NULL,
  `name` varchar(10) NOT NULL,
  `g_id` varchar(64) NOT NULL,
  `url` varchar(200) NOT NULL,
  `seq` int(11) NOT NULL,
  `description` varchar(200) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `is_show` varchar(2) NOT NULL DEFAULT 'Y',
  PRIMARY KEY (`m_id`)
);

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES ('700201904005', '主机组', '800201904003', '/index.html#/pages/admin/hostGroupManage', '1', '主机组', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904006', '修改密码', '800201904008', '/index.html#/pages/frame/changeStaffPwd', '1', '修改密码', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904007', '能力信息', '800201904009', '/index.html#/pages/admin/serviceSqlManage', '1', '能力信息', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904008', '主机资源', '800201904003', '/index.html#/pages/admin/hostManage', '2', '主机资源', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904009', '编码映射', '800201904009', '/index.html#/pages/admin/mappingManage', '2', '编码映射', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904010', '租户信息', '800201904009', '/index.html#/pages/admin/tenantManage', '3', '租户信息', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904011', '主机监控组', '800201904005', '/index.html#/pages/admin/monitorHostGroupManage', '1', '主机监控组', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904012', '监控主机', '800201904005', '/index.html#/pages/admin/monitorHostManage', '1', '监控主机', '2019-04-09 14:50:56', '0', 'N');
INSERT INTO `menu` VALUES ('700201904013', '主机监控', '800201904005', '/index.html#/pages/admin/monitorHostGroupManageAnalysis', '2', '主机监控', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904014', '监控事件', '800201904005', '/index.html#/pages/admin/monitorEventManage', '3', '监控事件', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904015', '租户设置', '800201904010', '/index.html#/pages/admin/tenantSettingManage', '1', '租户设置', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904016', '监控任务', '800201904005', '/index.html#/pages/admin/jobManage', '4', '监控任务', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904017', '主机详情', '800201904003', '/index.html#/pages/admin/hostDetailManage', '3', '主机详情', '2019-04-09 14:50:56', '0', 'N');
INSERT INTO `menu` VALUES ('700201904018', '服务信息', '800201904004', '/index.html#/pages/admin/appServiceManage', '3', '服务信息', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904019', '环境组', '800201904004', '/index.html#/pages/admin/appVarGroupManage', '1', '环境组', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904020', '环境变量', '800201904004', '/index.html#/pages/admin/appVarManage', '2', '环境变量', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904021', '应用版本', '800201904011', '/index.html#/pages/admin/appVersionManage', '2', '应用版本', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `menu` VALUES ('700201904022', '构建版本', '800201904011', '/index.html#/pages/admin/appVersionJobManage', '1', '构建版本', '2019-04-09 14:50:56', '0', 'Y');

-- ----------------------------
-- Table structure for menu_group
-- ----------------------------
DROP TABLE IF EXISTS `menu_group`;
CREATE TABLE `menu_group` (
  `g_id` varchar(64) NOT NULL,
  `name` varchar(10) NOT NULL,
  `icon` varchar(20) NOT NULL,
  `label` varchar(20) NOT NULL,
  `seq` int(11) NOT NULL,
  `description` varchar(200) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `group_type` varchar(12) NOT NULL DEFAULT 'P_WEB'
);

-- ----------------------------
-- Records of menu_group
-- ----------------------------
INSERT INTO `menu_group` VALUES ('800201904003', '资源中心', 'fa fa-globe', '', '1', '资源中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904004', '计算中心', 'fa fa-globe', '', '2', '计算中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904005', '监控中心', 'fa fa-globe', '', '4', '监控中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904006', '安全中心', 'fa fa-globe', '', '5', '安全中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904008', '个人中心', 'fa fa-globe', '', '6', '个人中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904009', '开发中心', 'fa fa-globe', '', '7', '开发中心', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904010', '系统设置', 'fa fa-globe', '', '8', '系统设置', '2019-04-01 07:55:51', '0', 'P_WEB');
INSERT INTO `menu_group` VALUES ('800201904011', '版本中心', 'fa fa-globe', '', '3', '版本中心', '2019-04-01 07:55:51', '0', 'P_WEB');

-- ----------------------------
-- Table structure for monitor_event
-- ----------------------------
DROP TABLE IF EXISTS `monitor_event`;
CREATE TABLE `monitor_event` (
  `event_id` varchar(64) NOT NULL,
  `event_type` varchar(12) NOT NULL,
  `event_obj_id` varchar(64) NOT NULL,
  `event_obj_name` varchar(128) NOT NULL,
  `tenant_id` varchar(64) NOT NULL,
  `threshold_value` decimal(10,2) NOT NULL,
  `cur_value` decimal(10,2) NOT NULL,
  `remark` longtext ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `notice_type` varchar(12) NOT NULL ,
  `state` varchar(12) NOT NULL ,
  `state_remark` longtext,
  PRIMARY KEY (`event_id`)
);

-- ----------------------------
-- Records of monitor_event
-- ----------------------------
INSERT INTO `monitor_event` VALUES ('006ea809-e089-452d-9103-795dd4665bb2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-07 09:06:04', '0', '2002', '000', '待告警');

-- ----------------------------
-- Table structure for monitor_host
-- ----------------------------
DROP TABLE IF EXISTS `monitor_host`;
CREATE TABLE `monitor_host` (
  `mh_id` varchar(64) NOT NULL ,
  `mhg_id` varchar(64) NOT NULL ,
  `host_id` varchar(64) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `cpu_rate` decimal(10,2) NOT NULL DEFAULT '0.00',
  `mem_rate` decimal(10,2) NOT NULL DEFAULT '0.00' ,
  `disk_rate` decimal(10,2) NOT NULL DEFAULT '0.00',
  `free_mem` decimal(10,2) NOT NULL DEFAULT '0.00' ,
  `free_disk` decimal(10,2) NOT NULL DEFAULT '0.00',
  `cpu_threshold` decimal(10,2) NOT NULL ,
  `mem_threshold` decimal(10,2) NOT NULL ,
  `disk_threshold` decimal(10,2) NOT NULL ,
  `mon_disk` varchar(128) NOT NULL ,
  `mon_date` datetime NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`mh_id`)
);

-- ----------------------------
-- Records of monitor_host
-- ----------------------------
INSERT INTO `monitor_host` VALUES ('9d29228a-80a1-48ed-b441-025f3b1526b9', 'd8a3f9d2-e1a5-4f14-a36e-0eaf4489a93e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '6904.00', '37730.00', '0.80', '0.70', '0.70', '/dev/mapper/centos-home', '2021-03-14 00:00:00', '2021-02-25 00:22:23', '0');
INSERT INTO `monitor_host` VALUES ('f5f4b2ef-e970-4d7d-a48e-a9c4803867ae', 'd8a3f9d2-e1a5-4f14-a36e-0eaf4489a93e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '7003.00', '141266.00', '0.70', '0.80', '0.80', '/dev/mapper/centos-home', '2021-03-14 00:00:00', '2021-03-03 23:37:46', '0');

-- ----------------------------
-- Table structure for monitor_host_group
-- ----------------------------
DROP TABLE IF EXISTS `monitor_host_group`;
CREATE TABLE `monitor_host_group` (
  `mhg_id` varchar(64) NOT NULL ,
  `name` varchar(64) NOT NULL ,
  `mon_cron` varchar(128) NOT NULL,
  `state` varchar(12) NOT NULL DEFAULT '3302' ,
  `mon_date` datetime NOT NULL ,
  `notice_type` varchar(12) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `remark` varchar(512) DEFAULT NULL,
  `tenant_id` varchar(64) NOT NULL,
  PRIMARY KEY (`mhg_id`)
);

-- ----------------------------
-- Records of monitor_host_group
-- ----------------------------
INSERT INTO `monitor_host_group` VALUES ('d8a3f9d2-e1a5-4f14-a36e-0eaf4489a93e', '生产组主机监控', '0 */1 * * * ?', '3302', '2021-02-24 22:36:32', '2002', '2021-02-24 22:36:32', '0', '测试', '512c369e-0642-41e5-9ea6-5fe737984ae6');

-- ----------------------------
-- Table structure for monitor_host_log
-- ----------------------------
DROP TABLE IF EXISTS `monitor_host_log`;
CREATE TABLE `monitor_host_log` (
  `log_id` varchar(64) NOT NULL ,
  `host_id` varchar(64) NOT NULL ,
  `tenant_id` varchar(64) NOT NULL ,
  `cpu_rate` decimal(10,2) NOT NULL DEFAULT '0.00' ,
  `mem_rate` decimal(10,2) NOT NULL DEFAULT '0.00' ,
  `disk_rate` decimal(10,2) NOT NULL DEFAULT '0.00' ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`log_id`)
) ;

-- ----------------------------
-- Records of monitor_host_log
-- ----------------------------
INSERT INTO `monitor_host_log` VALUES ('0053007d-09d5-4f21-8e67-fe6da16f68a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 15:52:04', '0');

-- ----------------------------
-- Table structure for privilege
-- ----------------------------
DROP TABLE IF EXISTS `privilege`;
CREATE TABLE `privilege` (
  `p_id` varchar(64) NOT NULL ,
  `name` varchar(10) NOT NULL ,
  `description` varchar(200) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  `resource` varchar(200) DEFAULT NULL ,
  `m_id` varchar(64) NOT NULL 
) ;

-- ----------------------------
-- Records of privilege
-- ----------------------------
INSERT INTO `privilege` VALUES ('500201904001', '主机组', '主机组', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904005');
INSERT INTO `privilege` VALUES ('500201904002', '修改密码', '修改密码', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904006');
INSERT INTO `privilege` VALUES ('500201904003', '服务信息', '服务信息', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904007');
INSERT INTO `privilege` VALUES ('500201904004', '主机资源', '主机资源', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904008');
INSERT INTO `privilege` VALUES ('500201904005', '编码映射', '编码映射', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904009');
INSERT INTO `privilege` VALUES ('500201904006', '租户信息', '租户信息', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904010');
INSERT INTO `privilege` VALUES ('500201904007', '主机监控组', '主机监控', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904011');
INSERT INTO `privilege` VALUES ('500201904008', '监控主机', '监控主机', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904012');
INSERT INTO `privilege` VALUES ('500201904009', '主机监控', '主机监控', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904013');
INSERT INTO `privilege` VALUES ('500201904010', '监控事件', '监控事件', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904014');
INSERT INTO `privilege` VALUES ('500201904011', '租户设置', '租户设置', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904015');
INSERT INTO `privilege` VALUES ('500201904012', '监控任务', '监控任务', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904016');
INSERT INTO `privilege` VALUES ('500201904013', '主机详情', '主机详情', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904017');
INSERT INTO `privilege` VALUES ('500201904014', '服务信息', '服务信息', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904018');
INSERT INTO `privilege` VALUES ('500201904015', '环境组', '环境组', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904019');
INSERT INTO `privilege` VALUES ('500201904016', '环境变量', '环境变量', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904020');
INSERT INTO `privilege` VALUES ('500201904017', '应用版本', '应用版本', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904021');
INSERT INTO `privilege` VALUES ('500201904018', '构建版本', '构建版本', '2019-04-01 02:24:53', '0', '/pages/admin/parkingSpaceFee', '700201904022');

-- ----------------------------
-- Table structure for privilege_group
-- ----------------------------
DROP TABLE IF EXISTS `privilege_group`;
CREATE TABLE `privilege_group` (
  `pg_id` varchar(64) NOT NULL ,
  `name` varchar(10) NOT NULL ,
  `description` varchar(200) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  `tenant_id` varchar(64) NOT NULL DEFAULT '9999' 
) ;

-- ----------------------------
-- Records of privilege_group
-- ----------------------------
INSERT INTO `privilege_group` VALUES ('600201904000', '管理员权限组', '管理员权限组', '2019-04-01 08:34:58', '0', '9999');
INSERT INTO `privilege_group` VALUES ('600201904002', '租户管理员权限组', '租户管理员权限组', '2019-04-01 08:34:58', '0', '9999');

-- ----------------------------
-- Table structure for privilege_rel
-- ----------------------------
DROP TABLE IF EXISTS `privilege_rel`;
CREATE TABLE `privilege_rel` (
  `rel_id` int(11) NOT NULL ,
  `p_id` varchar(64) NOT NULL ,
  `pg_id` varchar(64) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`rel_id`)
) ;

-- ----------------------------
-- Records of privilege_rel
-- ----------------------------
INSERT INTO `privilege_rel` VALUES ('1', '500201904001', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('2', '500201904002', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('3', '500201904003', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('4', '500201904004', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('5', '500201904005', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('6', '500201904006', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('7', '500201904001', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('8', '500201904002', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('9', '500201904004', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('10', '500201904007', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('11', '500201904007', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('12', '500201904008', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('13', '500201904008', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('14', '500201904009', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('15', '500201904009', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('16', '500201904010', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('17', '500201904010', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('18', '500201904011', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('19', '500201904011', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('20', '500201904012', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('21', '500201904012', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('22', '500201904013', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('23', '500201904013', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('24', '500201904014', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('25', '500201904014', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('26', '500201904015', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('27', '500201904015', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('28', '500201904016', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('29', '500201904016', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('30', '500201904017', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('31', '500201904017', '600201904002', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('32', '500201904018', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('33', '500201904018', '600201904002', '2019-04-01 08:18:29', '0');

-- ----------------------------
-- Table structure for privilege_user
-- ----------------------------
DROP TABLE IF EXISTS `privilege_user`;
CREATE TABLE `privilege_user` (
  `pu_id` INTEGER PRIMARY KEY  AUTOINCREMENT NOT NULL ,
  `p_id` varchar(64) NOT NULL,
  `privilege_flag` varchar(4) NOT NULL DEFAULT '0' ,
  `user_id` varchar(64) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  `tenant_id` varchar(64) NOT NULL ,
  PRIMARY KEY (`pu_id`)
) ;

-- ----------------------------
-- Records of privilege_user
-- ----------------------------
INSERT INTO `privilege_user` VALUES ('1', '600201904000', '1', '1', '2019-04-01 17:02:38', '0', '512c369e-0642-41e5-9ea6-5fe737984ae6');
INSERT INTO `privilege_user` VALUES ('2', '600201904002', '1', 'b8a0776d-521b-4393-8fe8-ae81b2587a05', '2021-02-23 21:54:20', '0', '13ba9db7-fcdd-496a-a9b0-abfd00bf24e3');
INSERT INTO `privilege_user` VALUES ('3', '600201904002', '1', '037eaaa6-1891-49fa-94d2-b9c3d7a96451', '2021-02-23 22:23:31', '0', '2d4451e6-83ea-4028-99d4-6f5a8e1d4864');

-- ----------------------------
-- Table structure for service_sql
-- ----------------------------
DROP TABLE IF EXISTS `service_sql`;
CREATE TABLE `service_sql` (
  `sql_id` INTEGER PRIMARY KEY  AUTOINCREMENT NOT NULL ,
  `sql_code` varchar(128) NOT NULL ,
  `sql_text` longtext NOT NULL ,
  `remark` varchar(256) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0'
);

-- ----------------------------
-- Records of service_sql
-- ----------------------------
INSERT INTO `service_sql` VALUES ('1', 'test', 'select * from u_user where 1=1 and name=#Name#', 'cd', '2021-02-22 21:22:52', '1');
INSERT INTO `service_sql` VALUES ('3', 'hostDao.SaveHostGroup', 'insert into host_group(group_id, name, description, tenant_id) VALUES \n(#GroupId#,#Name#,#Description#,#TenantId#)', '保存主机组', '2021-02-22 21:41:32', '0');
INSERT INTO `service_sql` VALUES ('4', 'hostDao.UpdateHostGroup', 'update host_group t set t.name = #Name#,t.description=#Description#\nwhere t.group_id = #GroupId#', '修改主机组', '2021-02-22 23:37:29', '0');
INSERT INTO `service_sql` VALUES ('5', 'hostDao.DeleteHostGroup', 'update host_group t set t.status_cd = \'1\'\nwhere t.group_id = #GroupId#', '删除主机组', '2021-02-22 23:38:28', '0');
INSERT INTO `service_sql` VALUES ('6', 'userDao.GetUser', 'select t.user_id ,t.username,t.real_name ,\nt.phone,t.email,t.state,t.create_time,t.tenant_id \nfrom u_user t\nwhere \nt.status_cd = \'0\'\n$if Username != \'\' then\nand t.username = #Username# \n$endif\n$if Passwd != \'\' then\nand t.passwd = #Passwd#\n$endif\n$if UserId != \'\' then\nand t.user_id = #UserId#\n$endif\nand t.state = \'100201\'\nlimit 1', '查询用户信息', '2021-02-23 00:10:30', '0');
INSERT INTO `service_sql` VALUES ('7', 'userDao.UpdateUser', 'update u_user t set \n$if Username != \'\' then\nt.username = #Username# ,\n$endif\n$if Passwd != \'\' then\nt.passwd=#Passwd# ,\n$endif\n$if State != \'\' then\nt.state = #State# ,\n$endif\n$if RealName != \'\' then\nt.real_name=#RealName# ,\n$endif\n$if Phone != \'\' then\nt.phone=#Phone# ,\n$endif\n$if Email != \'\' then\nt.email=#Email# ,\n$endif\nt.status_cd =\'0\'\nwhere t.user_id = #UserId# \nand t.status_cd = \'0\'', '修改用户信息', '2021-02-23 00:25:55', '0');
INSERT INTO `service_sql` VALUES ('8', 'hostDao.GetHostCount', 'select count(1) total from host t\nwhere \nt.status_cd = \'0\'\n  $if HostId != \'\' then\nand t.host_id = #HostId#\n$endif\n$if Name != \'\' then\nand t.name = #Name#\n$endif\n$if GroupId != \'\' then\nand t.group_id = #GroupId#\n$endif\n$if TenantId != \'\' then\nand t.tenant_id = #TenantId#\n$endif\n$if Ip != \'\' then\nand t.ip = #Ip#\n$endif', '查询主机总数', '2021-02-23 09:26:14', '0');
INSERT INTO `service_sql` VALUES ('9', 'hostDao.GetHosts', 'select t.*,hg.name group_name from host t\nleft join host_group hg on t.group_id = hg.group_id\nwhere\nt.status_cd = \'0\'\n  $if HostId != \'\' then\nand t.host_id = #HostId#\n$endif\n$if Name != \'\' then\nand t.name = #Name#\n$endif\n$if GroupId != \'\' then\nand t.group_id = #GroupId#\n$endif\n$if TenantId != \'\' then\nand t.tenant_id = #TenantId#\n$endif\n$if Ip != \'\' then\nand t.ip = #Ip#\n$endif\n$if Row != 0 then\n	limit #Page#,#Row#\n$endif', '查询主机', '2021-02-23 09:29:52', '0');
INSERT INTO `service_sql` VALUES ('10', 'hostDao.SaveHost', 'insert into host(host_id, group_id, name, ip, username, passwd, cpu, mem, disk, tenant_id) VALUES \n(#HostId#, #GroupId#, #Name#, #Ip#, #Username#, #Passwd#, #Cpu#, #Mem#, #Disk#, #TenantId#)', '保存主机', '2021-02-23 09:33:08', '0');
INSERT INTO `service_sql` VALUES ('11', 'hostDao.UpdateHost', 'update host t set\n                  $if Name != \'\' then\n                  t.name=#Name#,\n                  $endif\n                  $if Ip != \'\' then\n                  t.ip = #Ip#,\n                   $endif\n                  $if GroupId != \'\' then\n                  t.group_id = #GroupId#,\n                   $endif\n                  $if Passwd != \'\' then\n                  t.passwd = #Passwd#,\n                   $endif\n                  $if Cpu != \'\' then\n                  t.cpu=#Cpu#,\n                   $endif\n                  $if Disk != \'\' then\n                  t.disk = #Disk#,\n                   $endif\n                  $if Mem != \'\' then\n                  t.mem = #Mem#,\n                  $endif\n                  t.status_cd = \'0\'\nwhere t.host_id = #HostId#', '修改主机', '2021-02-23 09:37:32', '0');
INSERT INTO `service_sql` VALUES ('12', 'hostDao.DeleteHost', 'update host t set     \n                  t.status_cd = \'1\'\nwhere t.host_id = #HostId#', '删除主机', '2021-02-23 09:38:15', '0');
INSERT INTO `service_sql` VALUES ('13', 'userDao.SaveUser', 'insert into u_user(user_id, username, real_name, passwd, phone, email, tenant_id,user_role)\nVALUES(#UserId#, #Username#, #RealName#, #Passwd#, #Phone#, #Email#, #TenantId#,#UserRole#)', '添加用户', '2021-02-23 21:38:33', '0');
INSERT INTO `service_sql` VALUES ('14', 'userDao.SaveUserPrivilege', 'insert into privilege_user(p_id, privilege_flag, user_id, tenant_id)\nvalues(#Pid#, #PrivilegeFlag#, #UserId#, #TenantId#)', '保存用户权限', '2021-02-23 21:47:38', '0');

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `task_id` varchar(64) NOT NULL ,
  `task_name` varchar(100) NOT NULL ,
  `template_id` varchar(64) NOT NULL ,
  `task_cron` varchar(50) NOT NULL ,
  `state` varchar(12) NOT NULL DEFAULT '001' ,
  `status_cd` varchar(12) NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `tenant_id` varchar(64) NOT NULL ,
  `host_id` varchar(64) NOT NULL ,
  `notice_type` varchar(12) NOT NULL DEFAULT '2002',
  PRIMARY KEY (`task_id`)
);

-- ----------------------------
-- Records of task
-- ----------------------------
INSERT INTO `task` VALUES ('37054151-7bd0-4ace-9491-d7b9b9ba07d6', '进程监控1', '1', '0 */1 * * * ?', '001', '0', '2021-03-10 22:19:06', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '2002');
INSERT INTO `task` VALUES ('73d2eb7a-66d3-4c10-9d43-28919d67e859', 'sdfsdf', '1', 'sdf', '001', '1', '2021-03-10 21:56:23', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '2002');
INSERT INTO `task` VALUES ('c4bdca44-ea2f-40a3-9bc5-0ac7d729150b', '测试', '1', '123', '001', '1', '2021-03-10 21:46:38', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '2002');
INSERT INTO `task` VALUES ('d0088142-af06-41a7-af6f-3c335cb263ce', 'ggg', '1', 'ggg', '001', '1', '2021-03-10 21:59:00', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '2002');
INSERT INTO `task` VALUES ('ead5b80e-0255-4595-a045-a9a6564047a1', '测试主机', '1', '0 */1 * * * ?', '001', '0', '2021-03-11 22:18:11', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '3003');
INSERT INTO `task` VALUES ('eeefb52e-fcd4-4851-b33a-788932be22d3', '进程监控', '1', '0 */5 * * * ?', '001', '1', '2021-03-10 20:15:15', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '2002');

-- ----------------------------
-- Table structure for task_attr
-- ----------------------------
DROP TABLE IF EXISTS `task_attr`;
CREATE TABLE `task_attr` (
  `task_id` varchar(64) NOT NULL ,
  `attr_id` varchar(64) NOT NULL ,
  `spec_cd` varchar(12) NOT NULL ,
  `value` varchar(50) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(12) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`task_id`)
) ;

-- ----------------------------
-- Records of task_attr
-- ----------------------------
INSERT INTO `task_attr` VALUES ('37054151-7bd0-4ace-9491-d7b9b9ba07d6', '6b54c2be-1108-498d-a896-5578de14f0d2', '100101', 'Things.jar', '2021-03-10 22:19:06', '0');
INSERT INTO `task_attr` VALUES ('d0088142-af06-41a7-af6f-3c335cb263ce', 'a37e2f32-3145-4402-8a37-0d13d519cb69', '100101', 'ggg', '2021-03-10 22:02:56', '0');
INSERT INTO `task_attr` VALUES ('ead5b80e-0255-4595-a045-a9a6564047a1', 'd31ab9ee-cb33-4192-9037-bb247dfd8bfc', '100101', 'wuxuewen', '2021-03-11 22:18:11', '0');

-- ----------------------------
-- Table structure for task_template
-- ----------------------------
DROP TABLE IF EXISTS `task_template`;
CREATE TABLE `task_template` (
  `template_id` varchar(64) NOT NULL ,
  `template_name` varchar(100) NOT NULL ,
  `template_desc` varchar(200) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(12) NOT NULL DEFAULT '0' ,
  `class_bean` varchar(200) NOT NULL 
);

-- ----------------------------
-- Records of task_template
-- ----------------------------
INSERT INTO `task_template` VALUES ('1', '进程监控', '进程监控', '2021-03-09 22:33:09', '0', 'CheckProcess');

-- ----------------------------
-- Table structure for task_template_spec
-- ----------------------------
DROP TABLE IF EXISTS `task_template_spec`;
CREATE TABLE `task_template_spec` (
  `spec_id` varchar(64) NOT NULL,
  `template_id` varchar(64) NOT NULL,
  `spec_cd` varchar(12) NOT NULL ,
  `spec_name` varchar(100) NOT NULL ,
  `spec_desc` varchar(200) NOT NULL ,
  `is_show` varchar(12) NOT NULL DEFAULT 'T' ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(12) NOT NULL DEFAULT '0'
);

-- ----------------------------
-- Records of task_template_spec
-- ----------------------------
INSERT INTO `task_template_spec` VALUES ('1', '1', '100101', '进程名称', '必填，请填写进程名称', 'Y', '2021-03-09 22:34:38', '0');

-- ----------------------------
-- Table structure for tenant
-- ----------------------------
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant` (
  `tenant_id` varchar(64) NOT NULL ,
  `tenant_name` varchar(128) NOT NULL ,
  `tenant_type` varchar(12) NOT NULL DEFAULT '002',
  `address` varchar(256) NOT NULL ,
  `person_name` varchar(64) NOT NULL ,
  `phone` varchar(11) NOT NULL,
  `state` varchar(12) NOT NULL DEFAULT '2002' ,
  `remark` varchar(512) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`tenant_id`)
);

-- ----------------------------
-- Records of tenant
-- ----------------------------
INSERT INTO `tenant` VALUES ('13ba9db7-fcdd-496a-a9b0-abfd00bf24e3', '永梅', '002', '青海省', '王永梅', '18909711442', '2002', '永梅', '2021-02-23 21:54:20', '0');
INSERT INTO `tenant` VALUES ('2d4451e6-83ea-4028-99d4-6f5a8e1d4864', '学文', '002', '青海省', '学文', '17797173942', '2002', '', '2021-02-23 22:23:31', '0');
INSERT INTO `tenant` VALUES ('512c369e-0642-41e5-9ea6-5fe737984ae6', '梓豪平台管理员租户', '001', '青海省西宁市城中区', '梓豪', '18909711443', '2002', '管理员', '2021-02-19 03:28:29', '0');

-- ----------------------------
-- Table structure for tenant_setting
-- ----------------------------
DROP TABLE IF EXISTS `tenant_setting`;
CREATE TABLE `tenant_setting` (
  `setting_id` varchar(64) NOT NULL,
  `tenant_id` varchar(64) NOT NULL,
  `spec_cd` varchar(12) NOT NULL,
  `value` varchar(1024) NOT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`setting_id`)
) ;

-- ----------------------------
-- Records of tenant_setting
-- ----------------------------
INSERT INTO `tenant_setting` VALUES ('95c577fc-9ba3-4e19-a21e-14738d7b7161', '512c369e-0642-41e5-9ea6-5fe737984ae6', '300301', 'https://oapi.dingtalk.com/robot/send?access_token=b6de45c3ef202f1e34930547427efef3be1c3bb295f52c1855c76998142be6a5', '2021-03-07 15:25:43', '0');
INSERT INTO `tenant_setting` VALUES ('db0be22c-7c1f-4b50-a299-6c1c4094654d', '512c369e-0642-41e5-9ea6-5fe737984ae6', '300302', 'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=4a3ed4d9-f0b2-4827-b652-c6479bc8b125', '2021-03-11 21:42:58', '0');

-- ----------------------------
-- Table structure for t_dict
-- ----------------------------
DROP TABLE IF EXISTS `t_dict`;
CREATE TABLE `t_dict` (
  `id` int(11) NOT NULL,
  `status_cd` varchar(64) NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(200) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `table_name` varchar(50) DEFAULT NULL,
  `table_columns` varchar(50) DEFAULT NULL ,
  PRIMARY KEY (`id`)
);

-- ----------------------------
-- Records of t_dict
-- ----------------------------
INSERT INTO `t_dict` VALUES ('1', '3301', '运行中', '停止中', '2021-02-24 22:48:12', 'monitor_host_group', 'state');
INSERT INTO `t_dict` VALUES ('2', '3302', '已停止', '已停止', '2021-02-24 22:48:42', 'monitor_host_group', 'state');
INSERT INTO `t_dict` VALUES ('3', '1001', '短信通知', '短信通知', '2021-02-24 22:49:16', 'monitor_host_group', 'notice_type');
INSERT INTO `t_dict` VALUES ('4', '2002', '钉钉通知', '钉钉通知', '2021-02-24 22:49:33', 'monitor_host_group', 'notice_type');
INSERT INTO `t_dict` VALUES ('5', '300301', '钉钉地址', '钉钉地址', '2021-02-24 22:49:33', 'tenant_setting', 'spec_cd');
INSERT INTO `t_dict` VALUES ('6', '300302', '企业微信地址', '企业微信地址', '2021-02-24 22:49:33', 'tenant_setting', 'spec_cd');
INSERT INTO `t_dict` VALUES ('7', '3003', '企业微信通知', '企业微信通知', '2021-02-24 22:49:33', 'monitor_host_group', 'notice_type');
INSERT INTO `t_dict` VALUES ('8', '001', '数据库', '数据库', '2021-02-24 22:49:33', 'app_service', 'as_type');
INSERT INTO `t_dict` VALUES ('9', '002', '缓存', '缓存', '2021-02-24 22:49:33', 'app_service', 'as_type');
INSERT INTO `t_dict` VALUES ('10', '003', '计算应用', '计算应用', '2021-02-24 22:49:33', 'app_service', 'as_type');
INSERT INTO `t_dict` VALUES ('11', '10012', '停止', '停止', '2021-02-24 22:49:33', 'app_service', 'state');
INSERT INTO `t_dict` VALUES ('12', '10013', '启动', '启动', '2021-02-24 22:49:33', 'app_service', 'state');
INSERT INTO `t_dict` VALUES ('13', '10011', '升级', '升级', '2021-02-24 22:49:33', 'app_service', 'state');
INSERT INTO `t_dict` VALUES ('14', '1001', '未构建', '未构建', '2021-02-24 22:49:33', 'app_version_job', 'state');
INSERT INTO `t_dict` VALUES ('15', '2002', '构建中', '构建中', '2021-02-24 22:49:33', 'app_version_job', 'state');
INSERT INTO `t_dict` VALUES ('16', '3003', '构建失败', '构建失败', '2021-02-24 22:49:33', 'app_version_job', 'state');
INSERT INTO `t_dict` VALUES ('17', '4004', '构建成功', '构建成功', '2021-02-24 22:49:33', 'app_version_job', 'state');

-- ----------------------------
-- Table structure for u_user
-- ----------------------------
DROP TABLE IF EXISTS `u_user`;
CREATE TABLE `u_user` (
  `user_id` varchar(64) NOT NULL,
  `username` varchar(64) NOT NULL,
  `real_name` varchar(64) NOT NULL,
  `passwd` varchar(128) NOT NULL,
  `sex` int(11) NOT NULL DEFAULT '1',
  `phone` varchar(11) NOT NULL,
  `email` varchar(64) DEFAULT NULL,
  `state` varchar(12) NOT NULL DEFAULT '100201',
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  `tenant_id` varchar(64) NOT NULL ,
  `user_role` varchar(12) NOT NULL,
  PRIMARY KEY (`user_id`)
);

-- ----------------------------
-- Records of u_user
-- ----------------------------
INSERT INTO `u_user` VALUES ('037eaaa6-1891-49fa-94d2-b9c3d7a96451', 'xuewen', '学文', '9a55274928b959703835d8e6de1842f6', '1', '17797173942', 'xuewen@zihao.com', '100201', '2021-02-23 22:23:31', '0', '2d4451e6-83ea-4028-99d4-6f5a8e1d4864', '1001');
INSERT INTO `u_user` VALUES ('1', 'zihao', 'zihao', '25d346f8e979038a43a37e382524fe16', '1', '18909711443', '928255095@qq.com', '100201', '2021-02-18 18:20:05', '0', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1001');
INSERT INTO `u_user` VALUES ('b8a0776d-521b-4393-8fe8-ae81b2587a05', 'yongmei', '王永梅', 'e492c0c2fe55de16b05445668d0ec8e5', '1', '18909711443', 'yongmei@zihao.com', '100201', '2021-02-23 21:54:20', '0', '13ba9db7-fcdd-496a-a9b0-abfd00bf24e3', '1001');


create table business_package(
	`id` varchar(64) PRIMARY KEY  NOT NULL,
  `name` varchar(64) NOT NULL ,
  `varsion` varchar(32) NOT NULL,
  `path` varchar(128) NOT NULL,
	 create_user_id varchar(64) not null,
  `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `tenant_id` varchar(64) NOT NULL
);

create table business_dockerfile(
    `id` varchar(64) PRIMARY KEY  NOT NULL,
    `name` varchar(64) NOT NULL ,
    `version` varchar(32) NOT NULL,
    `dockerfile` longtext NOT NULL,
    create_user_id varchar(64) not null,
    `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
    `status_cd` varchar(2) NOT NULL DEFAULT '0',
    `tenant_id` varchar(64) NOT NULL
);
alter table business_dockerfile add log_path varchar(512);

CREATE TABLE business_images(
     `id` varchar(64) PRIMARY KEY  NOT NULL,
     `name` varchar(64) NOT NULL ,
     `version` varchar(32) NOT NULL,
     images_type varchar(12) not null,
     type_url varchar(512) not null,
     images_flag varchar(12) not null,
     create_user_id varchar(64) not null,
     `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')),
     `status_cd` varchar(2) NOT NULL DEFAULT '0',
     `tenant_id` varchar(64) NOT NULL
 );

CREATE TABLE `app_service_var` (
           `av_id` varchar(64) NOT NULL ,
           `as_id` varchar(64) NOT NULL ,
           `tenant_id` varchar(64) NOT NULL,
           `var_spec` varchar(64) NOT NULL ,
           `var_name` varchar(128) NOT NULL ,
           `var_value` varchar(512) NOT NULL ,
           `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
           `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
           PRIMARY KEY (`av_id`)
);

CREATE TABLE `app_service_hosts` (
         `hosts_id` varchar(64) NOT NULL ,
         `as_id` varchar(64) NOT NULL ,
         `tenant_id` varchar(64) NOT NULL,
         `hostname` varchar(128) NOT NULL ,
         `ip` varchar(64) NOT NULL ,
         `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
         `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
         PRIMARY KEY (`hosts_id`)
);

CREATE TABLE `app_service_dir` (
       `dir_id` varchar(64) NOT NULL ,
       `as_id` varchar(64) NOT NULL ,
       `tenant_id` varchar(64) NOT NULL,
       `src_dir` integer NOT NULL ,
       `target_dir` integer NOT NULL ,
       `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
       `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
       PRIMARY KEY (`dir_id`)
)

CREATE TABLE `app_service_port` (
        `port_id` varchar(64) NOT NULL ,
        `as_id` varchar(64) NOT NULL ,
        `tenant_id` varchar(64) NOT NULL,
        `src_port` integer NOT NULL ,
        `target_port` integer NOT NULL ,
        `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
        `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
        PRIMARY KEY (`port_id`)
)

CREATE TABLE `app_service_container` (
        `container_id` varchar(64) NOT NULL ,
        `as_id` varchar(64) NOT NULL ,
        `tenant_id` varchar(64) NOT NULL,
        `host_id` varchar(64) NOT NULL ,
        `docker_container_id` varchar(64) NOT NULL ,
        `state` varchar(12) not null ,
        `message` varchar(512) ,
        `update_time` datetime not null,
        `create_time` timestamp NOT NULL DEFAULT (datetime('now', 'localtime')) ,
        `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
        PRIMARY KEY (`container_id`)
)

create table app_version_job_images(
       job_images_id varchar(64) not null,
       tenant_id varchar(64) not null,
       package_url varchar(512) not null ,
       business_package_id varchar(64) not null ,
       business_dockerfile_id varchar(64) not null,
       create_time timestamp not null DEFAULT (datetime('now', 'localtime')) ,
       status_cd varchar(2) not null DEFAULT '0',
       job_id varchar(64) not null
);

create table faster_deploy
(
    deploy_id          varchar(64)  not null,
    app_name        varchar(128) not null,
    deploy_type        varchar(12)  not null,
    tenant_id      varchar(64)  not null,
    package_id        varchar(64) not null,
    shell_package_id          varchar(64) not null,
    create_time    timestamp    default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2)   default 0 not null,
    as_group_id    VARCHAR(64) not null,
    as_deploy_type VARCHAR(64) not null,
    as_deploy_id   VARCHAR(64) not null,
    open_port varchar(64)
);

create table db_link
(
    id        varchar(64)  not null primary key,
    name      varchar(10)  not null,
    ip        varchar(128)  not null,
    port      varchar(12)  not null,
    username  varchar(64)  not null,
    password  varchar(128) not NULL,
    db_name   varchar(64) not null,
    tenant_id varchar(64) not null ,
    create_user_id varchar(64) not null,
    create_time timestamp    default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2)   default '0' not null
);
create table log_trace
(
    id          varchar(64) not null
        primary key,
    name        varchar(64) not null,
    parent_id   varchar(64) not null,
    trace_id    varchar(64) not null,
    timestamp   varchar(20) not null,
    duration varchar(20) not null,
    service_name varchar(64) not null,
    ip           varchar(64) not null,
    port         varchar(20) not null,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);

create table log_trace_annotations
(
    id           varchar(64) not null
        primary key,
    span_id      varchar(64) not null,
    service_name varchar(64) not null,
    ip           varchar(64) not null,
    port         varchar(20) not null,
    value        varchar(12) not null,
    timestamp   varchar(20) not null,
    create_time  timestamp  default (datetime('now', 'localtime')) not null,
    status_cd    varchar(2) default '0' not null
);

create table log_trace_param
(
    id          varchar(64) not null
        primary key,
    span_id     varchar(64) not null,
    req_header  logtext,
    req_param   logtext,
    res_header   logtext,
    res_param   logtext,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);
create table install_app
(
    app_id         varchar(64) not null
        primary key,
    app_name       varchar(64) not null,
    version        varchar(64) not null,
    ext_app_id     varchar(64) not null,
    create_user_id varchar(64) not null,
    tenant_id varchar(64) not null ,
    create_time    timestamp  default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2) default '0' not null
);

create table app_publisher
(
    publisher_id varchar(64)  not null
        primary key,
    username     varchar(256) not null,
    email        varchar(64)  not null,
    token        varchar(128) not null,
    phone        varchar(11)  not null,
    state        varchar(12) default '001' not null,
    create_time  timestamp   default (datetime('now', 'localtime')) not null,
    status_cd    varchar(2)  default '0' not null,
    tenant_id    varchar(64)  not null,
    ext_publisher_id varchar(64) not null
);

create table business_images_ext
(
    id             varchar(64)  not null
        primary key,
    images_id           varchar(64)  not null,
    app_id        varchar(64)  not null,
    app_name    varchar(128)  not null,
    ext_images_id       varchar(64) not null,
    ext_publisher_id    varchar(64)  not null,
    create_time    timestamp  default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2) default '0' not null,
    tenant_id      varchar(64)  not null
);

create table resources_ftp
(
    ftp_id        varchar(64)  primary key  not null,
    name           varchar(64)    not null,
    ip             varchar(128)   not null,
    port             varchar(64)        not null,
    username       varchar(64)    not null,
    passwd         varchar(64)    not null,
    tenant_id      varchar(64)    not null,
    path      varchar(128)    not null,
    create_time    timestamp   default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2)  default '0' not null
);

create table resources_oss
(
    oss_id        varchar(64)  primary key  not null,
    name           varchar(64)    not null,
    oss_type       varchar(12) not null,
    bucket             varchar(128)   not null,
    access_key_secret             varchar(128)        not null,
    access_key_id       varchar(128)    not null,
    endpoint         varchar(128)    not null,
    tenant_id      varchar(64)    not null,
    path      varchar(128)    not null,
    create_time    timestamp   default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2)  default '0' not null
);

create table resources_db
(
    db_id             varchar(64)  not null primary key,
    name           varchar(64)    not null,
    ip             varchar(128) not null,
    port           varchar(12)  not null,
    username       varchar(64)  not null,
    password       varchar(128) not null,
    db_name        varchar(64)  not null,
    tenant_id      varchar(64)    not null,
    create_time    timestamp   default (datetime('now', 'localtime')) not null,
    status_cd      varchar(2)  default '0' not null
);

create table resources_backup
(
    id          varchar(64) not null
        primary key,
    name        varchar(64) not null,
    exec_time   varchar(64) not null,
    type_cd     varchar(12) not null,
    src_id      varchar(64) not null,
    src_object  longtext    not null,
    target_type_cd   varchar(12) not null,
    target_id   varchar(64) not null,
    tenant_id   varchar(64) not null,
    state       varchar(12) not null,
    back_time timestamp  not null,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);

create table business_images_ver
(
    id          varchar(64)  not null
        primary key,
    images_id   varchar(64)  not null,
    version     varchar(32)  not null,
    type_url    varchar(512) not null,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null,
    tenant_id   varchar(64)  not null
);

create table log_trace_db
(
    id          varchar(64) not null
        primary key,
    span_id     varchar(64) not null,
    db_sql  logtext,
    param   logtext,
    duration   varchar(64),
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);

create table waf
(
    waf_id           varchar(64) not null
        primary key,
    waf_name varchar(64) not null,
    port         varchar(64) not null,
    https_port varchar(64) not null,
    state varchar(12) not null,
    create_time  timestamp  default (datetime('now', 'localtime')) not null,
    status_cd    varchar(2) default '0' not null
);
insert into waf values (1,'waf',80,443,1001,'2020-01-01','0');
create table waf_hosts
(
    waf_host_id           varchar(64) not null
        primary key,
    waf_id varchar(64) not null,
    host_id         varchar(64) not null,
    create_time  timestamp  default (datetime('now', 'localtime')) not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_route
(
    route_id    varchar(64)  not null
        primary key,
    waf_id      varchar(64)  not null,
    scheme varchar(64) not null,
    hostname    varchar(128) not null,
    ip          varchar(128) not null,
    port        varchar(64)  not null,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);

create table waf_access_log
(
    request_id           varchar(64) not null
        primary key,
    waf_ip varchar(64) not null,
    host_id           varchar(64) not null,
    x_real_ip         varchar(64) not null,
    scheme        varchar(64) not null,
    response_code    varchar(64) not null,
    method    varchar(64) not null,
    http_host    varchar(64) not null,
    upstream_addr varchar(64) not null ,
    url    varchar(512) not null,
    request_length varchar(64) not null,
    response_length varchar(64) not null,
    state varchar(12) not null,
    message varchar(512) not null,
    create_time  timestamp  default (datetime('now', 'localtime')) not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_hostname_cert
(
    cert_id    varchar(64)  not null primary key,
    hostname    varchar(128) not null,
    cert_content      longtext  not null,
    priv_key_content     longtext  not null,
    create_time timestamp  default (datetime('now', 'localtime')) not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904037', 'waf监控', '800201904006', '/waf/index.html', '4', 'waf监控', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904034', 'waf监控', 'waf访问', '2019-04-01 02:24:53', '0', '/waf/waf.html', '700201904037');
INSERT INTO `privilege_rel` VALUES ('64', '500201904034', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('65', '500201904034', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `t_dict` VALUES ('35', 'sqli', 'SQL注入', 'SQL注入', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('36', 'xss', '跨站脚本', '跨站脚本', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('20', 'webshell', 'WebShell攻击', 'WebShell攻击', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('21', 'robot', '恶意爬虫', '恶意爬虫', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('22', 'cmdi', '命令注入', '命令注入', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('23', 'rfi', '远程文件', '远程文件', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('24', 'lfi', '本地文件包含', '本地文件包含', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('25', 'illegal', '非法请求', '非法请求', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('26', 'vuln', '漏洞攻击', '漏洞攻击', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('27', 'cc', '命中CC防护', '命中CC防护', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('28', 'custom_custom', '命中精准防护', '命中精准防护', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('29', 'custom_whiteip', '命中IP黑白名单', '命中IP黑白名单', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('30', 'custom_geoip', '命中地理位置控制', '命中地理位置控制', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('31', 'antitamper', '命中网页防篡改', '命中网页防篡改', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('32', 'anticrawler', '命中JS挑战反爬虫', '命中JS挑战反爬虫', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('33', 'leakage', '命中敏感信息泄露', '命中敏感信息泄露', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('34', 'followed_action', '攻击惩罚', '攻击惩罚', '2021-02-24 22:49:33', 'waf_access_log', 'state');

create table waf_rule_group
(
    group_id      varchar(64) not null
        primary key,
    group_name    varchar(64) not null,
    state varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table waf_rule
(
    rule_id     varchar(64) not null
        primary key,
    group_id varchar(64) not null,
    rule_name   varchar(64) not null,
    scope       varchar(64) not null,
    obj_id      varchar(64) not null,
    obj_type    varchar(64) not null,
    seq         int         not null,
    state       varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

INSERT INTO `menu` VALUES ('700201904038', '安全组', '800201904006', '/index.html#/pages/admin/wafRuleGroupManage', '5', '安全组', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904035', '安全组', '安全组', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafRuleGroupManage', '700201904038');
INSERT INTO `privilege_rel` VALUES ('66', '500201904035', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('67', '500201904035', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_ip_black_white
(
    id     varchar(64) not null
        primary key,
    type_cd    varchar(64) not null,
    ip   varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

INSERT INTO `menu` VALUES ('700201904039', '黑白名单', '800201904006', '/index.html#/pages/admin/wafIpBlackWhiteManage', '5', '黑白名单', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904036', '黑白名单', '黑白名单', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafIpBlackWhiteManage', '700201904039');
INSERT INTO `privilege_rel` VALUES ('68', '500201904036', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('69', '500201904036', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_area
(
    id          varchar(64) not null
        primary key,
    type_cd     varchar(64) not null,
    area_name          longtext not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904040', '地理位置', '800201904006', '/index.html#/pages/admin/wafAreaManage', '6', '地理位置', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904040', '地理位置', '地理位置', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafAreaManage', '700201904040');
INSERT INTO `privilege_rel` VALUES ('70', '500201904040', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('71', '500201904040', '600201904002', '2019-04-01 08:18:29', '0');


create table waf_cc
(
    id          varchar(64) not null
        primary key,
    path     varchar(512) not null,
    visit_count   varchar(12)    not null,
    visit_sec   varchar(12)    not null,
    block_sec   varchar(12)    not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904041', 'CC防护', '800201904006', '/index.html#/pages/admin/wafCCManage', '7', 'CC防护', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904041', '地理位置', '地理位置', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafCCManage', '700201904041');
INSERT INTO `privilege_rel` VALUES ('72', '500201904041', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('73', '500201904041', '600201904002', '2019-04-01 08:18:29', '0');


create table waf_accurate
(
    id          varchar(64)  not null
        primary key,
    action varchar(12) not null ,
    type_cd        varchar(12) not null,
    include varchar(512)  not null,
    include_value   varchar(512)  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

INSERT INTO `menu` VALUES ('700201904042', '精准防护', '800201904006', '/index.html#/pages/admin/wafAccurateManage', '7', 'CC防护', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904042', '精准防护', '精准防护', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafAccurateManage', '700201904042');
INSERT INTO `privilege_rel` VALUES ('74', '500201904042', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('75', '500201904042', '600201904002', '2019-04-01 08:18:29', '0');
drop table vpn;
create table inner_net
(
    inner_net_id      varchar(64) not null
        primary key,
    inner_net_port    varchar(64) not null,
    tun        varchar(64) not null,
    tun_name  varchar(64) not null,
    dns  varchar(64) not null,
    protocol  varchar(64) not null,
    state       varchar(12) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

insert into inner_net(inner_net_id, inner_net_port, tun, tun_name, dns, protocol, state) VALUES ('1','5555','192.168.1.0/24','virName','8.8.8.8','tcp','2002');
create table inner_net_hosts
(
    inner_net_host_id varchar(64) not null
        primary key,
    inner_net_id      varchar(64) not null,
    host_id     varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table inner_net_users
(
    user_id varchar(64) not null
        primary key,
    username      varchar(64) not null,
    password     varchar(64) not null,
    tel varchar(64) not null,
    ip     varchar(64) not null,
    login_time datetime not null ,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu_group` VALUES ('800201904014', '虚拟组网', 'fa fa-globe', '', '8', '虚拟组网', '2019-04-01 07:55:51', '0', 'P_WEB');


INSERT INTO `menu` VALUES ('700201904043', '中心节点', '800201904014', '/index.html#/pages/admin/innerNetManage', '1', '中心节点', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904043', '中心节点', '中心节点', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetManage', '700201904043');
INSERT INTO `privilege_rel` VALUES ('76', '500201904043', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('77', '500201904043', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `menu` VALUES ('700201904044', '组网电脑', '800201904014', '/index.html#/pages/admin/innerNetUserManage', '1', '组网电脑', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904044', '组网电脑', '组网电脑', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetUserManage', '700201904044');
INSERT INTO `privilege_rel` VALUES ('78', '500201904044', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('79', '500201904044', '600201904002', '2019-04-01 08:18:29', '0');

create table dns
(
    dns_id   varchar(64) not null
        primary key,
    dns_port varchar(64) not null,
    dns_ip            varchar(64) not null,
    state          varchar(12) not null,
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2) default '0' not null
);

create table dns_hosts
(
    dns_host_id varchar(64) not null
        primary key,
    dns_id      varchar(64) not null,
    host_id     varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table dns_map
(
    dns_map_id     varchar(64) not null
        primary key,
    host    varchar(64) not null,
    type    varchar(64) not null,
    value         varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


insert into dns(dns_id, dns_ip, dns_port, state) VALUES ('1','8.8.8.8','53','2002');

INSERT INTO `menu` VALUES ('700201904045', 'dns', '800201904014', '/index.html#/pages/admin/dnsManage', '1', 'dns', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904045', 'dns', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/dnsManage', '700201904045');
INSERT INTO `privilege_rel` VALUES ('80', '500201904045', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('81', '500201904045', '600201904002', '2019-04-01 08:18:29', '0');