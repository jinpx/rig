DROP TABLE IF EXISTS `tb_order_result`;

-- 创建结果表
CREATE TABLE `tb_order_result` (
  `id` bigint NOT NULL COMMENT '结果ID',
  `order_id` bigint NOT NULL COMMENT '对应的订单ID',
  `prize_at` bigint DEFAULT '0' COMMENT '派彩时间',
  `award_amount` bigint NOT NULL DEFAULT '0' COMMENT '中奖金额',
  `prize_amount` bigint DEFAULT '0' COMMENT '派彩金额',
  `win_lose_amount` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '订单盈亏情况',
  `prize_type` int NOT NULL DEFAULT '0' COMMENT '中奖类型',
  `settle_time` bigint NOT NULL COMMENT '结算时间',
  `prize_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '第一次开奖结果',
  `sec_prize_content` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '二次开奖结果',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `index_order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_bin ROW_FORMAT=DYNAMIC COMMENT='游戏结果表';
