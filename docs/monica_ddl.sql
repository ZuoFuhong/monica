CREATE DATABASE IF NOT EXISTS `monica` CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `t_service`;
CREATE TABLE `t_service` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL COMMENT '服务名称',
    `business` VARCHAR(100) NOT NULL COMMENT '业务名称',
    `owners` VARCHAR(100) NOT NULL COMMENT '负责人',
    `remark` VARCHAR(100) NOT NULL COMMENT '备注',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_name`(`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '服务表';

DROP TABLE IF EXISTS `t_service_namespace`;
CREATE TABLE `t_service_namespace` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `service_name` VARCHAR(100) NOT NULL COMMENT '服务名称',
    `namespace` varchar(30) NOT NULL COMMENT '命名空间',
    `token` varchar(100) NOT NULL COMMENT '鉴权凭证',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '服务命名空间';

DROP TABLE IF EXISTS `t_service_instance`;
CREATE TABLE `t_service_instance` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `namespace` varchar(30) NOT NULL COMMENT '命名空间',
    `service_name` varchar(100) NOT NULL COMMENT '服务名称',
    `healthy` tinyint(1) NOT NULL DEFAULT '0' COMMENT '健康状态：1-异常 2-健康',
    `isolate` tinyint(1) NOT NULL DEFAULT '0' COMMENT '隔离状态：1-不隔离 2-隔离',
    `ip` VARCHAR(30) NOT NULL COMMENT '实例IP',
    `port` int(11) NOT NULL COMMENT '端口',
    `weight` tinyint(4) NOT NULL DEFAULT '0' COMMENT '权重',
    `metadata` varchar(1000) NOT NULL COMMENT '元数据',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_namespace_service`(`namespace`, `service_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '服务实例表';