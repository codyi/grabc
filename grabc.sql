gCREATE TABLE `rabc_assignment_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rabc_assignment_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '授权时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rabc_assignment_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `route_id` int(11) unsigned NOT NULL COMMENT '路由id',
  `permission_id` int(11) unsigned NOT NULL COMMENT '权限ID',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rabc_permission` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(50) NOT NULL COMMENT '权限名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rabc_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(100) NOT NULL COMMENT '角色名称',
  `create_at` int(11) NOT NULL COMMENT '创建时间',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `rabc_route` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `route` varchar(100) NOT NULL COMMENT '路由名称',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;