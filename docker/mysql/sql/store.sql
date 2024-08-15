/*
 Navicat Premium Data Transfer

 Source Server         : roha24~25
 Source Server Type    : MySQL
 Source Server Version : 50739
 Source Host           : 8.135.237.23:19306
 Source Schema         : store

 Target Server Type    : MySQL
 Target Server Version : 50739
 File Encoding         : 65001

 Date: 15/08/2024 17:32:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for fund_expenses
-- ----------------------------
DROP TABLE IF EXISTS `fund_expenses`;
CREATE TABLE `fund_expenses`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `expenses_id` bigint(20) NOT NULL COMMENT '支出IID',
  `user_id` bigint(20) NULL DEFAULT 0 COMMENT '用户IID',
  `type` tinyint(1) NULL DEFAULT 0 COMMENT '支出类型：1-订单、2-红包',
  `before` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '扣除前',
  `after` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '扣除后',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_expenses`(`expenses_id`) USING BTREE,
  INDEX `idx_user`(`user_id`) USING BTREE,
  INDEX `idx_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '资金支出流水' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fund_expenses
-- ----------------------------

-- ----------------------------
-- Table structure for fund_income
-- ----------------------------
DROP TABLE IF EXISTS `fund_income`;
CREATE TABLE `fund_income`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `income_id` bigint(20) NOT NULL COMMENT '收入IID',
  `user_id` bigint(20) NULL DEFAULT 0 COMMENT '用户IID',
  `type` tinyint(1) NULL DEFAULT 0 COMMENT '收入类型：2-红包',
  `before` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '收入前',
  `after` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '收入后',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_income`(`income_id`) USING BTREE,
  INDEX `idx_user`(`user_id`) USING BTREE,
  INDEX `idx_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '资金收入流水' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of fund_income
-- ----------------------------

-- ----------------------------
-- Table structure for main_order
-- ----------------------------
DROP TABLE IF EXISTS `main_order`;
CREATE TABLE `main_order`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) NOT NULL COMMENT '订单编号',
  `status` tinyint(2) UNSIGNED NULL DEFAULT 1 COMMENT '订单状态：1-待支付、11-取消订单、12-失效订单、20-已完成',
  `pay_status` tinyint(2) UNSIGNED NULL DEFAULT 1 COMMENT '支付状态:1-待支付、2-取消支付、3-支付超时、11-支付失败、20-已支付',
  `pay_time` bigint(20) NULL DEFAULT 0 COMMENT '支付时间,毫秒',
  `pay_timeout` bigint(20) NULL DEFAULT 0 COMMENT '支付有效时间,毫秒',
  `pay_time_close` bigint(20) NULL DEFAULT 0 COMMENT '支付“取消|失效 ”时间,毫秒',
  `total` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '订单总价,入库*1000【1000 = 1元】',
  `quantity` int(8) UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品总数量',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '订单备注',
  `address_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '收货人姓名',
  `address_phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '收货人电话',
  `address_country` int(10) NULL DEFAULT 0 COMMENT '国家',
  `address_province` int(10) NULL DEFAULT 0 COMMENT '省',
  `address_city` int(10) NULL DEFAULT 0 COMMENT '市',
  `address_district` int(10) NULL DEFAULT 0 COMMENT '区',
  `address_detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '详细地址',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `order_id`(`order_id`) USING BTREE,
  INDEX `idx_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '订单列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of main_order
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pay_id` bigint(20) NOT NULL COMMENT '支付单编号IID',
  `order_id` bigint(20) NULL DEFAULT 0 COMMENT '订单编号IID',
  `type` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '支付来源：1-微信、2-支付宝、3-美团、4-第三方',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `pay_id`(`pay_id`) USING BTREE,
  INDEX `idx_order_type`(`order_id`, `type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '支付订单列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pay_order
-- ----------------------------

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` bigint(20) NOT NULL COMMENT '商品IID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '商品中文名称',
  `image` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '商品图片',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '状态:1-正在销售、2-新品、2-爆款、9-停止销售',
  `price` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '价钱',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_product`(`product_id`) USING BTREE,
  INDEX `idx_status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'sku商品表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------

-- ----------------------------
-- Table structure for purchase_order_detail
-- ----------------------------
DROP TABLE IF EXISTS `purchase_order_detail`;
CREATE TABLE `purchase_order_detail`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `detail_id` bigint(20) NOT NULL COMMENT '明细编号IID',
  `order_id` bigint(20) NULL DEFAULT 0 COMMENT '订单编号IID',
  `product_id` mediumint(8) NULL DEFAULT 0 COMMENT '商品IID',
  `total` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '商品总价入库*1000【1000 = 1元】',
  `quantity` int(8) UNSIGNED NULL DEFAULT 0 COMMENT '商品数量',
  `price` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '商品单价入库*1000【1000 = 1元】',
  `return_quantity` int(8) UNSIGNED NULL DEFAULT 0 COMMENT '退货数量',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '备注',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_detail_order`(`detail_id`, `order_id`) USING BTREE,
  INDEX `idx_product`(`product_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '订单商品明细表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of purchase_order_detail
-- ----------------------------

-- ----------------------------
-- Table structure for store
-- ----------------------------
DROP TABLE IF EXISTS `store`;
CREATE TABLE `store`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `store_id` bigint(20) NOT NULL COMMENT '店铺IID',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '店铺名',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_store`(`store_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of store
-- ----------------------------
INSERT INTO `store` VALUES (1, 1810940924055547904, 1, '国服', '2024-07-10 15:35:36.000', '2024-07-10 15:35:36.000');
INSERT INTO `store` VALUES (2, 1810941036622278656, 1, '欧服', '2024-07-10 15:36:03.000', '2024-07-10 15:36:03.000');
INSERT INTO `store` VALUES (3, 1810941555327660032, 1, '美服', '2024-07-10 15:38:07.000', '2024-07-10 15:38:07.000');

-- ----------------------------
-- Table structure for user_store
-- ----------------------------
DROP TABLE IF EXISTS `user_store`;
CREATE TABLE `user_store`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户IID',
  `store_id` bigint(20) NOT NULL COMMENT '店铺IID',
  UNIQUE INDEX `unx_user_store`(`user_id`, `store_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_store
-- ----------------------------
INSERT INTO `user_store` VALUES (1788408218839183360, 1810940924055547904);
INSERT INTO `user_store` VALUES (1788408218839183360, 1810941036622278656);
INSERT INTO `user_store` VALUES (1788408218839183360, 1810941555327660032);
INSERT INTO `user_store` VALUES (1788408218897903616, 1810940924055547904);
INSERT INTO `user_store` VALUES (1788408218897903616, 1810941036622278656);
INSERT INTO `user_store` VALUES (1788408218897903616, 1810941555327660032);
INSERT INTO `user_store` VALUES (1788408218960818176, 1810940924055547904);
INSERT INTO `user_store` VALUES (1788408218960818176, 1810941036622278656);
INSERT INTO `user_store` VALUES (1788408218960818176, 1810941555327660032);
INSERT INTO `user_store` VALUES (1788408219027927040, 1810940924055547904);
INSERT INTO `user_store` VALUES (1788408219027927040, 1810941036622278656);
INSERT INTO `user_store` VALUES (1788408219027927040, 1810941555327660032);
INSERT INTO `user_store` VALUES (1788408219090841600, 1810940924055547904);
INSERT INTO `user_store` VALUES (1788408219090841600, 1810941036622278656);
INSERT INTO `user_store` VALUES (1788408219090841600, 1810941555327660032);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户IID',
  `token` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT 'token',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '1=启用 0=禁用',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '昵称',
  `fund` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '用户资金,入库*1000【1000 = 1元】',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unx_user_token`(`user_id`, `token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 1788408218839183360, '2gDGQwDxsrX0UG8yRbophdHxHqD', 2, '蜻蜓队长(管理员)', 10000000, '2024-05-09 11:18:41.000', '2024-05-09 11:18:41.000');
INSERT INTO `users` VALUES (2, 1788408218897903616, '2gDGQugkyFF4MI10hK7WfT3W3Pe', 2, '和平星(管理员)', 10000000, '2024-05-09 11:18:41.000', '2024-05-09 11:18:41.000');
INSERT INTO `users` VALUES (3, 1788408218960818176, '2gDGQvEugR6Y5riFp2kVLdc7J0O', 1, '蜘蛛侦探', 10000000, '2024-05-09 11:18:41.000', '2024-05-09 11:18:41.000');
INSERT INTO `users` VALUES (4, 1788408219027927040, '2gDGQwhqJQczjkCikEvg3StOKSR', 1, '蝎子莱莱', 10000000, '2024-05-09 11:18:41.000', '2024-05-09 11:18:41.000');
INSERT INTO `users` VALUES (5, 1788408219090841600, '2gDGQvpg5xTE3Qn0SIzbyDXpdma', 1, '卡布达', 10000000, '2024-05-09 11:18:41.000', '2024-05-09 11:18:41.000');
INSERT INTO `users` VALUES (6, 1821087524543295488, '2kJyfAuHY2JwTFMMyPwGFi61YkQ', 1, '金龟次郎', 10000000, '2024-08-07 15:34:34.000', '2024-08-07 15:34:34.000');
INSERT INTO `users` VALUES (7, 1821087524656541696, '2kJyf7WxQ8N1US5nyJjbStxLtf8', 1, '丸子龙', 10000000, '2024-08-07 15:34:34.000', '2024-08-07 15:34:34.000');
INSERT INTO `users` VALUES (8, 1821087524765593600, '2kJyfA138eU7XLmWDetpSEqgCvZ', 1, '车轮滚滚', 10000000, '2024-08-07 15:34:34.000', '2024-08-07 15:34:34.000');
INSERT INTO `users` VALUES (9, 1821087524857868288, '2kJyf7EbM6FQWCyrkM36JDYTShv', 1, '蟑螂恶霸', 10000000, '2024-08-07 15:34:34.000', '2024-08-07 15:34:34.000');
INSERT INTO `users` VALUES (10, 1821087524929171456, '2kJyf703zSyhdKFBa2b9l2bjxWX', 1, '鲨鱼辣椒', 10000000, '2024-08-07 15:34:34.000', '2024-08-07 15:34:34.000');
INSERT INTO `users` VALUES (11, 1821857196679131136, '2kPybdu8GObZm5SVwHF1TLUrCE9', 2, '压测官', 10000000, '2024-08-09 18:32:58.000', '2024-08-09 18:32:58.000');

SET FOREIGN_KEY_CHECKS = 1;
