SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `sender_id` bigint NOT NULL DEFAULT 0,
  `receiver_id` bigint NOT NULL DEFAULT 0 COMMENT '私聊功能时使用,其他时刻为0',
  `channel_id` bigint NOT NULL DEFAULT 0 COMMENT '私聊功能时为0',
  `type` VARCHAR(20) NOT NULL DEFAULT 'text' COMMENT 'text、image、file,为image、file时content为存储路径',
  `content` TEXT NOT NULL,
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX idx_channel_id (channel_id),
  INDEX idx_receiver_id (receiver_id),
  INDEX idx_sender_id (sender_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
