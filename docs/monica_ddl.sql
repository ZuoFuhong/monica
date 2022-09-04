CREATE DATABASE IF NOT EXISTS `monica` CHARACTER SET = utf8mb4;

DROP TABLE IF EXISTS `t_service`;
CREATE TABLE `t_service` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL COMMENT '服务名称',
    `business` VARCHAR(100) NOT NULL COMMENT '业务名称',
    `owners` VARCHAR(100) NOT NULL COMMENT '负责人',
    `remark` VARCHAR(100) NOT NULL COMMENT '备注',
    `created_at` datetime NOT NULL,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_name`(`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '服务表';