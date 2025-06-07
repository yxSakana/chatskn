SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `guild`;
CREATE TABLE `guild` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `owner_id` bigint NOT NULL,
  `name` VARCHAR(50) NOT NULL DEFAULT '',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `guildMember`;
CREATE TABLE `guildMember` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `guild_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `role` INT8 NOT NULL DEFAULT 0 COMMENT '1: member普通成员;2: 管理员;3: owner拥有者;4: admin系统管理员',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `channel`;
CREATE TABLE `channel` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `guild_id` bigint NOT NULL,
  `name` VARCHAR(50) NOT NULL DEFAULT '',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `channelMember`;
CREATE TABLE `channelMember` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `channel_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `role` INT8 NOT NULL DEFAULT 0 COMMENT '1: member普通成员;2: 管理员;3: owner拥有者;4: admin系统管理员',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
