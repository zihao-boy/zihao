
CREATE TABLE `app_service` (
  `as_id` varchar(64) NOT NULL ,
  `as_name` varchar(128) NOT NULL ,
  `as_type` varchar(12) NOT NULL,
  `tenant_id` varchar(64) NOT NULL ,
  `as_desc` varchar(512) DEFAULT NULL ,
  `state` varchar(12) NOT NULL DEFAULT '10012' ,
  `as_count` int(11) NOT NULL DEFAULT '1' ,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`as_id`)
);

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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `job_shell` longtext NOT NULL,
  `tenant_id` varchar(64) NOT NULL ,
  `pre_job_time` datetime DEFAULT NULL,
  `cur_job_time` datetime DEFAULT NULL,
  `state` varchar(12) NOT NULL DEFAULT '10012' ,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`host_id`)
);

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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `id` int(11) NOT NULL,
  `domain` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `zkeys` varchar(100) NOT NULL,
  `value` varchar(1000) NOT NULL,
  `remark` varchar(200) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`id`)
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
INSERT INTO `monitor_event` VALUES ('02459493-59c8-487e-8c5b-f34db075c42a', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:07:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('02c282f3-acfa-4450-8fe3-728c991a8cf5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-07 09:09:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('03303857-fdaa-43a3-a845-191b3600f895', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:37:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0422b166-2116-4582-a583-bb828f4fe7d7', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:19:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('04739287-2a5f-42b8-809c-17e207244ea9', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-07 09:02:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('04e630d0-a045-4b54-8fb3-a4b34dd8dbe2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:56:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('06a51d96-b439-4fac-8680-f348f3c20417', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '14.60', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 14.6;', '2021-03-07 09:13:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('072797a3-99ed-4252-816b-a6c684c28f02', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:45:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('085d4a01-890f-4f69-a6dc-144f7ce342a2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:25:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('08e25c35-4a91-4d96-a774-5b7264ed9a78', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:13:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('094f4510-5339-4c88-89ab-22e3bbcfc2d6', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:07:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('099d0948-781d-4bc3-b656-e2323f817d4e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:30:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0aa329f8-714e-44ae-ad90-688ee104addb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:19:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0c26cdb7-6c82-4264-919a-53073dda343d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:28:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0c97588c-4484-4dd7-afe1-7591768e3600', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:49:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0ce9153b-9ba1-440e-b2d6-9d7412f3c3ad', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:25:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0d6275c8-cd29-43d6-a6b3-4d28ffd992a0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:24:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0d856537-49e7-4f0b-a717-8af07aa2ed1b', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:35:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0ecd9e85-b353-4fd2-ba0f-04a0579393bf', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,mysql进程不存在', '2021-03-11 00:37:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('0f5c4ace-d8ba-4ee4-93b8-ade8e43aab79', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:01:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1025eb8e-b97b-4d6a-8770-7ad637ee9cf7', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:55:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('10bca914-8dfc-4f6f-8ea2-6ab5748a17c0', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:31:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('13ceead4-e1f3-47d1-a117-4f3bd9ec1361', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.10', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.1;', '2021-03-06 23:34:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('152b9121-98cc-4e4c-a3be-fa657767c8bc', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:30:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('15e8e90c-8c6b-499d-9f74-e72b4e7886cb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:24:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('162b5d79-4dc2-47eb-a085-d07345bc5fed', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.60', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.6;', '2021-03-06 23:43:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('16bdeff2-9c7e-45d6-9a26-c942313b9b95', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:58:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('18d66183-4d29-48a6-a9be-dc28c4f9c885', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:36:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('193d42a4-4e00-48f2-8350-61364f43fc7a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:20:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1a62833a-6e66-4575-a0d8-3880829b6531', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:40:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1af4c7e4-e756-419b-a068-9f7584fb233a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:45:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1b20f919-1981-4b7d-94b6-0df187fa5ea0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '27.90', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 27.9;', '2021-03-07 09:12:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1b4c64dd-3b6e-4e84-8a7f-2b8ac0b4f3f9', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 08:55:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1c48789c-c272-43da-bd44-8b7f5cd58cab', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 23:56:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1ce947ba-8e7c-48c2-872f-e6f5b5f2fc51', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:12:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1d5ac02e-6063-4502-83ab-32d6175bdc21', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:18:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1d9e8e9d-593c-43e1-a97b-6a4a9f9ce9b3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '13.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 13.6;', '2021-03-06 23:02:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('1f29d34b-5984-4977-812f-c6d118eb5f8b', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', 'cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 18:08:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('208ece4c-d491-4acc-aa7f-35221c5ab696', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:45:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('212a2b7f-9a7e-4f6a-90e3-57e69e929e0e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:55:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('21cdb40e-ed6b-4a11-aea6-91dcd4590b79', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:16:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2493d0c7-239c-4b0c-ae4d-1be7364cf0f5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 08:59:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('265c8648-50cf-447a-ae35-cc3a589ae8e0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', 'cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 18:10:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('286c03aa-7495-4c81-820f-f7d2e670dc24', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '1.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 1.2;', '2021-03-06 23:54:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('29657a30-36b2-4ba0-8d86-eecc39741b37', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 23:04:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('29957705-b274-461d-a62c-94befeadcb9d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:01:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2a3bee19-f0f1-4f23-aec8-9f33f558bb1c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.5;', '2021-03-06 22:36:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2b52c9b6-316e-421b-b047-70facbbe8c64', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:05:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2be777cc-8f62-44ce-938a-05e664df17e1', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 22:55:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2e6fd3f6-d2e7-463c-9123-158ade1cff11', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:24:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2e70f830-bd72-460f-bf1f-f26523544743', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-07 05:36:39', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2ee33c49-6b00-4451-bdd9-99336562daae', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 08:51:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2f894f33-2867-4c79-b35b-5dcc3a889039', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:10:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('2fb4be0a-b67b-4fc9-8d31-7a193ff9537c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:14:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('30bd9481-be6b-45cf-90d6-8ea6d858e595', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 23:47:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3100027f-2e97-430b-8e90-c930b7929590', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 23:55:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3137e6e7-86a6-4bd0-87af-96ad0e53062e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:41:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('326b0669-c40a-44b3-b76b-d848b027ab08', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '10.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 10;', '2021-03-06 23:35:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('38689194-9957-4ece-bc46-566d5533ee18', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:50:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('38ccff44-aac5-496b-bf35-eebdb806cc6d', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:39:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('38f1949e-c507-4124-b26a-13e3658a16f3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 22:50:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3930adb4-f92d-4c89-9752-ea22e7407836', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-07 08:49:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('39a78f7a-de48-4e7b-9014-e84c3da9f5a2', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:03:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3a40eed7-35b0-4353-8ef6-9d2ee35271f3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:29:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3afbd309-ea47-4413-b93f-ef63a75ad0b3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:53:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3b58f02e-7fcf-4bc9-8451-6ce68b4d3f65', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-07 09:16:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3ccbecf7-5fbf-427f-90c1-34e7ddcea633', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.70', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.7;', '2021-03-06 22:39:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3db92431-70dc-4005-aba3-9ca6dcaa7a16', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-07 08:52:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3ea43ed2-6ceb-4ff4-abad-35a6d0daa406', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.90', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.9;', '2021-03-06 23:41:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3ebf15b9-64ca-44be-89be-8e06b19f748a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:26:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('3fedd1b7-c1ca-42f4-b0a7-d5b802b16d30', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 08:53:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('40fa5137-6deb-4f81-ad2a-36ce6ec215d0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:32:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('4147171f-1439-46b0-9448-88822403bf4c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:08:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('42d79eb8-6b60-4988-932c-decb4fb3dfca', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:01:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('42e5caa9-51ce-4a75-b55a-dec5981b6ab7', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:57:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('44db2a55-4b22-4707-be7c-73ccf2f179eb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-07 05:36:37', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('46137116-e5f3-4ce2-af6c-6eb350cba385', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 22:43:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('471633c6-7ed8-4c98-92da-07490e39a0db', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 22:38:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('484b5ca8-6ae0-4210-b217-e00470dcf73c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 22:59:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('489e5337-bb8f-4c9f-be88-f89c44ad4719', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:42:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('48df3b0f-ff0e-4393-8165-f038118f5756', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '0.01', '主机：HC测试主机1 内存使用率告警：阀值 0.50, 当前 0.5351473922902494;', '2021-03-07 19:52:05', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('49dead9f-3dfd-491b-bb1f-ad4a96908294', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:46:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('49e0b513-9392-4215-907d-1d2ac48473a7', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:47:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('4a8b681a-a7ce-4e64-8a31-2cb53f7aa064', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:56:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('4cd75a94-cba4-4a56-a55d-7be3467cd970', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', 'cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 18:05:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('4d15b93d-9011-44a8-9459-736afafec065', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', 'cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 18:11:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('4f38accc-b730-42dc-9a33-2065997176e2', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 23:36:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('5076f478-f1af-4b4f-b044-afc24bc5c577', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:59:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('507fb2d3-36fe-45f8-a2b3-eede2f1dfb69', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 23:32:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('50bc1ed6-a964-43f1-9113-7d35a99e13e8', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 08:54:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('52a7d7f0-1784-40bf-b91e-5b081103b960', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:46:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('533135d6-1e13-488f-bf6b-9b0abfd67bee', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 08:57:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('5382e79d-dcc3-4a02-9842-6e063e102fb7', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 08:55:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('53ac0cf0-368b-46f0-86ea-266d93d247cf', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 22:48:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('55f4f8b6-b072-47a1-848c-b6a8c4430835', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:17:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('56347f9c-366c-42fc-8e2b-550630503948', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:59:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('566efecd-c5f8-48df-816f-147a8abc7a14', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-06 23:21:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('5cf95a9b-a9d2-49b0-ac01-9dc90f4eb8fb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 00:01:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('5dabf30a-4b73-427e-b351-299e7c45871d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.4;', '2021-03-06 23:35:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('5fad1616-77b5-47c4-a6d2-e0b7a73d8ea5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-07 08:57:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('61ae3cce-94af-4c5e-b67d-21274172e061', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.70', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.7;', '2021-03-06 23:41:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('62ff4653-0359-45ed-9d2c-ee8de7c7eea8', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,mysql进程不存在', '2021-03-11 00:36:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('63d9f565-b2a6-481e-a96d-2d590b7667d8', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', 'cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 18:07:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('63fa513c-588d-449b-8e1a-be43a0555553', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:24:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('64330adf-2efd-459a-83e4-5d5f30ac14b5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:09:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('644a4e98-a105-43a2-a13a-f08c61b6ddd8', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-06 22:25:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('65f84488-2527-4786-baa1-450c67892994', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:11:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('67072f20-d834-4a55-8a99-c1045d41a61f', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:23:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6892f3c0-bc9b-4829-b9b5-2a6aa0289e07', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:26:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('68b28761-9346-4c56-854f-2614186c9f87', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:27:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6910f35d-18e1-41b7-b865-de3a398f024e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:35:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('69a58b01-7cee-44e2-a7a5-093cdebec420', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', 'cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 18:03:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('69b552bd-9c22-40a5-beac-a2dd4144c07d', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.3;', '2021-03-07 09:04:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6a380539-4f35-4b89-b6f3-7e798240dc10', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:28:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6ca78c92-4095-4464-bbb4-afe08a1bbb4a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:34:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6cdfa795-fa84-4d32-b314-8561a3b9212c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:47:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('6e826e54-5a59-43c7-bf0e-833948e42ee3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.90', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.9;', '2021-03-06 23:53:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7022f36b-9e5c-4372-bbb5-994041310508', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 23:11:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('70874456-95bb-474f-96e8-b06e7d0724ae', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '7.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 7.2;', '2021-03-06 23:53:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('72784d88-cb3a-413e-a352-6e32ea804193', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:58:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7285f61d-5181-4553-b04c-ddb11cce9f34', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'xxx进程不存在', '2021-03-11 00:30:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7397060f-18b7-47ee-9ddd-5a2d272cc2ab', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.2;', '2021-03-06 23:22:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('742bda4e-4998-42aa-8cb0-10709d8c8cd8', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:48:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7459908f-d0b2-4c67-ac91-39f683f05e52', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 08:53:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('74ac4c1c-61ed-41f4-a633-42a96db9ed70', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,Things.jar进程不存在', '2021-03-11 00:40:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('74d158e0-2ff4-43dc-8260-1dd9f6af0ce3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:29:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7573ba07-dff6-4842-a59c-021e15f7eb1c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:11:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('778fa3fe-3997-4849-8bb0-533c3602aff2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 08:50:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('786e60fe-3bbe-4e6c-9a0e-06fbe8fdd0e5', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 23:42:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7aa12eb9-e409-49ba-bcd9-c1ad4cadee88', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:03:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7b3a3565-8400-4d61-89a7-15acad8fc1cd', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:35:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7d54e0ba-9717-40d3-b730-2c51f948bdaa', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:45:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7dd425f4-ee42-4e89-b584-884d45591f0b', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:25:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('7ecfe7fd-9183-4be9-b834-89d168068677', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:26:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('80402208-99cf-4d69-8339-ba228d004872', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:33:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('80e3bb82-805e-4161-a21d-303e578dc24a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 23:20:01', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('810a5822-316a-430a-b852-17950ad9849e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:15:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('81801592-afdc-4bb7-b391-4e2c26eb520f', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', 'cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 18:09:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('81cf5139-130c-43c3-b419-5c135c1d3dad', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', 'cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 18:08:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('82514733-93d0-498c-9f2e-1c28f2e1e05b', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:54:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('82c52b04-18a3-40c0-b134-b1a1360f7e6c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', 'cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 18:04:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('841793bd-f281-42bf-b096-8516327676a7', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:47:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('86b3da4b-2715-441b-b579-40ee6f433a5a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:27:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('86be714a-c59f-4ed2-9cd6-22b5a0ca8157', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 09:24:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('86ea471b-22a3-4253-9ae6-c7417e61ffe0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 09:09:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8a2252ac-30a1-49df-8ccd-a144540e0bc6', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-07 09:08:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8a5a1a44-b40d-4a5c-9993-2c168a294f87', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:54:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8ae06d15-245b-4765-bf37-9b5d964727ba', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:35:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8c2439d4-059a-4df1-a945-a7e5a6e2a28f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:13:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8ca6a16a-27dc-4066-921c-5ff660736960', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:30:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8e35ec85-c54a-4a0a-a423-18350833413c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 09:03:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8e65cadb-7eaa-4937-8363-d4ad2164cc22', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'mysql进程不存在', '2021-03-11 00:35:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('8ea04635-a9f3-4e14-97cd-60d18232842c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 22:42:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('93cae34e-b617-4160-ae1a-976f86d8288c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-07 08:59:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('93f2d909-b624-48c7-a614-5a95499bc2fe', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 23:56:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('948172f3-c848-4135-be77-056510b388e4', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:42:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('96a6f21a-85e0-4c1d-86f8-9ae9b3968bd2', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:50:04', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('98292bb2-1446-4135-b4b5-5a8040cabb99', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:33:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9b586ee7-4bc6-4d1b-a945-21b01379d44f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 08:50:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9c35e4c7-ff9e-4070-aa90-7315bed0ffbb', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 08:54:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9d1b9d2c-bffb-4b5f-a477-663027683648', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 22:22:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9de235f8-e183-4123-af46-88ae103246ec', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '7.10', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 7.1;', '2021-03-07 00:00:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9e179019-f871-4f1b-9390-48a2025796bc', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 22:34:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9e1f5234-3604-4d01-b4d4-61fde85f6d0b', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-06 23:23:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9e38d1ae-9e47-470b-bb97-838cd37140d8', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:46:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9e672219-2130-4c53-8fb5-86d073ade642', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.5;', '2021-03-06 23:08:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('9f3d3f20-b8a5-4308-9876-e81c74ea319e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:51:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a02b1ad3-b7ff-4e52-ba01-2a6ccfa8ae04', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '26.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 26.2;', '2021-03-06 22:51:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a0f4ec32-0457-4736-bd74-786005df29cb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 09:16:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a294f198-afd1-40e5-8d7e-e3a79c6ef8a1', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.5;', '2021-03-06 23:38:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a2a45d90-d881-4232-aa42-8e2fcb0c4634', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.3;', '2021-03-06 23:36:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a36a7d85-7de1-48a5-8efb-81fab30b309d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:33:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a387024a-530a-4978-a4e8-4fa0af4bae7c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.5;', '2021-03-06 23:39:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a40a1dee-cd54-43b1-9221-b26dec8b996f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 22:27:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a430ac28-652b-4ddf-8c6b-1492b8d50765', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:14:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a4660624-d500-44ed-b91e-26394e08a8fb', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:31:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a527b0d6-83da-4016-b04a-449e3360264c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.90', 'cpu使用率告警：阀值 0.70, 当前 5.9;', '2021-03-06 18:03:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a566d602-0b02-43a3-8256-61e003ee2540', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 09:05:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a5ff27ec-44bf-4760-a7eb-137e86743b96', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:37:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a666ff8c-4e35-4c07-bb79-fd4e3415b98d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 09:14:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a667ded7-5658-4db4-81af-b73a2590a786', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:46:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a69c4b9d-6cbc-4202-8a7b-8d96b1877bef', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:34:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a83e5581-8222-4bd6-b6c0-2cd6e8504d44', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-07 08:58:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a8d78e64-5119-4f85-94a5-89b9009e8f57', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.2;', '2021-03-06 22:28:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a8e06ae8-55a8-4967-9c86-3ffdb297dcc3', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.20', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.2;', '2021-03-06 23:29:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a9958e78-1028-4e6a-9e14-650aae850c12', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-07 09:15:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('a9aaa597-1143-4a29-91c5-8f2fc9006f8c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 23:19:04', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ab485d0e-55c0-4050-b782-d5213d92b889', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:28:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ac75a2b1-6604-4a6b-8dd5-b5c0e84256af', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.70', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.7;', '2021-03-06 22:37:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('adaac595-9ac4-40a3-a1dc-6f3ef7951f22', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '23.80', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 23.8;', '2021-03-06 23:52:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('add1c7cb-31c5-4e20-a435-94b60dd76ecd', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:43:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ade22e73-f573-447f-9359-291220f839c5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 22:23:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ae6c9432-a4af-45c6-b76e-14404bd4f3f3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'mysql进程不存在', '2021-03-11 00:33:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('aebc6e9e-92ef-42e2-a019-91c958738ef7', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:20:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('af6cba8a-6aa1-4b4e-bb2a-d3e774626411', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:40:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('afe42e01-e507-48cd-823d-b7ae8c8f44f5', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:08:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b0e74019-d3db-4df0-b5cc-58856584401d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:27:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b12fb769-1e87-4895-ab5a-16045138235a', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.80', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.8;', '2021-03-07 09:00:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b16aa61a-15f0-4e41-a397-f5a1d626b5c9', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:31:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b344011d-0f8a-4ba0-adda-a2b6ea09de8d', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.5;', '2021-03-06 22:34:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b3d51f07-255f-4fb9-a00e-b960a7b8b78e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '0.02', '主机：HC测试主机1 内存使用率告警：阀值 0.50, 当前 0.5352733686067019;', '2021-03-07 19:51:05', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b4446185-3f9f-485c-bab0-d02916ebaa2f', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:24:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b4830628-4b6a-4219-a102-d768d0c0850a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:37:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b4abd8e3-adde-4114-af96-7d7cd54780fb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-06 23:26:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b4c773f9-0bbd-4021-8e32-6f7909fdf267', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:35:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b50ad7d8-75ce-406e-a7a7-34f1f199a3fa', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 00:01:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b60c4f28-44eb-43e8-a562-19d59b5b1c2f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-06 23:30:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b670809c-a5d1-44d4-9da5-53378618eedb', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '27.90', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 27.9;', '2021-03-06 23:10:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b6aedf5a-14af-4d46-af51-a73aafa35f00', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 22:30:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b790a334-917d-46af-83be-28d64cbb1e97', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '9.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 9.3;', '2021-03-06 23:58:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b82008cb-8324-4527-a3f8-3f3a0007f3e5', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.5;', '2021-03-06 22:30:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b929e706-b115-4268-b4d0-f80b166ace9c', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-07 09:15:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b9324666-5138-4eb8-9f41-49ce9d38e9d3', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:19:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b995d659-8a9d-43f3-8849-d6da776d8792', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 23:09:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('b9cc5456-25b2-4c07-9331-b856f97f74e1', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 23:40:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('bc52f207-6432-40a9-a11e-85fff5a27642', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.80', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.8;', '2021-03-07 08:56:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('bd8b37eb-9408-4f8a-a2ac-c72be25a8233', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', 'cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 18:11:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('bfab0b60-0b5d-4356-bb27-b41231e89020', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.60', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.6;', '2021-03-06 22:41:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c02d2e61-e0f5-458a-81ed-d2a94fccb876', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.70', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.7;', '2021-03-06 23:11:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c1150857-6d89-4b7e-a1d5-3ea4c013cc82', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.4;', '2021-03-06 23:21:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c2057bc1-f0cf-4d71-8936-af6741dbcaa2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:50:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c5b82b4e-93e2-4cea-a54f-50547f9f0bd9', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:39:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c73e9699-d46f-46f1-92b3-dc86a83d7b45', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:34:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('c761705a-d339-4548-b44f-add5f995fca9', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 23:05:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('cb16424d-7036-4f9c-9c72-0b632326457d', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:31:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('cc75f629-7892-443a-9a87-a4809ef4803c', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:23:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ccd15bb1-e04a-4f8c-974d-d9c23cae80f2', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:31:03', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('cdbbed1c-46f8-40a9-b927-9b63141aee6b', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:20:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('cec62b4e-944c-48bc-b093-0d8ebe7ed45a', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:48:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d297ac51-9b51-468a-9fcc-9b32f35a7d14', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:20:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d3daaa3a-1043-43cc-9c8d-019dbb30f9bc', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 22:28:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d45acd9a-3825-47ec-8bb9-344157a9cbdb', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-06 22:26:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d46c3c49-ad75-478b-a491-fa08ae31665f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 22:24:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d599181b-82f2-48fb-8d23-0c6bc80e92d5', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.5;', '2021-03-06 23:49:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d724f04b-3246-44fa-bc08-9a93ddda2e50', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '53.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 53.3;', '2021-03-06 23:12:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d8c20ba8-273d-4158-bc7a-3a6fc020c23f', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:53:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d98fb97b-46a1-4c9b-b800-04d76cdd53a3', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 08:58:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('d9c8329e-c73b-4668-aa21-73437f731418', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-06 23:13:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dac80ecc-b1dd-497b-9c42-0ecd4c2746b8', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:33:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dafe2fb2-ea42-4bb4-909c-ad5e79d31197', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 22:29:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('db0cbacf-1b66-4553-86b3-f80015da9af8', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:17:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('db1b4003-f040-45c7-9fca-e416f874c80e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:24:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dbc116cb-d769-4a04-beaf-aa3121d34ce2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:20:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dc708fa3-f112-4d78-9a86-339d59a16877', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:26:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dd180598-5677-473c-a6a2-adec9d0087b2', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', 'cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 18:04:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('dda915ad-706d-439e-aff6-49c89864991e', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', 'cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 18:06:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ddfd0369-80ed-40bf-9844-1f3f5b6126f0', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:23:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('de8280a2-32f1-4827-86f4-1fdef0913aa3', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:49:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e022ebae-3fd2-4548-b2b6-f7b454dae5a0', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:27:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e04925ef-da87-4048-a60d-55fae34e1337', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 22:44:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e20f5302-a02e-4809-a895-abc27c7d2c40', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:06:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e26979ab-77f2-4a22-8eec-c7aa7697d547', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-07 09:06:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e2717e1f-ab85-4e24-af6d-728b5806b42f', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'mysql进程不存在', '2021-03-11 00:34:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e2da4e55-d49e-40b0-92ce-f519235d33f1', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 09:10:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e49f0d1f-5082-430b-b9bd-c9931720b5ea', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'mysql进程不存在', '2021-03-11 00:32:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e4a107ba-f447-475b-9f33-26f82764852e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', 'mysql1进程不存在', '2021-03-11 00:31:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e539613c-1a91-4a23-994a-3c9b3fd6b754', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:19:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e61ef6dd-546a-4db1-90d1-e8cb003aa8a9', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:59:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e735a7e9-a252-41a1-811d-e6da83e81062', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-07 09:12:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('e8491cb9-0c90-4fc8-b807-86d379d358b3', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:25:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ea842e41-8cf3-48cb-a72a-a7ce23a3ab36', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:32:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('eb5cf687-d62c-4bbe-9b96-a4975ecf65c3', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:07:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('eb7404a6-7a34-4770-8ef0-f7b75219171e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-07 09:00:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ebbab0b8-469c-4fc4-b022-bc8939c43db1', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-07 08:52:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ebe685a2-ea9b-4c81-9c28-fc2ec479ed96', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:29:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ebef0a89-b402-4761-a841-2263cc294393', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:32:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ec4e4c67-d816-4294-8e33-a6f0929e3b1b', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.80', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 4.8;', '2021-03-06 23:04:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ed10ac65-cd93-48d5-9844-b31213305f6a', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', 'cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 18:07:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ed5afe52-0583-4237-9017-588dd4bc2a16', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.60', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.6;', '2021-03-06 23:50:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ed89f0f7-de36-47fc-b64d-8e8e7f045e00', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:05:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('edecbb8e-07cb-44ef-9cee-a57adfc44987', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.30', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5.3;', '2021-03-06 23:54:01', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ef6590c5-0870-43b2-b9eb-4c754569a02e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:29:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f0724c0c-5cea-4708-8dc1-59c55519086e', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:51:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f09ffa8c-4025-484f-8650-6e49308f2fd5', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 22:21:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f0bd93e3-66fe-49a1-b876-19be3e5852a9', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-06 23:44:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f2a44b64-cd32-4051-a230-06ed188d7096', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '4.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 4.4;', '2021-03-06 23:15:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f49ce569-309d-4c87-8a47-fe7149c18548', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.60', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.6;', '2021-03-06 23:57:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f49d1efe-069d-4023-bb94-38b72d1ea7aa', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.3;', '2021-03-06 23:16:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f67e9fba-5a46-46ae-bc55-de35a3336fcc', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '4.50', 'cpu使用率告警：阀值 0.70, 当前 4.5;', '2021-03-06 18:10:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f71f8d7a-94c2-4bb1-8d63-d0176483b146', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-07 08:51:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f7dd8375-ed64-404d-837c-90b153a373b1', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '5.00', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 5;', '2021-03-06 23:38:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f8289c4a-4e5c-4a16-81b5-bf9d3caf5636', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-06 22:21:02', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f82d67c9-fe26-4b4b-a64e-47cfceb22159', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-07 09:03:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('f8666b11-1cab-4b17-a8dd-bc855ae59c32', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.4;', '2021-03-07 09:04:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fa32ca6e-fe19-4edb-abc7-0d3cb4444679', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 22:52:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fab8053e-df60-4bd3-baa8-b45342e88748', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '0.04', '主机：HC测试主机1 内存使用率告警：阀值 0.50, 当前 0.5352733686067019;', '2021-03-07 19:53:05', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fbb68360-5438-43c9-8b0e-da910d2cce2a', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 23:21:01', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fc47529d-0caa-4c09-983c-f3cec7f459d2', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.50', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.5;', '2021-03-06 23:00:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fcd81238-4551-43ef-ac87-8884a79a2c98', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.00', '0.00', '主机【HC测试机】,ip=106.52.221.206:9004,wuxuewen进程不存在', '2021-03-11 22:33:02', '0', '3003', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ff757cdf-e53f-40d6-a0e9-c2fee7d0e014', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '2.30', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 2.3;', '2021-03-06 22:43:04', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('ffa876b6-8bcc-416b-99b6-4a87f1c49a92', '1001', '55dfa429-5f96-44f9-ace8-88870b6c3741', 'HC测试主机1', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.80', '3.40', '主机：HC测试主机1 cpu使用率告警：阀值 0.80, 当前 3.4;', '2021-03-06 23:02:03', '0', '2002', '000', '待告警');
INSERT INTO `monitor_event` VALUES ('fff88c3b-4845-4786-882e-edb6bdaba3d9', '1001', '0638e278-484c-4e49-b59f-9ed73ab418ed', 'HC测试机', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.70', '9.50', '主机：HC测试机 cpu使用率告警：阀值 0.70, 当前 9.5;', '2021-03-06 22:52:02', '0', '2002', '000', '待告警');

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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0' ,
  PRIMARY KEY (`log_id`)
) ;

-- ----------------------------
-- Records of monitor_host_log
-- ----------------------------
INSERT INTO `monitor_host_log` VALUES ('0053007d-09d5-4f21-8e67-fe6da16f68a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 15:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('00570f56-179e-4e35-90ae-b78f848fa31d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 21:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0077ae1a-909c-4eb0-ac3c-975b2e49a29d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('007a8430-83ed-4b17-9bd0-b05e9bc9e2db', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('00a9d518-9b45-4961-9961-cd7cf6ab4405', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('00ada2cc-adf2-48c9-a19b-a43a25d835fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('00b20bbf-3ef8-42fa-b84e-62adf973d64e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-14 12:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('00b98c8f-d57c-488c-93ef-64bf774bca0e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-06 23:56:03', '0');
INSERT INTO `monitor_host_log` VALUES ('00f8da97-72e7-4542-9d3b-0e2014f44ccd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-07 21:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('00fc5c56-9b46-42f5-a9e6-76a356debd80', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('01087123-6728-4712-87d9-157701f45504', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('011744de-f9ed-47c6-afae-87a5cee919ec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0143882a-7f7b-484e-8947-38bfa40e6129', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('014a43ca-573a-420a-abdf-16ca6187c289', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.30', '0.54', '0.10', '2021-03-07 19:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('015161fe-cb3e-4171-b2a7-36e74ed498aa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0176b0c7-b6e9-4cbe-92fb-fe61ee3d2eeb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('01932e04-8eb5-486d-a303-a701b3d63398', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 22:29:01', '0');
INSERT INTO `monitor_host_log` VALUES ('01af0b6b-56fa-4f47-bece-98f1b7923d87', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:00:06', '0');
INSERT INTO `monitor_host_log` VALUES ('01af9bed-91ab-4829-8a2e-345e5a8aca50', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.54', '0.10', '2021-03-07 10:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('01cbb53a-80a7-46ff-9045-4a9de0b34088', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.53', '0.10', '2021-03-07 09:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('01da8692-d45c-41d6-903e-d6e44fed380a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('02267dbf-bb6e-4648-abf4-378c961b860b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('024a0afa-d542-44b1-8dd8-17604ffa203f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('02764ad0-4a1e-4eca-ae48-97123ca74bc3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('028188fc-98d7-436c-84cd-4a016ab5d15d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 21:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('028fbc15-5c20-4315-9910-1e6b0d32f685', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 22:43:03', '0');
INSERT INTO `monitor_host_log` VALUES ('02981938-f733-42cb-bca5-3bb6e20206c1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('02a15b29-098c-474c-a177-9df45a20eb63', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 22:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('02b75cb3-1128-483b-97c1-c87695694f78', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('02bea921-b477-4ca5-b7a5-a47af7c93a8d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('02d4b953-c0d2-4eae-bb3c-2291641dbf41', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '27.90', '0.54', '0.10', '2021-03-07 09:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('031aa505-be73-4918-b98b-6efa33d0b3b8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.77', '0.56', '0.10', '2021-03-14 11:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('034098b6-7269-4e1e-bf45-32c05086d8b6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-07 21:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('035a1127-a6fa-454e-86db-a982bdea9b35', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('035a11b6-412d-4679-b62d-fe5267f38f4d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('035c7fc1-4bec-4805-8d2f-5ea603cf193a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:08:03', '0');
INSERT INTO `monitor_host_log` VALUES ('036f074d-6093-4e41-a7aa-7f97e8964de4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('038bd263-27bf-4d80-bd55-a9a7118d72f0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('039605f4-c7b9-454b-a9fb-878ac0d7562f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('03a1b579-ddba-425a-a51e-a4ed98dab4f2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('03a801a9-fa08-443f-a6b3-9391c5a3df5e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('03cf904a-29af-4020-a537-27f14b78e355', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('040a5d08-ad21-4fc8-9b9c-ed96495768c9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('048d9470-8b0a-41c0-b5e7-fa12b7df0713', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 21:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('04a7c76e-21ee-4bbe-b180-563bc46f3d2f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('04d265f1-af85-4758-a840-c26a9831a1bc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('04fa468d-ece8-421b-b967-b0c6073e5f22', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('05221c61-b75f-44e4-94b6-66b8943fa54b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('053bfbef-78f0-41fa-8570-c5d89f839e72', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0545e75e-7105-46ef-b1a4-b0520f83ba50', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 00:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('05c2367c-7c8b-496c-a957-d2e71bf526b6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 23:24:01', '0');
INSERT INTO `monitor_host_log` VALUES ('05e98232-fafa-4fd3-927f-e8dd3fbab089', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-10 21:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('062aae8f-9ce8-4479-890c-1a903ce883e1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 22:31:03', '0');
INSERT INTO `monitor_host_log` VALUES ('06363047-11be-40f3-aa6b-be5450ce2832', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('065c2ea9-7380-49fb-a48c-d8c2ffc59d11', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 10:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('06916af7-dafa-49c8-b00d-72259cb41646', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('069c8658-19d1-471a-ad49-b76a7a7f8b5a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('06a99180-7c50-4d88-b856-bda8f1c819b1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 23:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('06bee221-0c5f-43e1-ad1f-07ff87347e96', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('06e0dd60-7c47-4345-9dd8-09573c35f5d2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('06fc1b1b-9d7d-4549-8af5-ae8dd61f962b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0705b758-f0f1-473a-a4e5-f475ac430e43', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:34:01', '0');
INSERT INTO `monitor_host_log` VALUES ('07091907-065e-48b3-a425-707c91909949', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0712de19-7427-404a-b42b-a37c9524cdce', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('073e7b59-6992-40f1-a433-6cb266ff79cb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 17:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('075bd508-7feb-4785-880d-99237ab96ad2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 22:27:01', '0');
INSERT INTO `monitor_host_log` VALUES ('07879fe7-3e87-40f6-ba09-31d14bbccf42', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('07b6b87c-d05f-4f2f-9937-394dd2e10607', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 23:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('07beffed-0c7c-4d67-a329-b055d2c8221d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('07daebd9-36ca-424a-9a1c-b1fb5660c862', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('080d5cce-3437-46d8-a9e9-2293c025fc37', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('086c803d-2302-449c-b394-c56de841c60d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.56', '0.54', '0.10', '2021-03-09 23:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0899234c-8f1a-43a4-abce-c02217030b4d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:03:03', '0');
INSERT INTO `monitor_host_log` VALUES ('08cc01d1-22ad-4944-adf4-e792c2e96d70', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 16:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('08eb3c62-5265-4577-9106-4b078e9b1435', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0923143c-5f82-409b-966f-96ab5b67f4a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.02', '2021-03-07 23:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('09326279-9e4d-4a68-a51f-518c8ffefe86', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('093e3c18-2613-48b8-b5dd-149b0a1a06cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 15:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('09408dbd-db16-462c-a711-3f500374557f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('09686e14-5b5e-4330-9c0a-8eddb6cc1336', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('09751ecd-841a-4687-82f8-6a4b9589381c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0978fc0b-3440-459b-a70c-5e2e1d3473f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('097c0dc3-fe19-4b71-b69f-493fbacee4a7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('099bdadb-f8b4-4cce-b1f2-7f09088950e3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 23:31:01', '0');
INSERT INTO `monitor_host_log` VALUES ('09aa6728-466d-49db-af45-178c23d1d6a7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('09b12056-80e9-4a86-8c9d-a46624386ce6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:24:03', '0');
INSERT INTO `monitor_host_log` VALUES ('09d9d96a-94fa-4f49-a999-859604626d5b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('09dee98b-df16-41f3-bc87-66014419144b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('09fb43e3-cf74-49f5-8ea3-fff4b5093626', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0a13aa1f-50cf-4640-872f-cd64e0777a8c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0a240bfc-d45a-45b1-9a74-b7b9008d00b4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0a30ed5c-e71d-453b-8a80-ebb05c7e2eda', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0a628147-d179-44e7-bdc8-2c5e313f2c3e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0a66194d-cf04-4735-839c-f24e21d0cd08', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0a84f0a9-9f36-49e3-9f85-648ea577d16f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0aa216d4-c60b-4d9e-9926-774f7b2ac2b3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0ac0a844-4fc7-484b-ae24-a614de316449', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 14:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0ad97140-e89d-442b-8194-b673fa96e0ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0b02999c-c920-4150-9599-63edc37d9862', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 17:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0b19a050-5fa8-47d1-b27b-9f13f47eb6db', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.54', '0.10', '2021-03-11 22:22:05', '0');
INSERT INTO `monitor_host_log` VALUES ('0b1a97bc-753e-49d1-8d63-e6ecdfd43cce', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 23:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0b3664a1-a52e-43c2-a5c7-ebebed10c995', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0b5ac798-e995-474a-b4a8-ff894bfd99f7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.77', '0.57', '0.10', '2021-03-14 11:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0b613f05-ecce-46d1-870b-0d54af504770', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 21:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0ba07297-57b7-411a-9ec2-6e84a42b6c84', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0bb532b6-f0b6-44a8-b3ac-bc32d7a88f17', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0bd1c66e-d0bd-4bf6-a53e-a96f47cbbaf5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0bf31e71-fa71-4f4f-99c6-58d3343f9f0d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 15:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0bf9a4fc-726f-4c9a-883b-e29cf9a20c72', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0c3df3c0-2b9c-48b0-a958-492075f16ba4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 23:05:01', '0');
INSERT INTO `monitor_host_log` VALUES ('0c5af3c2-1937-47a2-be5b-84db37bd2afc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:05:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0c778f56-6bc4-451b-b6cb-9e15e635b8f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0c9564a7-c399-4732-8f0d-48c581332477', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 15:02:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0ca9f3a8-dda8-4cc9-a7fe-a6149cedc9ff', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0cbb0239-c7a8-4922-92f4-e1df5bcab77d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0cd1076c-4db5-4356-b175-5603d9dce7aa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0cde09df-1448-4b68-8587-9a4f65a85a8f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 15:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0cfb2e95-89ee-4703-bdb0-8930c86fdb57', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0d46428f-990a-445f-be39-d956ae0fc070', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0d6bd29a-7ee4-493c-b2f5-c1f5e4bae779', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0d74a797-4838-483c-b546-e7685ddcf1d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.76', '0.56', '0.10', '2021-03-14 11:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0d8746d6-3fd8-4aaf-8085-1a3200ec1325', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:46:06', '0');
INSERT INTO `monitor_host_log` VALUES ('0d88aa77-bf8e-49c9-b3b7-a30d4cd88304', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0db430e7-c9ef-47dc-b7c2-91fd3a347229', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 15:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0dc551d5-8479-4289-9d6c-849acbbb0ede', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0e3a0bbd-4856-44c7-98c6-1bbc1bcb51ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 23:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0e75ae1b-d373-4b9d-aa4a-138ffd63b532', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-07 08:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0e879c75-7ec2-4e4b-b29a-119a4d513783', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 14:43:18', '0');
INSERT INTO `monitor_host_log` VALUES ('0eae18e3-3b46-4856-ae51-d54e941b8fca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 22:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0ec562c3-f0e7-4b58-8f34-640930905c59', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0ec9f8ec-d624-4100-a4e2-bd4c61e88c2b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-13 22:47:08', '0');
INSERT INTO `monitor_host_log` VALUES ('0ed8d815-8f2a-4126-993b-156ff83b0e6c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 23:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0ee8f0fe-f17e-49cd-9b4d-70d2db20dcc2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0ee98d89-ab10-49e5-aef5-e304a9e52af3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.14', '0.14', '0.02', '2021-03-07 23:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0f0efe75-fa0c-40a5-8db1-029e56c5389b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.70', '0.54', '0.10', '2021-03-06 14:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0f1be54b-212b-4ef8-8a4a-81b560271939', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:31:01', '0');
INSERT INTO `monitor_host_log` VALUES ('0f20db6f-ca2f-4a7c-a9c9-f8f0e21bb3cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:33:03', '0');
INSERT INTO `monitor_host_log` VALUES ('0f2f3fc9-13a3-41eb-8138-81951d25000b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0f2f9617-b864-447a-a583-e0eafee46bdf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 17:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0f4d8b07-6577-4299-9e84-1b8cbeb5900d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0f5bc372-5a0c-4c5f-ae7f-7735170546ff', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0f99e84e-f8d1-478e-b1b2-9442fe263116', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('0fa85d74-9a0a-4d13-a1eb-6126892f4dac', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 17:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('0fdde151-d7e4-441a-8d00-52e7748141dd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.18', '0.10', '0.02', '2021-03-11 00:48:05', '0');
INSERT INTO `monitor_host_log` VALUES ('0fed1eba-1531-4693-ab5c-af77534e0c34', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('10062afa-206b-4e44-b35d-8fe891024e03', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('103a9db5-b8b3-4b01-83e5-062c3c371ba2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.11', '0.02', '2021-03-06 16:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('10483ad4-9b7c-46d2-a383-6f7a325c0728', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1061c06d-a141-487e-b968-e728c2612520', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.77', '0.57', '0.10', '2021-03-14 11:27:03', '0');
INSERT INTO `monitor_host_log` VALUES ('1066f0ef-0c76-4c67-82d1-e35bb593d594', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('107b791b-f6e0-494d-b6b4-a35ca061693a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('107e154c-150f-46f0-b4e6-436e737f8918', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('10875747-a7dc-4781-afe3-fd1ea5efc13a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.57', '0.10', '2021-03-14 13:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1095c3a5-1e8a-4c39-815e-7885dba51a42', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('10cd3f87-183f-4393-ae76-4014603ef909', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1140a09c-b10a-46f0-a3f6-58abd2b037d3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-11 22:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('11450811-d188-48c7-9747-78ebef9c6c7e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1190ccb5-c87d-4e9c-891c-b376da5946f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.03', '2021-03-10 20:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1199f911-42db-4df1-a9bf-94a3edfd4e50', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('11a6cbb6-d500-4ea9-8e1f-6e3b3674d3af', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('11e62c55-438e-458e-bad5-b71918f9f775', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('11ec62b1-44b7-43c4-a2ac-502441cb7bd9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:42:06', '0');
INSERT INTO `monitor_host_log` VALUES ('11ece459-3034-4253-a55d-4adf93202b82', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:47:06', '0');
INSERT INTO `monitor_host_log` VALUES ('12037946-75a8-4887-839e-be7fc82b9bee', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1244593e-acca-4df8-945d-e71be8cb13da', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('125c15d0-8a15-4e43-a097-3469a5ba036f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:46:40', '0');
INSERT INTO `monitor_host_log` VALUES ('127246cd-0e28-4091-aab1-f580e642a99d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1273181b-3773-4b8f-8993-6aa23f1f0307', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1275234e-619a-4faf-81f5-cd31d8a1d32a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:16:03', '0');
INSERT INTO `monitor_host_log` VALUES ('129febe5-8979-4e67-8f21-984bbc308d0a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:42:01', '0');
INSERT INTO `monitor_host_log` VALUES ('12b32b75-dca6-4323-b9ba-6d06f887e7ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 09:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('12d6f5c5-fdd6-41c9-a636-03c13840d8bf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('12d8c8fa-05f0-45cc-981c-c429c927b154', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('12e2f9e8-c41b-484d-b09b-f90f77a65c40', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:56:05', '0');
INSERT INTO `monitor_host_log` VALUES ('12f9c0be-bd9e-43f9-8921-b7f7efcc56fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 19:41:06', '0');
INSERT INTO `monitor_host_log` VALUES ('13010341-c053-4caa-8734-b372122e7577', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 15:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('130ad417-c634-4723-bc1e-e52a52ddc942', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 12:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('134e64e8-bf3e-4ce5-bb8a-fbc4000d82b7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 00:50:06', '0');
INSERT INTO `monitor_host_log` VALUES ('138b8402-ef30-4481-895e-f3c8e5f60942', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 10:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('13948a00-6fe6-482e-b86e-3fca8dee55f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 22:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('13bc7a9f-96f4-46f3-9010-9f879d4f639b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:49:03', '0');
INSERT INTO `monitor_host_log` VALUES ('14033169-a384-4110-8a28-e562cf9c5370', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('141d264f-8dc7-4811-bbfe-87f4530c72bb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-07 09:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('146760a4-e700-44f6-b4ab-f791744b48be', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('148819a2-4b02-40b8-86ae-cdb6592a2f10', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 19:38:06', '0');
INSERT INTO `monitor_host_log` VALUES ('14d22768-6209-4944-902f-c6cfb9188dbf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('14fe0848-05cd-4ede-857c-27b35b790d1c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.11', '0.02', '2021-03-11 22:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('150a1625-d72e-4025-9cac-7b3326b67ae0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.56', '0.10', '2021-03-13 22:39:08', '0');
INSERT INTO `monitor_host_log` VALUES ('1529a476-0918-4655-9d45-83b89795e98a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.14', '0.02', '2021-03-08 00:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1539a503-3790-4f8e-9f64-e4c90bdbf5a8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.90', '0.11', '0.02', '2021-03-06 23:53:01', '0');
INSERT INTO `monitor_host_log` VALUES ('153d3e35-7484-4b7e-9a94-b7e3746091cb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.54', '0.10', '2021-03-07 15:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('15b5caa2-b6e6-41df-a683-f72e10d5b32e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('15da7bb7-39b0-4a33-9a7f-d2757876a8f7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 17:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('15ddae04-771f-473a-af25-066f49a46993', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-08 00:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('15e7dd7c-63e3-4f1d-a70d-bf73c853b8fb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('16191f92-1872-4a78-a737-ed9d5bf66112', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('161be405-30e0-4fd5-b841-83dda0f29fe1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('168e1639-2f14-4781-981c-d8409126d06a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('172fed15-f77b-4bbb-86d1-c7e066dc4e98', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('174855be-1887-43f5-b7db-bfea8e3036a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 22:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('176e08cb-d60e-44bd-877d-3e6fcc308d33', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('179136fa-f97b-4bec-a033-f21b9a6fb974', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('17953b47-8c0e-4b51-9543-e8ee7ed432fa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('17b3d13f-7dd1-4e15-aaee-28f15608946b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('17c3ae68-851e-473d-879e-49ae60f6a3a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 14:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('17c6bfa1-5349-4ec1-b1ad-b8019ab6b758', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.75', '0.56', '0.10', '2021-03-14 11:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('17de1467-2eab-4f02-b0f4-23afca042969', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('17e219ce-7a95-4fa6-a824-80361ca45a85', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:04:06', '0');
INSERT INTO `monitor_host_log` VALUES ('17fc5ce0-c0d5-4626-b430-7160b4aeaeb5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('17fed935-ec91-4f7f-8bb1-629a8e6dd7a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1802e4ed-a383-4418-97a3-92f45d7e1765', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.54', '0.10', '2021-03-07 08:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1805cdf6-b638-436f-80ad-a6d106ff98e9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('18623a05-9b5f-4280-a193-fa21e7f69f22', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 21:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('18765161-c9df-422c-a08f-cd852c2c4dfa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 23:57:01', '0');
INSERT INTO `monitor_host_log` VALUES ('18945a1d-0bf7-49a6-beb2-b3d2a89e75fa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-14 12:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('18a24bf8-0314-4a57-9ca2-8f5a72019ac6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('18d64e0b-8da6-4091-8614-6c22f26c9fd2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('18dcc6fd-2109-48ef-885b-723562359095', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 23:25:01', '0');
INSERT INTO `monitor_host_log` VALUES ('1902d6e2-59d9-456f-8213-0f1b25410ba2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('19193cf5-8bb3-4da5-90e4-419f941524c8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('193be4f6-a988-4521-b368-0ddb7e72dd03', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 22:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('19887e00-7b2f-480c-a1b8-247525ab9b83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 18:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1998afbb-ca51-4550-b633-f80ef89cb39e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 15:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('199d86ee-a9d5-4900-8ef0-ae238e79a3fe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('19c78e22-e15b-4beb-8bdf-3b6013ec1107', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('19cac376-70b2-4721-9b65-16cb5606faa4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('19fc44ee-c2e6-4b7a-ba99-01b6d193c835', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-14 13:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1a042203-ce28-40b5-b6c5-1eeec92363f1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 22:34:03', '0');
INSERT INTO `monitor_host_log` VALUES ('1a19af3f-2c97-4f7c-a1f0-21fe1c42301b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1a4c6b44-43dd-4509-ad48-aec9c71233f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 09:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1a7a97f2-6f12-4419-a077-4413694f90d3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1af26cc0-f3ca-4c5d-b0c6-ecf752dbfc9c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1b16c49e-86d7-4168-846c-22e81188c0a4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 00:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1bc92489-9301-4494-bc75-60cc6a006f8e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1bda21d5-ebc5-4374-957e-2eff6aaecb2f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1be979d7-d3e6-4272-84bd-87d19deaa5ad', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.02', '2021-03-07 23:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1c13b10e-cce6-47ce-bdf0-3a1bb44e1efb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1c13bd36-1816-4af9-ae66-0244240878ec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1c23174a-8c65-4698-944b-6865158cde67', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 22:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1c3c6b5c-0db1-4003-8f04-104c9e52c8c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('1c6db868-cfbb-45b3-a748-3a6ce03fc6c7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1c7d382b-f9f6-442f-b79a-f8a019ca7a8b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1d0b6148-e70a-4726-b7c1-d2200ae2c71e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 23:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1d0dc59a-5d79-4f72-92fe-62262e79f27a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 13:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1d1bce2e-99dc-422e-ac69-706a00074869', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1d215ccc-7f5b-4432-8f8c-c4c66be30b15', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 15:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1d997239-3f23-4ed9-b61e-71ffcf807141', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1dc7ca2c-39d2-4096-b83f-f71c4092a20e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1dfbb788-95bb-4f58-b81b-1becf7ed6995', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1dfe7756-656b-488f-9add-99548b280c4a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-10 21:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1e02fdbd-87d7-4335-8ca1-88f22675a98f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1e176534-0c0b-4f40-bde5-ca266eff3d79', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1e50f986-c754-43b3-bd92-6a3482497292', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1e5b3432-c995-4025-8fb8-e315b5e4391f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('1e68ced7-598f-423a-8cca-adef7063780c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1e788701-6b52-4863-9df6-9bf7c5bfbd9f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1e7fdaa2-adac-48a6-bf39-f5e5a011817c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:09:06', '0');
INSERT INTO `monitor_host_log` VALUES ('1e8fa3d9-90a9-42bc-b556-7877b4755301', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:40:09', '0');
INSERT INTO `monitor_host_log` VALUES ('1eabd68d-5952-4182-8e11-8e70ddc810db', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.60', '0.54', '0.10', '2021-03-06 15:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('1ebb68d9-0019-4eb9-bc2e-98a924f2c746', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1ebd5b92-6c7c-4655-9802-e051479651b6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1ee1f9a3-b2b2-4519-83a1-6d1307184c21', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1f062a21-a81d-46e5-829b-aa5c589e87f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1f25460d-2186-452d-af05-e420e84b8a74', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.02', '2021-03-07 23:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1f2b88e0-f56f-422c-9c67-171c7f949e46', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1f31d5f2-7c38-4f94-a287-d0b6143011a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.03', '2021-03-09 23:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1f3ec8a9-a5a4-42d4-9ee9-5a7511dcc9f6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1f4de4ce-ca3d-4aa3-a079-63c4685549cb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 23:27:01', '0');
INSERT INTO `monitor_host_log` VALUES ('1f89f508-5fdc-4600-ae1c-0a2f55010438', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('1fa15b98-9c96-43be-b3bb-b757263675e8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 15:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1fb5f561-1e5d-4483-a551-daf91e3ace95', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 22:26:06', '0');
INSERT INTO `monitor_host_log` VALUES ('1fd7f52f-1441-45c8-93af-7e1ce36996b9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('1ff54f63-5e8f-4829-8ffd-98595f24375c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 15:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('20240b0b-6d1d-439d-81f5-22c834e85ed6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('204027af-3b4a-4572-811e-9fb73ef43f12', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20480d70-a753-437b-9f0c-6de1b1d4b611', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:32:01', '0');
INSERT INTO `monitor_host_log` VALUES ('20486a9b-89fb-40bf-a4bc-22dc4ca3f377', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.12', '0.03', '2021-03-09 22:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2074ea41-d322-443d-a2c6-6f9344c67707', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 12:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('207dd24b-362e-4ce7-8c3d-c81c9152e860', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:43:06', '0');
INSERT INTO `monitor_host_log` VALUES ('20a6b501-e7e2-43a5-ac34-3a77e672b4ef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20aa838e-bbbf-4b02-a5cf-1df9866d2d62', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20b15f3e-7a01-422c-80ab-ba2a17f6d892', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 23:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20bcbe24-eddb-4512-b3b6-0847947ca4c7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20e30619-54b0-4fbe-ac01-df6bfd916061', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('20ea58b0-4ce2-4f7e-8be9-098b305df1af', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.56', '0.10', '2021-03-14 11:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('20f09f02-e3b3-4084-9d39-852479f775dd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.70', '0.54', '0.10', '2021-03-06 15:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('20f467d8-2bb7-4b5f-8e88-bc18a6d78bb0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('210bfc16-3468-4fd0-aefc-47be70c2df46', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 12:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('212a7e5f-27a8-4e96-9530-22293a050b45', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('217475fa-60a0-4c27-86f3-721f9639b654', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2189fb76-f695-45a5-89ee-6d8cb681481e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.15', '0.02', '2021-03-08 00:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('21a06179-4365-4239-be3b-2c59fc1a1253', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 23:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('21cb65eb-8958-42b7-846c-15ce0603fff3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:01:01', '0');
INSERT INTO `monitor_host_log` VALUES ('21e7f369-dc6f-475a-b3c9-13939014d5c0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('21f670e6-0a87-4920-ba0c-d942a04d81ce', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('222f3f9a-12ff-4f75-9e0c-46ae29f70d19', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:56:03', '0');
INSERT INTO `monitor_host_log` VALUES ('22337907-db21-42a6-8712-4a17c672a21d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2278c647-28d7-4e06-b058-8162132b94a1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 08:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('228fe82d-b615-426c-8437-4a157c6e7a9c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('22de6bf2-074a-49d1-aa72-e26638f53b46', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('22e08801-eca5-4cea-a612-24fe1a8f6414', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 21:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('22e9edea-7cf8-4e9a-896b-bf52d437a597', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 21:49:03', '0');
INSERT INTO `monitor_host_log` VALUES ('230e2d21-fd8f-4f5e-8db3-5bd75bc32dbe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '7.20', '0.53', '0.10', '2021-03-06 23:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2359773b-c499-4ee4-aa3f-2fa6d10aa5c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('237d2125-f033-45ea-8105-c61d282c3ac7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 12:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('238770ad-d31b-4be2-b88b-a4ed9eb8d36c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('23bc812a-41a4-44bb-bc3d-e4e81bd08791', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('23cbdc58-858f-48e2-99f0-e22441579550', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 22:55:01', '0');
INSERT INTO `monitor_host_log` VALUES ('23e0e106-0e73-4648-ac8d-0f3975fcfff9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.90', '0.13', '0.02', '2021-03-06 18:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('23ecaceb-65aa-401c-91b7-a43b7fbcb4ef', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.54', '0.10', '2021-03-07 15:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('23f301cd-0a3e-4a5b-a320-14297b4f41fa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-07 14:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2403e64f-3b92-43c9-85f1-427bc3d2717e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:05:06', '0');
INSERT INTO `monitor_host_log` VALUES ('2440632a-11b3-4d2f-87e2-1d8878e8c65d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('24777aa0-3bbd-4972-af0a-50e378ad80f5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 22:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('249902cb-72da-414f-bab2-d00db915b41c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('24aa5d2d-ef9d-4734-80bd-1778d3651470', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:49:05', '0');
INSERT INTO `monitor_host_log` VALUES ('24aee45a-a06f-4df7-a96d-3665773cad2a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.03', '2021-03-09 23:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('24e39951-84d2-4845-81e7-a82e0fa70b07', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 17:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('24e66928-d5cc-435b-9db4-b7b909fd018d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 20:00:05', '0');
INSERT INTO `monitor_host_log` VALUES ('24edd8c4-1021-48ef-bdba-536935de7311', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('24f1b87d-b5f0-48f7-8281-9801decde180', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:36:06', '0');
INSERT INTO `monitor_host_log` VALUES ('25352c0c-4617-4d6e-a7f1-8928f6d4e8fc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2538ca15-57f9-471b-bc63-0ccc6dbd7534', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('253cf4f8-d9d5-4143-a5a5-7c6915228ef6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 22:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('25622148-649c-4dbf-8652-6cb1f688bd0d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 22:36:03', '0');
INSERT INTO `monitor_host_log` VALUES ('256dfc5a-f7cd-4a27-81ed-db8bd26d3b8b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2573ad4b-6a13-4af7-bac4-26b7e105e5ad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('259ba43c-b80a-4648-aab6-a67f35ea56d2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.54', '0.10', '2021-03-07 15:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('25a6e250-8dfa-4626-b8a0-33fd31f24be1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('261d6fbd-9917-4adf-9c6a-c6f0dbbd578b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('26229281-4ad5-4e5f-bad4-cd02b75453e6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2628c793-25e2-41e6-a4c1-c86ba92f7343', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.57', '0.10', '2021-03-14 13:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('265d9462-f288-4545-b675-2c929b9d1491', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2681476d-eb66-4b7d-a374-c10f95909046', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:41:06', '0');
INSERT INTO `monitor_host_log` VALUES ('269c7f3a-a794-4118-9d4a-aeb8ffe3b1ab', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 11:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('26b78199-1126-4390-9ce9-c0ae60a67d60', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('26fad8d2-c76a-4e0f-a690-2352045fb125', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('26fe5860-739f-4675-9ceb-4030374fbe24', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2700a704-ac16-428f-bbe8-2499e3dd9f54', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('270aac1d-ec12-41bd-8db9-7a474a9f93a4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 12:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('275df1c0-eb0e-46f9-89a5-370bd1c6dde9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 23:05:03', '0');
INSERT INTO `monitor_host_log` VALUES ('278e1450-5f9c-44a3-887f-886e81526d2b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('27995f2b-28ae-4ccb-8957-e6893187f4a5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2857019f-3ddf-4b5f-a171-57c8b25c85a2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 18:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('28695dfc-fcef-4e95-afb6-0081b4d3c5af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 15:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('28965b74-6e7c-4b89-bf1e-40e78bba297a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:35:06', '0');
INSERT INTO `monitor_host_log` VALUES ('289ed85a-a3f3-4e6a-b5a4-32ea6deec3d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('28bb261e-7003-42e3-96cb-32eb2d7a8444', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('28c1d51c-47a7-46c1-9d13-3ffbe871ddd2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 11:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('28da9626-5298-4ecc-8c18-7d482677e8e5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:12:03', '0');
INSERT INTO `monitor_host_log` VALUES ('28e17f50-e9ff-48d4-866d-590d6656dbe9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('29053b5a-2274-4dad-b65b-327ed67308fb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('29097c36-d1f4-4200-9d62-34705b749bdb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2918eec7-a5c1-40ca-a1ad-16295b8be524', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 19:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('292e7ef9-9643-4901-8e30-90b4da8c6cbd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:19:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2931857d-960c-441c-bdad-258f755f4c3d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:48:03', '0');
INSERT INTO `monitor_host_log` VALUES ('29523cb8-9cfa-4da9-9e8e-b9c76d752db4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2966798f-822b-44b2-8eca-bd4753a40c52', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('297c8e28-a79d-4c15-9a03-a4fe5a48e46d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.53', '0.10', '2021-03-06 23:57:03', '0');
INSERT INTO `monitor_host_log` VALUES ('297f736e-c463-42ba-b128-fd63454fea18', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:38:01', '0');
INSERT INTO `monitor_host_log` VALUES ('29dd8b61-684b-4860-af8a-5d7a55ffec6d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 16:37:01', '0');
INSERT INTO `monitor_host_log` VALUES ('29f26ab1-ffdf-4cdd-8af1-a664f46a7c93', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2a1475f8-b3f0-4dd2-bf8a-f3d70c9cf50e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2a175efe-b4f1-4274-9cec-21dd1b925fcf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-07 09:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2a1ea8de-6e0b-4e19-b1fa-1b4889ca094a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2a2b6dfe-3155-4886-b996-15f05926d138', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.15', '0.54', '0.10', '2021-03-07 21:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2a32e9fe-ccdf-4679-b90c-4866064a743e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.70', '0.54', '0.10', '2021-03-06 15:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2ae825df-ec47-49ed-8fb4-9382eb80f464', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2af69791-5fff-4498-9aad-fe67a8af6650', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2afd8b65-6a15-4fcf-8b78-3667e2a44a4a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2b13b9f6-9d78-4575-97bb-085ee640ccfd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b196de4-3312-415d-8ebd-0a4d8498e855', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-09 21:49:01', '0');
INSERT INTO `monitor_host_log` VALUES ('2b445465-535a-4490-9544-118201a5d487', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b4f18dc-b251-4f82-b00e-48f73b9fb220', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b66e7d3-5667-447c-b774-bb8f0eb2095c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b718ef2-d02e-4656-91c2-10dbb4344390', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 22:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b74f70c-1978-464a-813e-a6637b354ab8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2b7c8513-ed57-4b54-965b-1e851681974c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2b8210c5-e351-44d0-8a40-ca85428c61fe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2bbee936-d5b2-49c3-be35-f3552ee3daac', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.03', '2021-03-10 21:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2bdee246-1f2d-408a-a9ef-0cd5cfe418aa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:57:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2be3bd04-c755-4717-ac88-e39bc53b2b57', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2bfdf06f-e64a-4f7f-9897-8700e2b48063', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2c3b1e6d-76a8-4ccf-9a00-a125d5d16aa8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:43:06', '0');
INSERT INTO `monitor_host_log` VALUES ('2c50125a-fa24-4797-96bd-a1c1eaaecd5d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2c61079a-9df9-4aed-a40e-42b70cf137d6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2c6a9ceb-0127-4cfb-960f-9d79a43c6f46', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 23:56:01', '0');
INSERT INTO `monitor_host_log` VALUES ('2c867949-da8f-4374-9e32-9464241da26b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 14:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2cc1fcef-fc3e-40a0-b3c9-265f2249ae6a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.24', '0.54', '0.10', '2021-03-07 10:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2cd321df-48e1-46c7-9d48-8a3e8893a90c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2cec4809-53f3-4096-9843-ebb21ea05417', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2d412ee0-6207-470b-a755-2ffaf308b516', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:42:07', '0');
INSERT INTO `monitor_host_log` VALUES ('2d524eec-1da8-497f-9c0a-f84726b98601', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2d8d096e-c552-43e8-b9af-55b7f9d0dd0b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2e11ab49-6680-407b-9d00-9219b16e516d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2e149a4b-b6b0-4821-b087-93c4d2b5b10a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2e4d3636-961b-4c1c-8e41-e345bd4df686', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '12.00', '0.54', '0.10', '2021-03-06 15:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2e5cf780-22a5-413f-82f3-811faecb7549', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2e7ecdcb-535c-43cc-894c-63fd5c555452', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 16:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2e89e276-0fa2-4513-bf10-4955621455b2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2e98e1a5-ce3e-413c-af1e-638ffd6cb375', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2ea3370f-f2c6-4f30-9b40-c08aa0f30537', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.11', '0.02', '2021-03-06 17:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2eba3d97-1636-40fc-9f20-824f11d3c706', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 22:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2ec867c1-1a51-4407-9d72-22ffae2080fb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2efb548c-5af0-4a1d-856b-9609693691fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 17:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2eff87af-c17a-4480-b8fd-688217d6674d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2f1b0a19-afaa-4724-a384-c99b33b7a2af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:44:01', '0');
INSERT INTO `monitor_host_log` VALUES ('2f407a3b-e6fe-4bd8-957a-7c87e91ec5e6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2f52918a-f765-4c09-9a53-f1cc92a30b92', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.78', '0.56', '0.10', '2021-03-14 11:24:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2f65c4ce-a729-4594-8399-ab0c08fcffee', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2f7354c6-c69d-40f9-93fc-fca848a0efff', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2f9a526d-cd3e-4a47-9e3b-35cf54cabb7a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.11', '0.02', '2021-03-11 22:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2fa752f3-1d37-43cf-8bf2-a93e92b35865', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('2fab0b7f-2cff-49ff-9e50-9e2568a3e0fe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2facbd43-6c54-407d-ba1e-be9a98e1c9b3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2fb595c4-c5ba-4cc2-b492-0b194aa9de82', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 23:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('2fd6b576-bdb1-4752-8d5d-bb154578b62e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:25:06', '0');
INSERT INTO `monitor_host_log` VALUES ('2fdd1b56-8eac-48d0-9f32-36f3285d1ecf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('2fde8770-7650-426a-a555-ae4a7d84c8df', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:02:06', '0');
INSERT INTO `monitor_host_log` VALUES ('2ffef999-ec26-4887-bab6-0eccae9d27fe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('30038560-3b71-4496-86a3-fc21c593e614', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:27:06', '0');
INSERT INTO `monitor_host_log` VALUES ('300a7f38-1572-4482-8ed7-3b1f85ac728d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 22:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3060166d-bb6f-4dcc-87ee-6f2c55e62537', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 15:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('30773415-f816-4568-9cfc-ebebf4c1e0b7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-07 15:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('307dee04-3c9d-4dc0-a186-1ffdfbaa063d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 13:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('30a24c65-7acb-41b3-bc12-0e90cc91cbdf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 15:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('30b1539e-5da0-4484-9938-f2b5f4bfd72e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 23:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('30cb119c-1f68-4e58-9284-9ba854ff2e98', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('30e88e80-d9f2-4c1a-871c-5eef4b87e44d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('30eb379d-ffe2-42c2-94ff-f04641621845', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:38:06', '0');
INSERT INTO `monitor_host_log` VALUES ('31133976-c1d7-4e2e-aaac-f48e31409f5a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 19:57:06', '0');
INSERT INTO `monitor_host_log` VALUES ('31136ddc-aacb-4ac2-8628-bbf0f72174ca', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('31256a9f-3fbd-431f-b5d5-78de3691354f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('312e5e17-bda9-4d02-98af-c3f608e522d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('313c41ac-8bcb-4dd9-aaa6-d6b06105de8f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-09 22:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3154708e-bf6a-485b-9a95-f89378c1fc4e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 12:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('316c5a67-a5c1-49d0-a768-65c23f4d683b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3173b130-6320-48aa-9bad-a654c7e4dbfb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 20:02:03', '0');
INSERT INTO `monitor_host_log` VALUES ('317b2215-730e-4662-9b3a-4a21e34a57d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:49:08', '0');
INSERT INTO `monitor_host_log` VALUES ('319c8da6-7928-4936-a486-13226e407886', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 20:02:06', '0');
INSERT INTO `monitor_host_log` VALUES ('31d3e78f-e252-4ea6-8bf7-84dd198f5e32', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:38:03', '0');
INSERT INTO `monitor_host_log` VALUES ('31dcb6ca-aafe-40a3-ac39-87a895422cb4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('324cdfb4-2061-4842-b1b7-045b11d05099', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('32651e25-fb14-409e-9e28-9d16c286f43b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('32682aff-6e71-46ed-8a25-6332c8523502', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3295d67c-413e-4a88-9dcf-a26e8261d119', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('32c85823-cd4b-4129-afcf-0bd340c913d0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:00:01', '0');
INSERT INTO `monitor_host_log` VALUES ('32c9d0c7-2b68-40a1-8c68-7681105e1747', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('33395ea3-7cf2-4e1c-a0e6-0044eabce601', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('334c5095-ee16-49d7-a650-513e853ec12b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 23:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('336be4af-6fe9-4fb9-96e4-0658676f268d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3374b2da-8417-41a4-9549-b2e1aeccc3b2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.76', '0.57', '0.10', '2021-03-14 11:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3376e277-5892-4038-aed3-3833cefd128e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 17:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('337b5bfe-c5c7-4505-8b58-ae47ef35e727', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-14 11:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('337e1c11-6fb2-496c-9f5d-e2f82ffa12fd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('337e573c-b492-4250-90d4-204bf264ae7c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 22:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('33c92e93-0a32-4559-82ec-dd6ece96e5f8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('33ccc697-712b-4c68-8048-1fc4c331d406', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:46:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3404caa7-37ed-47fc-a735-23cee3063eee', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('34238712-ec1e-4e8b-90ca-9162113f40fd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-07 09:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('34598f2d-a8bf-46fd-afbe-fa01c44ea196', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:21:06', '0');
INSERT INTO `monitor_host_log` VALUES ('34a4e05a-4bd1-4e2c-ad66-e67342bd87a5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 17:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('34c04280-004f-4278-9b40-28d02a4d498a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('350be1f3-6941-45ce-8152-b13537a9c382', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '13.60', '0.12', '0.02', '2021-03-06 17:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3521f35b-c1ed-42ba-bc65-c5d04fa54653', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:46:01', '0');
INSERT INTO `monitor_host_log` VALUES ('354c67c5-8143-4467-917d-429320b78f57', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 17:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('35635e10-a909-49f0-90ba-d3827dd1d4ba', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 14:42:22', '0');
INSERT INTO `monitor_host_log` VALUES ('35671451-886c-4727-ba2d-c1215ecf7ed5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-14 13:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3574f02a-efb1-4271-ac3c-adc92b6425fc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:31:06', '0');
INSERT INTO `monitor_host_log` VALUES ('3582555d-8355-4602-bc7a-a216388e698a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.35', '0.53', '0.10', '2021-03-09 21:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3586a022-0a28-417e-bd26-5fb40baf355c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3595d72b-6063-4dd7-ba01-4a73ea06ef26', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('35d85d15-0945-4c12-bfd3-ff2c439eb3e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 12:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('35e01652-8971-463b-9955-6f54ef478668', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('35e2e174-77ac-489f-83a4-094d9adde50f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('36282fb7-a129-41f1-8346-7e75701965fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('362bc888-8471-4b2b-948a-3de5760088fd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 18:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3666aecd-3691-448d-9d43-79d3e1bbf4a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3669f428-cd43-477d-a222-b46a520788c4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3674bbb0-9ad3-44c7-9657-f2d84cbefcf5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('368533cc-4334-402c-aad2-a4a351433204', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('369346ca-168e-4356-8f92-f9560c093d5d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('36acfa4d-cc08-4c66-88c1-d086a11b5c20', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 23:00:03', '0');
INSERT INTO `monitor_host_log` VALUES ('36aff5ec-e2d0-43dd-af38-c78fa0ebc694', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('36bedd69-ffda-4f59-b982-943a923eecb7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:38:05', '0');
INSERT INTO `monitor_host_log` VALUES ('36bf683e-bf4f-4332-a0e2-9078ad78e7f0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('36db62d2-8339-4009-8c8f-d163a549e449', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('36f6c36e-32bb-41d6-a750-dfed9151309a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('372506b2-334f-4545-94c7-844678bf9a9c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 23:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('372fae6b-1313-4194-bdde-7974db1cef53', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('372fe5d6-a1ef-47ea-9bee-fe3fdc3f327f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('375d4eba-5d7f-4b0b-bba9-9aac1181e90a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3770ec98-ed57-4f09-9625-9964a09a9163', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 23:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3790d309-83ef-45cd-a2d4-6f7362b7628f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 15:48:03', '0');
INSERT INTO `monitor_host_log` VALUES ('379b3bd8-7d3d-446b-adf3-51eddebfc93d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 21:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('37aa9870-fc11-43aa-9eca-2277d78138ed', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('37afb43f-224f-46b0-9571-37fb7af1aaa6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('37bcfcd2-3140-4da7-b8b0-ac18790983ca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 10:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('37d4ed35-6ff8-4a72-ad1f-effdc6f08d83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('37f39eb8-8344-4d18-b905-ec4ed2d52427', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('37f5d1c0-c80d-4ea9-88e2-b95496c8312d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-07 05:36:37', '0');
INSERT INTO `monitor_host_log` VALUES ('3826f67e-f1cd-4e65-9700-feecff767bdc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('389a1431-3bae-4213-a353-fdc5e4811dcd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('38c0d3af-43f7-4808-bd6d-b82ac71d8763', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 09:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('38ca88db-b4bc-41bc-be2b-47313f6b0118', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('38d82ee5-f4f5-400d-89ca-ed2301f18a54', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('38e0f73d-f88e-46a4-bf70-9d2d7fa513a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('38e20559-a66e-43bf-ae1f-578458ac79f0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('38e2bf82-a7b8-45b0-8aff-0b5178e5adbb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.70', '0.54', '0.10', '2021-03-06 15:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('39089680-a7df-41c5-a14d-18d2339f5a8c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 14:45:26', '0');
INSERT INTO `monitor_host_log` VALUES ('394b5c38-c068-4b4a-86f2-5ec1e4af4c7e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.55', '0.10', '2021-03-07 23:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('39576530-7d34-492e-acba-67701b1aa019', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('398786e9-cdee-40ee-90f7-3e429b8458cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 17:20:10', '0');
INSERT INTO `monitor_host_log` VALUES ('39883ba2-5715-412c-bd02-3d9dd841d47c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3991b518-9184-4fc5-9786-f2ebc176fb22', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('39bc8519-84a2-496b-97ea-b9d7f2908d7d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('39e6a970-01d9-4ed6-92fb-4c74642890bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3a1fd6cb-39f4-42f8-a71e-1e68841c30d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3a207737-e7aa-42b8-ad6a-fcbe133b52a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 15:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3a397f24-06f0-4623-a04c-96a9cf52d64b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3a3c8fbb-510c-42bc-97e5-0a40652b9bb2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3a56db17-73d9-4f0b-ad00-558692308445', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:18:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3a666c9e-1cc4-4461-94f5-4dcd6532fd3c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 23:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3a6bb5a8-cde1-47e9-aff5-75e878b3afe7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.20', '0.11', '0.02', '2021-03-11 22:24:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3a71c3b0-8c7d-4025-90ec-061f669b3b2a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 14:46:34', '0');
INSERT INTO `monitor_host_log` VALUES ('3a7809b6-41d5-48db-a111-ee20692f0ab7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:01:07', '0');
INSERT INTO `monitor_host_log` VALUES ('3a800744-106a-401f-bad1-131f80195ee5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 23:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3acf9176-3aac-44e1-a94b-0a7dc7827f05', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3afbda67-70b4-4483-9954-78c037a50a3f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3affd475-f873-45f5-8b19-934adf5fa5a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-08 00:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3b10083a-435a-4246-b4c9-52def2ebffcd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:03:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3b29b17a-1768-4314-be30-c959e25b6a53', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3b2c7027-4323-4cc0-a0dc-e086001f90e7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 16:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3b3d8d43-8ca8-43c0-9715-461b6fd3a0d1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.21', '0.53', '0.10', '2021-03-10 20:40:06', '0');
INSERT INTO `monitor_host_log` VALUES ('3b4b8c49-8755-4147-bd12-182a0f642e3b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-07 21:40:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3b505fe5-10bf-462f-bcda-70c3909fb570', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3b5c98b0-feeb-4183-b46d-a8af63019ee9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.23', '0.54', '0.10', '2021-03-07 15:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3b615928-3343-4c17-8e20-cc8d31d8805d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3baae645-4652-40b8-a5a0-ea48de685e5c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3bb6e896-4141-481f-a032-252131029952', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3bbb076e-f8c1-43ea-bd1c-e3c9c3ddd1b8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3be707f7-9cf1-49e0-a0a6-f60c6b1a07e2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3be92881-34e0-42a7-bb95-fa078bc8b2e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3bfb9b98-d712-4c1e-bd02-c23a8460c6ba', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3c040e31-7cc5-4866-907f-c93d3f757e45', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.16', '0.54', '0.10', '2021-03-07 10:27:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3c19db60-28ed-4aa4-88ec-bb8b13c6bc58', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 20:01:05', '0');
INSERT INTO `monitor_host_log` VALUES ('3c42ed23-650b-4c02-b7a2-fe9ae0b6df90', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3c498698-ba3a-4a7a-b380-89d945eddc42', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 23:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3c84e818-a04d-4a27-ae19-b042ea9ab825', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 17:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3c9200b0-ef92-4681-8f41-f81e7729c7b1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:07:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3c949e91-27e6-4cd7-9ccb-8d0b8e12b467', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 22:48:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3cb2215d-1262-448a-95b5-dd7703278cc3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3cbca478-1f3a-43ce-bf60-95f4bc0a6a99', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3cbd0242-f361-4093-a56d-86c91b6cc687', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:01:06', '0');
INSERT INTO `monitor_host_log` VALUES ('3ce40c6b-d1b0-467f-be96-319d2995b08e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 16:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3d0dfd0a-ac2b-4ef8-88e8-350ab0b88aef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3d3cda78-4e76-4670-bbbb-4dda787a7fc2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3d4232a1-e5be-4d6b-b5fc-fbca2f9e1a44', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:19:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3d7f3e49-debb-4311-a530-afe039d9f6c2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '15.30', '0.54', '0.10', '2021-03-06 17:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3db37964-d357-4c63-9d5d-ab44b3a25dbe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3dbba58e-c0a4-400a-9c00-55ff1341eaa7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.56', '0.10', '2021-03-14 12:17:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3dc13acd-457d-4e3e-9f48-f998a92d2342', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3dc86619-2a74-44f8-b990-83a9c0f54583', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 18:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3dda2f08-bc7a-4d13-98d7-d3bfb0b4896d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 12:06:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3deb7801-3649-4dde-b79d-e56d14cbccaf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3e0aabb6-668c-4c82-a458-e734b0839f9f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:56:05', '0');
INSERT INTO `monitor_host_log` VALUES ('3e42fb4e-8765-474f-9649-5a03d4e26dea', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 23:19:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3e488eab-9bef-49d5-98e4-86af580f84b2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:27:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3e513c55-02bd-4e07-b288-2d855d146478', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3e607053-b974-4379-865b-8c7629216cd4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 10:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3e6c81f8-03b1-4ff9-ac8f-5920a793d78d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:34:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3ecc5f30-a823-4edd-a8be-8f1eea2624eb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3ee1d3fc-a46a-48af-9bfc-cf3d8dfa37fc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('3f1353d2-c720-40d2-8aaa-d874c1de285c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3f2000f5-54c3-4121-98bd-d8a6c6b57818', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3f335405-cae9-485f-93c9-5fe93c3d41bb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 16:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3f57fa35-a0a4-4c6c-92b7-f543cca3ba48', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 22:59:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3f7a073c-b60d-413d-902a-decae06a720d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3f893fcd-b780-46f9-a433-c625a25fbdc0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-11 21:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('3f9d08ae-1530-499f-9e78-0cce7ed95bca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:26:01', '0');
INSERT INTO `monitor_host_log` VALUES ('3fb62705-069c-4fa8-99ab-688476a8a862', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('3fe59cbb-0985-4f4e-8652-c181dc11ff7f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('400ae656-99cc-46e5-88de-3237678107d8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-07 21:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('401170b4-9fa7-4000-ab0b-f59aee126965', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:46:03', '0');
INSERT INTO `monitor_host_log` VALUES ('40189857-a7ac-4bc6-9d0a-936ba89b6faf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 21:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4058f832-c22f-482b-8cde-a23c1861e050', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:51:03', '0');
INSERT INTO `monitor_host_log` VALUES ('407dee03-1112-42ca-a4e4-f2e61dc73812', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4087e8ac-8166-4e2b-818c-bc642b9ec360', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 12:20:01', '0');
INSERT INTO `monitor_host_log` VALUES ('408ccdee-6a43-4227-9aa9-49599fda6d42', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('40b9068b-f507-4628-a7e2-d10d73a3f048', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('40ed21a6-b0c8-4d35-aad0-68aba3000e8b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('40f09c24-25d0-41a2-a161-8fbc8c68c422', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 09:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41021bfc-7772-4026-9f87-4fda225f4d32', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 18:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41297308-3eff-416e-aebf-13cdbdb3d7b5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 19:42:06', '0');
INSERT INTO `monitor_host_log` VALUES ('41447668-1424-4870-bf4a-64657b090311', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41496b51-a1c8-402d-b23d-36b9706df287', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 21:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('416936c4-2ebc-4104-bac1-e22a372389e7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 19:56:06', '0');
INSERT INTO `monitor_host_log` VALUES ('41aded10-d145-403a-a52e-35d9b155ae3b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('41b1533b-4768-4fc8-8585-ba9cf8becdd7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41bfbd17-31b6-4d5f-9745-b7c0f1d2f68c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('41c24eb8-ef84-466e-af06-45dc65196625', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41d94b32-5e08-403a-8561-870112309b59', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 15:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('41de1a27-be02-4368-85c1-fa2bcfebd60d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-14 11:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4211d9b2-cf80-4399-b7a6-537b08e64639', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('421b5d7d-8de3-42b9-a82b-a7b0a5099196', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 23:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('422f53f5-623e-4a55-956c-490d88c703ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('423eaa3d-91f0-4fb8-bad7-6cd188914650', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.70', '0.54', '0.10', '2021-03-06 16:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('423ebeb2-dcfb-4cac-8f1c-d6101db3eafe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('423fddb3-085f-4b90-9c28-f088ae2c419f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-09 23:19:01', '0');
INSERT INTO `monitor_host_log` VALUES ('424ba46a-4cfa-4d2d-b94c-b370d912e7d5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4271f395-900b-4a8b-8b12-aad23de31df6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:11:01', '0');
INSERT INTO `monitor_host_log` VALUES ('42738c82-2e42-4412-8f3c-0216084775ec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-10 21:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('427ed47a-cec5-41af-929b-6ea6b08b4410', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 15:33:01', '0');
INSERT INTO `monitor_host_log` VALUES ('429e6f43-02d2-4cb1-b0c1-1febbcc6fa28', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('42ac9453-aa95-4512-92bf-171e98888adc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:47:27', '0');
INSERT INTO `monitor_host_log` VALUES ('42bca319-7424-4b3d-bd2c-d8071987a9b5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.16', '0.11', '0.02', '2021-03-11 21:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('42d34abe-d86b-46a2-95ef-c83cd6d19bf8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-07 22:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('42f41612-57a0-4ada-ac30-e2fa7c43606f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('42f4acc0-076c-4839-bed7-95da1a718710', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.55', '0.10', '2021-03-07 23:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('42fe08ec-7d8e-4ee8-b8bc-d1b0ba1b24ea', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('43014994-2ea6-453e-ac60-b584292411c6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 18:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('431fc747-58b6-412b-8a1c-6353efc1bf79', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 19:52:06', '0');
INSERT INTO `monitor_host_log` VALUES ('4322dd26-0b4f-45c0-97cd-77f0c9e5fb67', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('434bc12d-3d58-45cf-a536-ab190c73cf89', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('435e22dc-ecbd-426d-8843-56a87e34e68e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 22:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('436566ef-c7f6-4ab6-931e-e291584ee957', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-07 09:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('43661000-6127-4d47-82c6-8bcb6817d6cc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('43708888-f554-4d61-b0b4-53b3790d25fb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('43be0c9c-f4de-487d-a47b-64c2703095fa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-10 21:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('43f993ad-5cda-4533-b64a-fc8a17c1d0dc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('43fc3e8b-694e-49b9-8c5c-c1e06fe19834', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('442b5856-218a-41c7-a983-08ec9ae009ea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('443cc8e0-1eb1-4da7-b645-3a96d0ff0078', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.54', '0.10', '2021-03-11 22:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4441cabe-802b-46e7-b3a1-004617e444a8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:39:06', '0');
INSERT INTO `monitor_host_log` VALUES ('447609a7-4351-4b46-99d1-b387ca035551', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:17:06', '0');
INSERT INTO `monitor_host_log` VALUES ('4477a445-88e8-4dc2-88fb-d23fcee88c17', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 21:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4482c2d1-a552-4718-af08-52cc9589ef90', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('449b3276-e465-47e7-aa05-75b352f0667a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('44bb059b-787d-4a1c-b03a-e37a26129411', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('44c9aee4-f6b5-498b-9468-2cd3a60daafe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 22:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('44febaa3-cb9a-4237-95ce-6faba56d3b5b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 10:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('45063c91-af5e-4a14-94e4-605405f2f3df', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('450d3ded-6e99-46ff-903a-ca72cfe34622', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('451ff4d0-8552-4485-bd3e-099b68bd3f0d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:46:29', '0');
INSERT INTO `monitor_host_log` VALUES ('45438cda-dca7-4249-ade7-1e1d86f05afe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('45571481-d088-47eb-ab15-09f9aa43545c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4561f43f-abac-41fc-bcab-8b69ff60db2c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.13', '0.02', '2021-03-06 23:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4563f0ea-62de-4dbf-b903-65bf83e1a1f4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 12:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4567e1cc-e0da-44bf-bdd0-aa2f732f7de7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4578f65d-9815-4ee0-bd8f-05e524971a39', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:38:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4580ae49-0527-40f7-a0a1-bb769c22e23a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('458c4110-3787-4528-8ce9-0a4522011384', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 10:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('45931d7f-53c1-48fa-a273-620f2050dca5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4593a76c-22fb-4c1b-834f-840e2b8c1e92', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:42:08', '0');
INSERT INTO `monitor_host_log` VALUES ('45a71d54-4eaf-42ac-bbad-bc436a30bd81', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('45cb638b-5af8-4c0a-a28a-20ca2f920086', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 14:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('45ccf0b3-66e4-4d39-b82c-14ad22d1311d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('46128176-0f43-42ff-99f2-83b30fee5c71', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.13', '0.02', '2021-03-06 17:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4628db36-dbd1-4030-801f-754fe279b8df', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-11 22:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('463fceef-9808-4d65-8a30-f5b37d582622', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('464c3d76-7cb9-434b-b8f5-ece4e142ffb8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('46630488-8aff-4ff8-8e68-af983b246b09', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4683a23c-23fe-4cd7-b722-ae9d6c4756df', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 18:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4684f92d-7f4e-4cdf-8056-c2bf81fd8d00', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-07 08:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('46d50f52-f7a5-4622-9d23-1c201d13d08a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.03', '2021-03-10 20:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4709cf52-5a12-42f0-9ba6-c7d91d8df72b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:15:01', '0');
INSERT INTO `monitor_host_log` VALUES ('470b3bcf-efb5-46c5-9e59-3317b08961b9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-07 09:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('470d9cf4-4c39-4446-bfa1-c73f38270dc9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-11 21:48:05', '0');
INSERT INTO `monitor_host_log` VALUES ('47671e29-3179-4184-a58a-d12583da85d5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-13 22:38:08', '0');
INSERT INTO `monitor_host_log` VALUES ('47675316-223d-4587-ae02-70b47e7f269b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('476be771-8ad1-4d8a-97ec-15ac3a2ee334', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 17:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('477d6c6b-4dff-4c93-8a3b-7e9c04e57a96', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-07 09:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('477de332-dfac-4295-b10f-dfa2ea1b7056', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('47b2e500-e9fb-42c0-80b9-1d856564861d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('47bc397e-c91d-4fcb-92f2-f325f55c895e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('47cd8f15-d584-4cc1-b0e0-5968d63f28af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 17:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('47d8925a-41ff-4c56-8896-d279fcc5d15c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 17:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('47db8fdf-1faf-4c31-9c9e-43e0f18de089', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('47dbab4a-c196-4f1d-8fa4-19137a828121', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('48154603-cfac-4992-b247-c080e4fab2e8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.12', '0.02', '2021-03-06 17:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('483fbbef-9f6c-4fee-8652-790d71ca02ff', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 17:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4845afdd-dbfe-4b00-99eb-b12a718d755b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '7.10', '0.53', '0.10', '2021-03-07 00:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4860227f-2535-406b-b62c-0e3e6aa6eceb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 23:54:01', '0');
INSERT INTO `monitor_host_log` VALUES ('487261d4-87ea-45e6-8155-ebd0e603deb0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.54', '0.10', '2021-03-07 21:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('48772914-0a4c-4a3f-8158-19c441b1b222', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 13:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('48788f05-6fce-4574-a3a5-422afda53e8b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('48803090-dcea-4501-b3be-90821ea78697', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('48a75622-5c00-45d4-bf86-fae0ca05eddf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:07:03', '0');
INSERT INTO `monitor_host_log` VALUES ('48b2711d-43a6-4b15-9dc1-6bc485936e34', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:10:03', '0');
INSERT INTO `monitor_host_log` VALUES ('48b7b62f-3de6-4bc8-8da7-3001cd4424a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('48d5402d-ac22-43d0-b9f6-ed8736f696c0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 17:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('48d88584-7b65-48af-83bb-53fd2423676d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.53', '0.10', '2021-03-08 20:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('48e782c7-0ebd-4c52-9152-e2b628afc20d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:10:06', '0');
INSERT INTO `monitor_host_log` VALUES ('48fa540f-fd52-4f35-bfa4-1e50d575d433', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 17:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4905c32f-5901-4437-a72b-a727d53338da', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4918179c-34d9-4dff-8a0b-6b0b896bb936', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('49231b46-0e7c-46dd-b730-1cbee2f01fd6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4928c264-7217-48d6-9cc8-1d140a5f4583', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('494860e9-8987-430c-bdc9-d97f3d30cbc8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4964e29f-8621-4d77-b1f6-b01b451f2f17', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 09:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4992c660-fe57-46a3-bced-c069fd945e4c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-09 22:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('49b60abe-8387-4332-b546-b4d77a98877f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-10 21:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('49d6cdd5-9cef-4cfd-95b9-afe1b0161dba', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:40:06', '0');
INSERT INTO `monitor_host_log` VALUES ('49fe121d-3102-429d-89e4-a16bb875a03d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 14:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4a0c1cce-795f-425c-a302-8692e45f75d3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4a25d13f-cb6d-46ee-a497-20eadb220389', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4a31e85e-5286-4df3-bfa6-da7266004c6a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 22:59:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4a503472-33ed-48e4-8df5-74870f50b3ee', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4a57fb60-5812-417d-9cf9-38c91682d623', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4a8a4c1c-a9da-4db1-9401-8852c1d84b0b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:01:05', '0');
INSERT INTO `monitor_host_log` VALUES ('4b16ac26-d919-489c-88c7-3c73f7a97e64', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4b1c95f3-9411-42ee-9a2b-ba82f7ab3c04', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:14:01', '0');
INSERT INTO `monitor_host_log` VALUES ('4b2a124e-5d45-45b6-aaf5-d86f6c913f7a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4b4b81b7-6d2e-4169-a536-00641a7c65c8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-08 20:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4b4c035c-34da-44d7-9908-fc0d7ce5a530', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4b4d3a7a-35b5-44de-abad-9b0303c26702', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '8.70', '0.12', '0.02', '2021-03-06 15:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4b792cbe-5911-4335-a7b9-9ba5ad2bc161', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 14:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4b7a8465-4145-4f25-8f1a-1c023087a61c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4bc4e8f6-e9e5-4218-af2a-f9796e13c8df', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 17:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4bde59bd-c1a0-42e2-935f-e7995ad7eee3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4c414279-804e-474e-ae1e-c453b53e3c97', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '8.70', '0.12', '0.02', '2021-03-06 16:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4c45bfd9-ca65-492b-83db-b8f6c3212e76', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:49:07', '0');
INSERT INTO `monitor_host_log` VALUES ('4c7c6609-f30c-4747-97ab-d00b96f57082', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4c82839b-3397-46dc-b2be-5c69e331c713', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 18:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4c94ff85-788e-4ce7-beae-c3f40d9d7fc1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4ca727c1-8be3-45e3-b7e3-afbd724eed0f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 15:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4cac2687-ffc9-4959-9970-684a55f13883', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4cd4c5ec-b144-4f26-b87a-ba8a5d149b81', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 11:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4cdf76ae-73e5-485f-a0c4-8d81a46a8a22', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.53', '0.10', '2021-03-07 08:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4ce698cd-3fbe-4193-b108-2d254130ba7c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4d62fee1-caf7-456d-b075-900649363680', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.54', '0.10', '2021-03-10 21:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4d6e6c7e-88f7-49c8-841d-127768a1704e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4d739dc2-e3d1-4ae7-a619-7ffbc4416194', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4d7f2ef0-de20-45bb-a5b4-b9b89e924cb4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.57', '0.10', '2021-03-14 13:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4db95380-753f-4e6b-ba4f-8b5517185c94', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4dc342ae-c8b6-45fc-8f13-2f5ae7c3f59f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 10:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4dd09b23-f72b-415f-b868-01c7e632d96d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4dffb028-af81-442e-a3e1-08123f350ff4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4e227093-0e2a-4e75-886c-4cd411f1124e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 11:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4e24ea9d-dfe3-4517-a1ee-edd6e1aa3f85', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4e35f263-e400-4854-998a-0365627c2c48', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 17:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4e3c6fe8-7494-4c08-9374-913fe053fc75', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4e4070ac-84a6-4bee-af26-d8370422ade5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4e4684af-b963-4c36-9517-8cd750384aca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 16:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4e52f690-5c33-4d35-a4fb-cc7379381768', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-07 21:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4e83295b-e768-4b96-9f5f-4e40cb867738', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.56', '0.10', '2021-03-14 11:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4e9b867c-1dcd-4904-ac58-ddf7530db85c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-07 09:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4ea5f0a8-d52f-40f9-bc72-803a12a83fb4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.02', '2021-03-07 23:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4f07ae1e-51ec-4ba0-bdbc-f2add15d6baa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4f087199-a490-4732-acbb-c23f79a93b4a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-11 22:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4f4c0c78-1ba8-4f33-9f73-9cd45871a26c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 21:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4f5a87a6-68be-43d0-b8da-9bc1d3073266', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 15:11:05', '0');
INSERT INTO `monitor_host_log` VALUES ('4f5c1b0c-20b6-413d-a0ff-0f772e13ea4e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 08:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4f704f1a-adc4-47e7-8ade-34ac94cd5c97', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4f7ad5a9-3b6e-4c06-bb3e-cdfe288078cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4f9f8def-64f8-434a-a762-8d5c5f88a5ee', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4fa306ec-535f-4e9e-9ba0-3f47420b225a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.60', '0.54', '0.10', '2021-03-06 15:00:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4fad8d04-335c-411c-a6c2-b6214fc2c5bc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:50:03', '0');
INSERT INTO `monitor_host_log` VALUES ('4fb7b365-8f73-483d-9764-c674b9bfbaa0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4fc31dfa-c188-438c-acbb-fb289d76088c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('4fc4d24e-2946-4313-98f7-f18d518c1713', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.54', '0.10', '2021-03-07 19:52:05', '0');
INSERT INTO `monitor_host_log` VALUES ('4fc58ced-d56a-4443-92da-dff16ee98bc7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('4fdfdc77-2a5f-4ddb-9c3d-246fe7793b0e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:47:03', '0');
INSERT INTO `monitor_host_log` VALUES ('50063fb6-6a85-4717-a456-aad6772ca30e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.12', '0.54', '0.10', '2021-03-07 10:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('500eb44a-0256-41f2-be4e-bae3833f0cdd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 18:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('501993a5-e7b8-45ce-9684-d3b6e8e66e6a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('50228a6b-2167-4651-b816-f95f43cd3f2d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('50229419-06e2-407c-ae99-999854ff85d4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.02', '2021-03-07 23:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5032abdf-01ea-48a5-ac57-d04727a0e536', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5045e1d5-7bd2-4f0b-9314-a639653ebc0d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 16:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5065cacb-856b-450b-a32b-7fe27c1d4007', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 17:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('50727e77-4e43-4108-84df-f6f666ffa51f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5073733f-c1f7-481b-b5a7-54f2867b5aaa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5082315d-1e3b-4d81-ae45-24db86ab7884', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5098fd3d-5191-4a02-8314-66303a90c711', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('50c39d5a-1c53-43d9-bd57-f84e3ab47517', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('50f0eec1-1d75-43b3-8db7-89b2abc7071b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('510b9b9a-bdfa-4732-9d03-962fa560bb83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('514d4e78-01f7-4df6-afe5-e64c59d7198d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 00:51:06', '0');
INSERT INTO `monitor_host_log` VALUES ('515e12ed-a18d-4aa5-83f1-25e961d836f5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:32:01', '0');
INSERT INTO `monitor_host_log` VALUES ('51628b30-97d5-4852-a0d4-f081551a7791', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 19:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('516a6e23-2456-4aa2-ab5f-075472de8bad', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('516ace86-7e8b-4cc1-8451-8880f90ea2d1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('519333b0-10e8-46df-bbf9-f9c93df2879d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('51961d72-082e-46eb-bf40-28e7bb81d839', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('51b3f8c5-3a6a-455d-a7a8-4ddb22f5e4d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:46:06', '0');
INSERT INTO `monitor_host_log` VALUES ('51bc5bfc-3558-4d8a-ad44-267e9f07b2b2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('51cdf262-a657-41c1-b33b-cbd0614836dc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-07 09:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('51e5caed-04f2-44a1-bb6e-a7892bf940ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:47:01', '0');
INSERT INTO `monitor_host_log` VALUES ('51f320e5-b21d-422b-ae14-cb5e5aecb6c7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 22:30:01', '0');
INSERT INTO `monitor_host_log` VALUES ('520275ed-2290-44b7-916d-7bb4f6ce5dae', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('520bb611-470f-4450-bca5-62756f9b7344', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:08:06', '0');
INSERT INTO `monitor_host_log` VALUES ('520fe149-22cc-478b-a2a6-c7df06f535e4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5213ea49-e36f-49af-87a8-a47adb613cbb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('52173a5e-642c-459e-b9c4-647ed8e932ea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('521d2955-3bc6-4534-abd2-819fed1bdbcb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.56', '0.10', '2021-03-14 11:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('523d182c-370a-4863-b83c-9aa7ec44915d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:31:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5245d700-c987-4f38-baca-528588bd7816', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('524759ce-cc95-4c98-a410-eac7af2fc3d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('525164ea-b497-4514-80a1-664de6f5f183', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '20.20', '0.54', '0.10', '2021-03-06 16:56:03', '0');
INSERT INTO `monitor_host_log` VALUES ('525a04e6-3431-4771-afac-03516248c8df', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:57:05', '0');
INSERT INTO `monitor_host_log` VALUES ('526199bb-43ae-426b-b5e8-0e41ad0ee8d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 17:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('52bb3706-b8a3-4ee1-af81-0e89ba921af5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 21:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('52cce2af-169c-4683-ad83-e8caa22d8da4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.18', '0.11', '0.02', '2021-03-11 22:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('52d1f874-0edf-4c4b-9e60-39e6f56c5931', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('52eea305-3cc3-4c7d-8dee-61561500cc87', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('52fdc196-3cb8-42a0-aa20-65abf32cfb58', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:38:06', '0');
INSERT INTO `monitor_host_log` VALUES ('530b4267-9153-4c53-bfbf-da22ec63ce65', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('53429e43-6545-4a33-988d-be84ce55ce21', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('53e285cf-9330-47a7-82ab-4168b05a86b1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.15', '0.02', '2021-03-08 00:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('53fed598-e19c-4e54-b21a-73f60f196648', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.57', '0.10', '2021-03-14 13:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5426be3b-9e62-474a-8530-f6cff28c9aaa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.00', '0.12', '0.02', '2021-03-06 15:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5437771d-bc8b-4c93-982e-e282e43b15e9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('544f5209-0d0f-4a07-864e-3644b1f3d636', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5464f499-0cb3-4537-98c6-033bb9e7327a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('547cdd71-5556-4711-aadc-b5e5aa657a07', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '32.60', '0.54', '0.10', '2021-03-06 16:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('548d502c-5dae-41e5-9be1-c8985729dd7d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('54ae5260-3e85-484e-94e2-0b9c54e8bfd9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('54c07051-d5c0-42a9-90ff-59841a7b2b2c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 17:16:06', '0');
INSERT INTO `monitor_host_log` VALUES ('54d58718-5379-4613-ae97-2632af5902da', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:03:01', '0');
INSERT INTO `monitor_host_log` VALUES ('55142151-fa27-418f-a06a-1d578c26c7ad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:47:06', '0');
INSERT INTO `monitor_host_log` VALUES ('555011cb-9961-457f-a944-efb072f2171b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 15:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('55622f4b-46a7-4385-87db-ecbd9b2cee33', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('557559da-8ba6-483e-b227-af43cf5da91f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 11:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('558cb84c-b49b-4799-a1d9-89c91117231f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('55b34b60-e96c-44fc-8e68-42da7fca3d49', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 09:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('55c1d27b-83bb-4ed4-ba5b-b8cdee1f9369', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('55d0cdea-f14a-4bb1-9dae-2cb670c7609f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 19:33:03', '0');
INSERT INTO `monitor_host_log` VALUES ('55ddea6f-4f06-4a90-b16a-6e381bdf4607', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('55e4d375-736c-4373-b07b-83ad8326b35c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('55f4bc6f-bc29-4a75-9a87-2361f9fae01e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5646768c-e79f-462a-a6cd-9d7c6cf0b601', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 23:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('565231f9-1320-4798-b2bf-76bfe5e51e54', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('566c85ce-39f0-4c1c-9126-07c0dafa8467', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('567504d8-b049-4419-bac9-05e29d762b4d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-11 22:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('568c8d56-c6f6-44b8-93b9-fce8d37ec99b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-09 22:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('56aa46f7-02cf-47cf-92ed-62d20a732d72', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('56bceff8-7b96-4cf0-b69c-bad85808f9a8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('56c93ee1-b70f-4640-8c11-ccbc9c6a8149', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 21:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('56cb0ff1-f730-4824-8e14-2ecd76c33136', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.57', '0.10', '2021-03-14 11:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57119836-60d9-4345-b4d5-9255b2e6c7f5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57653a16-abc7-4265-88e1-e8c9a3e8a355', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:56:03', '0');
INSERT INTO `monitor_host_log` VALUES ('578774c0-f4bc-4230-8132-466ceec4b8d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57a974f2-60e8-4ccd-87d1-fed9aaec07d2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.56', '0.10', '2021-03-14 12:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57aa89e3-d2ec-4253-9afb-2d6f18f88b4e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57b3992e-4c20-417d-8ac8-18603d3cd5c9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.20', '0.53', '0.10', '2021-03-06 23:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('57caa2ba-205c-4e6b-8282-b0b0387659ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 09:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('57d6de5f-853d-4e1b-8fdf-f04a34015999', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('57e00b9c-7b01-4f09-95ce-79ef9f4dc94b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('57e74971-b8bd-440a-ae61-41137239b761', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('57f10112-3f0b-4f56-9242-bafd1eba62a7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('57fbba9b-c608-46d8-aef5-e704c8796735', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('583151cc-fad2-459f-b793-7f3bc2c7fc41', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 22:49:01', '0');
INSERT INTO `monitor_host_log` VALUES ('583f1a16-203c-4946-9b49-e22f06133d0d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('586a97f2-f4e3-474d-85c0-fe0678022169', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:37:06', '0');
INSERT INTO `monitor_host_log` VALUES ('5874d5e1-48b5-4deb-bd82-37c5b4ee407d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.75', '0.57', '0.10', '2021-03-14 11:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('58753650-74d1-4024-8401-ed4b81486647', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:37:05', '0');
INSERT INTO `monitor_host_log` VALUES ('5884a580-01ba-43d8-9b6a-589fbdea3d9f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('58c6284b-a97b-44c6-8372-f072930a2150', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 16:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('58f4a509-9095-4d4f-a003-3ebf54e6b41a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('591841a4-0960-47c9-a758-b0d565f83ca8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 19:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5918423d-fc0c-4206-a970-d47dbd77ef5d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.30', '0.53', '0.10', '2021-03-06 23:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('59521dcf-b2f4-4fdf-9bc3-df9eb0860e5c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:36:05', '0');
INSERT INTO `monitor_host_log` VALUES ('5958ccfe-c85a-4930-bd7d-ab822b3de246', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('595b38f6-0116-4e21-b707-29b83605377d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:16:06', '0');
INSERT INTO `monitor_host_log` VALUES ('595fd910-0e3b-4746-9b32-9c91857a4bee', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 22:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('59736451-a964-4e38-9250-b979783d07ae', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 17:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('598722fb-5aef-42b7-a8b5-da37b3ee1268', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('59ca95dd-1591-4a99-825e-5653647471bc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('59e3299e-cc2b-4038-a74f-254b51bd4817', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:29:06', '0');
INSERT INTO `monitor_host_log` VALUES ('59f2b885-9981-4007-af7c-a0678eb5ede6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5a2a022c-0604-4aa7-8713-3b98c1ac0255', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.78', '0.57', '0.10', '2021-03-14 11:29:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5a939b4d-8392-49d5-83f9-4f666e111b09', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:49:31', '0');
INSERT INTO `monitor_host_log` VALUES ('5abec8b4-ee5a-41e0-a227-f46f4214e535', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 12:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5ac01281-9a90-4a3b-8518-1bb3866f15cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5ade11da-d65e-4ee1-8461-6760a433d1d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5ae89537-029d-47cd-b914-2c455398fa81', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-07 08:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5ae8f2ce-daa0-4b82-b566-ab1518a16fe8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 10:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5af9cea0-6bd6-4efe-9dc8-04448248f845', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5b53b4a1-a7aa-4812-be0c-13814cb2334c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.53', '0.10', '2021-03-09 22:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5b7c2884-98e7-44be-b32c-0ce7f26083ca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5b92afff-9b43-484b-99ef-7ad4a29e14c5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 18:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5bca8650-195e-4ddf-9b57-bda90c8c706a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5c0012e4-b9eb-4b6b-a10e-39d3316d3e9b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 21:46:42', '0');
INSERT INTO `monitor_host_log` VALUES ('5c0f77e2-93ab-4b5f-9e3c-f27599fd470a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5c15e6ad-8b0a-43a1-9d40-0b7349acce27', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 22:42:01', '0');
INSERT INTO `monitor_host_log` VALUES ('5c236ad9-1987-48af-be7a-903de7391a2a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5c3b2220-df15-45d4-b2f3-ddb0fd3db34f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:14:06', '0');
INSERT INTO `monitor_host_log` VALUES ('5c4b8d5a-2b70-4783-b6c4-21278202fb06', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5c68179e-ad03-4c9f-9820-7b963f5c921f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5c6b1010-fd9d-49ee-97e8-e75fd827aa73', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5c7b0a19-4a51-4198-9748-ff8426770003', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 16:26:01', '0');
INSERT INTO `monitor_host_log` VALUES ('5c91060f-b52b-42fa-b02d-4264af4ebd20', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 17:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5ce05404-6755-4221-a61d-808986933195', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5d1a15dc-f0da-42d2-a791-da85763f896c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 22:39:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5d5d57ed-0855-4aec-b340-4077da5a5c72', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5d72db3c-254a-4f69-8456-002969580fb8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:33:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5d8e23a2-94c5-4d82-84f9-be0be59e54d7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5dea8abf-f5f0-42b6-9e81-e1a59896c48f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 11:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5e4487e4-703b-4cbc-8982-6567a8557ef5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.54', '0.10', '2021-03-07 19:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5e833e37-2378-4575-b991-1bd6329ad343', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 17:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5e9456d8-a1ad-457a-8ef8-5ba8e7319030', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-14 12:03:01', '0');
INSERT INTO `monitor_host_log` VALUES ('5eae1b09-4796-4887-b96e-585dcc5cae2f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 23:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5ed66b1b-7880-4f5f-ba84-516b1e799f17', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5ee52c53-7d96-4c2d-a0bd-a7880f77aca1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 22:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5f2f9f30-b25f-421d-8b27-fbdc913000f7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:34:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5f658741-6a62-42de-b54a-d719590f35ae', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5f761114-bdc6-41c7-b0c0-1a7a2ecbf5f1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5fa9ac26-4acb-4862-9a18-62ff6c021fe8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5fc2bac4-7dca-4882-8c21-fb93103e2a31', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('5fcba8f3-b9ad-4814-8f8b-26aa34dd579e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 17:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('5fcdb9e7-fa36-4eab-8bc5-82653c7fdbe5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:15:03', '0');
INSERT INTO `monitor_host_log` VALUES ('5ff66a98-3fb8-40ef-850f-978e9c65aac4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.15', '0.02', '2021-03-08 00:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('60241630-0dfd-438b-ad38-15fcddd7743c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.12', '0.02', '2021-03-06 17:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6045c693-c8c7-4534-aa86-ba3e3dbce8e9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 11:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6054e6fd-ae7b-4738-9181-572c803414a4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6090998b-8748-4432-9189-d913b2decc07', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('60a088e1-c96d-4e5e-ab57-db2b49b6129b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.11', '0.02', '2021-03-11 22:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('60bf458b-e0d0-4979-a5bf-6039cefff64d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '13.50', '0.54', '0.10', '2021-03-06 16:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('60f9c71e-316e-4acc-b93a-0f0c303826ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6110391a-5281-4715-9493-319e29334a01', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:12:01', '0');
INSERT INTO `monitor_host_log` VALUES ('61471f1b-58f2-45e6-8d04-529fb30644f3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '14.60', '0.54', '0.10', '2021-03-07 09:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('61987e26-68c1-4f97-8c8c-54fbbcf3c8b0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 23:15:03', '0');
INSERT INTO `monitor_host_log` VALUES ('61b19ec5-667f-4ab3-a073-e74a52d03d92', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('61cdd1bc-17f1-4cc3-ba74-1c51786a14c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 15:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('61d0b402-9068-40ae-acf7-763ff50e2a16', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('61ddde10-40a3-4d42-a48f-8dd3705a4715', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:15:01', '0');
INSERT INTO `monitor_host_log` VALUES ('622ffc9a-0c2c-41fb-96b1-565016015a88', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 16:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('629f77dd-b084-4f9c-ab0b-636de07988a5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 23:48:01', '0');
INSERT INTO `monitor_host_log` VALUES ('62a371c0-ba8d-42cf-afe7-ab6aebf674b4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('62ad4922-d499-4851-887d-0a7afc0b6c53', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('62b087ef-ced8-4cdd-ac4a-2a185b528e80', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('62f34417-36b5-447c-b36d-703a303986f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('62fbf860-6e73-4ad1-8fb9-5f2d1789ee53', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 13:02:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6305d1d4-fd56-4d9f-81ed-3b074fd38a4c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-10 21:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('631f79e1-3ad7-4a22-8edb-e5cb70c7cbdd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:52:01', '0');
INSERT INTO `monitor_host_log` VALUES ('6320133e-5eaf-4576-8a77-5ed8fdb62eb7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:54:05', '0');
INSERT INTO `monitor_host_log` VALUES ('632b0f99-b45b-4fa6-bc75-75b44b4d3a50', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:39:06', '0');
INSERT INTO `monitor_host_log` VALUES ('63333257-86bc-49ec-a589-95fcff6c498c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:22:06', '0');
INSERT INTO `monitor_host_log` VALUES ('633e7e9d-fc22-4803-b42b-7bbe173e09af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:12:01', '0');
INSERT INTO `monitor_host_log` VALUES ('6343fee6-c70c-4ae6-a99e-b3b8875b883d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('634fc156-d5a8-4899-951b-1b1e93daecc3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 19:37:06', '0');
INSERT INTO `monitor_host_log` VALUES ('63be5359-4cbd-4cf4-aea8-1e95bede515d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.30', '0.56', '0.10', '2021-03-14 12:13:03', '0');
INSERT INTO `monitor_host_log` VALUES ('63c734b9-7866-49fe-8b22-3e0cc59d77e9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('63c7c0e7-930e-4639-9bfc-51139683b9d4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 18:03:03', '0');
INSERT INTO `monitor_host_log` VALUES ('63eee462-6292-4380-8bd6-f2a97dff4dd7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 17:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6415f422-43e9-4239-809e-5467b53ccc9a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 18:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6419bb5c-2ffa-4a50-95d1-895428cbff7a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-08 00:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('643f6e14-887e-4063-9701-047e84a5537a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:38:06', '0');
INSERT INTO `monitor_host_log` VALUES ('6443f16f-b9ef-40de-b46e-1bc5801194e3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('644bc341-767b-473d-bb41-5aed0fa9875a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 18:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6464f800-3d65-48f3-bf1c-0ec91c86e5f6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('646d7856-e4fd-43c2-86e5-22873f3c9b09', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6488cf2d-178b-41f2-8af8-34b72a53a136', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 18:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('648aa907-c188-4d48-a8aa-6f65c49c4938', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:24:06', '0');
INSERT INTO `monitor_host_log` VALUES ('648d684e-ef35-46d5-9979-3dc61aab798e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 10:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6506b15c-eaa7-4845-9b94-502693a0961c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6510a7df-9b2d-4551-b920-098168c0163c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6511fb08-ae73-48cc-98ab-042a0c7fe87a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:34:03', '0');
INSERT INTO `monitor_host_log` VALUES ('65190f68-3192-4400-90fa-5da51d6e5c85', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6524a2fa-cde3-4904-8e9d-5583a457b9c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 19:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6546fe08-df3d-4204-9472-6be665ac47f2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('656f91d7-5250-4195-8412-211ec581c4dd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6577c8b4-79ea-4c71-9c75-877281d816b4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.11', '0.02', '2021-03-06 17:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6579f7a8-41a7-41ba-9033-544989c05d67', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.76', '0.57', '0.10', '2021-03-14 11:31:03', '0');
INSERT INTO `monitor_host_log` VALUES ('65a09399-e47f-4a69-a092-589e175ae527', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:48:05', '0');
INSERT INTO `monitor_host_log` VALUES ('65dd01c9-f3f4-4c5b-aa7a-91325595ac5b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('65e1795e-0177-47e2-b6f7-67ac06b32cc1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('65e4b3f5-cb5c-4359-bec8-c4c32c729599', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('65eeca98-e8c8-4a74-897e-45e27769fc17', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 18:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('660308b0-87b9-4715-b45b-d990ac6361f8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-08 20:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('660ae171-79ce-40fb-8985-5b4fe0e52424', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('660ba3a2-8702-48a2-b769-65b9d9b9c740', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6617f40c-72eb-4403-b43f-3d894e3bb8f8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '14.30', '0.54', '0.10', '2021-03-06 16:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6624d019-1978-49f4-af63-f09aaf51efba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.00', '0.11', '0.02', '2021-03-06 15:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6662905d-9101-46c7-8a0d-26714fbf1a69', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('667cb427-3c14-48fe-bb43-8d05a3630da8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('669b496c-323d-4ef8-8cdb-f81c5abf072b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('66a3c283-9a2d-439f-af14-c06e88c1a2ad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.33', '0.56', '0.10', '2021-03-14 12:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('66ae5656-3ac7-498c-83f2-d40a0184f559', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-11 00:48:07', '0');
INSERT INTO `monitor_host_log` VALUES ('66e00c47-41f3-425c-8f3e-81f1a5b6fc44', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('66e28a23-ab25-43e0-b97b-ffa7dfb00323', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 18:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6700e424-6817-4b08-8e6d-d2aacea41f56', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:20:01', '0');
INSERT INTO `monitor_host_log` VALUES ('67309bbf-8973-471e-822b-215e1ad9f118', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6731c036-edbb-45db-8287-82c1b9498f7e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('67546aff-6f6f-41a8-ac40-0f39543b7511', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.03', '2021-03-10 21:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('675a3075-07b5-4301-9f83-5fa9176407a1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:20:08', '0');
INSERT INTO `monitor_host_log` VALUES ('6778bb1e-bc8d-425d-8eeb-16fd77535f03', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.13', '0.02', '2021-03-07 23:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('67940b0e-7eb4-4210-b8f9-ac316eb383bb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.56', '0.10', '2021-03-14 11:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6799e344-e4ca-4803-bbcd-82f3a299c754', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 23:23:01', '0');
INSERT INTO `monitor_host_log` VALUES ('67c34dba-0ff8-4f94-95c1-164d1b0e6e1b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('67d1946c-ba66-4140-a8f2-b4520827918e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('68195310-dad7-4c56-83d2-21df089b0c36', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 13:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6820f580-c0a5-4daf-b550-2390dfbd6a74', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 21:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('685ba5e5-d1a8-4d22-98ba-c5e09b935557', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('68629d8a-79b2-48d4-af08-05e973dd2ba0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('687be880-51ad-400a-8e3e-3c1eef694f4c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('68d167b9-ef9e-4820-add1-c6ff3226d482', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('68f8ad49-dd6a-410c-8b6f-3a41b1f84836', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('690c4739-db87-4bd0-b0fd-e99377f5dd9e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('69313620-583d-4a32-8b13-96c203be09d5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6932a811-5801-45f8-8b7f-c01cf59b7d42', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('694ada98-6c3f-4773-a688-51840d736ad5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('694fa294-d1d7-459b-8e1a-8b19e77ab4bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.14', '0.02', '2021-03-07 23:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('698b4a47-b993-4e21-a1a1-afcf3f23d61c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-06 23:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6990090b-e7f5-4d49-a08c-59a6773549b1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('69c2b86d-4962-4cef-a0c3-7130a975d7be', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 15:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6a0be653-b04a-49a3-8a27-ef71b7a57da2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 21:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6a307424-1e41-4bb3-8eef-39ac25742bc9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6a3ee761-36f1-4b89-9f5a-ce16d3b6b142', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 18:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6a5540a0-5bfa-46b9-afe1-127680404ade', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6a6bba96-510d-403a-b3d2-3652c26c5f37', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.13', '0.02', '2021-03-06 23:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6a8f31ae-87df-421d-a298-b07d5f5c5b8e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6aaf3811-2561-4ed3-8ccb-0f5409a0c19c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6abf166a-85a9-4875-8f72-bf6240e78bb9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6ad49fcc-dfcc-4357-8320-13b55dbeaf67', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6b32744e-3f7a-4f95-a5e3-1d4516414423', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 14:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6b36e016-0842-440e-85d9-1624368cef90', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.02', '2021-03-07 23:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6b3976e1-b140-46b5-802f-269dcb03528b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6b3c2a29-e93d-484f-9338-40785dd35068', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-13 22:48:08', '0');
INSERT INTO `monitor_host_log` VALUES ('6b67594b-b3fa-422d-9df3-efb4a0f221bf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6b7d5d76-8073-450b-92c8-8d82d5a81a9b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6b864b24-2f6b-457f-bde5-a456ac99329f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 17:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6ba73a31-6906-401b-814b-2c91e0a06ace', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6bab98cd-e7b1-437e-b992-d7fc46527a6b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6baf9c6a-82d3-4a6c-b641-e2086abc1590', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.08', '0.14', '0.02', '2021-03-07 23:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6bb64141-30cf-4163-b11d-a8139c392e66', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6bca2e26-f60d-4ee1-ac22-3073a3fee7b2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6bfcaeeb-745e-4af7-9579-755bfc76a2bf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:16:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6c03e69e-88f0-4b67-93f5-243ff9ec1364', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6c26aec8-0594-4503-b7f4-8964f2e2fd99', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6c3ed363-a6cf-4c3b-bca3-774a91bb1141', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:46:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6c64758d-bfe8-482f-be8d-ee6a27b6317f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 08:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6caad52e-91bc-459b-8c7b-ee28f5c2be36', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 17:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6cb967ea-d0cd-461a-a31d-44882dd8b68f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 23:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6cc91b95-6f2f-4b21-a0d3-8630cb8069ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6cd954b9-a1ff-4e2c-97af-40955896b197', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6cee9e7a-0f45-4005-92a4-4c05397e16de', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 15:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6d58400a-8cee-4084-acc1-b034d9c0d699', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.54', '0.10', '2021-03-07 22:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6d5b89fe-4cb0-44b7-b7e0-94bd667670db', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:57:01', '0');
INSERT INTO `monitor_host_log` VALUES ('6da418b4-11a3-4f4b-9b80-ed90e3542405', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.54', '0.10', '2021-03-07 23:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6ddd9966-caaa-411e-a6da-278761b17e7b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6e05561b-1b00-411d-acbf-0d30b5d66913', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 19:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6e10f5d8-7d9f-487d-99ab-e2c898cfeb34', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 23:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6e13c48e-a6f1-4804-b4c3-27fac0e7f4c7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6e2dc7e4-e63f-4a22-899b-5d5030c67834', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 15:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6e361141-9ad1-46f7-8be8-e480ddd0d731', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 22:48:03', '0');
INSERT INTO `monitor_host_log` VALUES ('6e3ded33-22ab-45c3-925f-c072ffa11c95', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6e95d9d7-d510-48a5-bd34-a6752045753d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6eda9cd0-6500-443e-a01e-f53f43322ac6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6ee29bdf-f465-4679-8160-cd0e6d9105fe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6f30f36d-f212-4ecb-b8c8-897b90f48987', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6f4ed33a-a3b9-41a9-bacf-d934b5aa3343', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6f599e03-8738-41b1-8082-ba19d1d56eb0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6f5ee64b-0928-4e4d-9f60-ca2f900eb2a9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6f86b53e-ad9a-4fb0-bcf7-2b68beda1e42', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 22:23:01', '0');
INSERT INTO `monitor_host_log` VALUES ('6f90be7a-194d-4fe1-9686-aa803b4f6ceb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6fad2b3f-2637-4929-a13e-38f0dc9593bf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6fc7157d-7204-43e9-9121-100ed4a08eca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 20:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6fc8f793-7d74-435a-a147-bab196498685', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6fcdc548-db2b-4d74-8431-46d619a068f8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-08 00:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('6fcfd18f-0214-482c-a513-6f8a9fc86262', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 14:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('6fd2ced3-ac82-4584-881e-7fb09b4bcc23', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 23:04:01', '0');
INSERT INTO `monitor_host_log` VALUES ('6fdbc9d8-7a60-4386-8025-1e83d9db3f67', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.74', '0.56', '0.10', '2021-03-14 11:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7019aed9-d261-44aa-a00a-c7966a0f5211', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('708de8c7-2f43-40d5-a2cb-87a011c80643', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7094656c-eda6-4f0c-bae4-dfb2a0808efa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('70a7ef2f-84b1-4989-a23b-5ed67eefbdc8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('70cf2b91-c29e-42f4-a232-ec837125f2a1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 23:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('710259ed-47fd-4562-8b83-2c1ea3287abb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.50', '0.53', '0.10', '2021-03-06 23:38:03', '0');
INSERT INTO `monitor_host_log` VALUES ('71133dfb-788b-4014-a056-83b9c64e2399', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7117d919-f7a4-43f8-82f8-ca32a85d82ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 17:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('71292799-e408-430c-abf3-81d52d232802', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:04:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7137a3d1-60a9-4bca-aa6f-a011d86f3e69', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('714f6de5-0d45-4a69-b696-61c32422aaea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:30:06', '0');
INSERT INTO `monitor_host_log` VALUES ('71809121-48be-4795-947d-7e010ebc4dc4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.57', '0.10', '2021-03-14 13:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('71b0eaea-53ea-4cb2-b322-1aad55a0aa75', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('71b58358-8273-4efe-937f-3fa9d8053611', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:23:06', '0');
INSERT INTO `monitor_host_log` VALUES ('71dd3ae9-2811-4a5a-b505-9257fbc5d337', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7233a5f5-85c9-4b56-a0cf-c2de00c07b6b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('72713b34-4d02-4933-b6b1-a0e4df9e7a6c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('72912b9a-6087-4374-89b6-35881ded2cd6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('729489e2-8a1d-4b6c-bb9d-d84fb02c9471', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('72b60816-8ca2-4831-8dfb-812b437f01b3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 22:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('72b793a6-08aa-428a-b22b-e113b39edff6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('72d229dd-da08-4322-b595-48e9951bcc56', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('72d58ac5-0514-45c5-b7cc-31a83e9053ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 21:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('72e63513-13ab-4adc-b6ea-9c5b5cfe1933', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('72e7ffee-355f-4aec-a639-95900c1ed197', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:32:05', '0');
INSERT INTO `monitor_host_log` VALUES ('72f05fa8-3b22-4e41-8df4-c0f31955aec0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7309db4f-b41b-4e22-bafc-7fe68c0eac7d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:35:03', '0');
INSERT INTO `monitor_host_log` VALUES ('73527b9a-7de1-497c-a9f0-6752bade7fd1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7373a881-3542-4b3e-ba63-12ace98b6c6b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:33:06', '0');
INSERT INTO `monitor_host_log` VALUES ('73a34bfc-564f-4642-ba8e-9cb27338f9c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('73c2087b-3e0f-48cc-bb93-ed8abb0d16c9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('73c3c3a3-a726-4f49-9781-9333d85b90c9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('73e4b6cb-b3a4-4d4d-903d-2735c03604a7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('740ed030-7a01-48ec-a392-655b4f04ba09', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 14:49:31', '0');
INSERT INTO `monitor_host_log` VALUES ('74220988-55c5-409b-8041-c285d31e140a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-08 00:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('74287eb4-7cda-4698-9674-0c00b8d8f103', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7445e7ba-619b-4630-891a-439a741c6a28', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('746ca33a-1871-4090-8083-e96bbf1bace5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('749127cb-e4c2-409b-a70a-d499afa174d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('74bbf0ae-cd98-4033-beee-186969236a09', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-08 20:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('74cacf13-9ad3-4755-be0f-248c20e358a7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('74cb81c0-dcdd-480b-a4f4-d5dd18bee524', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('74de4977-45af-434e-bd8b-a16bb6877c08', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('74f05b2f-909e-4ba1-abed-f71d00823b50', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('750e937f-0060-4cdc-a53c-04ab158dc629', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('751fff06-1c55-488e-a517-64c22c2b0b9a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-11 21:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('752a89e5-0083-4404-bea7-0bae78db4db9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 21:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('752d5793-3347-4d71-97e9-08e93860955c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('75481ed4-1fbe-4317-a4d8-c381d0c18999', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('75588ce1-eaa5-41c1-991d-807a5cb8f704', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 13:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('75777d7b-14c2-4b62-8499-4cbbf0e0d91a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7593bc21-9bd8-48c0-acdb-d64e458e29e7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('75aa9990-9ace-4174-b5fb-44b41cdf066d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:41:15', '0');
INSERT INTO `monitor_host_log` VALUES ('75b18f32-8cd5-4b85-8eb3-d7278fc780d7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('75c75f1a-a876-4123-ab76-e8c618d4e8d3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('760da4da-aa12-42c1-8835-0effe33b13db', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('762fae6b-4a4d-4560-a39d-22054e9e1299', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 15:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('76475b2b-a365-4990-9294-552a339f4be4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7676f916-ee79-4c3a-b0a6-b51e7420a45d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7681a707-1201-45d8-af35-3636b230e5f0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 12:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('76905d2a-7050-4fc0-86cc-ca26a11d8728', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('76b231a1-d7cc-4f79-a4c9-e404ec902575', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('76b2fd22-0521-4e6b-b99b-cbb9079144dc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-07 08:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('76c2894d-9b73-4096-9619-91fdb633a7c0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 17:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('76dbf1fb-27f3-448b-9bc3-843bde6b962d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.30', '0.11', '0.02', '2021-03-07 09:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('76deafb0-e635-47a4-9abe-052fc9d86694', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('76ef4689-2c09-412c-8d80-071b8f3a67d9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 17:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('76f2f4fb-0f56-4c4c-be66-018b444c73e2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 19:33:06', '0');
INSERT INTO `monitor_host_log` VALUES ('770484ab-db89-4678-ad48-58c27e2fb36a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.54', '0.10', '2021-03-07 14:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7715d0ec-38bd-4613-8a9c-b60680f905ab', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:31:03', '0');
INSERT INTO `monitor_host_log` VALUES ('77165596-bad8-47ca-815f-267fc7cb25bc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('77208381-fbc5-4442-8e11-4bacb6407a3e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('773bdc0f-be3e-491f-84d7-259d35178a6b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('77423a9e-bd00-4e78-b75f-fa939a61e1e0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('774885da-9a4d-4870-9431-f642a59bc27e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 17:38:01', '0');
INSERT INTO `monitor_host_log` VALUES ('775ad4bf-cd43-4d69-9733-1a26a7a914b4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('77665c3b-c6a7-432b-a5b1-a1e7d80a1911', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-07 19:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('77813536-db77-4be0-b85c-028f82dfd1cf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('778f1e90-70e7-4a0e-b9d4-46dcb4a673a0', '', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('779f0d62-c108-4f0c-bb78-0f6f52e12bd5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('77a3f8d0-b732-45e1-a3fc-1089ad074211', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 23:08:01', '0');
INSERT INTO `monitor_host_log` VALUES ('77b4e42d-caa3-41e2-b498-bc5618d936e3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 16:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('77c43290-f997-4424-a191-cb4b2af3ac8b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:34:03', '0');
INSERT INTO `monitor_host_log` VALUES ('77eeee0d-14b0-4c01-b692-cbabf1c3f6cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('77f93907-7a6f-4d17-8030-0770f7933282', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 18:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7803e96a-70e3-4754-aaf3-261b88645f23', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.12', '0.02', '2021-03-06 23:13:01', '0');
INSERT INTO `monitor_host_log` VALUES ('781033aa-7c28-4502-b4d6-dcdc1418246f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7813f208-0bb4-408e-bf22-161c52e3cae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 23:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('784e4582-da16-47d8-8993-9d3c88d17769', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 08:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7853fb70-8dd0-464a-9319-638f953036d4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.30', '0.57', '0.10', '2021-03-14 12:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7862d343-3e16-4e6e-873f-901f705a41be', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('786db638-84f9-4b4b-a90d-b53296e07056', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:45:01', '0');
INSERT INTO `monitor_host_log` VALUES ('78704f17-4824-4abc-9078-f23ae3433e79', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('787a442d-af4b-4e7c-acb1-4450c56c2841', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7893e8a1-27cf-4a1c-92e1-1a3093c678d9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:51:05', '0');
INSERT INTO `monitor_host_log` VALUES ('78a2f968-cd2b-4221-afb8-08b3a975ee20', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 21:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('78aae47d-ae3b-4160-b7c7-41ddc9fbfa1a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 21:46:03', '0');
INSERT INTO `monitor_host_log` VALUES ('78f4990c-d5de-4f18-a33f-909812ffe426', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 17:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('78fc8b45-3ce4-45ff-97e4-968b68cee604', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-06 23:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('79136794-64e3-477d-ad40-cb667999c667', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7936fa46-c12f-473e-9b0b-4b46c103e166', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('79403d29-b323-403b-b3fc-3cdba89dec81', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.50', '0.54', '0.10', '2021-03-06 22:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('795e09f8-1bb7-4de4-96ae-6bcf694ce2c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.03', '2021-03-08 20:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7969f620-571b-4ed2-ab18-8e49427128d8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 18:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('797fcd5d-07a7-42d4-abcc-52abba849bbd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:56:01', '0');
INSERT INTO `monitor_host_log` VALUES ('7987271e-c2a5-4bce-9a9b-c100c843a5d8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7987d4ce-77b1-4d67-ae43-7de1a88a8a33', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-09 22:17:01', '0');
INSERT INTO `monitor_host_log` VALUES ('7995f734-ed9c-4347-a143-6a9170ad2659', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 18:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('79998970-7d40-4b3e-a7ab-167d45de6da2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:41:01', '0');
INSERT INTO `monitor_host_log` VALUES ('79a0159f-a5e5-42bf-8919-9f350e5394ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 22:25:01', '0');
INSERT INTO `monitor_host_log` VALUES ('79ba77a5-d01f-44e7-9be9-15e1f449f877', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-06 23:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('79d35af1-39df-4208-ae6a-21579d1096f5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('79fc9aae-1799-4bbc-aa6a-013b4f7df4b3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 16:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7a10c652-fbc9-42d8-be1c-2123efb12f85', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7a1a440d-1ad6-496d-bfec-8fc74b74499b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7a27e77b-8203-4692-80f8-15f8db1e54b6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7a5c34c8-430e-4253-b31f-f3ca74dc9d62', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7a6856d7-f249-47a4-8f3d-71778f74cfeb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7a7583cb-d41b-4a73-8b1a-f883d299aacf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7a9d02e9-5866-4612-90db-8cfa09ab03e2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7b34e124-715f-409f-a7c0-bfecb3436364', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7b3cb2e4-304e-4df4-aed7-c4fe1faf35bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7b5ff684-2eca-4ace-94f6-f3eb1475bea3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-08 20:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7b6fdd3d-4ec9-482f-8f17-989d6e9268f8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 13:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7b8cac8c-b8b1-497c-92f5-eb1bdc9b67f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7ba1f85f-6684-45d7-8f58-71f2d4423e07', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.00', '0.12', '0.02', '2021-03-06 15:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7ba6bfc8-3891-47b5-ad08-44e918c2abb1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7bcde869-a015-4238-880b-209f98e781d5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7c042c89-5cfe-48f1-95ab-2b0762d454bf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7c1e8f20-8d59-4c20-82f9-7593c1fe76b5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:47:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7c262723-8746-417c-b313-4dc03adcbcfd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.56', '0.10', '2021-03-14 12:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7c269160-ca0e-4c4e-80a6-cec9d2253461', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7c4cfaed-b119-4c9a-b98b-06b2c2c478c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7c63bdd5-c1b4-4e1d-99b8-f936b290bd97', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7c792570-bec8-4921-a027-00c3d50e2512', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 09:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7c9f0cab-cabb-4031-842b-29ed31d5e4b2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7cdccfb9-b337-4278-8954-e6df8383411f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.54', '0.10', '2021-03-07 09:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7cddabd1-697e-45e4-bf73-540614ac9edd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7cf37312-df49-4539-a82f-84700fc25025', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 23:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7d2c3c34-6fd3-4503-8db0-6179e953e61f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7d34087c-4388-4690-a00a-f421ac89cf46', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:45:06', '0');
INSERT INTO `monitor_host_log` VALUES ('7d5e591e-20b2-4076-80f7-c9395befc511', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:39:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7d66166a-795d-433f-a04f-c5487bb391e5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7d6caddf-618f-47a0-b1ab-47558e633ce2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7d6f966b-58fe-4ca7-afed-cf924612f732', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-09 22:31:01', '0');
INSERT INTO `monitor_host_log` VALUES ('7d822aac-256d-4485-bea9-e2435c1fdaf2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 10:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7d87a9c1-eb53-4355-9d37-4c5afe3fdc7b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7da1597e-b31c-48da-92c5-fa7a24166414', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7db72296-a8ab-4f1f-9aec-279e68c1114f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 08:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7dc5ebfb-27a4-4e7b-a6eb-52df4c96965e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-07 18:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7dd09660-4df4-419b-97fe-fd47f7bab9f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7e0c877d-95e5-4a6a-8760-f6f111176b18', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7e4d5637-784b-4881-b8c9-c5357af826db', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:35:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7e79563c-019d-499e-a05e-b283bc6f3748', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7e798f44-1183-4c69-9371-c35857af7677', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-11 22:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7ea6687b-b4fc-41a0-b773-a824809afb92', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '6.10', '0.54', '0.10', '2021-03-06 15:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7eab9eee-59f6-4201-9cce-baf0c1effc53', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7ec93daf-68e2-4feb-8da3-4315cc1b3469', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7ed6b04b-0164-4587-9610-778d78a66a4f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7ef04bdf-98b5-4775-b4ad-1213d03479c3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7ef491a5-15b1-4bb0-aae5-26512d4e607e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-07 08:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7f1565a1-dc33-44a0-aae4-0cf73b67db0a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f1b8273-b123-4d34-af8c-7df618a4d522', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 17:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f468fdb-22cf-46b7-a678-a39840c99d07', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f46af06-d830-4e7b-8840-f5f703a8b9a8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 12:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f4d3787-8c43-43ee-b155-7c259a5e573a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f6a400f-e469-4f54-8379-0b8ca80360a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 18:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7f7368bc-d6ca-442b-9548-2b99beaae7e5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.31', '0.57', '0.10', '2021-03-14 12:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('7f791412-d4bd-418e-9682-55a6b669cc30', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7f8e4250-97ee-425a-838c-d47a6d48c2ad', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7f9aea28-cb60-45d0-8c6b-a15ae5f5bb85', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:02:06', '0');
INSERT INTO `monitor_host_log` VALUES ('7fce02c9-a1bc-4b07-8acf-47523afa7084', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('7fd5d331-01e9-467d-9e8c-0ea7deafb8ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 16:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('7feda27b-3762-4fee-89e2-01d91181a26e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:39:05', '0');
INSERT INTO `monitor_host_log` VALUES ('801f99b8-cf9d-48f9-8f01-2c9e4efc719d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 13:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('804219ec-d46e-4e42-9a78-f6c6eba46a4d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 21:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('80a00c6e-449a-4f6e-a830-e43dadc94527', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('80bbfcaf-b384-4163-bdf3-72f2a69d2718', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('80efc7d7-600b-4a78-bbe4-b2901e968c3e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('80f78e3c-df21-4912-be90-a196f852eeeb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-08 20:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('812c826f-2ac3-4389-bc11-06b1480c1d83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('81308040-e943-4f91-a89a-68d85f85638a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 21:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('813417ad-7163-4be9-9831-bf0f563dfc1f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.12', '0.02', '2021-03-06 15:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8177cad6-7a6c-4445-b1f7-9a9c25de8a96', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 15:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8181850b-6047-46e2-aad9-0c5d18b4a77a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-10 21:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('81933239-c0d8-404c-86ef-6908366d5629', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 18:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('819de4a6-0ee6-46ef-8b35-796f22b7bb9c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:49:06', '0');
INSERT INTO `monitor_host_log` VALUES ('81d6bc46-c0f1-42b0-9530-e0fcb7884eb3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('81eff437-224c-4c63-ad12-c0b1e567540c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 22:43:01', '0');
INSERT INTO `monitor_host_log` VALUES ('821058c4-136c-4fd7-b5b8-2860195c53e4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8265c824-fcce-496f-b7f8-e4f724d7fc77', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('82ce9e87-8c28-4020-8a9b-48ee63eb26f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('82ed06cc-b8b0-4e02-a5ed-a7340fe34ca7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 20:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('82fad61a-08ab-4a50-9faa-f27ec3c935ad', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('831913ce-d3e7-4ebb-b99d-998fb0a2b9cb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 20:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('83c9947b-fb63-4e67-bfd4-c8ba7badde9d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 22:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('83e9b6d3-3409-4a5f-9f42-7a7cb546df30', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('83ee46c2-5c35-4675-8f78-200a3abcd250', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-11 00:51:05', '0');
INSERT INTO `monitor_host_log` VALUES ('83f59c31-53eb-445f-98ea-ba418ed55466', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 09:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('840e9cbe-e400-4d02-b5bd-8eff7c16de45', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8439df6b-9c23-432a-a30d-d6f736df5986', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 21:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('843c9c26-6713-424b-9628-a3f12d7f0b98', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('844014d2-a135-4bbe-ba71-31a335b68122', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('848d33ad-c21f-46a7-b7d0-76dac344be63', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 18:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('84a57f47-6957-4ed4-8bb4-69f1c302f511', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('84aa1aee-e6ec-4bed-9a93-3f2f67dd230f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:44:06', '0');
INSERT INTO `monitor_host_log` VALUES ('84f02451-fddc-4a26-a988-9abaeeb2a5a5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8503fa60-0c73-4fc3-8a59-6625fa1e54de', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('85081e6f-9274-47d2-84e9-85b7a62987c1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8519c328-f0bc-4d86-937d-6b09ba8f5352', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 09:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8531cb78-d13a-4f7e-8010-8ee21b71f802', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.13', '0.02', '2021-03-14 13:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('856e45d5-1fbf-42c5-a882-e4a346ff8319', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('858eb486-bd5f-4d9e-ac4d-48d716a75e4a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:44:06', '0');
INSERT INTO `monitor_host_log` VALUES ('858fcf9b-82b6-4ec7-89b4-30a4a891e9d2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('859f40a8-eb93-4a73-a603-946b450f47b1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.24', '0.53', '0.10', '2021-03-10 20:12:06', '0');
INSERT INTO `monitor_host_log` VALUES ('85db1ef5-eb2f-424d-88b1-533f1ce51cd7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('85e33c81-1436-44bb-8819-151c9759ab0b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('85eb4cb7-a3a5-4960-b67f-10ab2abf8767', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.50', '0.13', '0.02', '2021-03-06 17:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('862f8497-7e05-41b9-aaba-46993d044439', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 09:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8637654e-5cd4-4bc6-a69a-f7e34e38af70', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8641f864-01ae-4e74-876d-61d0976937b4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:45:01', '0');
INSERT INTO `monitor_host_log` VALUES ('864ca863-0311-48db-9eca-bbdcd5147e74', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 14:40:19', '0');
INSERT INTO `monitor_host_log` VALUES ('868cab95-32d4-4502-8ebb-52b9925b295e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('868d5e98-eac8-4ac9-af21-764a391a7680', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 19:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('86974312-2694-4df9-b1ff-4ff698517d17', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 09:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('869a6b82-272d-4e61-8cd3-5bdad9bfc1c9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 10:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('86a2579c-9767-4a80-8571-795234236972', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.11', '0.02', '2021-03-06 16:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('86a403b3-95ca-4c47-9a1f-12a002290f71', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('86c8d636-b1c7-4a5f-9d89-cd1b95cae629', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.53', '0.10', '2021-03-07 08:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('86cd7bcc-6f5d-4840-8417-20b1b21b6695', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:28:06', '0');
INSERT INTO `monitor_host_log` VALUES ('86d7fab3-e7e3-4d59-9e70-ef6af6b3dbd4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:21:01', '0');
INSERT INTO `monitor_host_log` VALUES ('86da3a7b-e164-4ea1-afc3-50c390245a28', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 13:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('86e3a002-30cb-469f-b0c4-dad57b8005c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:29:03', '0');
INSERT INTO `monitor_host_log` VALUES ('86f8d965-ab8c-403c-901b-bd41e0dbc595', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8732b778-6a24-4694-bd50-faa000aaefea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:26:06', '0');
INSERT INTO `monitor_host_log` VALUES ('87589e3e-d7fe-4ef9-bd50-c6da35e29747', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-06 23:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8796e033-1f94-47cc-be7a-4b77ee1a3c0c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-08 00:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('87a1770a-acd1-4342-8527-8cb232e88881', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:51:06', '0');
INSERT INTO `monitor_host_log` VALUES ('87bfaf43-b24e-481b-8c04-9c7ffe0dd41f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('87c9320f-713e-445a-bb47-891c38103233', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 23:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('87cdb615-9059-481c-b8b9-73931200ab19', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:02:05', '0');
INSERT INTO `monitor_host_log` VALUES ('87f0d70c-a817-40fb-bf9b-6d805fb56a24', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('88210c32-fcc7-416f-9cd2-d2e9050b4719', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('888d6b9d-f699-423f-b445-0571faf677d6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 23:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('888fd35b-2aba-49f6-8442-ea597dcb63bf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8897922c-5fa8-4d0f-ad79-2617d90d90cc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('88a7f651-3fad-484e-aad6-5060d0982042', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 22:53:01', '0');
INSERT INTO `monitor_host_log` VALUES ('88db12fb-cb8c-420c-a42d-abeef463d0c2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('88e876d5-24c8-4ce6-8e0f-80c49d55358b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 23:20:01', '0');
INSERT INTO `monitor_host_log` VALUES ('88fcc7e4-e8b5-4647-acc0-bb174f27f254', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('89127cbd-b927-415f-8b9c-70c81bec732e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 18:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('891d04c4-246c-421a-8154-56eaaf338310', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-11 21:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('892d5f13-2f06-42e1-9ec0-308f43615b80', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.76', '0.56', '0.10', '2021-03-14 11:19:03', '0');
INSERT INTO `monitor_host_log` VALUES ('893aea16-f11d-449a-9d9f-344159fc2209', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:33:03', '0');
INSERT INTO `monitor_host_log` VALUES ('893ddbe5-1615-455d-ba26-c726286ed63c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('893f9c65-84fb-42af-919b-8e8d6259377a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('898e4e3f-4d1c-4d6a-add3-4e71b3cae6c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('899a7a4d-4896-4a20-b3bd-a553393e06fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('89ab42c5-5e30-4517-8852-90357c37a1ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8a022894-af15-4ebc-a88b-ff90f2ef8aeb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:50:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8a6f6187-c907-4915-9649-f236c2298b73', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:53:05', '0');
INSERT INTO `monitor_host_log` VALUES ('8a7c6b4e-d4eb-404e-bc60-123f47b9015a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8a80c9a7-5827-4fbd-837b-46e209822d2d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:33:01', '0');
INSERT INTO `monitor_host_log` VALUES ('8a973f52-524f-47b4-829c-a3bc2ce5840a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 14:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8aa9427c-f30a-40e1-be78-323d40090338', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 22:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8ab12682-650c-4ae7-adcb-9e7eeed145ba', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.15', '0.02', '2021-03-08 00:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8ab4079c-c7eb-4d08-aa6a-3de74d6b6f16', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.53', '0.10', '2021-03-07 09:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8ad2118f-f054-4c2c-a81c-9b02eb48abd2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 19:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8af0c8c8-7f81-4d6d-bfee-4787f5e589c3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 21:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8afe825a-5efe-41e2-b71d-91827550b538', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 12:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8b182983-3d1e-4a51-9e73-99d573e18b14', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 22:51:01', '0');
INSERT INTO `monitor_host_log` VALUES ('8b2fc03a-0d2f-4047-b9f5-c303146d24cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8b2fe08d-9c59-49da-991a-5fbf8f6b2337', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8b6c7095-e0f8-4843-a2b8-6a257470c6fc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.12', '0.02', '2021-03-06 15:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8bb27fd8-e8a3-409c-9622-0a729d3155eb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.14', '0.15', '0.02', '2021-03-08 00:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8bf21398-68f9-4447-9366-9bd33fd870dd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 08:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8c0d9a30-a877-49d7-b5ba-3aad7e29e874', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8c1f7081-9aed-4374-b187-68e8bb94b79b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 13:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c203675-7917-4107-9c72-899e1d9918ff', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c2cf18e-84df-4ba9-a3a8-7b64a9336770', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 17:59:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8c2fcfcf-3965-43c8-b2da-ebcd3fdae643', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-10 21:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8c400347-e6da-4880-ba03-1b5dbae4feef', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '23.80', '0.53', '0.10', '2021-03-06 23:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8c449cc3-76bd-4618-b244-5e9fc9d1ebb7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:58:06', '0');
INSERT INTO `monitor_host_log` VALUES ('8c46b90f-b2cf-4900-892f-8712584d13c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 16:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8c47397b-e37e-40d3-9f37-9edb0d3e65c7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 15:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c6bf4f6-0ca4-444c-8d4b-ff01b208fd32', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.90', '0.11', '0.02', '2021-03-06 18:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8c7ddaf1-61f3-4cf2-9205-bc3ac1143ddf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.54', '0.10', '2021-03-07 21:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c8fc820-a015-4c04-b531-bc4b90cdaa68', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c935827-c074-4ba4-8252-f907bf5cabd4', '', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 14:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c93ebaf-5dad-47cc-8676-849a321b8510', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.10', '0.02', '2021-03-11 00:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8c9bc546-d769-446b-a237-80a1dbbd3e12', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8ca1d2e6-efa5-4345-b63c-aadbe6480a83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8ca606a3-2fb0-43f4-b058-7dc5ce3116df', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8ca8930a-5927-43df-8d98-4437b860370b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 16:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8caf22c5-e009-4e52-950d-01e53a422c5f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:06:01', '0');
INSERT INTO `monitor_host_log` VALUES ('8cbc6a1f-5db4-436d-b62f-ac13f26e3913', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8d14ffe1-b22a-48bf-82e6-ef2876bf0469', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8d17b539-37d5-45d3-88f1-254058273faf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.11', '0.02', '2021-03-07 14:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8d1dd43b-7af4-4bda-b9bf-1cebd13ec69f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 19:41:05', '0');
INSERT INTO `monitor_host_log` VALUES ('8d4baafb-ffa3-451f-a199-2278347114b9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8da0ea4b-e028-47aa-ad26-3696d42265ca', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:46:31', '0');
INSERT INTO `monitor_host_log` VALUES ('8dc4046b-860e-4816-be63-aa4979f4a0ab', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-08 00:05:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8dc543bb-495b-4b97-bc57-972b65572e2b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 23:58:01', '0');
INSERT INTO `monitor_host_log` VALUES ('8dd3408b-f926-4adc-98f1-4d2d4a56f93c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 18:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8dd5820a-3a62-4dcf-b32a-58c35ac3d60b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8de4a12e-6ad3-41cf-aed5-9136162e81a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.54', '0.10', '2021-03-07 18:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8e0ae885-bb3c-4b1d-847b-9f0c72ad0828', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.15', '0.02', '2021-03-08 00:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8e1ea3ac-bfaf-4bd1-a9f3-a306daee6fdf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 15:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8e1fe5a0-294f-4c63-98f6-a48d67a70a9f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:54:05', '0');
INSERT INTO `monitor_host_log` VALUES ('8e29296e-cc48-443f-9244-c2864d9d8ed4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8ea76b75-4772-440e-a014-d0ed59b54947', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8ef1651f-0239-4072-91d5-31d5324f3a76', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-14 13:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8efe52d8-7b23-474d-a508-75206b2f00bf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 12:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8f1e97f5-1e4d-4ff9-8cb0-e8e22691de87', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 08:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8f4f097e-9b5a-4b8c-99f7-cc7f932059f5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 17:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8f5f54cd-19e4-4c98-9cba-a3d160a6b321', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:22:01', '0');
INSERT INTO `monitor_host_log` VALUES ('8f63ba05-7f90-4a6d-8414-0969f5791950', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.30', '0.54', '0.10', '2021-03-06 23:36:03', '0');
INSERT INTO `monitor_host_log` VALUES ('8f8955e3-5e5d-4852-87e3-35c85082de34', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8f8bf56e-f3fd-4061-beb6-6914b1b4be5c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8fa6197f-2e75-48ca-8527-841129d93f86', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8fb3018b-77d5-4963-be68-70ba180f44b1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8fb7f754-18ed-4741-961b-42b8d9ad8539', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8fc5c0c4-da31-4236-9f23-6f7ce9c7d5c1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8fc63ba7-d800-4c40-a815-dab36c4d1cb7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 15:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8fd24692-e45d-4681-82a6-6ccd59c0a2f1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8fe9ac54-a40b-434e-a1b5-57302d9a7ad6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('8fef7535-8595-42f9-996b-1776103ffbd6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-11 21:59:05', '0');
INSERT INTO `monitor_host_log` VALUES ('8ff7d2a9-c61a-4792-a2a0-a99c1c0c21ba', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 09:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('8fff0b33-dc71-4a45-9679-d95cf76c8154', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:17:03', '0');
INSERT INTO `monitor_host_log` VALUES ('90001c1a-f81d-42ad-9351-d1d51e4e1fdc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 10:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('90008c8a-9924-4377-a5a7-45115a5002f6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.12', '0.02', '2021-03-06 17:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('900bf7b2-b142-43b9-b36c-60a935c20aff', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('900d05b6-3dd4-4220-b512-a7e37c18b080', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('90103af4-3aa7-4dc5-8402-78b28536e534', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('905d256d-52ed-4d59-9de2-6494c4802a90', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-11 22:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('906c352c-b530-454c-80d0-2dae5f4f5ae1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('906e8411-4be6-4380-92a5-ca16e85a709e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:39:03', '0');
INSERT INTO `monitor_host_log` VALUES ('907646bf-f61f-4520-8c1f-1f2ad98520ef', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('90b18e2a-f83b-413f-aa26-3dac174ce5ed', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:14:03', '0');
INSERT INTO `monitor_host_log` VALUES ('90c86cac-2708-4a0c-8170-10fe38dad24c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('90cc27ba-b9ec-42f6-8c27-bdff8bde50db', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('90cdb65c-cb1d-4ceb-bb55-554b3281f40f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('90fcc113-c96d-47f4-84c6-f0865d9c3733', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9126e2c8-bc7d-4d3d-89a2-d6af29e7b029', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('914704c8-15f6-45fc-a8fc-c64b0f31c367', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('915dbcc5-87ba-4449-b159-1aa945db489b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.13', '0.02', '2021-03-06 17:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('91909782-bdd4-4cfa-aab6-f6e783344b6c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('91982701-a4c1-4223-80df-c11e367685d9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:46:08', '0');
INSERT INTO `monitor_host_log` VALUES ('919ff5e0-0c67-4c84-b3f1-b6f6c3ab9dff', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-07 09:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('91d08bbe-1896-4a21-be8e-829d97fc0492', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('91f8371b-89ca-429b-9143-947f52286502', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:49:06', '0');
INSERT INTO `monitor_host_log` VALUES ('9204f7f0-b0d2-49cc-b344-d967df81bf29', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('92120281-9f5e-49f9-a187-5cdfbe5d6126', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9230d305-fbc3-4c36-81a4-091094093704', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.13', '0.02', '2021-03-06 23:21:01', '0');
INSERT INTO `monitor_host_log` VALUES ('92364480-910d-4586-85ae-77ee860155ce', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 10:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('923935cf-0ca0-4435-a380-d6e655b2c74a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('92666772-8699-4d25-96f3-c76251b69f4c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('92677f3a-bf76-44cb-9dbc-a53b12f06574', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('926dd120-f243-4e9c-90d8-3900142f72ca', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('927f7725-0859-47f2-8ab1-a6604720e8ab', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('92c14d0a-e007-4417-8037-13ebef8baf81', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 23:45:01', '0');
INSERT INTO `monitor_host_log` VALUES ('92e14269-f8d9-4d62-819f-7818a20395f7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-08 00:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('92e7b537-ebc4-43aa-851a-851024b22d8f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('931086f6-317b-4b23-9271-4d9fa1c765cb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:55:06', '0');
INSERT INTO `monitor_host_log` VALUES ('9330d2ce-9d2d-4461-bf1b-78965ffb8134', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('93577be2-76f8-47a6-a6c1-ebe64134346c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.77', '0.56', '0.10', '2021-03-14 11:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9359191a-86e8-43e3-9a69-8678a981f43d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 15:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('93831e71-740c-4750-948c-d0664e7ea604', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 16:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('93957be4-a207-4506-8dba-72963cdb92aa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 14:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('93aaafdf-e9b2-44d6-8a63-3bd7a8a3d90c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 14:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('93cd9088-f746-4b34-801f-7742ead74552', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 21:57:03', '0');
INSERT INTO `monitor_host_log` VALUES ('93d1461a-9854-41f5-9d91-9cd3c9f1c7bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('93dbb59b-8b4c-4770-a472-cc6a4971e7c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 16:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('940983c2-1835-4140-b746-31d3f2ebe230', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 10:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('94126c6f-5f57-4a43-ab16-5e7bb210d190', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 20:34:07', '0');
INSERT INTO `monitor_host_log` VALUES ('9437f28f-e9ba-4806-bd2a-7570dba27158', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:37:06', '0');
INSERT INTO `monitor_host_log` VALUES ('943de617-a762-4e66-a1c3-99831e22160f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('94533bf2-613f-4d02-ad53-c4c6d8f2ac1b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 15:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('94624d26-bc51-4247-a80b-22a3ae95a2af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('94669a07-bec4-4a53-a785-743b1b1ca7d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 23:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('94823cb8-06d4-40e0-b1ac-4e28eae1f585', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-13 22:50:08', '0');
INSERT INTO `monitor_host_log` VALUES ('9484eb5a-e967-4d51-ac2a-a899ca156202', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('94a58748-b52e-4c53-973e-c49b2e0f2fa3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:33:03', '0');
INSERT INTO `monitor_host_log` VALUES ('94eb09f7-29b6-4292-9479-2c72f4b83743', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('94f0c0a5-d43c-4cc5-b54b-586a9e2418ac', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('94f35e2a-69a3-4073-89ac-72a3af8debbf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('94fd86b1-a210-4254-bf27-09d8f969a6e3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 15:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('95801fe1-2e77-4182-975f-679aef2055db', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9595eafb-6ff4-4011-bbce-6028c4aa1dd8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-11 22:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('959b5cfa-5289-4875-8c10-c6925ff310a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 17:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('95dc057b-a4b6-4c55-90b3-562248a2cd38', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('95fb0774-8d0a-43ae-9c70-dceae005f1c9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.16', '0.11', '0.02', '2021-03-11 21:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('95fd18fa-f3a6-497c-bbb3-5882c24dee89', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.25', '0.54', '0.10', '2021-03-07 21:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('96079cae-d955-4dff-9ef5-0dfe221f746d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:53:44', '0');
INSERT INTO `monitor_host_log` VALUES ('96619789-f2fd-4460-b5cf-652ba5d0c44a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 19:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('966bc825-b721-40ac-a703-78839793bae9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.11', '0.02', '2021-03-07 09:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('968a6f0e-6e3f-48a4-942b-ec5440a60346', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('969b4e83-d6b6-4e7a-880d-2cb258c93687', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('96a35d91-26db-48e4-bf15-2732e15b918f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.03', '2021-03-09 23:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('96cff2aa-bde9-47d9-9996-ec7d11a9cd8a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 23:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('96d16fa4-6b05-4e74-95db-16bb05da60b3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 22:36:01', '0');
INSERT INTO `monitor_host_log` VALUES ('96ddf28d-5c38-4068-b582-f2e3a5858f3a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.56', '0.10', '2021-03-14 11:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('96e91c0d-7073-46ae-b8a8-26ae2c53c638', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('97086012-4e93-45d5-bd93-6450be3af93b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:53:05', '0');
INSERT INTO `monitor_host_log` VALUES ('97800c04-680c-4481-97a4-c21820b5cd97', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9786919e-0607-4bfa-bfe1-d9a5daffaf35', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.16', '0.11', '0.02', '2021-03-11 22:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('979d6162-b34f-4bb8-a1e1-181ae77078dd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:18:01', '0');
INSERT INTO `monitor_host_log` VALUES ('97dc2d16-3b8b-465a-82f2-28dc5525fd27', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-07 09:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('97f18c47-7d9e-4d74-a42a-631f2167da5c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('981f45e2-8e85-476d-a513-cb13a2e7198f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.54', '0.10', '2021-03-07 22:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('98288a10-6979-4bc9-8208-1e55d4523074', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:07:03', '0');
INSERT INTO `monitor_host_log` VALUES ('98495063-6d7d-438b-8446-60dcc1afa447', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('98528c77-2e7a-4bd6-8383-5add53962dac', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('98766764-3c68-40ae-90e9-76487a6a70ac', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:34:05', '0');
INSERT INTO `monitor_host_log` VALUES ('98782b1d-21b5-4078-8aa9-c37df537fe22', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 17:49:05', '0');
INSERT INTO `monitor_host_log` VALUES ('987e32bf-31d9-4cd6-9de0-0171df081a29', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 11:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('989a90cf-78d9-4a67-9eac-55058b64b889', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('98a6cf1a-25a4-483e-9479-933c513ba2d7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('98b63421-0e25-4099-bd9b-cbfff5b36655', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 18:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('98c3662f-2955-43a6-9b28-4277cecc895b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:27:07', '0');
INSERT INTO `monitor_host_log` VALUES ('98c93cf7-a076-4862-9910-3dfd984843d0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('98ce76f6-a3f3-4c0d-83e2-31d70d4303f1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('99314c50-b810-47dc-9c5d-87e6839a90f5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('999cc4a8-93f7-43ee-bd6b-6f906c001e66', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 15:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('99c9c459-f8db-48e1-8386-3ddad411a17e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 14:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('99dc350c-555d-4cb6-8e15-f058add19a8c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9a4f4234-1e36-441b-97d1-ae93be20b569', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9a557bfe-4929-4cd7-8fe4-2cd2ec3cc52b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9a618138-fb98-4a3d-8032-4c469e066d37', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9a7beafd-8df4-4150-88d3-735458cceb56', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:55:06', '0');
INSERT INTO `monitor_host_log` VALUES ('9a83c709-5cfc-4d9f-ae55-01f7594a7cbb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 21:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9a88ef58-368d-493b-a8e4-9eef05ce77a0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9a9927bd-b483-4898-ba6f-de72e01c6c8e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 15:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b06f536-b131-4a25-baee-bc7b375a73cb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9b0da0fd-e207-47a6-90c4-e290939eb397', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:14:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9b0f6882-325f-473c-8e34-a4e7ad7e03de', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-09 23:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b2fd594-978b-438a-bcdb-85d5c496efc8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9b3864b4-67eb-44b3-959b-90db317cbf59', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9b4b2774-0652-4062-a5ff-49d880c62a2c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9b5387f7-07f9-4145-98c0-46991ee655a2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b61abca-1668-48a6-9abd-b25aabc180d4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 19:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b64acd0-8720-464f-a291-1fc8b3c3e0e5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b680bf2-77fa-4e8d-855e-5b658f47a55a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9b681d2a-1239-440b-8ef1-11534da7cbed', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.55', '0.10', '2021-03-08 00:00:05', '0');
INSERT INTO `monitor_host_log` VALUES ('9b81e482-2706-4e7b-83b4-f0caa300621c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9ba94bf6-e955-4b6f-b3c9-b1b3cb3b0928', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9bcefb88-ef99-45e6-b8ad-81ec0a1747cc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9c28021f-6515-444c-bf2a-2fa0c52fff60', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 22:25:05', '0');
INSERT INTO `monitor_host_log` VALUES ('9c4039d0-41f1-438f-957b-0a9e6c788c53', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9c97ad23-d72e-46c6-848d-502fcb9b5241', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9c988e75-2eb7-4b70-a55e-5f402829b826', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 18:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9c9ae7d2-5f46-487a-a7ec-83019d4ce001', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9d378ad5-8498-4631-adb6-794c0ceb6a68', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.20', '0.54', '0.10', '2021-03-06 16:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9d4d7d2f-ef0c-410d-8787-7191a3d24916', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9d67b08f-a1ea-4f10-a75b-ca25e0fbe5fc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9d906c32-a72e-4c76-9671-39d20471de1b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 23:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9da26db3-fdb8-4e31-b579-f4cd47ceac61', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.12', '0.02', '2021-03-06 22:52:01', '0');
INSERT INTO `monitor_host_log` VALUES ('9dadec0b-ccad-4d2b-979e-cf603e44d81a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9dc17f39-5aba-4702-ad4a-7d33c4e9c938', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:16:05', '0');
INSERT INTO `monitor_host_log` VALUES ('9dc495a3-b7db-4ccb-8baf-ca62f89550ea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9dcd1713-3a7d-474b-92a9-ae8dd568048d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9dcd983a-255b-4627-b048-cc06c72cd313', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:48:06', '0');
INSERT INTO `monitor_host_log` VALUES ('9dd9c4a2-05d0-405b-88dc-57d4c36942c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9de323cf-36f6-4c19-ab00-c0e56ddaf15e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 17:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9e0b9f53-32b2-4d5c-8d94-853a5d395ccf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 23:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9e207124-9668-4d9b-87f8-d0a3ee8e87bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9e29244e-8223-4809-8a0c-6cb6e9ddbd8a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 14:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9e30a954-353a-4027-9e2a-81137f7fb7f0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.56', '0.10', '2021-03-14 12:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9e3b6c9d-f1d6-416e-906d-88ebe5cf4f1c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:46:05', '0');
INSERT INTO `monitor_host_log` VALUES ('9e49a7db-818d-405a-8c1e-cc2e9d0d2b6f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9e5a938e-9f00-4ffe-9fd5-bab9c2f517a7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.32', '0.54', '0.10', '2021-03-10 21:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9e628c88-3125-4186-b1d2-90e7dd063a43', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9e7789d2-24f1-4afe-8fd9-ce4cfa1c2a18', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-09 22:38:01', '0');
INSERT INTO `monitor_host_log` VALUES ('9eac1352-616d-418b-9bbb-56f43ccbdcb1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 18:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9ecc768c-a416-48f0-9ed7-100d9cec88a3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 19:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9eeee9b3-8506-4fa4-8792-d711c751bdce', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9f327c88-81c0-4f60-9f13-df23cc86cbe0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9f36b0b7-c42e-4c6b-9fec-42efcac08e17', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9f46e219-81bb-4d93-b16a-9f428c51afa3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-11 22:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9f62884e-a78c-4663-b205-690bf754ef14', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.53', '0.10', '2021-03-09 21:48:03', '0');
INSERT INTO `monitor_host_log` VALUES ('9f7cbe02-d7dc-49cc-9c14-cf9caa318734', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9f87f291-5dee-4cc6-b8eb-4afe09d1c12d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.25', '0.56', '0.10', '2021-03-14 11:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9f97910a-99b1-4428-8ae6-d2b39f3fb2ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('9fad153c-17bb-47fd-addc-cbc963c306d1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('9fb887a2-b510-457c-b1fc-5b13021ce075', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 15:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a0019663-6d7e-446f-8b14-aff55e9fb41e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 14:47:31', '0');
INSERT INTO `monitor_host_log` VALUES ('a0171605-3401-45c2-9b10-0fd59cf210b6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a03abbe7-3d38-42dc-b97e-79ec4cd7cc70', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a07cfe7e-8644-45cc-b5a1-359a9ad0b6bb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a0b32d7f-bc08-419b-ac13-e52aad2a8b13', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a0bdbadb-f1d0-4b29-a593-5040f27a967c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a0c6a259-9c7e-441c-a5b1-e1c4a8a2928f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 19:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a0fd9af4-1cb5-4cbf-a7d9-4672d80bda12', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 13:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a1151fcf-79de-4b3e-b1a4-3ba6526626f3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:47:06', '0');
INSERT INTO `monitor_host_log` VALUES ('a19cffd7-c819-49b3-9b05-119908d46b70', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a1a4bb9b-2308-4a8a-b4f5-9c29a33331f4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:04:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a1ca6afc-8318-4959-8fbf-7430d0f7fd54', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a1d6e37f-cfbf-463b-a8c2-c9612bf1d067', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a1d89521-75f5-4729-8a1f-cc6a53700c83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a1ea0bf2-4d0f-4865-80a9-e03102b2dd5b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.56', '0.10', '2021-03-13 22:43:08', '0');
INSERT INTO `monitor_host_log` VALUES ('a1ec2c39-467b-44ec-93ff-ae3050c6608d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a1f7c661-029e-4788-8dd3-b44184f40bfc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-09 22:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a20773a4-ff36-4182-b3c1-c7799e2f8b8e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a211b249-473f-473c-a4b9-399190e520eb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a262e863-54ad-4a61-acec-a5bba2ee7bca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 22:34:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a2657580-2a9f-48db-9fd1-9fb5de72c1d5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a2831e27-4df1-486a-9fe0-07a5a0b915e3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a2b7d9f5-c8b1-4967-8101-b43473ed0efb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a2bdb2ef-558c-43ae-a476-8fa7de7cd840', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 14:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a2f70069-d7da-401e-b223-781a61ba4cb5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.57', '0.10', '2021-03-14 13:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a2f7fe5f-5be6-4c7b-b59a-35175275e163', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 13:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a2f83d7c-15c5-4f60-af34-f5fb800e08a9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a3111ce1-1719-4cdc-ae33-6243fee1f205', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 12:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a3133560-2c11-4db3-a453-a2883792c3b2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a313f951-3d89-4005-9be4-9cd58586cb1a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 15:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a322e12c-a0c3-41ca-961e-824541dce720', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 13:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a340ae21-8cf9-4399-ab02-6b2ee9b7b58b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a357d49e-8f19-4c94-97f7-17f831f496fc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a39eeaa4-c134-41a5-9135-bd996fc63704', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:36:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a3b4657a-bb31-4ae6-ba3d-4cb312e93666', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a3bee009-62b9-48f9-90db-2d8718ea028c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 15:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a3c24e1d-6adb-458b-b470-801dccfc9f36', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:32:06', '0');
INSERT INTO `monitor_host_log` VALUES ('a3c64e3f-d92e-4f76-9203-0f658d0d835f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a3fb88e1-44cf-4043-81c4-a858c8bbfaa1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a4008117-5eb6-4858-986b-bc32c49ded65', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 15:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a4414709-c768-470b-95d6-8d980cf41060', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a465d9c3-4279-438d-adce-51d9541d0f6a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 19:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a4703e73-cbc3-4ada-bb98-6750b95630c2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-14 11:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a4872231-9425-4271-8373-6c4da67c45d1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a4abd8a5-87cc-4eec-b52f-67ac17e57d86', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '13.60', '0.12', '0.02', '2021-03-06 23:02:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a53b5c74-9e5c-424a-9d45-b4fcf1481cc8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.54', '0.10', '2021-03-06 14:38:11', '0');
INSERT INTO `monitor_host_log` VALUES ('a540fa50-240e-488a-a26e-44a30faf2386', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a5481b2d-a247-4261-8d2a-9a81166d84a8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 14:49:33', '0');
INSERT INTO `monitor_host_log` VALUES ('a57d76e2-674e-4ccc-9989-f41990e45685', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.00', '0.54', '0.10', '2021-03-06 15:50:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a58461e3-1e8b-4e31-8b1d-33b434896259', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a58a83af-2075-4ea6-8407-1aafc643cf7b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 22:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a5c41759-11df-4e22-be0a-b80f98ad89fa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:29:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a5c4ebbb-b5d4-4eca-9e5e-666d9c2901b7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.53', '0.10', '2021-03-09 22:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a5c878e4-48d8-4f62-85af-114a8006412e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a5cbeecf-3154-4aa9-a97e-f8cb59a82a34', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.12', '0.02', '2021-03-06 15:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a5ed266b-c220-4bea-b1d1-5d6d1e374328', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 15:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a5f5a7a2-6d02-41ff-8bcb-6b0fa5ebfd1e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a6023244-1a4d-4497-ad65-bb0ed0a6040e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a60960ba-8008-484e-90ac-7ad633789138', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 12:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a61cbf3d-ed33-45ba-9035-8a956ef66978', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 17:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a63c5088-f210-4700-9065-9ddc492cd580', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a64cf52e-93f1-4466-af7b-cb8d39001b67', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 22:15:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a66d7a43-c8f8-417a-96ef-6ded07eef67d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a67e2e3e-cb35-4c44-b6a0-e812c8a462ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a68f8e83-3c42-4e0c-81bc-39c9f0e17096', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a691df38-b617-4651-a3fe-9c90b089a1b3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a6a0450e-87b3-4e09-9985-892927d9cd6a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6bbecbb-ff0d-4780-a382-18b2663b1111', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6c87af7-e29f-41dd-96e3-a5e326b60e4c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6cb5dc3-1849-4723-b218-ba4f8312fa08', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6d230fc-77b0-4d04-a454-a0d90a2b4976', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6d7b633-cc9f-4ec9-9258-6a0c8fcf8f48', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a6e6f2e5-401d-41ef-8434-63e2f024a139', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a6eb63c2-d9ee-44ae-8815-cf8f86cf0a76', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a6f6b590-9747-4f60-8ed4-470abda93387', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 18:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a6f750d3-2fea-42ca-9256-3d141e28f90a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a7014f92-81bc-47da-a946-d873baf81328', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a7051fd1-ebcc-4c35-aac2-afc44f3cd550', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a75b7d9c-83f4-4561-8ef6-00bdd2285bc6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.53', '0.56', '0.10', '2021-03-14 11:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a762fdc8-b62d-44f7-85b2-048a73bbf742', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.53', '0.10', '2021-03-07 05:36:39', '0');
INSERT INTO `monitor_host_log` VALUES ('a7681bad-2c8c-441c-b72b-756bf609e973', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a7909d15-abe1-402f-9de9-5f3afe96a11c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a7bd31c5-ee33-4461-91df-113c36c8d72a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a7d01c3a-7bb0-4339-9502-703f725055b8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 23:04:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a7d6489c-c402-4d5d-9863-442eca31bde3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a8343b1a-22e8-469f-8e09-e9998d795612', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 22:57:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a839b89c-4dc0-43fa-82c1-a2cfdfe98188', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a860bb59-6e97-4e9b-b95f-1bbdb99ebcc0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '12.80', '0.54', '0.10', '2021-03-06 14:44:15', '0');
INSERT INTO `monitor_host_log` VALUES ('a86f0812-ee08-472e-9289-d18a78335ce9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 23:11:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a872feeb-9e37-4c04-8fa7-5dc2dab965bd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a8c33e94-ecbe-4557-8e8f-1dae9cfcc2b7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 12:01:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a8e54ec6-d27b-48f4-aa9b-b470d9e7bab6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a90a866d-bccf-4e5e-8767-269ab11ed254', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a916dd4c-37ac-4c2e-88ef-911c5d5ba6a4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a92b0fa1-aac3-4142-ac6c-f6cad57a021e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 22:37:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a93ae2a2-5109-45c5-a42f-427ac287a3cb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a943c975-37ad-4b89-b0a3-d28a25f8d508', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a944746d-67b1-45b9-a390-30aa6625cea3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 16:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('a9464753-7116-467f-bff8-a5df524362d6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 15:17:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a95a063e-df54-4408-9732-5f0ad7ec54c3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 22:58:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a97855dc-6ed3-4146-b6e0-bdbf9f788ee4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 23:32:01', '0');
INSERT INTO `monitor_host_log` VALUES ('a97af9ac-a54d-41dc-81b8-068cacddf1ca', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:19:03', '0');
INSERT INTO `monitor_host_log` VALUES ('a989a971-1405-4708-a149-e2d44589056b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 22:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a99f4536-808b-4ae1-bbd6-a4adece76645', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('a9c846f8-4e5c-454d-8aa2-b58a8482a2cb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 15:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('aa0802b4-7be0-46b1-a234-e34106327d09', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aa293b36-3db5-412c-ac79-7e23b748902f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 08:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('aa3bbc13-a36a-4fb2-849b-c72ab4659599', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 22:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('aa4212fc-2040-4c62-80fa-3495c0bd6769', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aa47f445-79d6-44f1-a44f-ccb220c7d6be', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.13', '0.02', '2021-03-06 17:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('aa539e7a-623f-427f-8b71-87116072b467', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aa57d563-59e6-4568-a978-7d005de43410', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '11.50', '0.54', '0.10', '2021-03-06 15:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aaa92bc8-91c8-4b32-ad4c-649070d36f6f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aacb6d2a-f976-498a-b0e7-c684b9cb6cfe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 17:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aae5c60b-a740-42d3-9c1f-f40229f17780', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ab0ca5b2-936a-4653-b1a8-47bdb16fb902', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.54', '0.10', '2021-03-06 16:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ab2d58e9-f699-4bcb-94af-3fe1c0dd4f88', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:53:06', '0');
INSERT INTO `monitor_host_log` VALUES ('ab337dd5-4a87-49aa-86e1-0e2b748bd89a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:15:06', '0');
INSERT INTO `monitor_host_log` VALUES ('ab4992e7-eb69-43d9-be7f-88aff9b012ef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ab63e6de-458e-4037-8cce-3d228f30c726', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ab717735-cbe2-4658-831e-4bb37e1d716f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 22:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ab97ca23-525f-44f3-8718-ac15a2afd1b9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.03', '2021-03-09 23:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('abc32694-f79d-4d0e-9bf3-46338e667a91', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('abd19d01-7d67-48ff-a860-cc71133f6a64', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('abd61a2c-f753-4c9b-b606-7bdcde870394', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('abd8adc7-e40b-42f4-8a4f-96da23ebad94', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 22:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('abf8395c-e44c-423d-a765-d77c979f458c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.03', '2021-03-10 20:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ac20e32e-7f59-43a9-8b8c-62b3240be2db', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:03:06', '0');
INSERT INTO `monitor_host_log` VALUES ('ac51030a-06b9-4f65-a76f-28afc67bb215', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 14:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ace32a28-0d67-4b17-8a3f-c6c7328d4136', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ad024ea0-4e80-4e84-a47a-9f4d61ff2148', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ad13ea90-7503-4553-b455-85f7167d0425', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 23:22:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ad5b64bf-585b-40ef-8487-311e0b201212', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ad7dfe9f-980f-4599-b206-04ff5ff8eb08', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-07 15:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ad84b918-f913-4d06-a3a1-fcbaed08312f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ad8dd03a-d703-458c-90fa-3802dca67b9b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ad9e3484-b53f-4193-89ec-03c5a2165ccd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('addd68be-4730-4730-97c4-ba14592e6cd2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 19:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('addef5df-b765-403d-a447-741cbfb6249d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:53:03', '0');
INSERT INTO `monitor_host_log` VALUES ('adfb5b87-9dc8-45cd-a9ea-eea6d4c2455a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ae02cdd2-9919-458b-aae0-6426d5df9470', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 23:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ae0e63d9-de58-46db-8ab4-66a4cdb67898', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ae19c92f-08c4-42bd-ae60-237f19eac1af', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ae296b10-e804-42c6-bfdf-17c3c6af258c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ae520461-e545-46a5-b309-857100aeb881', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ae5b4b10-dbab-483e-869d-b010f45fc21a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ae7195a0-f6b4-45e6-be0f-9f7ecad84c47', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 00:52:06', '0');
INSERT INTO `monitor_host_log` VALUES ('aea9dae0-4df3-44d4-8b66-ea317a87a16e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '17.80', '0.54', '0.10', '2021-03-06 15:12:05', '0');
INSERT INTO `monitor_host_log` VALUES ('aebb22d9-a852-4477-a45e-bab75c6daa3e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('aed0f7b7-82f8-45c6-be10-003678d9f340', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('aed133e4-cc3f-48d4-8385-ec639c66a65f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.13', '0.03', '2021-03-09 22:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('aee40853-09a8-4aa6-84cf-b8d1dc9007cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('af2197b0-debc-4a71-afc9-228436162937', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('af2a6661-616f-4c78-ad74-acad7174ec60', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 20:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('af503bc9-7b72-44a7-8930-06b5ed609704', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 23:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('af7f1d99-c6f2-4d3d-af7b-8c67cbd3d9ae', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('af8a386c-0a02-4e9f-bdfc-62b275b7451d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('af8b9f93-7e93-4cf6-b298-c61679897c80', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:44:08', '0');
INSERT INTO `monitor_host_log` VALUES ('afbe6b19-43ea-4bb9-a6b7-5f67c741995d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 22:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('afeff30f-d73a-4806-a98b-a097512a23d7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 14:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b039c9d5-3fab-4bc5-9831-5b8a74accf91', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b04d8fc2-d3fa-41e0-977c-105b5729a063', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 13:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b07ab1f4-e7b7-4f3c-bb2e-8c934b4e4169', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.54', '0.10', '2021-03-11 22:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b08d57c2-331e-4639-a9ab-042f0a2dd163', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 14:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b0ca10eb-4cf3-4fc8-a78a-aed77ed38122', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 23:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b0f20017-6446-48cb-84eb-c7f7a7ac492a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 22:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b0f91cdf-8604-4585-ad6b-2d311596b2ee', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 22:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b1209dbb-b8cc-46cc-ba57-30f6bc0f9329', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 21:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b1284e4b-bbd5-4ada-aba9-77ae63edc377', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b164577b-1958-4669-8c05-6d043502155a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '14.00', '0.54', '0.10', '2021-03-06 16:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b172891d-42a3-4ecf-8482-481c06464570', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 11:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b175e5f4-e9d6-4dd2-8792-0e76d42c9216', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 11:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b1818b29-cfb3-4dd3-b950-b02b7cef92de', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b195b95f-3c70-4c76-aa6e-b19309e3656c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 22:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b199877c-dc95-46ef-9940-af6e8ac34852', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b1a8da50-9e66-4817-854f-868539d04afb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:24:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b1ab862f-e76a-4b70-8cb7-61a8e15965ad', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 14:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b1d0cc71-0cfa-4bdf-a256-cebcca6fdddf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b1df54da-7118-4319-8e4b-929723e08f09', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b1ea8847-c914-422b-8965-9d8364ff3711', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b21069fd-f234-4653-8e55-015a9db5e3e2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 14:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b2147fc9-79ce-4f23-b224-fcae0f2c5fea', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 15:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b225060e-1d68-46ba-b798-cd180fa991cb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b22b4ec6-c9a5-4e08-9029-2890f391dd3b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b23869b0-d3ce-44d4-a0ed-d4db0fea4e4e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b260da62-ece7-4654-a9b7-131abc65141a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b26f7809-a73f-4711-934a-82eb1c83e35d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b280f5a6-12dd-4785-9fd1-de4d118d5828', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 15:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b284b391-b854-413f-a6cc-2e2975a5b140', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-14 13:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b2b1dcb3-b127-4f9f-9378-2253294b753d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b2dd1662-9eb6-4873-9fd5-b06e939f84e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b2e1e574-0bfa-4c5d-b1c4-1a50806f6157', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b2ebc5f3-34da-4bed-8981-93e1cb937543', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b315919a-650a-4c53-b5e6-28b29b026ab9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 16:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b31cf7ea-d90f-4174-ab24-af93ae3cc78c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b34a9b0e-f549-4eea-82ba-d590b16281ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.02', '2021-03-07 23:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b34eff06-1505-4e1a-afed-7fd8e68652e0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 14:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b355f127-e643-4ebc-a284-53c5b2d4ba68', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:17:01', '0');
INSERT INTO `monitor_host_log` VALUES ('b36bd503-8d66-4263-8d5f-635dad96d3b0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 15:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b371e723-7dfb-4d0d-8442-6e234dce581d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 15:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b37ec961-b054-479b-893d-55a288412ebf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b3ac9910-540d-420e-aebc-5290e0ee7ae5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.54', '0.10', '2021-03-07 22:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b3c748f2-abb5-457f-afa9-c5c5c5c94edb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 23:33:01', '0');
INSERT INTO `monitor_host_log` VALUES ('b44ab49e-695d-483b-8644-6fc1cf2a20dd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:27:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b454848e-fabc-4c21-a311-fe9ed2bff502', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b4586816-fef1-4813-9643-6dfb36001782', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:54:01', '0');
INSERT INTO `monitor_host_log` VALUES ('b473b0f2-fa85-4b03-b56a-33455a1c7125', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b4a525d4-2026-4c6e-9706-f2312eda7aa1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 22:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b4ebbc40-c3a2-4661-b838-d6d94b3e9573', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 23:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b511c7ef-e524-449e-b77d-d1cbcb96402c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:45:06', '0');
INSERT INTO `monitor_host_log` VALUES ('b548a976-0efc-4910-95c9-0c3e3a042aa1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b5732990-12cf-40ca-845d-d3089cd1bc67', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b581f84b-e872-4898-b8ac-d3e708aa48ec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:43:16', '0');
INSERT INTO `monitor_host_log` VALUES ('b5906654-3895-4b3e-9ef6-08e77808712d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.57', '0.10', '2021-03-14 11:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b5d28bbe-92bb-4b4c-b93c-1f1808241e1d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 11:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b5e4dfc1-cca5-44b3-b275-a542d68b2a1c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 21:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b5edeadc-9137-4eef-9fbb-2ebd2c87e756', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b67fcedc-e1d8-46ba-a734-9c6974314b81', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b68c5826-1f98-46d7-9679-2a5e6d638a58', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b697eeca-d0fa-4b5c-8f96-9b0ec82fec3a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.50', '0.54', '0.10', '2021-03-06 17:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b6aafdc4-2b08-48ef-81bf-4d3d54f37858', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b6ca2f19-7cfe-4bd5-967a-dba8c2be5ebc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b6ff8af9-7644-4b3b-9a9e-0ad85cf7978d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b700577b-fbdc-4d8e-95d1-82a9438a1be5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.03', '2021-03-10 20:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b764002c-2753-4766-85f5-430d640350a9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-08 20:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b794ba6c-7fdc-4731-a9e7-7bc1b6a7b8cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b7b2483f-d501-49fe-a66a-3a49e20bc03f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 11:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b7ba163e-d905-4438-82e4-2a1769029511', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-11 21:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b7cd8aff-f369-4504-be8c-078a9217cd9c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b7d4a623-6e9b-42fe-b426-fcd408a8d4b4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b7d6d078-99ac-4a08-ac0b-b8af80456668', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 23:00:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b8012731-ac65-48c8-979f-fadf6772c69c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 11:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b8296d67-99df-4287-b665-6eb51cc88eb2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 19:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b8319d32-1389-4961-92cd-ce7a4b384ab6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:16:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b8725669-0c97-4a9a-bfbc-5edea37538b6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b88b4c3a-0572-4b71-b5d4-9db10ba266c3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b8984b6e-8236-4696-916d-ed444fa61847', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b899c106-642c-43e0-9277-fc9cd5882684', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.11', '0.02', '2021-03-11 22:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b9310388-e8d1-4414-9f6d-9e0d3ed09f4d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.78', '0.56', '0.10', '2021-03-14 11:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('b97794c8-6c78-4ca7-a661-0e4c231f59c8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 17:16:08', '0');
INSERT INTO `monitor_host_log` VALUES ('b98275b9-45fd-4f36-8dc5-1e22a8dee1e2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.03', '2021-03-09 22:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b9921f2e-b231-48d0-aef7-19e12b34e278', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b99a4f02-e037-43aa-9aa4-72d3cb2a0e13', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b99fcf55-5c8e-47a7-8356-a2929237e2d5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 16:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b9ab5da1-964d-43e4-8a90-4cc5617bcca4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 10:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b9b887fe-87bc-49bd-964d-a6da3e0d3f10', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b9c7d339-6eed-4c08-80c9-0cea1e9cd05b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 13:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b9cd1d08-5862-4563-bb73-c65af2c08af9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 18:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('b9d8822f-0cab-4022-80a1-3188950899b6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-11 22:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('b9ea3e1b-518f-43e7-81f2-218407e45b2e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ba030f1b-6b0a-41bc-bc47-81bb02cf1d78', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ba154f49-11de-4fba-8ff2-8b9f2584ca49', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 15:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ba1a4bcf-e81e-46f4-8837-dc7a2115d9af', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 14:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ba41ec88-72a8-4af5-b675-9358ab6cdab6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:27:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ba84d0cb-5d77-458d-b2b7-3234af60c342', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bad007ac-31a1-48f0-b600-88655e182834', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 13:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bae87a77-e4c0-44c5-b71c-e3698d2a468e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bae9f10e-53bd-4fe9-9140-93fc1f9db573', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 16:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bb34e284-c552-4c38-bf62-7ed69b2f638b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:05:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bb4d48ac-bf99-4d73-a78f-58379ebcebad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 00:49:06', '0');
INSERT INTO `monitor_host_log` VALUES ('bb5110d1-1f9f-4ac6-921f-03da5f2c8fc2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 18:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bb6537a1-f2e0-43fe-a83a-f181913fb8c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bb676837-8ae3-4488-bc17-967334510a42', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 19:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bb6cbd9f-d888-4b5e-a353-7e03c8767e40', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bb7baa20-a281-43c0-8909-2b53037996d3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-09 21:40:01', '0');
INSERT INTO `monitor_host_log` VALUES ('bb7c7103-3f27-443f-9700-d3f9c9b3c0e7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bbaee438-a8cf-43f4-93b3-f4fbd26e1ac3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bbb30567-fdc0-4970-a125-1ba5df8f1d40', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:31:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bbcd67d8-a866-4409-8db8-ce282e5390e5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.14', '0.02', '2021-03-07 23:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bc0b9fe2-ee50-4d04-99f1-b6a53e905d98', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bc2378bb-43d3-417c-a0fa-0b6ff597f30e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bc436dc1-b503-4f93-9393-74d09721583a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 23:42:01', '0');
INSERT INTO `monitor_host_log` VALUES ('bc4c336f-12ad-404b-a7b7-99d0aaa1a5ac', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 22:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bc593639-791a-4968-b1b8-52e4e4b95a54', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:04:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bc6f9a07-a0fb-41ee-8567-a79c985b0938', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bc8205b2-166e-4dcc-9b50-509092910350', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 17:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bc868518-e73c-486d-b8d9-c8e45a1090fe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bc94a269-49bd-4753-bc0b-992c2cc51926', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bcb2580a-fd7f-405b-9cf3-807c8c5dbbeb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bcdbf4ce-c71b-4ffd-a18e-dac31d134368', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-13 22:40:06', '0');
INSERT INTO `monitor_host_log` VALUES ('bce64465-0784-48b2-aeee-0fed2db7f3c4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 15:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bce7d859-b3e0-4c5a-bb86-1c60c9ca2662', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bd15f8d4-bdb2-4087-8bf7-c52a743d861b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 23:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bd2d7971-38fd-465c-9213-d5643e6cf6bf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-08 20:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bd3f29da-3020-4349-9406-54fa25b655ed', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '1.20', '0.54', '0.10', '2021-03-06 16:50:05', '0');
INSERT INTO `monitor_host_log` VALUES ('bd5d290a-e467-4898-9dcc-9e9e9039be0c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 19:39:06', '0');
INSERT INTO `monitor_host_log` VALUES ('bd6cdf5b-3a36-4c85-9e53-fb6e2d0131a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bdb208fc-1a28-4c6c-9db0-b07d22c9b1e7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:30:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bdd30af1-e332-4233-831d-2d74f67d2ebe', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:13:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bddd23c9-1743-4b64-9838-88fb0ae0ff54', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:22:01', '0');
INSERT INTO `monitor_host_log` VALUES ('bde3a1f3-4a51-4b3a-b1b9-f7a27a666cd6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.00', '0.13', '0.02', '2021-03-06 23:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bdedd73e-90bd-41c8-8397-ea49921ca94e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 14:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('be07e3cb-af96-4b25-8efb-945083b25a1a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '10.00', '0.11', '0.02', '2021-03-06 17:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('be14c79b-17db-45a9-96d3-c60b1b338e20', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('be1bcd6a-78fa-46b8-b1f9-fa45edc9735e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:47:05', '0');
INSERT INTO `monitor_host_log` VALUES ('be77a30d-2c52-4ddb-81c4-bda82500eb04', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('be899a8b-ebb4-43e7-b0df-7c9ce9a9e4c1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:35:01', '0');
INSERT INTO `monitor_host_log` VALUES ('bf3fa17b-0ee4-468e-a541-b800cb7dc691', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bf5980ec-3a32-4c7d-b79f-ea9b5f3aaf75', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('bf8b4554-12f9-4bf6-b2d3-97e4b19a21d6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bf99bf0b-75a3-4af6-9cc2-59e981e1c947', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.14', '0.03', '2021-03-09 23:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('bf9ea6ee-df2d-41e5-848e-f3915a23eeea', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:54:06', '0');
INSERT INTO `monitor_host_log` VALUES ('bfa5210e-1bd4-48d6-b274-cb65be69256a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:36:01', '0');
INSERT INTO `monitor_host_log` VALUES ('bfa6cc32-dfe4-4cbb-83f0-963ecbde7c5e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('bfb6bf3d-069b-4d55-a01f-803696b5c736', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.13', '0.02', '2021-03-06 17:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c0263df8-94be-49ff-8c56-8170efd3f4be', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:02:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c083bc8a-805f-48c6-beaa-5003ad163d99', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:49:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c094318b-c2b0-490b-9262-bc38ac88fc59', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c0e79089-279a-459d-933e-1d21da884f5e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c0f43591-a5c2-4472-a6a0-4d74ce3bf62a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c11370a4-e91f-4b2b-8038-bc7c766b76be', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c16fcce3-0c93-4a17-819a-662c2af1f8da', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 23:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c1c791fa-4c34-4528-9518-13b42a16da05', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 16:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c1d00950-2ca9-463b-ab51-49f437051c3b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c1f7820b-cd93-4bf5-9080-92628007e692', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 22:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c2093c82-d850-461c-bd36-f64af2e759d9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:41:06', '0');
INSERT INTO `monitor_host_log` VALUES ('c229d6ab-08b1-4f64-8f80-6b846e6fb821', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 13:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c2340465-d20d-4c62-86d3-b1a401f6d094', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c24e9888-efaa-4039-8681-f30bd801c97e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c270b601-dc31-4494-9431-53e7d1ee591c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c28b2ff7-f5e3-44fa-b73b-85ee3cb52e1e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 14:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c28c37b7-695c-4d69-bb09-da1dae3a7003', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 23:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c2b51f07-9b05-495a-a910-634f0d4fcec5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:59:05', '0');
INSERT INTO `monitor_host_log` VALUES ('c2c1edcb-5bf4-41c6-a15c-65f295e39bec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c2cd002d-d6e1-45f7-9757-9f707855a908', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c2cd34a1-748a-4ea6-85c1-47fee8ceef93', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:57:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c2d4e4ea-ea66-4d42-b668-12f4fdbdc3cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-06 23:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c3069d39-6e91-4732-b8f4-7589d19a7714', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-08 20:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c3163b59-fafc-4581-b210-750700428733', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:08:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c31fa99f-f5a4-4aae-90d0-5ca641242122', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-11 22:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c3302aff-ebaf-4b51-96b7-6c37bd36c34e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c36393bd-01d9-4165-a948-c563b69dd9a3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '20.20', '0.54', '0.10', '2021-03-06 16:36:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c36da08a-77b7-4d56-a854-fd53861cc85d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c3b705fc-38ac-4476-a34e-32febe8269bb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c3ba4b73-bb6f-43cd-a3b5-f9eb11d2750b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c3c01a32-7edf-431f-873e-6f7980e47915', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.57', '0.10', '2021-03-14 13:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c3c75eb7-8162-4d46-879f-4d6a6a4c4413', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 22:56:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c3f9c3c4-cddc-4831-8d41-8b159b96341c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c3fd9670-bea2-4da8-aa48-32d503b89c2d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c4011cbf-8e38-456c-a149-37fcba59be3a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:57:01', '0');
INSERT INTO `monitor_host_log` VALUES ('c4031d06-b93d-4c08-aae8-c9d0fb954b4d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c4166f25-33b5-4184-8cf5-c4b6063024d2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.16', '0.54', '0.10', '2021-03-11 22:19:06', '0');
INSERT INTO `monitor_host_log` VALUES ('c4551d43-b900-412a-b276-67843d7b6496', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c464a45e-948f-4100-b68b-5d597cd0b203', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.90', '0.13', '0.02', '2021-03-06 17:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c47a4bb1-cb53-4566-9a85-35656a10ecf1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 22:38:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c48c35d1-3513-44e1-8235-94cc22d61548', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 12:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c49c818e-2837-42eb-adec-828c2b36acd7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 22:20:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c4ccface-beff-4dc9-8709-9339d4c8ef0f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c511ae8f-750e-484d-b46a-6fecd089530f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c517bcb4-4561-42b3-976f-d54b691bca28', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-09 23:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c5557733-5ef5-46ff-b488-7ba400087dce', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-07 09:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c56c8d16-69c7-49fd-ad0b-792702874569', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c57a038a-8f66-4328-9484-5e70d1fe6ffa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 18:00:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c58dfcb9-a7c4-4bda-ba3c-99562ea5e0dc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c5a6a423-94e3-4ed2-b81b-21c4f96e325d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 19:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c5d80486-35bf-439e-9181-4fb3f1d682e8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '26.20', '0.54', '0.10', '2021-03-06 22:51:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c5f45b98-049c-4233-a4fa-9eaa5ce6faf2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c6459cee-f135-4598-9f44-28d42ab822f0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:48:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c6476144-18a9-44ae-ad23-3e0961af86c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.13', '0.02', '2021-03-06 17:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c64772d9-22dd-4d77-8a01-917a798623be', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c662dfb6-4a03-4fc1-a8ad-286778937152', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c683bdfa-754d-429e-a757-0508a44b2584', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c6b2b0f9-2baa-452e-8a67-8d34371dce71', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c6cabea0-37d1-4a63-8d6e-bc675125f553', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c6cff721-df21-479b-b2a6-f1c8d6192dcd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:45:06', '0');
INSERT INTO `monitor_host_log` VALUES ('c6d0c94a-56a1-4d44-a4f3-3cde98500b59', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-08 20:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c6e4dc45-be21-4632-a858-2b659e9f8b10', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 23:30:01', '0');
INSERT INTO `monitor_host_log` VALUES ('c6eec08d-c435-496f-b41d-903b91987d11', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c6f16193-211b-4555-81b3-842f82c13a6f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c6fcb63f-a876-4fb2-b2d3-7888f09c68bf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 14:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c7057648-72aa-4c2f-92b3-2936c88dabae', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 18:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c7090315-990e-4313-9d28-3da7c5651288', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.13', '0.02', '2021-03-06 18:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c70c33e5-83af-4ba5-b060-e039e7119350', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 19:39:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c70ec205-5d42-48ab-a665-42301c7e775a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-09 22:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c734b515-f2b7-4098-8a28-8d6da99dce0d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c73f1892-4d26-46b4-97e2-f3812a1b71a4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c76dc6a0-eb37-40ff-b5c9-d5f6fed07425', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c78db1cd-37f6-4ac4-aca6-6368f423ce28', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 19:59:06', '0');
INSERT INTO `monitor_host_log` VALUES ('c793b4f1-a95c-4b74-8782-5171d5118444', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 21:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c7b7c475-9153-417c-b23d-041f801f573f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.55', '0.10', '2021-03-08 00:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c7bb6b77-8540-4bef-bb02-2c370b8be5a6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c7e68f0e-d027-4a99-8760-4beab77715e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c7f4dc27-a408-4769-8c42-bd4f179cd43d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '12.90', '0.54', '0.10', '2021-03-06 16:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c809e7af-5049-4b92-b870-c285ba78ee04', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c812c461-a180-4dc6-a369-a74387191d68', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c83874eb-0fdb-45a0-9d88-694f65db96f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 21:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c86841de-e3e1-4aa2-b5fd-391acf07402c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c8ba910d-8374-41d3-a4ca-c80eee51eb8a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 18:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c8e53131-f67b-416a-88fb-d97787289d49', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c8e7aac7-97a5-4e7b-9962-213e757b37b7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c92db4dc-097e-485b-b117-40aba1819595', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 11:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c949f203-1fbd-4a9c-a765-6f6b573ea121', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c96496b7-255c-4aab-ac26-898e720f2f82', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 16:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c96bfa32-0864-4317-90a8-f36a9b572a02', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('c994d827-5c0a-426f-9068-94933fb2dbef', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.50', '0.54', '0.10', '2021-03-06 14:41:21', '0');
INSERT INTO `monitor_host_log` VALUES ('c9a2289c-52a1-47b2-b8a9-8d7ad2f25c06', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('c9c6455c-af1d-45a8-ab42-b75df7b2351c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('c9d57e17-369c-409f-82f7-06fd3ff70bc3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ca0d9643-df29-4e05-b6ae-424cca89f1ab', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ca1b43cd-f530-4862-8203-24527292aa18', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ca1b62e2-cc9a-4bb7-b6c3-31ce202812ab', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.07', '0.54', '0.10', '2021-03-07 18:45:05', '0');
INSERT INTO `monitor_host_log` VALUES ('ca24fc9b-8b96-461c-9852-cfd5f4b44502', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.03', '2021-03-08 20:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ca33db94-3dbb-42ef-89f4-967f73111e17', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ca4d4f45-aec3-4c77-85ed-f21055b2ab55', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ca539b8f-5c4a-4c03-b0cc-d788782473b0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 15:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ca8d988e-2ebe-4a2e-8818-9732820540da', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ca92cdc6-05f7-414c-8563-71489d60b70d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 15:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('caef12db-c9fc-409f-b5ff-c07af5f2d365', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cafc4b55-1a94-4a30-8348-b06218a4c287', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:18:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cb383e47-8d3a-4640-825d-d3bf31807478', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 18:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cb3a5759-51c0-457a-bc9b-e954e1cbef57', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.12', '0.12', '0.02', '2021-03-11 22:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cb796b35-b1ea-4941-b126-a40364bbfe64', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 21:11:01', '0');
INSERT INTO `monitor_host_log` VALUES ('cb7d9f78-2d18-45bf-b963-640d780deb91', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cb7eb437-48b9-417b-b8d0-ad3dd4183a9f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:06:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cb915bbe-c88b-4ba2-ad0e-993a8bd5cfe5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cbc812a0-79a5-4064-a27a-0d12137b67f4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 08:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cbe094da-366a-47e7-9541-6811b4af86be', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cbec3894-ecf0-4290-ba02-941e9db56aef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cc05dbac-7299-47d1-8546-349d99cadd98', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:50:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cc137001-77af-46e6-aeb3-6b3582fdbac0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 10:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cc2d4ef3-2bfd-4417-bf11-76076d2b340e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 19:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cc67fbba-e8f8-4f8d-8035-6d9b5359d083', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 21:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cc724931-876d-41b6-ac09-b71188e7e319', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cc81725f-b73f-4984-8cb6-8d9235672341', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '12.80', '0.54', '0.10', '2021-03-06 16:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cca851eb-d14f-4321-8236-fcf4efa253e5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:06:06', '0');
INSERT INTO `monitor_host_log` VALUES ('ccccac0a-517a-4008-9a14-a3588adb24b6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cd0c4ea6-2a49-468b-b067-421b78b94c8a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 13:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cd1f7cd0-f09b-4150-9603-7da232682fbe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cd2016e9-70b4-4056-a471-17c4137b401f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cd435309-046c-4552-9f66-c45fd52e9847', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 17:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cd76c7ad-5179-495a-af69-e0064c386959', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:42:07', '0');
INSERT INTO `monitor_host_log` VALUES ('cdba2a68-e2e8-4e7a-aa76-bfd15b22b547', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cdc25375-7afb-4a40-a963-0b08d00e30e8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 10:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cdcf39c3-1667-46d8-a11b-5ca1be99fb69', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 15:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cdd2ec02-f7a8-43dc-9bf4-1ce6e2d506e7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cdd5819e-d44f-4534-baa4-2c58f25be4b4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cde2dc47-cc13-402c-8122-eb3d1e841d58', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.20', '0.54', '0.10', '2021-03-11 22:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cdfae3ff-9a4c-4975-86bc-956a26a116b7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:55:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ce00b8c0-9376-4268-a733-7db3d15ab4d8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.13', '0.02', '2021-03-06 17:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ce264c41-1dc4-4a7d-a348-a7a10498222e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ce2a5fc2-7c59-4b09-84e3-cff95893fea4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ce5655ee-095b-49f5-b205-772543ac3317', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cea9a7e1-b083-4ee6-9238-656773bc952f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 23:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cedc86b6-5f1e-4839-999e-4e7cbbe70762', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.20', '0.54', '0.10', '2021-03-06 22:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cefd6e72-9534-48d0-aa6c-f3dedfefeae4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cefe8627-dbd7-4266-8ff5-35c6625cd06c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cf0707fd-e640-4f21-816f-a7f5e4cfacec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 15:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cf17d6bb-8e09-4a22-b55c-64a59da5acbf', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 23:44:01', '0');
INSERT INTO `monitor_host_log` VALUES ('cf489c26-f1aa-4f31-817a-cdf258e83392', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:46:12', '0');
INSERT INTO `monitor_host_log` VALUES ('cf4cb32e-af97-4006-9550-4972e8f901e9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:29:03', '0');
INSERT INTO `monitor_host_log` VALUES ('cf6c505f-f591-4822-a578-2f9549c2cf93', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 19:44:06', '0');
INSERT INTO `monitor_host_log` VALUES ('cfece88b-73cb-4340-b6ae-d2e874d8a1ab', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:29:02', '0');
INSERT INTO `monitor_host_log` VALUES ('cfed3e9c-94b8-40a0-b570-d76e00524763', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('cffd7a7e-acb3-43e0-b810-6008039ed42a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d03f06e0-336c-46cf-9c2a-78ad80f40d2b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d0881400-ed75-4af0-92db-251acc78c62e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d0941a74-4d06-49a6-9f44-fccab0a0053a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d094a5a9-2d6e-4d20-a365-e193d5e1d9d7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:11:06', '0');
INSERT INTO `monitor_host_log` VALUES ('d0adcf59-ffe5-4e31-8c2a-7dc29264f5b4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d0b37b8b-60c8-4218-a682-f1993116cd12', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d0beaa86-8c5f-4df4-bd94-2f402da2d46d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:25:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d0c31b1d-37aa-41c3-a3d7-9fc3a27683ed', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 14:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d0e4f5bb-257f-4ed2-bdb6-e8b5df111f2b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d0f9ac04-a710-47fa-949e-9870d4ffd15f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.11', '0.02', '2021-03-14 12:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d0fd4c33-d22e-434d-91bc-ae44c9c834aa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 22:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d1174045-317e-4ca7-b5a9-4a3a44bbb82e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.15', '0.02', '2021-03-08 00:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d11a2113-d6a9-40e9-b68b-8c1cbefb66d9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 16:36:01', '0');
INSERT INTO `monitor_host_log` VALUES ('d12bf175-1afa-45f4-aee2-79a34cebc5b1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d19dbe48-943e-489b-8bb5-74d9374a2e80', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d1abfd90-137d-45c4-9c3f-041d58bbb63e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-13 22:42:06', '0');
INSERT INTO `monitor_host_log` VALUES ('d20bd98c-dbe9-4ba7-aa7a-92e15b82b075', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d2287cee-eaa6-4cee-acfa-04823e06311f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d22ec6ca-534a-4a28-8395-103e8c04757b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d266e103-a756-40bf-840e-32142613bf73', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d27704f8-1a0c-4d7d-8ecc-dc4738b688c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d2846e89-c433-498d-a395-1d99f6d64fd0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d28c06f9-b222-46f8-a8d9-390a6deb9ae6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d2a52fcf-75f9-44fe-85c3-4a65187af4c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-07 09:10:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d2d611cd-cb6c-4177-a985-249b4a8c2d70', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.56', '0.10', '2021-03-14 12:16:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d2df6596-2a96-4773-b960-e85258b5ed3e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d305256f-3fe6-43fa-b670-5b6aa8e6bc0e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d305f6b8-e72f-4a25-84c1-b5870d937e45', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d31d1c3c-fd42-4ce2-b4f7-56423c7a090f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.03', '2021-03-10 21:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d32e118e-b3e7-4bb8-928d-c48f367b4a19', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d3358d59-ee14-438d-9cd4-f423add3d583', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.54', '0.10', '2021-03-07 15:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d352feb8-f606-4390-bce2-a96a24133016', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 23:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d3b30420-6244-438a-b93b-6424f99a3692', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 22:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d3e17553-0b7e-4919-b066-cd6d1f34be87', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:21:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d3fd7cd0-9e00-4432-8722-885ac0c55212', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d3fe2e35-312d-4611-8b4b-b0708349639b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:00:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d40660d5-9c11-494c-b919-1d405569cb5f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:21:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d420433a-ee5d-4998-a5fb-cfc671403e7d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.57', '0.10', '2021-03-14 13:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d43ca235-7b26-427a-a859-eff43e13f163', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d43f99e6-7ec7-4315-9b4e-ad2d6cd529fb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 16:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d445a31b-7915-4e14-8d72-6c9daa80c523', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 19:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d466a0e5-42fe-48c4-8bb3-77dafc80ae85', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 17:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d4d17c87-1f12-4384-9514-13ff6fa4ce06', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d4d4234b-9aad-4c58-9c47-ed8094b33899', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d4eec2f9-df1a-4b19-bc5f-171af6769319', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d4f69053-b0cf-44df-bed6-3206f58587ca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.14', '0.02', '2021-03-07 23:35:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d4fb9079-e54e-4f9e-9479-74217c952442', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 10:46:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d506d468-732d-4688-984d-e726bf582bbe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d515332d-0a38-4941-a46f-aea3398524c7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:44:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d52c3b23-3e91-4af2-953b-ee43a0b2cb98', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d535e556-51b9-48fd-8522-3a154d2c69c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-14 12:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d556c8a7-f081-4c2c-9151-381054653160', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-07 00:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d55ce735-b0c7-4406-be9e-6b6e6fed4aa2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:28:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d56940e1-ac23-4c75-9838-cc153babcc22', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:13:06', '0');
INSERT INTO `monitor_host_log` VALUES ('d56ed8e3-c165-4c0a-88cc-18221f1ad94c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 18:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d5849071-dd9d-464a-93ae-6a3571f036c4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-08 20:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d595ab36-6b36-49af-93be-afddcc25293c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 13:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d5b20dc0-659f-462b-a2a0-ea00ad3b8500', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.40', '0.54', '0.10', '2021-03-06 23:21:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d5ce3424-f314-4376-8254-83dc297e5886', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 19:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d5d44fd1-96b6-45d9-ac63-2524c706dd07', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d6015e01-1d08-44d0-869c-7d4a0b2a7cd1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 15:59:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d60321e1-1198-4177-9abe-994c2f445390', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d604b9b9-3dd5-414d-adcf-4eab45c7b26f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.76', '0.57', '0.10', '2021-03-14 11:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d605104d-9c39-422a-bc2a-994b0b8b5f46', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-14 13:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d60ebf9d-eaf5-48ff-8431-a213f47ee629', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d6435d94-9ef8-4eda-9774-449a7dba2c62', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d677b011-c68d-46aa-9339-78535313b062', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 23:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d67f6ba6-4d16-4d6f-a07c-416a11f6ed4e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 09:34:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d6903cdc-b628-426e-a06b-96aec91972f8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-14 11:09:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d6d440ed-b80e-4273-96c4-9fd75761069f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:47:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d6fdacab-e145-429d-975c-e1e805ae286a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d740af2f-6914-4728-b246-d4dfde4aff27', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.54', '0.10', '2021-03-06 23:08:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d7596ecb-a949-4701-8c39-0facc1cad048', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 15:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d76e971f-a3fa-4257-9141-bd745d4889d6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d77377ff-fecf-4842-81c6-cf3d2427e97b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:36:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d78488a9-d236-4259-9fc8-f10d86f8dca8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d7db6c4b-fac3-40cf-8c79-4b6c5332abf1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 15:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d7f3173b-eb42-4bd8-90a6-ef912887f92c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.13', '0.02', '2021-03-06 15:50:01', '0');
INSERT INTO `monitor_host_log` VALUES ('d7f441fe-a0d0-4e38-97f9-0fab886db827', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d7fbd44a-d723-4504-86d7-50d85a6c0c61', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d81cf2ae-7dbf-47d9-9a79-055789f52e7a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 23:09:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d8336874-59f9-408e-b96d-edaf3a48232e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:30:03', '0');
INSERT INTO `monitor_host_log` VALUES ('d89ad877-dc43-42a9-920d-81d989386bd7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 11:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d8a90096-0daa-4e63-96c7-2f907d43b3b8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:15:01', '0');
INSERT INTO `monitor_host_log` VALUES ('d8aa8a81-3673-4f8b-b510-6b6b014bf528', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d8e52dbc-306f-4cfb-9d3e-a4b9f229f1d0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d901d006-36fb-4ac1-9c02-9fc951d6571f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 09:32:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d973a947-f1cd-4f93-95bd-b47e5f6599de', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-11 00:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d9797965-c3fb-4e37-a8d8-96d8e515f437', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d9a17b66-326f-45ad-ac06-ad219a818659', '', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 14:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('d9da0685-bec6-432f-8591-26aa4b8dc712', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 23:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d9e94c44-fe12-4049-8541-8492435caf1c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 22:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('d9ecfbb1-5aab-43f9-908d-4b09206d2aab', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 19:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('da49a401-d7a7-4006-921b-2e2afa1f0a1c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 14:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('da6a635d-ca8d-40e4-b834-61b0ee763796', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('da736ed2-e194-4c0b-9f71-4e7a832ef242', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('da7a3840-db93-4b2b-b2d5-23f268fa3c01', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.02', '2021-03-07 23:20:02', '0');
INSERT INTO `monitor_host_log` VALUES ('daa972e8-dd7b-4f27-a902-088e6a68e513', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dab357a1-01e7-4b0b-9d53-ec69e2806f1c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dabba1d6-cebc-4214-8f45-079d6dbb2c8a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 15:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dacd6c6e-c7a7-4f42-8882-5d440010de97', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 23:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('dad47c12-6def-43f6-84c5-8d890c276ac6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:41:01', '0');
INSERT INTO `monitor_host_log` VALUES ('db31329c-404a-49e6-b597-d1e20adf9107', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('db92cd01-12ed-4ac9-9be6-73504f1f01b2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-10 21:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dba89e20-de35-493b-b3db-acb102d38092', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:47:03', '0');
INSERT INTO `monitor_host_log` VALUES ('dbaf3dc5-d0cc-45e7-8d12-9b4803e5d790', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dbb83251-3464-44bd-b284-73fb465cb6a6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dbdf5a23-1d91-46c8-960a-ae1dc7ff4398', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dc10891f-8218-4638-b98e-aee9a3fdb5e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 10:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dc22c762-6bfd-4188-be42-bb8e47608620', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dc342963-a887-4fcd-a292-3200e0794a16', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 19:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dc67e0ce-937c-464f-875d-e7390bbc1e4f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dc818efd-56e6-45de-88c4-583f161d8c7a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 13:00:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dcc30c61-2b02-4ed3-bdff-2e02bcb65058', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 15:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dcc882ae-a1f7-4454-88ca-b434ff15a179', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 22:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dcd8bf6b-dc29-4bc6-a94d-5944c296d77c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.02', '2021-03-07 23:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dcf0608a-ab71-4649-b06b-aae1e8fff83f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 16:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dd0bdd56-ac89-49b8-a1e4-e04dd630357b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('dd4681e1-9839-4987-a685-5d6b8ddcf1ee', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.18', '0.54', '0.10', '2021-03-07 14:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dd542ed7-3d9f-4752-9b49-999a329aa407', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dd65770d-05fc-4e2d-bdda-05b79fb228ce', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dd72ebe4-45af-4727-b781-3912b3fe00e2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-09 22:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dd8aa85a-bc69-40e3-a0c4-52ec653ab104', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dd9b985c-13cd-4c1a-8458-ed5059aaa2bf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ddab724e-2410-4649-b184-0bb19b0bfdc6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.54', '0.10', '2021-03-07 21:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ddac999f-0bb7-4198-b03d-343aed894879', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 20:00:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ddea106a-9159-472e-b2a9-92e35780f08d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 22:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('de01cfd3-f0bb-490c-95d1-b11db12cb461', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.18', '0.54', '0.10', '2021-03-09 23:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('de0af3f2-31cd-4f3e-b837-564d9175e804', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('de426c00-87ba-405d-a238-2ee169abd59f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.54', '0.10', '2021-03-09 23:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('de476394-c211-4ccf-9640-2ef312fbd0d4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 13:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('de7b3cc6-cd98-4435-91a0-8778a5d73357', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:50:06', '0');
INSERT INTO `monitor_host_log` VALUES ('deb8a336-d1d1-431f-98a8-aa17e27c263b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('dec0fc64-25b3-4e28-88e5-2bd0f5a9849d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 16:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('def4e910-5386-4858-8929-3c2c75abacb7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:31:04', '0');
INSERT INTO `monitor_host_log` VALUES ('def57766-2563-4ffe-81f7-00c57be7fb31', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:29:04', '0');
INSERT INTO `monitor_host_log` VALUES ('df2bb5dd-d76f-4b67-b5a0-6b1df0e17e47', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 14:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('df420118-56d9-48f5-b738-d8bed6b9406c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:24:04', '0');
INSERT INTO `monitor_host_log` VALUES ('df6c0177-f356-4e63-9820-a4b01498aa8e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 23:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('df8acdd0-06c8-43f3-ba51-7fd480767708', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-14 13:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('dff2a155-f00d-4d2d-9103-44e4207b8b10', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:54:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e0108b1f-b94c-470e-a2b8-062823690681', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.56', '0.10', '2021-03-14 12:06:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e03b09bb-5e6f-4d3a-bb06-eb47b0ce1680', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:37:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e051c93c-1ee7-4d47-b6ce-7b28c4e12f7a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e05e238d-b294-40d9-a31c-93c397a5fa44', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:34:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e05e6919-f6ef-4d00-aa22-369e14dfd5ed', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.50', '0.54', '0.10', '2021-03-06 23:49:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e0727121-57d0-4d11-9a20-efe0db0e909d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 23:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e086e435-98b2-40ae-a310-02be627f6d77', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-09 21:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e09f5121-b63f-4791-ab5d-1a7c9ee15b49', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 20:18:06', '0');
INSERT INTO `monitor_host_log` VALUES ('e0b6525a-3140-4c9f-b58f-67885a938c23', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 22:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e0c0b519-c19a-4ada-9578-106fb5d9f09d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e0e27515-7ca3-44e8-b9dc-9c8c9d9a8298', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:07:06', '0');
INSERT INTO `monitor_host_log` VALUES ('e109092b-f8bb-4d79-b5ce-7862f370ab0d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e113b014-c9b2-4265-b54f-cb2aef2ff650', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 23:07:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e123a9ee-e57f-444b-9231-41d67001b0bc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-07 09:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e1271705-acb8-41c5-988a-468f32bdf6b8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-10 21:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e12acf1f-9293-4cbd-a573-8b0d1ad07917', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 23:46:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e1334fae-ddfe-4918-becb-c56e21e01770', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 10:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e1459072-1f0c-43d1-afd2-d7eb7ff021e1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e14d07ae-f7cf-4056-b4b7-f25a8dbf54c6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:43:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e17b3e16-50d6-4f4d-affe-4d9d42675c47', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:13:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e1822a78-ac48-46cd-82be-ee9ace9de26c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.19', '0.54', '0.10', '2021-03-07 19:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e18de9da-af8a-4e73-9633-8f43709cc502', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.15', '0.02', '2021-03-08 00:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e19a3067-c1e8-4366-8ad9-a1f51e314e94', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-08 20:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e1b97570-026d-4d0b-adb3-c96ad16c06e6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:50:06', '0');
INSERT INTO `monitor_host_log` VALUES ('e1c10c2f-526f-4596-8882-925f0d9c7079', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 13:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e1d21b84-55e1-43b0-83b3-cc0603a1b9f5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 11:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e1ef2d9d-54d4-4386-9c00-c7a817cdf7c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e20973ad-06d1-4b4b-bd7a-caad25268d4a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:44:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e21c59b6-6364-4a00-a167-dd1070bab357', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '25.00', '0.54', '0.10', '2021-03-06 17:12:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e21de83c-a853-4b6d-8ec6-f7f2815659a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e2423131-fd03-4ab6-ab38-e5688b303feb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:50:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e25a53a2-b80d-45be-853c-39cd6bf23d00', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:19:06', '0');
INSERT INTO `monitor_host_log` VALUES ('e267dc4a-f547-4234-b7aa-8c93a3819c0c', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-09 22:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e275a74c-b6f5-4cf2-931e-8f159f3d69ad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-09 23:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e2a41742-b9b5-4046-ad03-33e1fc39375e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:39:07', '0');
INSERT INTO `monitor_host_log` VALUES ('e2aa3ef8-b624-411b-981b-3541c416ae19', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:44:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e2b185a4-ae95-425c-ae46-0f1a3083154e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.13', '0.02', '2021-03-06 17:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e2ce5c7d-70dc-488c-bc24-80d2dc6361f7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-07 23:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e31247af-a50a-4c01-b099-c55e90be6e90', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 23:22:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e31a4011-a09d-4928-b89a-0e2455ecf051', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.13', '0.02', '2021-03-06 15:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e341654c-6fba-4c5d-a9f3-b2641fa02c64', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-09 22:17:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e3467b9c-87bc-46f0-b999-c5ec0ecf92f7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e34957a6-b20e-4db1-b94e-6a75ceb2a26c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 21:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e35f16a5-85b1-457c-af4c-e3ba2e499776', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 22:58:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e378af33-e011-46ea-bdf9-abde4eb1faec', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.14', '0.02', '2021-03-08 00:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e37c3823-f89a-4783-98c1-6dd27673fc01', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.56', '0.10', '2021-03-14 11:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e38d6702-cd51-4c12-b805-b5e8b7ac6a3e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e391446b-3ebc-4d3f-bf0b-a6a8a0802e89', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-09 22:14:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e3c9740a-d8cf-4527-ae1f-db6bc7f2b6b8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.17', '0.11', '0.02', '2021-03-11 22:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e3dedb05-025d-4484-80f2-5def77744874', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-10 20:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e40bb87f-9ef8-4747-b99b-cf5e1d9299ba', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 22:18:13', '0');
INSERT INTO `monitor_host_log` VALUES ('e41ebffa-a2e1-4e14-9d10-c9412efc49f9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.15', '0.02', '2021-03-08 00:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e4380909-0780-48aa-90f6-fe67aa9a6733', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-07 09:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e45dd571-0b39-4ba8-98bb-052a882f1ed4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.60', '0.53', '0.10', '2021-03-06 23:43:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e48db9eb-5e47-48b3-a8be-b7c49d403b06', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:46:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e4b49659-f8c8-423a-a500-9a212c606ade', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e4c6b0cd-2157-4bf7-8f36-8e507ffd6faf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.60', '0.56', '0.10', '2021-03-14 12:12:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e4d6307b-d921-45e1-b1b2-c982480631fe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e50888eb-2ec6-49c6-98df-88dd8a5c6cc0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.54', '0.10', '2021-03-07 22:41:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e5096fea-df65-46aa-a88b-9acc036ff86b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-14 12:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e5103d55-4e28-478a-8f9a-52b87d0ba97d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:09:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e51c6fa5-0719-4a0a-be5b-12a3546367ea', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e52d2f96-2b1d-492c-a92b-5777bc6ac8cf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e53015d7-d30d-42d4-b4d4-2afd61c99c48', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e533737a-51d0-4234-8f5a-2124a30f743b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e5a357e9-d371-4106-9b28-0a18d7fecee0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 10:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e5a6cade-89e1-4c30-905b-bb55805c8a98', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 20:20:06', '0');
INSERT INTO `monitor_host_log` VALUES ('e61168f8-1b6f-4a88-ae85-8010e59be699', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-13 22:41:08', '0');
INSERT INTO `monitor_host_log` VALUES ('e63b7f88-a4ff-437d-93c1-beb8a24364c2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 21:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e6561530-7bd2-4bc0-be1a-e72a4269d2e4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 14:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e65924a7-feb4-49a2-a7d3-4100e36148e7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:40:08', '0');
INSERT INTO `monitor_host_log` VALUES ('e660b28e-8b77-4254-a99f-2d908c81ad70', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.53', '0.10', '2021-03-07 08:53:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e673ef18-0577-412a-bba6-9b5fddb4f727', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-07 21:33:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e681921d-1457-4333-a862-3c544b980160', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:39:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e7088394-6ce7-43a5-96ac-a1217070d2fa', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-07 09:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e71d55b8-0cc7-4c03-91f2-9aacdf8194d6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e736449a-bd51-4adf-bd37-2a50906a5519', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '8.70', '0.12', '0.02', '2021-03-06 16:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e74ce2ec-719f-4ed0-a69a-4ef5c86ce407', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 15:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e755e858-a1aa-4549-97d3-e408ef51b93e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e763f561-ea0d-4b6a-91dc-65459bbf0435', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-10 20:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e7953cc3-02c0-4023-86f2-89c4495a397b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 15:55:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e798cd3c-bd16-4399-81e3-d780fabdac96', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 18:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e7b7e007-f0f6-41d2-adb9-a664baa3bf7d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-14 11:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e8338688-9143-4fdc-a561-08e01fec2daa', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:18:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e84dfc8d-8ceb-4795-95c7-1c998e1516f7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 21:42:03', '0');
INSERT INTO `monitor_host_log` VALUES ('e87d33f9-626c-44cb-bbdb-800f071e1360', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 21:47:05', '0');
INSERT INTO `monitor_host_log` VALUES ('e882a893-9a5f-4724-bf5a-0888a2100b12', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 22:26:01', '0');
INSERT INTO `monitor_host_log` VALUES ('e8cd15d7-a821-4184-9909-99cb4df25cdd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-10 21:58:04', '0');
INSERT INTO `monitor_host_log` VALUES ('e8edeeed-0013-4f9e-b5ff-aa05154fec10', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 21:02:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e90ecd94-427d-4129-9d48-cf6ab1d8ca8d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.14', '0.02', '2021-03-07 23:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e94ef4f7-7e42-4b5b-802d-e3578e31a986', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 10:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e9b04cfa-76b3-4631-bb5d-5aab076d1d39', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-07 10:43:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e9c184d4-1812-456c-a455-da307c93d123', '', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:33:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e9ccffd6-a94b-41b8-b3ab-f4375bbec0e5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-07 08:50:02', '0');
INSERT INTO `monitor_host_log` VALUES ('e9ce528c-4c90-4d52-a011-f54d61e5951a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-09 21:53:01', '0');
INSERT INTO `monitor_host_log` VALUES ('ea20ab08-4b8f-4755-8ba7-84e9bcafd9c5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 16:42:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ea3909a3-8698-4d0a-812c-d2b0089b8e42', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-07 09:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ea3aa6f2-169a-4908-ae5d-db0091828c2c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 19:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ea411bc4-108f-4f8c-8e58-62cfafab65e2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 10:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ea43aa46-70a1-4a89-ba32-1ec7148f9ffe', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 15:18:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ea6038ce-cd6b-4c2e-ae21-bdc894eb1b9f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.50', '0.53', '0.10', '2021-03-06 23:39:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ea6b6d37-e18a-4dc5-8c17-1d692a5273eb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 19:54:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ea7cc902-52fa-4f48-bf79-b826a1548dd8', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:57:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ea9a7cb1-2429-4461-9a45-0021b7f2d865', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eac3edb4-0185-48e7-851d-d1e19622c881', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.11', '0.02', '2021-03-06 14:45:18', '0');
INSERT INTO `monitor_host_log` VALUES ('eaf078a4-325e-4107-8841-a223627848e4', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 22:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eaf5a744-ce87-4608-a6a3-b89c9d2b7000', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('eb16e3e3-317e-4df3-ac24-c9e86151a4ca', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-08 20:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eb1ed5a3-f683-48b5-b09d-3ae1f535cab4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 16:11:04', '0');
INSERT INTO `monitor_host_log` VALUES ('eb256254-d14d-4444-b154-5848572e49d9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.13', '0.03', '2021-03-10 21:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eb2b5ad8-a023-453e-878d-766b4ea1e7bc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 12:08:04', '0');
INSERT INTO `monitor_host_log` VALUES ('eb8f3682-49be-45e7-9d82-aaf1c43c283a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ebc61dbd-b610-4893-abeb-a424242363b7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 20:26:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ebd1f6c1-3b35-48ab-920e-55b7ccb10e26', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.56', '0.10', '2021-03-14 12:15:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ebdea333-4b97-4585-96b2-275108892dd1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 15:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ebedc590-c8c5-4c70-b7aa-01b56e84e73d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ec032334-95cd-4368-9c67-9a595e695720', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 22:40:01', '0');
INSERT INTO `monitor_host_log` VALUES ('ec084593-e22d-40c6-8f51-cc4dc1a5a82e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-07 21:35:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ec279e32-766d-4586-9275-6a0380990c56', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ec60ec8d-4088-4101-8646-ed26fbe099c3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-14 11:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ec7f3b1f-1321-4640-84b6-bce97b522b37', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 22:38:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ec89a6f5-4a7a-47bf-ad4a-616b6ba31383', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:49:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ecafaeb6-9fb7-4c02-9fa4-4e633abb36e7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 20:51:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ecbf40cd-8227-448f-ab1a-71e7ff53e0f2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 16:48:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eccb4c23-8c50-4d65-85f6-84983ee9fe41', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 22:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ed090727-646f-41d8-aaaa-f00d0bb9803d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '12.00', '0.11', '0.02', '2021-03-06 16:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ed17d8f6-70a2-49a3-8049-51e32e4a8a28', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ed1a5f90-b6b2-4a4f-a787-bbb74d3bf8d1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 16:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ed561126-e1ff-4704-839f-1f9e91c4c4b9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 19:45:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ed6a4e04-e51b-47d5-8925-c48585e6ffb4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:43:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ed8af22d-f186-4149-94ea-3ce5cd815c79', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 17:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('edd2a264-e2fa-4162-ab3f-1038c4c01ce0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-07 10:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ee033c4f-df73-4140-a560-04170e2c8cd3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.56', '0.10', '2021-03-13 22:37:08', '0');
INSERT INTO `monitor_host_log` VALUES ('ee084117-65f0-4ce7-930b-b0feb61ec675', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ee2c3507-f586-4648-9447-ad5dca8aaf8f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.11', '0.02', '2021-03-07 08:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ee4a1c6b-f8fa-4730-8990-c2e7b35c8623', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.53', '0.10', '2021-03-07 08:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ee66a4a3-5ad8-48b6-9603-ebf0664af07f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.55', '0.10', '2021-03-07 23:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ee687fb8-dcc3-40bb-88bf-6f7850ee2e6b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 19:58:06', '0');
INSERT INTO `monitor_host_log` VALUES ('eec036ad-097b-4219-95d8-3593f868ad0b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 23:38:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eee28d84-95bb-4f50-ab76-c24e4bc7d37a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-14 13:27:02', '0');
INSERT INTO `monitor_host_log` VALUES ('eef2b11a-8b34-48b8-975a-70d0a4e652d1', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ef1f0638-0bd4-4ff3-a1d1-06cd72eb263d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ef3323c8-3bec-47d2-8b9a-ed87b0bf47a3', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.54', '0.10', '2021-03-11 22:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ef56d71c-e016-43b7-8ccb-89338e581a0a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 21:19:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ef7b09de-dfd5-4209-944c-7b61ad829e79', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-09 23:40:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ef805aa3-56d0-4f59-a3b9-ee8d2b6ef3a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.56', '0.10', '2021-03-13 22:45:08', '0');
INSERT INTO `monitor_host_log` VALUES ('ef9f690e-3e51-44ca-b7ac-311c993d726f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-14 11:09:02', '0');
INSERT INTO `monitor_host_log` VALUES ('efa0c81d-6cf0-4d0a-9313-900efd3a814f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 22:37:03', '0');
INSERT INTO `monitor_host_log` VALUES ('efade3a1-a726-46bf-a73c-7b4418e039ae', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:16:04', '0');
INSERT INTO `monitor_host_log` VALUES ('efb68449-a603-4a0f-acd6-f84ba0f0f94d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 23:17:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f00f67dc-5513-463b-9593-c372ffafddb6', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f0247d55-78ff-41bc-acaf-daa588c36655', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.15', '0.54', '0.10', '2021-03-11 22:19:06', '0');
INSERT INTO `monitor_host_log` VALUES ('f0290023-3ee9-41a7-84a4-47d39dfb28b0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 19:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f02cbeaa-778c-4896-a1eb-4dff776385d2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-08 20:44:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f05fee3d-77cb-4f31-b34f-d280acd0339e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-10 21:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f0956ae1-78ed-44d2-b123-f7cd7d4289c0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.30', '0.54', '0.10', '2021-03-06 15:48:06', '0');
INSERT INTO `monitor_host_log` VALUES ('f0a71395-b1c7-49a3-affc-a16bf2d85257', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.55', '0.10', '2021-03-08 00:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f0df1e3b-5b98-422a-b12f-5121ae6ec655', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.63', '0.56', '0.10', '2021-03-14 12:03:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f0e4f8be-0fd6-4ed1-bd15-6ee525ea923d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 08:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f0ee4f07-6cc6-4bb2-8503-73d79f567b68', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.11', '0.02', '2021-03-06 14:53:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f0f20aba-d33c-4023-8f58-982ce1b2c195', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:04:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f1640608-1b1c-4d29-a6ca-761b6ad94508', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.90', '0.11', '0.02', '2021-03-06 23:41:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f183b901-a924-4bcb-8f1a-871c6f47183a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:52:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f1a4baa2-2dcc-49ef-864f-6dcaa7f92060', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 22:54:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f1b706e9-491c-4588-8cf0-44343df7d27e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.12', '0.02', '2021-03-06 16:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f1c70538-8059-4401-a7f5-516070e98fe9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.12', '0.02', '2021-03-06 17:11:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f1c77af1-df63-4d95-b608-0f45ee9d17f6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-09 21:47:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f1e947a2-de2f-438f-98dd-dcee18877668', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 14:49:33', '0');
INSERT INTO `monitor_host_log` VALUES ('f1ebff22-decb-4832-a412-3a826c1ea4a0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f219b4d3-2461-4c9a-99af-aeb48072480d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 18:56:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f2395c76-40af-4c2a-8907-aa442b47b853', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.12', '0.03', '2021-03-10 20:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f2464cdf-cc39-4f7f-bafa-4f2b0fabf87a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f24e85de-06d8-41eb-8859-ad139be9fcef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 14:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f26c6505-a3ff-44a8-8fdb-1e592a50b46d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-09 23:16:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f29355da-6976-4121-830e-724c3c436abc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-10 19:43:06', '0');
INSERT INTO `monitor_host_log` VALUES ('f2bcde3b-97ea-4455-81bf-47935ec1627b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f2e5959c-ce53-4fbb-aced-186a1fa93c55', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.10', '0.02', '2021-03-13 22:48:06', '0');
INSERT INTO `monitor_host_log` VALUES ('f2f159b9-aaf7-4c85-a54b-b54a17123291', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:06:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f3243702-ef0a-4d99-81fe-a818799dc24b', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '53.30', '0.54', '0.10', '2021-03-06 23:12:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f338ebf0-03b1-42aa-b5bb-c8e3160db3c9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:50:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f35a4b63-7050-49e2-9a89-bc9815acf2ff', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-10 19:52:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f38c5b70-7070-4873-bd98-754e0fb8caf9', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-11 21:55:05', '0');
INSERT INTO `monitor_host_log` VALUES ('f3f7ba3c-dc66-474c-9578-e12004379e1c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.03', '2021-03-08 20:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f4069988-b361-496a-ad90-007fbf1ee926', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.02', '2021-03-11 21:52:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f4121c5b-2033-4149-8880-95d8a0a4be82', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 15:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f42b0d36-397d-4f9d-8bd6-916f5c5b622d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 22:56:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f44e84ea-0fb3-4ced-bca4-4c16da8ad718', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-07 10:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f45eda2b-769f-4491-ae8d-be7893232956', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f47608a0-915b-4b81-8ac5-df4b6e51e2a2', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 16:03:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f4b9c92a-b332-4b23-b54f-841adb57dc09', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-08 20:46:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f4c45758-134f-4ae0-8f14-1e194d3553c2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.53', '0.10', '2021-03-11 21:51:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f5081a7b-361a-4b1f-b03c-fa2d54817834', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f50eba34-9f0e-4a48-abf4-7721101ead0f', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '27.90', '0.54', '0.10', '2021-03-06 23:10:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f512094a-7025-44b8-a743-fff10c7641ec', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-10 21:17:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f51dc963-17ce-448f-a605-00a75c14bfdc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-11 21:58:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f529c705-396f-4fe5-b0e8-e5709a2c5895', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.10', '0.02', '2021-03-13 22:46:06', '0');
INSERT INTO `monitor_host_log` VALUES ('f55a1789-d065-43d1-b46c-e43ccf26510e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 17:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f567475b-6561-4879-8cec-647e20e695f7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.12', '0.02', '2021-03-07 10:41:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f571167d-dab5-4ccd-9ede-ada57d652688', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:56:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f57587dd-153b-4f4e-a2d0-7f773be097c8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.30', '0.13', '0.02', '2021-03-06 15:40:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f5839b1a-abdc-4303-b4ae-afa8780d02b0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:39:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f5aa290d-fcb2-4cfe-ac0c-f00e539af0ad', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-11 22:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f5db2de6-604e-44b7-b4b6-6bb6332196b5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 23:28:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f5e6c65f-60c1-444c-b833-5dd21fd55a03', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-09 22:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f60299eb-b4aa-4a35-8369-a7b4a2c7070f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 15:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f67752ec-76c9-45b9-9c6e-dbf7dc755350', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-09 22:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f67e759b-a940-48c9-b224-bf595ea6a8d5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.02', '2021-03-07 19:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f68dd427-129c-440a-b489-1e54519cbfe6', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 22:57:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f6acc484-dc56-49fa-8852-3b950c306cd7', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.57', '0.10', '2021-03-14 13:14:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f6be3849-8492-46f6-8f80-4bab363b5890', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.27', '0.57', '0.10', '2021-03-14 12:23:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f6cca07f-2711-4aa3-92a8-1b09d8e0ddfd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f6dd0652-7005-43bc-a59e-e7ec7ac531a8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.50', '0.13', '0.02', '2021-03-06 23:26:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f73903f4-4baa-4a4b-b427-c58f4fd656cd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 19:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f76993b5-efe1-4d83-8ad7-0689a96e5b8d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 22:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f7969945-0cee-40d4-9728-c4f1b1d4c967', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-06 17:32:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f7af9b3a-83aa-43da-a28d-6e25e7b3cf9d', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.13', '0.03', '2021-03-10 21:30:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f7b3ffb8-cf51-498b-b9c9-e70100bbddc8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:17:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f7cd0000-fdad-4ad7-b41c-23606498de70', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 19:44:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f7d0e0ad-541e-49a8-ae6a-76a6c23f2dbf', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.29', '0.56', '0.10', '2021-03-14 12:02:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f7d29046-d75a-4460-af1d-7a532ed73ccb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 19:35:05', '0');
INSERT INTO `monitor_host_log` VALUES ('f7da9f3b-e2bd-41af-9fef-d03b2dca060e', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.54', '0.10', '2021-03-07 23:16:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f7dd2556-cb8a-43c9-ac59-2068a9317d29', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.53', '0.10', '2021-03-06 23:41:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f7ebd038-2af2-45e1-be20-9256e611ffd0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-14 12:25:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f82a1af1-6a9e-4129-96e5-b1d51c700451', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.12', '0.02', '2021-03-11 22:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f82dc111-b09e-4998-a7e7-6366ca1e46a7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 22:50:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f83e0c3a-027c-471d-90e7-05b06c6f7242', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.28', '0.57', '0.10', '2021-03-14 12:38:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f8458499-4fe8-46ef-b11d-af69469669fb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 22:23:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f869a817-477f-4171-906c-64450486614a', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:10:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f881b329-9e92-4d42-8ba1-34a815b61923', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 21:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f88fe439-53be-4749-a564-acfc2c29fd24', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.36', '0.57', '0.10', '2021-03-14 13:12:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f893a6d8-faf1-4446-b793-1e646fe90d49', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-09 23:25:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f8daa75a-a70c-4ce0-b8c8-f949bd69ba35', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.03', '2021-03-09 22:30:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f90e52fc-f8bc-4d5c-9bf1-e7bf6cc4b2d8', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.11', '0.02', '2021-03-11 22:26:03', '0');
INSERT INTO `monitor_host_log` VALUES ('f91623fc-7da7-4114-a0ea-1105766b6647', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f9240ec3-c53f-4859-b085-02d5dd5fed2a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.26', '0.57', '0.10', '2021-03-14 12:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f92c7f2f-19c8-42e0-b615-367e49d424d0', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.03', '2021-03-09 23:26:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f93236e9-dc70-42c6-81eb-587ad7dead31', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 23:16:01', '0');
INSERT INTO `monitor_host_log` VALUES ('f95fed10-3c9b-4ca7-a6f7-bd22bf6d1040', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.10', '0.12', '0.02', '2021-03-14 13:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f980a9ce-50d4-4861-bd06-b9c154f05291', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.55', '0.10', '2021-03-07 23:40:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f989ed6e-15c7-40e5-b587-7dba18d8e3e0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-06 16:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f98d2b68-fbd0-4a42-a7a3-0935a901dcfc', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.30', '0.54', '0.10', '2021-03-06 16:20:04', '0');
INSERT INTO `monitor_host_log` VALUES ('f994e8e1-d735-475e-abb5-98aac7768ceb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.12', '0.02', '2021-03-07 10:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('f9b5fc70-32c2-435d-9e8b-f26956bfeb79', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 22:31:01', '0');
INSERT INTO `monitor_host_log` VALUES ('fa11685d-bc21-42a4-9e5f-988130453298', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.50', '0.54', '0.10', '2021-03-06 18:06:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fa2bf396-5964-4e8f-b3fa-ed7e2ba550ab', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.01', '0.54', '0.10', '2021-03-07 18:57:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fa48b518-f079-4406-a858-73a87ce00cdd', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.30', '0.54', '0.10', '2021-03-06 23:24:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fa508dad-0a22-4037-9c6a-8b97cc8d84e2', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 21:28:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fa864edc-399d-4a1a-9b8a-899365887e60', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.56', '0.10', '2021-03-14 11:13:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fa99c9a9-45f9-431d-b275-a185b8290400', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 21:45:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fab21a12-8c8e-4031-a39d-89fc15591470', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.54', '0.10', '2021-03-07 15:05:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fafcb710-68a4-45e0-a7ad-f86b03ab322a', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 09:58:55', '0');
INSERT INTO `monitor_host_log` VALUES ('fb101c70-9918-48e4-898d-217b7f0f6ce9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.50', '0.13', '0.02', '2021-03-06 17:42:05', '0');
INSERT INTO `monitor_host_log` VALUES ('fb31b20e-57b3-4446-8cbe-da6cd7d079cb', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.04', '0.53', '0.10', '2021-03-09 21:45:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fb5f6caf-797c-4e54-b47b-f3ffdd73b5dc', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 14:49:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fb741d20-e739-4e26-8204-275013f54a35', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.13', '0.13', '0.02', '2021-03-07 23:24:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fb867eb0-4384-4860-8612-e43d67e0c62b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 14:16:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fbeedb18-52da-4c62-810d-3417026b0a7b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '9.10', '0.12', '0.02', '2021-03-06 17:14:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fbff4877-513c-4234-81aa-0f3c2f84be80', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.03', '2021-03-10 20:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fc1683a5-8569-46a8-a34d-550154f2bc46', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 15:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fc292ce9-2334-4720-b2d9-9d340c189a83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.12', '0.02', '2021-03-06 23:09:01', '0');
INSERT INTO `monitor_host_log` VALUES ('fc9fef66-22e9-4358-9c66-401bb650571b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 19:01:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fca7727e-eb2d-4151-b7de-d5d72a39dc1f', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.11', '0.02', '2021-03-06 18:07:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fcb09b7c-9121-4e65-8d39-b1479831b646', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:32:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fcf404dd-d094-4e8f-ae31-8312775c4db0', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-11 22:24:05', '0');
INSERT INTO `monitor_host_log` VALUES ('fcfb5947-1e8e-4bba-9a1e-e6836b600486', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 18:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fd159dac-edba-4e24-b670-914edd084fda', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '3.40', '0.54', '0.10', '2021-03-06 17:23:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fd215346-ceb4-4671-8942-fa59e5bffd83', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:15:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fd290fed-979a-4ca3-9794-f5e7a53a58a7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-07 15:05:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fd4ded7d-2454-483b-947e-66c6a21e5eda', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 14:37:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fd5632ff-5550-4a87-9a7a-1bf12d3a290b', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.03', '2021-03-10 21:36:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fd5e76cf-3460-4e31-ba55-09f0901433c1', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 23:27:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fd784c3c-3ab6-47f1-949c-2e181300a773', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.00', '0.11', '0.02', '2021-03-06 14:44:06', '0');
INSERT INTO `monitor_host_log` VALUES ('fd7f6a31-07b5-4b86-ba06-80bb2eaca7dd', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.03', '2021-03-10 20:07:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fd8838d6-b847-4d33-819d-a8cdf95f5016', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '2.40', '0.54', '0.10', '2021-03-07 09:11:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fd8d91b5-3a0c-409a-ab3b-53bb339a7d14', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 19:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fddcbc4d-5384-4762-b479-fcd1c2e8c0e5', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 14:54:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fdf56869-aaab-4f82-ae9c-189e420a901d', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 10:01:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fe039c25-a019-4db3-a8ed-a251e4912143', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.13', '0.02', '2021-03-07 23:12:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fe35db4f-c011-432e-848e-6bbbfaa99e8c', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.11', '0.11', '0.02', '2021-03-11 21:55:03', '0');
INSERT INTO `monitor_host_log` VALUES ('fe5a76e9-7ba5-4da4-b0d5-346c4458d7ef', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.09', '0.13', '0.02', '2021-03-07 23:11:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fe79ebec-ff9d-45de-9b71-ae69d5a9cd47', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.02', '2021-03-07 10:22:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fea53d97-f838-4125-b787-ff0807be9586', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 15:04:04', '0');
INSERT INTO `monitor_host_log` VALUES ('febde92f-c0e9-4282-994e-f80d596d41d7', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.11', '0.02', '2021-03-14 12:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('febe725c-7b44-418f-9b73-d20cbcec86c9', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.14', '0.03', '2021-03-09 23:28:01', '0');
INSERT INTO `monitor_host_log` VALUES ('fece4ac3-b89f-4cd6-be3c-773f74d12e48', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.12', '0.03', '2021-03-10 19:59:04', '0');
INSERT INTO `monitor_host_log` VALUES ('fef2bda3-050a-4c93-ab51-70edc49c50bb', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.12', '0.02', '2021-03-07 22:31:02', '0');
INSERT INTO `monitor_host_log` VALUES ('feff3f14-8022-4027-b805-a3997dbdf084', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.54', '0.10', '2021-03-07 22:43:03', '0');
INSERT INTO `monitor_host_log` VALUES ('ff1ef967-4c59-4bc6-8000-818622c15ca4', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.02', '0.53', '0.10', '2021-03-08 20:22:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ff5340f4-ff05-458f-b4d5-170dce68c1b3', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.10', '0.02', '2021-03-11 00:50:05', '0');
INSERT INTO `monitor_host_log` VALUES ('ff683698-708f-4f4c-9fe3-0d54e46c6159', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.06', '0.11', '0.03', '2021-03-10 20:15:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ff6c96fe-68d5-44ab-b9b4-f081f2e493ae', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.05', '0.53', '0.10', '2021-03-10 21:48:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ff9d17b2-10ee-4efe-82fd-a99c896d3b0e', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.08', '0.13', '0.02', '2021-03-07 23:01:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ffb17bab-8dd8-4849-bb7b-57c886dc3646', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.11', '0.02', '2021-03-06 17:07:01', '0');
INSERT INTO `monitor_host_log` VALUES ('ffbaba8b-67f4-438f-8ff1-4f68f244ec69', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.80', '0.12', '0.02', '2021-03-06 15:19:02', '0');
INSERT INTO `monitor_host_log` VALUES ('ffbd8f38-885c-4c03-a770-c8f9d6eaa537', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.03', '0.54', '0.10', '2021-03-07 10:42:04', '0');
INSERT INTO `monitor_host_log` VALUES ('ffcbd478-b7bf-4a9d-b5ff-794c2e595356', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.11', '0.02', '2021-03-06 15:53:01', '0');
INSERT INTO `monitor_host_log` VALUES ('ffd0b8dc-f5fb-41a2-a40b-341a60263b99', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '5.60', '0.12', '0.02', '2021-03-06 17:40:02', '0');
INSERT INTO `monitor_host_log` VALUES ('fff64e49-2bf9-4af6-89d6-fff7b8b4f735', '0638e278-484c-4e49-b59f-9ed73ab418ed', '512c369e-0642-41e5-9ea6-5fe737984ae6', '0.00', '0.13', '0.02', '2021-03-06 17:43:02', '0');

-- ----------------------------
-- Table structure for privilege
-- ----------------------------
DROP TABLE IF EXISTS `privilege`;
CREATE TABLE `privilege` (
  `p_id` varchar(64) NOT NULL ,
  `name` varchar(10) NOT NULL ,
  `description` varchar(200) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `pu_id` int(11) NOT NULL ,
  `p_id` varchar(64) NOT NULL,
  `privilege_flag` varchar(4) NOT NULL DEFAULT '0' ,
  `user_id` varchar(64) NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `sql_id` int(11) NOT NULL ,
  `sql_code` varchar(128) NOT NULL ,
  `sql_text` longtext NOT NULL ,
  `remark` varchar(256) DEFAULT NULL ,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  PRIMARY KEY (`sql_id`)
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
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
