DROP TABLE IF EXISTS `tb_order`;

-- 创建注单表
CREATE TABLE `tb_order` (
  `id` bigint NOT NULL COMMENT '订单ID',
  `created_at` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_at` bigint DEFAULT '0' COMMENT '更新时间-订单取消的时候用',
  `merchant_id` bigint NOT NULL DEFAULT '0' COMMENT '商户id，自营为1',
  `merchant_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '商户名称，自营为bingo',
  `merchant_user_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '商户用户id',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `order_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '订单号',
  `issue_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '期号',
  `game_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '游戏名称',
  `game_code` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '游戏code',
  `play_code` int NOT NULL COMMENT '玩法code',
  `multiples` int DEFAULT '0' COMMENT '倍数',
  `amount` bigint NOT NULL COMMENT '投注金额',
  `status` tinyint DEFAULT '1' COMMENT '订单状态: 1未开奖 2已开奖 3未中奖 4撤单 5等待二次开奖 6二次结算',
  `order_type` tinyint DEFAULT '0' COMMENT '订单类型 1普通下注 2自动投注',
  `ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '下单IP',
  `bet_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '投注内容',
  `second_bet_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '二次投注内容',
  `jackpot_ratio` bigint DEFAULT '0' COMMENT 'Jackpot奖池注水比例',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
  `device_type` int NOT NULL COMMENT '设备类型',
  `currency` varchar(10) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '币种',
  `bet_time` bigint NOT NULL DEFAULT '0' COMMENT '投注时间',
  `game_type` int DEFAULT '0' COMMENT '游戏类型',
  `site_id` int NOT NULL DEFAULT '0',
  `free_status` tinyint NOT NULL DEFAULT '2' COMMENT '状态：1 - 免费，2 - 自费',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_order_number` (`order_number`(25)) USING BTREE,
  KEY `index_merchant_user_id` (`merchant_user_id`(20)) USING BTREE,
  KEY `index_issue_number` (`issue_number`(20)) USING BTREE,
  KEY `idx_created_at` (`created_at`),
  KEY `idx_game_type_created_at` (`game_type`,`created_at`),
  KEY `idx_user_id_order_number` (`user_id`,`issue_number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC COMMENT='注单表';