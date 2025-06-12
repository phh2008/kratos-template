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
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
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
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
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
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_by` bigint NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint NOT NULL DEFAULT 0 COMMENT '更新人',
  `deleted` tinyint NOT NULL DEFAULT 1 COMMENT '是否删除：1-否，2-是',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1000, 'test', 'test', 'test@test.com', '$2a$10$ITxtKZMlLHEqVQU7x5C62OGyDPiduBNGxKBEZRRJ/jkJnFG2.TSi.', 1, 'finance', '2025-06-12 00:00:00', '2025-06-12 00:00:00', 0, 0, 1);
