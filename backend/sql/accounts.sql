DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` varchar(50) NOT NULL,
  `market` varchar(20) NOT NULL,
  `access_key` varchar(100) NOT NULL,
  `secret_key` varchar(1000) NOT NULL,
  `status` smallint(2) NOT NULL COMMENT '0已删除 1 不统计 2统计中',
  `update_at` bigint(20) NOT NULL,
  `pass_phrase` varchar(100) DEFAULT NULL,
  `label` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

