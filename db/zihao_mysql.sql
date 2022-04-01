/*
Navicat MySQL Data Transfer

Source Server         : 梓豪
Source Server Version : 50623
Source Host           : dev.db.java110.com:3306
Source Database       : zihao

Target Server Type    : MYSQL
Target Server Version : 50623
File Encoding         : 65001

Date: 2021-11-23 10:22:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app_service
-- ----------------------------
DROP TABLE IF EXISTS `app_service`;
CREATE TABLE `app_service` (
  `as_id` varchar(64) NOT NULL COMMENT '服务ID',
  `as_name` varchar(128) NOT NULL COMMENT '服务名称',
  `as_type` varchar(12) NOT NULL COMMENT '服务类型，001 数据库 002 缓存 003 计算应用',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `as_desc` varchar(512) DEFAULT NULL COMMENT '应用名称',
  `state` varchar(12) NOT NULL DEFAULT '10012' COMMENT '状态  10012 停止 10013 启动 10013 升级',
  `as_count` int(11) NOT NULL DEFAULT '1' COMMENT '容器数',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`as_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `av_id` varchar(64) NOT NULL COMMENT '主键ID',
  `avg_id` varchar(64) NOT NULL COMMENT '组ID',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `var_name` varchar(128) NOT NULL COMMENT '环境变量名称',
  `var_type` varchar(12) NOT NULL COMMENT '系统变量 1001 普通变量2002',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `var_spec` varchar(64) NOT NULL COMMENT '变量编码',
  PRIMARY KEY (`av_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `avg_id` varchar(64) NOT NULL COMMENT '组ID',
  `avg_name` varchar(128) NOT NULL COMMENT '组名称',
  `avg_type` varchar(12) NOT NULL COMMENT '组类型，001 普通',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `avg_desc` varchar(512) DEFAULT NULL COMMENT '应用名称',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`avg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `av_id` varchar(64) NOT NULL COMMENT '主键ID',
  `name` varchar(128) NOT NULL COMMENT '应用版本名称',
  `remark` varchar(512) NOT NULL COMMENT '版本描述',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`av_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_version
-- ----------------------------

-- ----------------------------
-- Table structure for app_version_attr
-- ----------------------------
DROP TABLE IF EXISTS `app_version_attr`;
CREATE TABLE `app_version_attr` (
  `attr_id` varchar(64) NOT NULL COMMENT '主键ID',
  `av_id` varchar(64) NOT NULL COMMENT '版本ID',
  `version` varchar(32) NOT NULL COMMENT '版本',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`attr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_version_attr
-- ----------------------------

-- ----------------------------
-- Table structure for app_version_job
-- ----------------------------
DROP TABLE IF EXISTS `app_version_job`;
CREATE TABLE `app_version_job` (
  `job_id` varchar(64) NOT NULL COMMENT '任务ID',
  `job_name` varchar(128) NOT NULL COMMENT '服务名称',
  `job_shell` longtext NOT NULL COMMENT '构建shell',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `pre_job_time` datetime DEFAULT NULL COMMENT '上次构建时间',
  `cur_job_time` datetime DEFAULT NULL COMMENT '本次构建时间',
  `state` varchar(12) NOT NULL DEFAULT '10012' COMMENT '状态 1001未构建，2002构建中，3003构建失败，4004构建成功',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`job_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `detail_id` varchar(64) NOT NULL COMMENT '明细ID',
  `job_id` varchar(64) NOT NULL COMMENT '任务ID',
  `log_path` varchar(256) NOT NULL COMMENT '日志路径',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `state` varchar(12) NOT NULL DEFAULT '10012' COMMENT '状态 3003构建失败，4004构建成功',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
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
  `host_id` varchar(64) NOT NULL COMMENT '主机ID',
  `group_id` varchar(64) NOT NULL COMMENT '组ID',
  `name` varchar(64) NOT NULL COMMENT '主机名称',
  `ip` varchar(20) NOT NULL COMMENT 'ip',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `passwd` varchar(64) NOT NULL COMMENT '密码',
  `cpu` int(11) NOT NULL COMMENT 'cpu核数',
  `mem` decimal(10,2) NOT NULL COMMENT '内存',
  `disk` int(11) NOT NULL COMMENT '磁盘大小',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`host_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `group_id` varchar(64) NOT NULL COMMENT '组ID',
  `name` varchar(64) NOT NULL COMMENT '组名称',
  `description` varchar(200) DEFAULT NULL COMMENT '组描述',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `domain` varchar(50) NOT NULL COMMENT '域',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `zkeys` varchar(100) NOT NULL COMMENT 'key',
  `value` varchar(1000) NOT NULL COMMENT 'value',
  `remark` varchar(200) DEFAULT NULL COMMENT '描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

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
  `m_id` varchar(64) NOT NULL COMMENT '菜单ID',
  `name` varchar(10) NOT NULL COMMENT '菜单名称',
  `g_id` varchar(64) NOT NULL COMMENT '菜单组ID',
  `url` varchar(200) NOT NULL COMMENT '打开地址',
  `seq` int(11) NOT NULL COMMENT '列顺序',
  `description` varchar(200) DEFAULT NULL COMMENT '菜单描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `is_show` varchar(2) NOT NULL DEFAULT 'Y' COMMENT '菜单页面是否显示 Y显示N不显示',
  PRIMARY KEY (`m_id`),
  KEY `index_m_id` (`m_id`) USING BTREE,
  KEY `index_g_id` (`g_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `g_id` varchar(64) NOT NULL COMMENT '菜单组ID',
  `name` varchar(10) NOT NULL COMMENT '菜单组名称',
  `icon` varchar(20) NOT NULL COMMENT '菜单图片',
  `label` varchar(20) NOT NULL COMMENT '菜单标签',
  `seq` int(11) NOT NULL COMMENT '列顺序',
  `description` varchar(200) DEFAULT NULL COMMENT '菜单描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `group_type` varchar(12) NOT NULL DEFAULT 'P_WEB' COMMENT ' 菜单类型',
  UNIQUE KEY `g_id` (`g_id`) USING BTREE,
  UNIQUE KEY `g_id_2` (`g_id`) USING BTREE,
  KEY `index_m_id` (`g_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `event_id` varchar(64) NOT NULL COMMENT '事件ID',
  `event_type` varchar(12) NOT NULL COMMENT '事件类型，1001 主机事件，2002 应用事件',
  `event_obj_id` varchar(64) NOT NULL COMMENT '事件对象ID',
  `event_obj_name` varchar(128) NOT NULL COMMENT '事件对象名称',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `threshold_value` decimal(10,2) NOT NULL COMMENT '阈值',
  `cur_value` decimal(10,2) NOT NULL COMMENT '当前值',
  `remark` longtext COMMENT '描述值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `notice_type` varchar(12) NOT NULL COMMENT '通知渠道 1001 短信通知 2002 钉钉通知',
  `state` varchar(12) NOT NULL COMMENT '001 告警成功 002 告警失败',
  `state_remark` longtext COMMENT '告警成功失败描述',
  PRIMARY KEY (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `mh_id` varchar(64) NOT NULL COMMENT '主键ID',
  `mhg_id` varchar(64) NOT NULL COMMENT '监控主机组主键',
  `host_id` varchar(64) NOT NULL COMMENT '主机ID',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `cpu_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT 'cpu 使用率',
  `mem_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '内存使用率',
  `disk_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '磁盘使用率',
  `free_mem` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '空闲内存，单位为G',
  `free_disk` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '空闲磁盘单位为G',
  `cpu_threshold` decimal(10,2) NOT NULL COMMENT 'cpu 阈值',
  `mem_threshold` decimal(10,2) NOT NULL COMMENT '内存 阈值',
  `disk_threshold` decimal(10,2) NOT NULL COMMENT '磁盘 阈值',
  `mon_disk` varchar(128) NOT NULL COMMENT '监控磁盘',
  `mon_date` datetime NOT NULL COMMENT '监控时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`mh_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `mhg_id` varchar(64) NOT NULL COMMENT '监控主机组主键',
  `name` varchar(64) NOT NULL COMMENT '名称',
  `mon_cron` varchar(128) NOT NULL COMMENT '监控表达式',
  `state` varchar(12) NOT NULL DEFAULT '3302' COMMENT '状态 3301 运行中 3302 停止中',
  `mon_date` datetime NOT NULL COMMENT '监控时间',
  `notice_type` varchar(12) NOT NULL COMMENT '通知渠道 1001 短信通知 2002 钉钉通知',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `remark` varchar(512) DEFAULT NULL COMMENT '备注',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  PRIMARY KEY (`mhg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of monitor_host_group
-- ----------------------------
INSERT INTO `monitor_host_group` VALUES ('d8a3f9d2-e1a5-4f14-a36e-0eaf4489a93e', '生产组主机监控', '0 */1 * * * ?', '3302', '2021-02-24 22:36:32', '2002', '2021-02-24 22:36:32', '0', '测试', '512c369e-0642-41e5-9ea6-5fe737984ae6');

-- ----------------------------
-- Table structure for monitor_host_log
-- ----------------------------
DROP TABLE IF EXISTS `monitor_host_log`;
CREATE TABLE `monitor_host_log` (
  `log_id` varchar(64) NOT NULL COMMENT '日志ID',
  `host_id` varchar(64) NOT NULL COMMENT '主机ID',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `cpu_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT 'cpu 使用率',
  `mem_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '内存使用率',
  `disk_rate` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '磁盘使用率',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of monitor_host_log
-- ----------------------------
INSERT INTO `monitor_host_log` VALUES ('0053007d-09d5-4f21-8e67-fe6da16f68a5', '55dfa429-5f96-44f9-ace8-88870b6c3741', '512c369e-0642-41e5-9ea6-5fe737984ae6', '4.70', '0.54', '0.10', '2021-03-06 15:52:04', '0');

-- ----------------------------
-- Table structure for privilege
-- ----------------------------
DROP TABLE IF EXISTS `privilege`;
CREATE TABLE `privilege` (
  `p_id` varchar(64) NOT NULL COMMENT '权限ID',
  `name` varchar(10) NOT NULL COMMENT '权限名称',
  `description` varchar(200) DEFAULT NULL COMMENT '权限描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `resource` varchar(200) DEFAULT NULL COMMENT '资源路径',
  `m_id` varchar(64) NOT NULL COMMENT '菜单ID',
  UNIQUE KEY `p_id` (`p_id`) USING BTREE,
  UNIQUE KEY `p_id_2` (`p_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `pg_id` varchar(64) NOT NULL COMMENT '权限组ID',
  `name` varchar(10) NOT NULL COMMENT '权限组名称',
  `description` varchar(200) DEFAULT NULL COMMENT '权限组描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `tenant_id` varchar(64) NOT NULL DEFAULT '9999' COMMENT '租户ID',
  UNIQUE KEY `pg_id` (`pg_id`) USING BTREE,
  UNIQUE KEY `pg_id_2` (`pg_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `rel_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限关系ID',
  `p_id` varchar(64) NOT NULL COMMENT '权限ID',
  `pg_id` varchar(64) NOT NULL COMMENT '权限组ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`rel_id`),
  UNIQUE KEY `rel_id` (`rel_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8;

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
  `pu_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '权限用户ID',
  `p_id` varchar(64) NOT NULL COMMENT '权限标志 是 1是权限组 0是权限',
  `privilege_flag` varchar(4) NOT NULL DEFAULT '0' COMMENT '权限标志 是 1是权限组 0是权限',
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  PRIMARY KEY (`pu_id`),
  UNIQUE KEY `pu_id` (`pu_id`) USING BTREE,
  KEY `index_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

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
  `sql_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'sqlID',
  `sql_code` varchar(128) NOT NULL COMMENT 'sql 编码',
  `sql_text` longtext NOT NULL COMMENT 'sql内容',
  `remark` varchar(256) DEFAULT NULL COMMENT 'sql 备注',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`sql_id`),
  UNIQUE KEY `sql_code` (`sql_code`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

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
  `task_id` varchar(64) NOT NULL COMMENT '任务ID',
  `task_name` varchar(100) NOT NULL COMMENT '任务名称',
  `template_id` varchar(64) NOT NULL COMMENT '模板ID',
  `task_cron` varchar(50) NOT NULL COMMENT '定时时间',
  `state` varchar(12) NOT NULL DEFAULT '001' COMMENT '状态 001 停止 002 运行 ',
  `status_cd` varchar(12) NOT NULL DEFAULT '0' COMMENT '数据有效状态 0 有效 1失效 ',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `host_id` varchar(64) NOT NULL COMMENT '主机ID',
  `notice_type` varchar(12) NOT NULL DEFAULT '2002' COMMENT '通知类型',
  PRIMARY KEY (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `task_id` varchar(64) NOT NULL COMMENT '任务ID',
  `attr_id` varchar(64) NOT NULL COMMENT '属性id',
  `spec_cd` varchar(12) NOT NULL COMMENT '规格id,参考task_template_spec表',
  `value` varchar(50) NOT NULL COMMENT '属性值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(12) NOT NULL DEFAULT '0' COMMENT '数据状态',
  PRIMARY KEY (`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `template_id` varchar(64) NOT NULL COMMENT '模板ID',
  `template_name` varchar(100) NOT NULL COMMENT '模板名称',
  `template_desc` varchar(200) NOT NULL COMMENT '描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(12) NOT NULL DEFAULT '0' COMMENT '状态 0 有效 1 失效',
  `class_bean` varchar(200) NOT NULL COMMENT '模板处理spring bean 类'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of task_template
-- ----------------------------
INSERT INTO `task_template` VALUES ('1', '进程监控', '进程监控', '2021-03-09 22:33:09', '0', 'CheckProcess');

-- ----------------------------
-- Table structure for task_template_spec
-- ----------------------------
DROP TABLE IF EXISTS `task_template_spec`;
CREATE TABLE `task_template_spec` (
  `spec_id` varchar(64) NOT NULL COMMENT '规格ID',
  `template_id` varchar(64) NOT NULL COMMENT '模板ID',
  `spec_cd` varchar(12) NOT NULL COMMENT '规格',
  `spec_name` varchar(100) NOT NULL COMMENT '规格名称',
  `spec_desc` varchar(200) NOT NULL COMMENT '规格描述',
  `is_show` varchar(12) NOT NULL DEFAULT 'T' COMMENT '页面是否展示，T就是展示 F 不展示',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(12) NOT NULL DEFAULT '0' COMMENT '状态 0 有效 1 失效'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of task_template_spec
-- ----------------------------
INSERT INTO `task_template_spec` VALUES ('1', '1', '100101', '进程名称', '必填，请填写进程名称', 'Y', '2021-03-09 22:34:38', '0');

-- ----------------------------
-- Table structure for tenant
-- ----------------------------
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant` (
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `tenant_name` varchar(128) NOT NULL COMMENT '租户名称',
  `tenant_type` varchar(12) NOT NULL DEFAULT '002' COMMENT '002 普通租户 001 管理员',
  `address` varchar(256) NOT NULL COMMENT '联系地址',
  `person_name` varchar(64) NOT NULL COMMENT '联系人',
  `phone` varchar(11) NOT NULL COMMENT '联系电话',
  `state` varchar(12) NOT NULL DEFAULT '2002' COMMENT '状态 1001 待审核 2002 审核完成 3003 拒绝审核',
  `remark` varchar(512) DEFAULT NULL COMMENT '租户说明',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `setting_id` varchar(64) NOT NULL COMMENT '设置ID',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `spec_cd` varchar(12) NOT NULL COMMENT '租户规格',
  `value` varchar(1024) NOT NULL COMMENT '属性值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据状态，详细参考c_status表，0在用，1失效',
  PRIMARY KEY (`setting_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status_cd` varchar(64) NOT NULL COMMENT '表字段状态 取值',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `table_name` varchar(50) DEFAULT NULL,
  `table_columns` varchar(50) DEFAULT NULL COMMENT '表字段说明',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

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
  `user_id` varchar(64) NOT NULL COMMENT '用户ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `real_name` varchar(64) NOT NULL COMMENT '真实名称',
  `passwd` varchar(128) NOT NULL COMMENT '密码',
  `sex` int(11) NOT NULL DEFAULT '1' COMMENT '1 男 2 女',
  `phone` varchar(11) NOT NULL COMMENT '手机号',
  `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
  `state` varchar(12) NOT NULL DEFAULT '100201' COMMENT '用户状态 100201 在用',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status_cd` varchar(2) NOT NULL DEFAULT '0' COMMENT '数据有效',
  `tenant_id` varchar(64) NOT NULL COMMENT '租户ID',
  `user_role` varchar(12) NOT NULL COMMENT '用户角色 1001 管理员 2002 普通用户',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_cd` varchar(2) NOT NULL DEFAULT '0',
  `tenant_id` varchar(64) NOT NULL
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
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time  timestamp   default CURRENT_TIMESTAMP not null,
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
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
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
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
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
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
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
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table business_images_ver
(
    id          varchar(64)  not null
        primary key,
    images_id   varchar(64)  not null,
    version     varchar(32)  not null,
    type_url    varchar(512) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);
insert into waf values (1,'waf',80,443,1001,'2020-01-01','0');
create table waf_hosts
(
    waf_host_id           varchar(64) not null
        primary key,
    waf_id varchar(64) not null,
    host_id         varchar(64) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_hostname_cert
(
    cert_id    varchar(64)  not null primary key,
    hostname    varchar(128) not null,
    cert_content      longtext  not null,
    priv_key_content     longtext  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
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
