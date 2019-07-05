/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : gosports

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 05/07/2019 17:11:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_match
-- ----------------------------
DROP TABLE IF EXISTS `t_match`;
CREATE TABLE `t_match` (
  `id` int(20) NOT NULL COMMENT '比赛id',
  `league_id` int(20) NOT NULL COMMENT '联赛id',
  `league_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '联赛名称',
  `match_status` int(10) NOT NULL DEFAULT '0' COMMENT '比赛状态:0.未开赛 1.进行中 100.已结束',
  `match_date` datetime NOT NULL COMMENT '比赛时间',
  `home_team_id` int(20) NOT NULL COMMENT '主队id',
  `home_team_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '主队队名',
  `away_team_id` int(20) NOT NULL COMMENT '客队id',
  `away_team_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '客队队名',
  `half_time_home_score` int(10) NOT NULL DEFAULT '0' COMMENT '主队半场得分',
  `half_time_away_score` int(10) NOT NULL DEFAULT '0' COMMENT '客队半场得分',
  `home_score` int(10) NOT NULL DEFAULT '0' COMMENT '主队全场得分',
  `away_score` int(10) NOT NULL DEFAULT '0' COMMENT '客队全场得分',
  `round` int(10) NOT NULL DEFAULT '0' COMMENT '轮次',
  `match_result` int(10) NOT NULL DEFAULT '0' COMMENT '0.无结果 1.主队胜 2.打平 3.客队胜',
  `created_time` datetime NOT NULL COMMENT '创建时间',
  `updated_time` datetime NOT NULL COMMENT '更新时间',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`),
  KEY `t_match_index_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
