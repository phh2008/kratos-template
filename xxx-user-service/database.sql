-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) NULL DEFAULT NULL,
  `v0` varchar(100) NULL DEFAULT NULL,
  `v1` varchar(100) NULL DEFAULT NULL,
  `v2` varchar(100) NULL DEFAULT NULL,
  `v3` varchar(100) NULL DEFAULT NULL,
  `v4` varchar(100) NULL DEFAULT NULL,
  `v5` varchar(100) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC) USING BTREE
) ENGINE = InnoDB;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (1, 'p', 'finance', '/api.helloworld.v1.User', 'DeleteById', '', '', '');
INSERT INTO `casbin_rule` VALUES (2, 'g', 'test', 'finance', '', '', '', '');

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `perm_name` varchar(50) NOT NULL DEFAULT '' COMMENT '权限名称',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT 'URL路径',
  `action` varchar(50) NOT NULL DEFAULT '' COMMENT '权限动作：比如get、post、delete等',
  `perm_type` tinyint NOT NULL DEFAULT 1 COMMENT '权限类型：1-菜单、2-按钮',
  `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父级ID：资源层级关系',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '权限表';

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES (1, '测试', '/api.helloworld.v1.User', 'DeleteById', 1, 0, '2025-06-12 15:17:42', '2023-06-28 16:08:18', 2, 2, 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_code` varchar(50) NOT NULL DEFAULT '' COMMENT '角色编号',
  `role_name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 'finance', '财务', '2025-06-12 15:07:25', '2025-06-12 15:07:25', 1, 1, 1);

-- ----------------------------
-- Table structure for sys_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` bigint NOT NULL DEFAULT 0 COMMENT '角色编号',
  `perm_id` bigint NOT NULL DEFAULT 0 COMMENT '权限ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '角色权限表';

-- ----------------------------
-- Records of sys_role_permission
-- ----------------------------
INSERT INTO `sys_role_permission` VALUES (1, 1, 1);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `real_name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `user_name` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `password` varchar(200) NOT NULL DEFAULT '' COMMENT '密码',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 1-启用，2-禁用',
  `role_code` varchar(50) NOT NULL DEFAULT '' COMMENT '角色编号',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_by` varchar(50) DEFAULT NULL COMMENT '创建人',
  `updated_by` varchar(50) DEFAULT NULL COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1000, 'test', 'test', 'test@test.com', '$2a$10$ITxtKZMlLHEqVQU7x5C62OGyDPiduBNGxKBEZRRJ/jkJnFG2.TSi.', 1, 'finance', '2025-06-12 00:00:00', '2025-06-12 00:00:00', 0, 0, 1);

CREATE TABLE `sys_access_key`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `access_id`  varchar(100) NOT NULL COMMENT 'access id',
    `access_key` varchar(100) DEFAULT NULL COMMENT 'access key',
    `remark`     varchar(500) DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_access_id` (`access_id`)
) ENGINE=InnoDB COMMENT='访问密钥';

CREATE TABLE `sys_file` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `file_path` varchar(500) NOT NULL DEFAULT '' COMMENT '文件路径(保存后路径)',
    `file_size` bigint unsigned DEFAULT '0' COMMENT '文件大小(字节)',
    `file_md5` varchar(100) DEFAULT NULL COMMENT '文件md5',
    `media_type` varchar(200) DEFAULT '' COMMENT '媒介类型：image/jpeg 等',
    `original_name` varchar(200) DEFAULT '' COMMENT '原始文件名',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `created_by` varchar(50) DEFAULT NULL COMMENT '创建人',
    `updated_by` varchar(50) DEFAULT NULL COMMENT '更新人',
    `deleted` tinyint unsigned DEFAULT '1' COMMENT '是否删除：1-否，2-是',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_file_path` (`file_path`) USING BTREE,
    KEY `idx_file_md5` (`file_md5`) USING BTREE
) ENGINE=InnoDB COMMENT='文件信息表';
