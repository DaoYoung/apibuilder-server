-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.13-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.5.0.5196
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 apibuilder 的数据库结构
CREATE DATABASE IF NOT EXISTS `apibuilder` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `apibuilder`;

-- 导出  表 apibuilder.apis 结构
CREATE TABLE IF NOT EXISTS `apis` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL,
  `module_id` int(10) NOT NULL,
  `author_id` int(10) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `request_url` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `request_method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'get',
  `request_param` json DEFAULT NULL,
  `request_header` json DEFAULT NULL,
  `response_content` json DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.apis 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `apis` DISABLE KEYS */;
INSERT INTO `apis` (`id`, `task_id`, `module_id`, `author_id`, `title`, `request_url`, `request_method`, `request_param`, `request_header`, `response_content`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 1, '2018-07-29 18:10:30', '2018-08-02 13:03:11', '2018-08-02 21:03:11'),
	(2, 1, 1, 1, 'sss', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"name": "json-dft", "Content-Type": "application/json-dft"}', 1, '2018-07-31 13:38:27', '2018-08-02 21:25:34', NULL),
	(3, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-07-31 15:04:35', '2018-07-31 15:04:35', NULL),
	(4, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-07-31 20:43:49', '2018-07-31 20:43:49', NULL),
	(5, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-01 18:26:30', '2018-08-01 18:26:30', NULL),
	(6, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-02 10:41:48', '2018-08-02 10:41:48', NULL),
	(7, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-02 21:03:01', '2018-08-02 21:03:01', NULL);
/*!40000 ALTER TABLE `apis` ENABLE KEYS */;

-- 导出  表 apibuilder.api_commits 结构
CREATE TABLE IF NOT EXISTS `api_commits` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `api_id` int(10) NOT NULL DEFAULT '0',
  `author_id` int(10) NOT NULL DEFAULT '0',
  `task_id` int(10) NOT NULL DEFAULT '0',
  `changes` json NOT NULL,
  `commit_message` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.api_commits 的数据：~22 rows (大约)
/*!40000 ALTER TABLE `api_commits` DISABLE KEYS */;
INSERT INTO `api_commits` (`id`, `api_id`, `author_id`, `task_id`, `changes`, `commit_message`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 0, 0, '{"Title": {"after": "bbb", "before": "aaa"}, "ResponseContent": {"change_json": null}}', '', '2018-07-30 17:22:24', '2018-07-30 17:22:24', NULL),
	(2, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 17:26:10', '2018-07-30 17:26:10', NULL),
	(3, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 17:27:14', '2018-07-30 17:27:14', NULL),
	(4, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 17:33:36', '2018-07-30 17:33:36', NULL),
	(5, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "AuthorId": {"after": 1, "before": 1}}', '', '2018-07-30 17:34:17', '2018-07-30 17:34:17', NULL),
	(6, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "AuthorId": {"after": 1, "before": 1}}', '', '2018-07-30 17:34:54', '2018-07-30 17:34:54', NULL),
	(7, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "AuthorId": {"after": 1, "before": 1}}', '', '2018-07-30 17:35:14', '2018-07-30 17:35:14', NULL),
	(8, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "TaskId": {"after": 1, "before": 1}, "AuthorId": {"after": 1, "before": 1}}', '', '2018-07-30 17:35:50', '2018-07-30 17:35:50', NULL),
	(9, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "CommitTaskId": {"after": 2, "before": 0}, "CommitMessage": {"after": "修改bug", "before": ""}, "CommitAuthorId": {"after": 2, "before": 0}}', '', '2018-07-30 18:18:14', '2018-07-30 18:18:14', NULL),
	(10, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 18:20:25', '2018-07-30 18:20:25', NULL),
	(11, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 18:21:10', '2018-07-30 18:21:10', NULL),
	(12, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 18:21:37', '2018-07-30 18:21:37', NULL),
	(13, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}}', '', '2018-07-30 18:22:29', '2018-07-30 18:22:29', NULL),
	(14, 1, 0, 0, '{"Title": {"after": "bbb", "before": "bbb"}, "CommitJson": {"name": {"after": "json", "before": "json-dft"}}, "ResponseContent": {"name": {"after": "json", "before": "json-dft"}}}', '', '2018-07-30 18:32:12', '2018-07-30 18:32:12', NULL),
	(15, 1, 2, 2, '{"Title": {"after": "bbb", "before": "bbb"}, "CommitJson": {"name": {"after": "json", "before": "json-dft"}}, "ResponseContent": {"name": {"after": "json", "before": "json-dft"}}}', '修改bug', '2018-07-30 18:35:55', '2018-07-30 18:35:55', NULL),
	(16, 1, 2, 2, '{"Title": {"after": "bbb", "before": "bbb"}, "ResponseContent": {"name": {"after": "json", "before": "json-dft"}}}', '修改bug', '2018-07-30 18:37:12', '2018-07-30 18:37:12', NULL),
	(17, 1, 2, 2, '{"Title": {"after": "bbb", "before": "bbb"}, "ResponseContent": {"name": {"after": "json", "before": "json-dft"}}}', '修改bug', '2018-07-30 19:18:17', '2018-07-30 19:18:17', NULL),
	(18, 1, 2, 2, '{"Title": {"after": "bbb1", "before": "bbb"}}', '修改bug', '2018-07-31 20:47:26', '2018-07-31 20:47:26', NULL),
	(19, 1, 1, 1, '{"ID": 1, "title": "bbb1", "status": 1, "task_id": 1, "CreatedAt": "2018-07-29T18:10:30+08:00", "DeletedAt": null, "UpdatedAt": "2018-07-31T20:47:26+08:00", "author_id": 1, "module_id": 1, "request_url": "/doc/api/", "request_param": {"Content-Type": "application/json-dft"}, "request_header": {"Content-Type": "application/json-dft"}, "request_method": "get", "response_content": {"name": "json-dft", "Content-Type": "application/json-dft"}}', 'rebuild', '2018-07-31 20:48:08', '2018-07-31 20:48:08', NULL),
	(20, 1, 2, 2, '{"Title": {"after": "bbb", "before": "test"}}', '修改bug', '2018-08-02 10:42:16', '2018-08-02 10:42:16', NULL),
	(21, 1, 1, 1, '{"id": 1, "title": "bbb", "status": 1, "task_id": 1, "author_id": 1, "module_id": 1, "created_at": "2018-07-29T18:10:30+08:00", "deleted_at": null, "updated_at": "2018-08-02T10:42:16+08:00", "request_url": "/doc/api/", "commit_param": null, "commit_header": null, "request_param": {"Content-Type": "application/json-dft"}, "commit_content": null, "commit_message": "", "commit_task_id": 0, "request_header": {"Content-Type": "application/json-dft"}, "request_method": "get", "commit_author_id": 0, "response_content": {"name": "json-dft", "Content-Type": "application/json-dft"}}', 'rebuild', '2018-08-02 10:42:19', '2018-08-02 10:42:19', NULL),
	(22, 2, 2, 2, '{"Title": {"after": "sss", "before": "bbb"}}', '修改bug', '2018-08-02 21:25:33', '2018-08-02 21:25:33', NULL);
/*!40000 ALTER TABLE `api_commits` ENABLE KEYS */;

-- 导出  表 apibuilder.api_logs 结构
CREATE TABLE IF NOT EXISTS `api_logs` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL DEFAULT '0',
  `type` smallint(3) NOT NULL DEFAULT '0' COMMENT '1发布 2修改 3对接 4测试 5变更作者 6重构 7注释字段 8注释model',
  `entity_id` int(10) NOT NULL DEFAULT '0',
  `entity_type` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.api_logs 的数据：~8 rows (大约)
/*!40000 ALTER TABLE `api_logs` DISABLE KEYS */;
INSERT INTO `api_logs` (`id`, `user_id`, `type`, `entity_id`, `entity_type`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, 0, '0', '2018-07-30 19:18:17', '2018-07-30 19:18:17', NULL),
	(2, 2, 2, 0, '0', '2018-07-30 19:18:17', '2018-07-30 19:18:17', NULL),
	(3, 1, 1, 0, '0', '2018-07-31 14:00:24', '2018-07-31 14:00:24', NULL),
	(4, 2, 2, 0, '0', '2018-07-31 20:47:26', '2018-07-31 20:47:26', NULL),
	(5, 1, 6, 0, '0', '2018-07-31 20:48:09', '2018-07-31 20:48:09', NULL),
	(6, 2, 2, 0, '0', '2018-08-02 10:42:16', '2018-08-02 10:42:16', NULL),
	(7, 1, 6, 0, '0', '2018-08-02 10:42:19', '2018-08-02 10:42:19', NULL),
	(8, 2, 2, 2, '', '2018-08-02 21:25:33', '2018-08-02 21:25:33', NULL);
/*!40000 ALTER TABLE `api_logs` ENABLE KEYS */;

-- 导出  表 apibuilder.api_models 结构
CREATE TABLE IF NOT EXISTS `api_models` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `author_id` int(10) NOT NULL,
  `model_code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_models 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `api_models` DISABLE KEYS */;
INSERT INTO `api_models` (`id`, `author_id`, `model_code`, `model_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(14, 1, 'merchant', '商家1', '2018-08-02 22:50:41', '2018-08-02 22:51:46', NULL);
/*!40000 ALTER TABLE `api_models` ENABLE KEYS */;

-- 导出  表 apibuilder.api_model_maps 结构
CREATE TABLE IF NOT EXISTS `api_model_maps` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `model_id` int(10) NOT NULL DEFAULT '0',
  `author_id` int(10) NOT NULL,
  `target_type` tinyint(1) NOT NULL COMMENT '1belongs to 2has many',
  `target_id` int(10) NOT NULL,
  `deleted_at` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_model_maps 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `api_model_maps` DISABLE KEYS */;
/*!40000 ALTER TABLE `api_model_maps` ENABLE KEYS */;

-- 导出  表 apibuilder.api_model_notes 结构
CREATE TABLE IF NOT EXISTS `api_model_notes` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `model_id` int(10) NOT NULL DEFAULT '0',
  `author_id` int(10) NOT NULL,
  `model_key` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `level` smallint(3) NOT NULL DEFAULT '1',
  `parent_id` int(10) NOT NULL DEFAULT '0',
  `note` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fkey_parent` (`level`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_model_notes 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `api_model_notes` DISABLE KEYS */;
INSERT INTO `api_model_notes` (`id`, `model_id`, `author_id`, `model_key`, `level`, `parent_id`, `note`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(15, 14, 1, 'name4', 1, 0, '商家名', '2018-08-02 23:18:28', '2018-08-02 23:50:31', NULL),
	(16, 14, 1, 'cid', 1, 0, '城市', '2018-08-02 23:55:57', '2018-08-02 23:55:57', NULL);
/*!40000 ALTER TABLE `api_model_notes` ENABLE KEYS */;

-- 导出  表 apibuilder.api_notes 结构
CREATE TABLE IF NOT EXISTS `api_notes` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `api_id` int(10) DEFAULT NULL,
  `author_id` int(10) NOT NULL,
  `fkey` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fkey_parent` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fkey_token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `note` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fkey_parent` (`fkey_parent`(191)),
  KEY `api_id` (`api_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_notes 的数据：~14 rows (大约)
/*!40000 ALTER TABLE `api_notes` DISABLE KEYS */;
INSERT INTO `api_notes` (`id`, `api_id`, `author_id`, `fkey`, `fkey_parent`, `fkey_token`, `note`, `model_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, '$root.data.merchant', '$root.data', 'mdOG/quDax4mA', '商家', 0, '2017-06-26 08:45:32', '2017-06-26 16:45:32', NULL),
	(2, 1, 1, '$root.data.merchant.list.0', '$root.data.merchant.list', 'mdOG/quDax4mA', 'Merchant', 1, '2017-06-26 08:50:06', '2017-06-26 16:50:06', NULL),
	(3, 0, 1, 'Merchant.id', 'Merchant', 'mdYTYD2vJ4qw6', '商家ID', 0, '2017-06-26 08:52:16', '2017-06-26 18:55:33', NULL),
	(4, 0, 1, 'Merchant.name', 'Merchant', 'mdYTYD2vJ4qw6', '商家名称', 0, '2017-06-26 08:52:16', '2017-06-26 18:55:33', NULL),
	(5, 0, 1, 'Merchant.logo_path', 'Merchant', '0Merchant.logo_path', '头像', 0, '2017-06-27 02:39:13', '2017-06-27 10:39:13', NULL),
	(6, 0, 1, 'Merchant.user_id', 'Merchant', '0Merchant.user_id', 'id', 0, '2017-06-27 02:42:45', '2017-06-27 10:42:45', NULL),
	(7, 0, 1, 'Merchant.active_works_pcount', 'Merchant', '0Merchant.active_works_pcount', '擦拭地方', 0, '2017-06-27 02:44:19', '2017-06-27 10:44:19', NULL),
	(8, 0, 1, 'Merchant.sign', 'Merchant', '0Merchant.sign', '标签', 0, '2017-06-27 02:46:06', '2017-06-27 10:46:06', NULL),
	(9, 0, 1, 'Merchant.fans_count', 'Merchant', '0Merchant.fans_count', '粉丝', 0, '2017-06-27 05:27:59', '2017-06-27 13:27:59', NULL),
	(10, 0, 1, 'Merchant.fans_count', 'Merchant', '0Merchant.fans_count', '粉丝', 0, '2017-06-27 05:36:40', '2017-06-27 13:36:40', NULL),
	(11, 0, 1, 'Merchant.fans_count', 'Merchant', '0Merchant.fans_count', 'fens', 0, '2017-06-27 06:01:58', '2017-06-27 06:01:58', NULL),
	(12, 0, 1, 'Merchant.bond_sign', 'Merchant', '0Merchant.bond_sign', '保证金', 0, '2017-06-29 06:36:10', '2017-06-29 06:36:10', NULL),
	(13, 0, 1, 'Merchant.grade', 'Merchant', '0Merchant.grade', '等级', 0, '2017-07-06 11:55:33', '2017-07-06 11:55:33', NULL),
	(14, 2, 1, 'user', 'test', 'test.user', '用户', 0, '2018-08-03 00:33:14', '2018-08-03 00:34:36', NULL);
/*!40000 ALTER TABLE `api_notes` ENABLE KEYS */;

-- 导出  表 apibuilder.departments 结构
CREATE TABLE IF NOT EXISTS `departments` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `department_name` varchar(100) NOT NULL DEFAULT '0',
  `pid` int(10) NOT NULL DEFAULT '0',
  `lead_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.departments 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;
/*!40000 ALTER TABLE `departments` ENABLE KEYS */;

-- 导出  表 apibuilder.modules 结构
CREATE TABLE IF NOT EXISTS `modules` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `spid` varchar(255) DEFAULT NULL COMMENT '所有父级ID列',
  `pid` int(10) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `author_id` int(10) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.modules 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `modules` DISABLE KEYS */;
INSERT INTO `modules` (`id`, `title`, `spid`, `pid`, `author_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, '用户中心', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(2, '版本', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(3, '模块', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(4, '接口', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(5, 'kai', '', 0, 1, '2018-07-29 16:57:27', '2018-07-29 16:57:27', NULL),
	(6, 'kai', NULL, 0, 1, '2018-07-31 13:39:13', '2018-07-31 13:39:13', NULL);
/*!40000 ALTER TABLE `modules` ENABLE KEYS */;

-- 导出  表 apibuilder.users 结构
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `role_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.users 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password`, `avatar`, `email`, `phone`, `status`, `role_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'yidao', '123123', 'sql.png', 'yi_dao@hunliji.com', '13567119103', 1, 0, '2018-08-05 23:33:43', '2018-08-05 23:34:06', NULL),
	(2, 'yidao', '123123', 'sql.png', 'yi_dao@hunliji.com', '13567119103', 0, 0, '2018-08-05 23:50:05', '2018-08-05 23:50:05', NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- 导出  表 apibuilder.user_roles 结构
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL,
  `lead_id` int(10) NOT NULL,
  `department_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.user_roles 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
