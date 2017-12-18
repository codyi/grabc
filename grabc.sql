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

  CREATE TABLE `rabc_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(128) NOT NULL COMMENT '菜单名称',
  `parent` int(11) DEFAULT NULL COMMENT '父级菜单ID',
  `route` varchar(255) DEFAULT NULL COMMENT '菜单地址',
  `order` int(11) DEFAULT NULL COMMENT '菜单排序',
  `create_at` int(11) unsigned NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `parent` (`parent`),
  CONSTRAINT `ac_menu_ibfk_1` FOREIGN KEY (`parent`) REFERENCES `ac_menu` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8