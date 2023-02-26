/*
Navicat MySQL Data Transfer

Source Server         : defDB
Source Server Version : 50741
Source Host           : 175.178.26.250:3307
Source Database       : dousheng_db

Target Server Type    : MYSQL
Target Server Version : 50741
File Encoding         : 65001

Date: 2023-02-26 20:31:11
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `tb_comment`
-- ----------------------------
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `video_id` bigint(20) DEFAULT '0' COMMENT '视频id',
  `content` varchar(255) DEFAULT '' COMMENT '评论',
  `is_deleted` tinyint(4) DEFAULT '0' COMMENT '删除评论',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='评论表';

-- ----------------------------
-- Records of tb_comment
-- ----------------------------
INSERT INTO `tb_comment` VALUES ('1', '3', '4', '666', '1', '2023-02-21 16:37:57');
INSERT INTO `tb_comment` VALUES ('2', '1', '2', '哈哈哈', '1', '2023-02-22 09:42:49');
INSERT INTO `tb_comment` VALUES ('3', '1', '4', '111', '1', '2023-02-22 15:52:17');
INSERT INTO `tb_comment` VALUES ('4', '2', '2', 'hello', '1', '2023-02-22 16:31:09');
INSERT INTO `tb_comment` VALUES ('5', '1', '6', '**', '1', '2023-02-23 12:40:31');
INSERT INTO `tb_comment` VALUES ('6', '1', '6', '哈哈哈', '1', '2023-02-23 12:40:39');
INSERT INTO `tb_comment` VALUES ('7', '1', '6', '哈哈哈', '1', '2023-02-23 12:40:40');
INSERT INTO `tb_comment` VALUES ('8', '1', '6', '死', '1', '2023-02-23 12:40:48');
INSERT INTO `tb_comment` VALUES ('9', '1', '6', '**', '1', '2023-02-23 12:40:54');
INSERT INTO `tb_comment` VALUES ('10', '1', '1', 'hello word', '1', '2023-02-23 22:09:25');

-- ----------------------------
-- Table structure for `tb_favorite`
-- ----------------------------
DROP TABLE IF EXISTS `tb_favorite`;
CREATE TABLE `tb_favorite` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '点赞id',
  `username` varchar(40) DEFAULT '' COMMENT '用户名',
  `user_id` bigint(20) DEFAULT '0' COMMENT '用户id',
  `video_id` bigint(20) DEFAULT '0' COMMENT '视频id',
  `is_deleted` tinyint(4) DEFAULT '0' COMMENT '取消点赞',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='点赞表';

-- ----------------------------
-- Records of tb_favorite
-- ----------------------------

-- ----------------------------
-- Table structure for `tb_message`
-- ----------------------------
DROP TABLE IF EXISTS `tb_message`;
CREATE TABLE `tb_message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '消息id',
  `to_user_id` bigint(20) DEFAULT '0' COMMENT '接收方id',
  `from_user_id` bigint(20) DEFAULT '0' COMMENT '发送方id',
  `content` varchar(255) DEFAULT '' COMMENT '消息内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='消息表';

-- ----------------------------
-- Records of tb_message
-- ----------------------------
INSERT INTO `tb_message` VALUES ('1', '3', '1', '哈哈', '2023-02-22 16:18:41');
INSERT INTO `tb_message` VALUES ('2', '3', '1', '哈哈', '2023-02-22 16:22:00');
INSERT INTO `tb_message` VALUES ('3', '3', '2', '你好', '2023-02-22 16:26:55');
INSERT INTO `tb_message` VALUES ('4', '1', '3', '笑死', '2023-02-22 17:35:22');
INSERT INTO `tb_message` VALUES ('5', '1', '3', '谁是谁是谁', '2023-02-22 17:35:24');
INSERT INTO `tb_message` VALUES ('6', '1', '3', 'SSD', '2023-02-22 17:35:27');
INSERT INTO `tb_message` VALUES ('7', '1', '3', '四年都还行', '2023-02-22 17:35:32');
INSERT INTO `tb_message` VALUES ('8', '1', '3', '啊哈还有', '2023-02-22 17:35:37');
INSERT INTO `tb_message` VALUES ('9', '1', '3', '就是傻子', '2023-02-22 17:35:39');
INSERT INTO `tb_message` VALUES ('10', '3', '1', '六六六', '2023-02-22 17:36:23');
INSERT INTO `tb_message` VALUES ('11', '1', '3', 'dd ', '2023-02-22 19:22:05');
INSERT INTO `tb_message` VALUES ('12', '3', '1', '哈哈哈', '2023-02-22 19:22:05');
INSERT INTO `tb_message` VALUES ('13', '1', '3', 'fd ', '2023-02-22 19:22:16');
INSERT INTO `tb_message` VALUES ('14', '3', '1', '姐姐', '2023-02-22 19:22:20');
INSERT INTO `tb_message` VALUES ('15', '1', '3', 'gg ', '2023-02-22 19:22:47');
INSERT INTO `tb_message` VALUES ('16', '3', '1', '刚刚', '2023-02-22 19:23:05');
INSERT INTO `tb_message` VALUES ('17', '1', '3', 'hh', '2023-02-22 19:31:10');
INSERT INTO `tb_message` VALUES ('18', '1', '3', 'jj ?', '2023-02-22 19:32:53');
INSERT INTO `tb_message` VALUES ('19', '3', '1', '行行行', '2023-02-22 19:33:00');
INSERT INTO `tb_message` VALUES ('20', '3', '1', 'df', '2023-02-22 19:33:06');
INSERT INTO `tb_message` VALUES ('21', '1', '3', 'yy', '2023-02-22 19:33:08');
INSERT INTO `tb_message` VALUES ('22', '1', '3', '78', '2023-02-22 19:35:14');
INSERT INTO `tb_message` VALUES ('23', '3', '1', '等等', '2023-02-22 19:39:24');
INSERT INTO `tb_message` VALUES ('24', '3', '1', 'ed', '2023-02-22 19:39:31');
INSERT INTO `tb_message` VALUES ('25', '3', '1', '二等', '2023-02-22 19:39:33');
INSERT INTO `tb_message` VALUES ('26', '1', '3', 'hh', '2023-02-22 19:45:17');
INSERT INTO `tb_message` VALUES ('27', '1', '3', 'hu', '2023-02-22 19:45:34');
INSERT INTO `tb_message` VALUES ('28', '1', '3', 'yy', '2023-02-22 19:45:37');
INSERT INTO `tb_message` VALUES ('29', '3', '1', '信息', '2023-02-22 19:45:41');
INSERT INTO `tb_message` VALUES ('30', '3', '1', '等等', '2023-02-22 19:45:45');
INSERT INTO `tb_message` VALUES ('31', '1', '3', 'uuuuuuuiii', '2023-02-22 19:45:46');
INSERT INTO `tb_message` VALUES ('32', '3', '1', 'ddd', '2023-02-22 19:45:52');
INSERT INTO `tb_message` VALUES ('33', '1', '3', 'uuu', '2023-02-22 19:45:52');
INSERT INTO `tb_message` VALUES ('34', '1', '3', 'nnheushdu', '2023-02-22 19:49:33');
INSERT INTO `tb_message` VALUES ('35', '1', '3', '4', '2023-02-22 19:54:25');
INSERT INTO `tb_message` VALUES ('36', '3', '1', 'dd', '2023-02-22 20:23:49');
INSERT INTO `tb_message` VALUES ('37', '3', '1', '等等', '2023-02-22 20:23:54');
INSERT INTO `tb_message` VALUES ('38', '3', '1', '傻逼', '2023-02-23 12:49:59');
INSERT INTO `tb_message` VALUES ('39', '2', '1', 'hello world', '2023-02-23 22:24:39');

-- ----------------------------
-- Table structure for `tb_relation`
-- ----------------------------
DROP TABLE IF EXISTS `tb_relation`;
CREATE TABLE `tb_relation` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '关注id',
  `follower_id` bigint(20) DEFAULT '0' COMMENT '粉丝id',
  `following_id` bigint(20) DEFAULT '0' COMMENT '博主id',
  `isdeleted` tinyint(4) DEFAULT '0' COMMENT '取消关注',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='关注表';

-- ----------------------------
-- Records of tb_relation
-- ----------------------------

-- ----------------------------
-- Table structure for `tb_user`
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(40) DEFAULT '' COMMENT '用户名',
  `follow_count` int(11) DEFAULT '0' COMMENT '关注总数',
  `follower_count` int(11) DEFAULT '0' COMMENT '粉丝总数',
  `background_image` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '用户个人页顶部大图URL',
  `signature` varchar(255) DEFAULT '' COMMENT '个人简介',
  `total_favorited` int(11) DEFAULT '0' COMMENT '获赞数量',
  `work_count` int(11) DEFAULT '0' COMMENT '作品数量',
  `favorite_count` int(11) DEFAULT '0' COMMENT '点赞数量',
  `password` char(40) DEFAULT '' COMMENT '密码',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES ('2', 'demouser', '0', '0', '', '', '0', '5', '0', 'e10adc3949ba59abbe56e057f20f883e', '2023-02-21 11:42:45');

-- ----------------------------
-- Table structure for `tb_video`
-- ----------------------------
DROP TABLE IF EXISTS `tb_video`;
CREATE TABLE `tb_video` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户作者id',
  `play_url` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '视频URL',
  `cover_url` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '封面URL',
  `favorite_count` int(11) DEFAULT '0' COMMENT '点赞总数',
  `comment_count` int(11) DEFAULT '0' COMMENT '评论总数',
  `title` varchar(255) DEFAULT NULL COMMENT '视频标题',
  `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='视频表';

-- ----------------------------
-- Records of tb_video
-- ----------------------------
INSERT INTO `tb_video` VALUES ('1', '2', 'http://rq9lt9dry.bkt.clouddn.com/a%20happy%20day%21.mp4?e=1708487128&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:mpRB6amz7jKyRMYrtdn6w-mBiYw=', '', '1', '1', 'a happy day!', '2023-02-21 11:45:29', '2023-02-23 14:09:25');
INSERT INTO `tb_video` VALUES ('2', '2', 'http://rq9lt9dry.bkt.clouddn.com/trivial.mp4?e=1708487300&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:Tg1TpXydcWIO-iOg4S2InGWxQHw=', '', '1', '1', 'trivial', '2023-02-21 11:48:21', '2023-02-22 14:16:13');
INSERT INTO `tb_video` VALUES ('3', '2', 'http://rq9lt9dry.bkt.clouddn.com/%E9%87%91%E9%93%B2%E9%93%B2%E4%B9%8B%E6%88%98.mp4?e=1708487305&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:C20AB3uP1Jy7lTwg0xUDXX4acY4=', '', '1', '0', '金铲铲之战', '2023-02-21 11:48:25', '2023-02-23 07:58:40');
INSERT INTO `tb_video` VALUES ('4', '2', 'http://rq9lt9dry.bkt.clouddn.com/dragon.mp4?e=1708487334&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:dvwibDl9TqbCP-uRjcB5nsWfr4Y=', '', '6', '1', 'dragon', '2023-02-21 11:48:54', '2023-02-23 07:58:42');
INSERT INTO `tb_video` VALUES ('13', '2', 'http://rq9lt9dry.bkt.clouddn.com/%E6%98%9F%E6%9C%9F%E5%9B%9B.mp4?e=1708700541&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:V1huu-E-gK_4rBj5e2PglANdDaU=', 'http://rq9lfs4ld.bkt.clouddn.com/%E6%98%9F%E6%9C%9F%E5%9B%9Bcover.jpeg?e=1708700542&token=XuigBGSCJ7vpAtRtpu04NqLGLXpEROCaqgOxTZ0W:MKhzieJTG_1BV0O7hLLoa6l4BD4=', '0', '0', '星期四', '2023-02-23 23:02:23', '2023-02-23 23:02:23');
