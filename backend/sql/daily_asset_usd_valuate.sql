DROP TABLE IF EXISTS `daily_asset_usd_valuate`;
CREATE TABLE `daily_asset_usd_valuate` (
  `id` varchar(50) NOT NULL,
  `account_id` varchar(50) NOT NULL,
  `update_at` bigint(11) NOT NULL,
  `asset_usd` decimal(12,6) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

