/*
 Navicat Premium Data Transfer

 Source Server         : 字节青训营
 Source Server Type    : MySQL
 Source Server Version : 50741
 Source Host           : 43.138.127.168:3306
 Source Schema         : douyin-easy

 Target Server Type    : MySQL
 Target Server Version : 50741
 File Encoding         : 65001

 Date: 28/01/2023 09:37:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_comment
-- ----------------------------
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `father_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '评论父id',
  `to_user_id` bigint(20) NULL DEFAULT NULL COMMENT '回复谁',
  `video_id` bigint(20) UNSIGNED NOT NULL COMMENT '视频id',
  `from_user_id` bigint(20) NOT NULL COMMENT '评论者',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `video_id`(`video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '视频评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_comment
-- ----------------------------

-- ----------------------------
-- Table structure for tb_file
-- ----------------------------
DROP TABLE IF EXISTS `tb_file`;
CREATE TABLE `tb_file`  (
  `id` bigint(20) NOT NULL COMMENT '文件id',
  `tag` tinyint(1) NOT NULL COMMENT '文件存储类型 0 路径 1 网络地址',
  `suffix` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件后缀',
  `prefix` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件前缀',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件地址',
  `size` bigint(20) NOT NULL COMMENT '文件大小，字节',
  `auther` bigint(20) NOT NULL COMMENT '上传者',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_file
-- ----------------------------

-- ----------------------------
-- Table structure for tb_follow
-- ----------------------------
DROP TABLE IF EXISTS `tb_follow`;
CREATE TABLE `tb_follow`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '关注id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `follow_id` bigint(20) NOT NULL COMMENT '关注对象id',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_fan`(`user_id`, `follow_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '关注表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_follow
-- ----------------------------

-- ----------------------------
-- Table structure for tb_like_video
-- ----------------------------
DROP TABLE IF EXISTS `tb_like_video`;
CREATE TABLE `tb_like_video`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '喜欢视频id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `video_id` bigint(20) NOT NULL COMMENT '视频id',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_like_video`(`user_id`, `video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '喜欢视频表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_like_video
-- ----------------------------

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` bigint(20) NOT NULL COMMENT '用户id',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `fans_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数目',
  `follow_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注数目',
  `recive_like_count` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '收到的赞数目',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_id`(`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------

-- ----------------------------
-- Table structure for tb_video
-- ----------------------------
DROP TABLE IF EXISTS `tb_video`;
CREATE TABLE `tb_video`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `auther_id` bigint(20) NOT NULL COMMENT '作者id',
  `play_id` bigint(20) NOT NULL COMMENT '视频文件id',
  `cover_id` bigint(20) NOT NULL COMMENT '视频封面id',
  `favorite_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频喜欢数目',
  `comment_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频评论数目',
  `create_at` datetime(0) NOT NULL COMMENT '创建时间',
  `update_at` datetime(0) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `video`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '视频表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_video
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
