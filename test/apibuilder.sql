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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.apis 的数据：~9 rows (大约)
/*!40000 ALTER TABLE `apis` DISABLE KEYS */;
INSERT INTO `apis` (`id`, `task_id`, `module_id`, `author_id`, `title`, `request_url`, `request_method`, `request_param`, `request_header`, `response_content`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 1, '2018-07-29 18:10:30', '2018-08-02 13:03:11', '2018-08-02 21:03:11'),
	(2, 1, 1, 1, 'bbb', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"name": "json-dft", "Content-Type": "application/json-dft"}', 2, '2018-07-31 13:38:27', '2018-08-15 19:46:16', NULL),
	(3, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-07-31 15:04:35', '2018-07-31 15:04:35', NULL),
	(4, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-07-31 20:43:49', '2018-07-31 20:43:49', NULL),
	(5, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-01 18:26:30', '2018-08-01 18:26:30', NULL),
	(6, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-02 10:41:48', '2018-08-02 10:41:48', NULL),
	(7, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-02 21:03:01', '2018-08-02 21:03:01', NULL),
	(8, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-15 16:22:27', '2018-08-15 16:22:27', NULL),
	(9, 1, 1, 1, 'test', '/doc/api/', 'get', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', '{"Content-Type": "application/json-dft"}', 0, '2018-08-15 19:24:50', '2018-08-15 19:24:50', NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_models 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `api_models` DISABLE KEYS */;
INSERT INTO `api_models` (`id`, `author_id`, `model_code`, `model_name`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(14, 1, 'merchant', '商家1222', '2018-08-02 22:50:41', '2018-08-15 19:49:34', NULL),
	(15, 1, 'merchant', '商家', '2018-08-15 19:50:47', '2018-08-15 11:50:56', '2018-08-15 19:50:57');
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
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_model_notes 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `api_model_notes` DISABLE KEYS */;
INSERT INTO `api_model_notes` (`id`, `model_id`, `author_id`, `model_key`, `level`, `parent_id`, `note`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(15, 14, 1, 'name4', 1, 0, '商家名', '2018-08-02 23:18:28', '2018-08-02 23:50:31', NULL),
	(16, 14, 1, 'cid', 1, 0, '城市', '2018-08-02 23:55:57', '2018-08-02 23:55:57', NULL),
	(17, 0, 1, 'cid', 1, 0, '城市', '2018-08-15 19:56:39', '2018-08-15 19:56:39', NULL),
	(18, 14, 1, 'tescid', 1, 0, '城市', '2018-08-15 20:02:34', '2018-08-15 20:02:34', NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.api_notes 的数据：~15 rows (大约)
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
	(14, 2, 1, 'user', 'test', 'test.user', '用户', 0, '2018-08-03 00:33:14', '2018-08-03 00:34:36', NULL),
	(15, 0, 0, 'user', 'test', '', '用户', 0, '2018-08-15 19:53:24', '2018-08-15 19:53:24', NULL);
/*!40000 ALTER TABLE `api_notes` ENABLE KEYS */;

-- 导出  表 apibuilder.bugs 结构
CREATE TABLE IF NOT EXISTS `bugs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL DEFAULT '0',
  `appoint_user_id` int(11) NOT NULL DEFAULT '0',
  `task_id` int(11) NOT NULL DEFAULT '0',
  `priority` tinyint(1) NOT NULL DEFAULT '1',
  `title` varchar(50) NOT NULL,
  `description` varchar(500) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.bugs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `bugs` DISABLE KEYS */;
/*!40000 ALTER TABLE `bugs` ENABLE KEYS */;

-- 导出  表 apibuilder.bug_proofs 结构
CREATE TABLE IF NOT EXISTS `bug_proofs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bug_id` int(11) NOT NULL DEFAULT '0',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1请求 2响应',
  `target_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1代理测试 2场景测试',
  `target_id` int(11) NOT NULL DEFAULT '0',
  `error_marks` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.bug_proofs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `bug_proofs` DISABLE KEYS */;
/*!40000 ALTER TABLE `bug_proofs` ENABLE KEYS */;

-- 导出  表 apibuilder.containers 结构
CREATE TABLE IF NOT EXISTS `containers` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL DEFAULT '0',
  `status` int(10) NOT NULL DEFAULT '0' COMMENT '1默认激活 0非默认',
  `last_author_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.containers 的数据：~33 rows (大约)
/*!40000 ALTER TABLE `containers` DISABLE KEYS */;
INSERT INTO `containers` (`id`, `title`, `status`, `last_author_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'test', 0, 0, '2018-08-09 17:09:51', '2018-08-09 13:02:19', NULL),
	(2, 'test', 0, 0, '2018-08-09 17:17:01', '2018-08-09 13:02:19', NULL),
	(3, 'test', 0, 0, '2018-08-09 17:26:56', '2018-08-09 13:02:19', NULL),
	(5, 'test', 0, 0, '2018-08-09 12:14:26', '2018-08-09 13:02:19', NULL),
	(6, 'test', 0, 0, '2018-08-09 12:14:31', '2018-08-09 13:02:19', NULL),
	(7, 'test', 0, 0, '2018-08-09 12:15:37', '2018-08-09 13:02:19', NULL),
	(8, 'test', 0, 0, '2018-08-09 12:16:11', '2018-08-09 13:02:19', NULL),
	(9, 'test', 0, 0, '2018-08-09 12:17:37', '2018-08-09 13:02:19', NULL),
	(10, 'test', 0, 0, '2018-08-09 12:17:52', '2018-08-09 13:02:19', NULL),
	(11, 'test', 0, 0, '2018-08-09 12:18:11', '2018-08-09 13:02:19', NULL),
	(12, 'test', 0, 0, '2018-08-09 12:18:28', '2018-08-09 13:02:19', NULL),
	(13, 'test', 0, 0, '2018-08-09 12:57:21', '2018-08-09 13:02:19', NULL),
	(14, 'test', 0, 0, '2018-08-09 13:00:06', '2018-08-09 13:02:19', NULL),
	(15, 'test', 0, 0, '2018-08-09 13:00:31', '2018-08-09 13:02:19', NULL),
	(16, 'test', 0, 0, '2018-08-10 16:03:19', '2018-08-10 16:03:19', NULL),
	(17, 'test22', 0, 0, '2018-08-10 16:05:01', '2018-08-10 16:05:01', NULL),
	(18, 'test22', 0, 0, '2018-08-10 16:11:29', '2018-08-10 16:11:29', NULL),
	(19, 'test22', 0, 0, '2018-08-10 16:12:31', '2018-08-10 16:12:31', NULL),
	(20, 'test22', 0, 0, '2018-08-10 16:13:05', '2018-08-10 16:13:05', NULL),
	(21, 'test22', 0, 0, '2018-08-10 16:15:58', '2018-08-10 16:15:58', NULL),
	(22, 'test22', 0, 0, '2018-08-10 16:24:11', '2018-08-10 16:24:11', NULL),
	(23, 'test22', 0, 0, '2018-08-10 16:25:13', '2018-08-10 16:25:13', NULL),
	(24, 'test22', 0, 0, '2018-08-10 16:25:13', '2018-08-10 16:25:13', NULL),
	(25, 'test22', 0, 0, '2018-08-10 16:25:13', '2018-08-10 16:25:13', NULL),
	(26, 'test22', 0, 0, '2018-08-10 16:26:40', '2018-08-10 16:26:40', NULL),
	(27, 'test22', 0, 0, '2018-08-10 16:26:40', '2018-08-10 16:26:40', NULL),
	(28, 'test22', 0, 0, '2018-08-10 16:26:40', '2018-08-10 16:26:40', NULL),
	(29, 'test22', 0, 0, '2018-08-10 16:27:49', '2018-08-10 16:27:49', NULL),
	(30, 'test33', 0, 0, '2018-08-10 16:27:49', '2018-08-10 16:27:49', NULL),
	(31, 'test33', 0, 0, '2018-08-10 16:41:38', '2018-08-10 16:41:38', NULL),
	(32, 'test33', 0, 0, '2018-08-10 16:41:39', '2018-08-10 16:41:39', NULL),
	(33, 'test33', 0, 1, '2018-08-14 20:31:53', '2018-08-14 20:31:53', NULL),
	(34, 'test33', 0, 1, '2018-08-15 14:35:05', '2018-08-15 14:35:05', NULL);
/*!40000 ALTER TABLE `containers` ENABLE KEYS */;

-- 导出  表 apibuilder.container_deploys 结构
CREATE TABLE IF NOT EXISTS `container_deploys` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `container_id` int(10) NOT NULL DEFAULT '0',
  `team_id` int(10) NOT NULL DEFAULT '0',
  `deploy_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0无需部署打包 1开发人员自己 2按指定角色 3按指定人员 ',
  `deploy_user` int(10) NOT NULL DEFAULT '0',
  `last_author_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.container_deploys 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `container_deploys` DISABLE KEYS */;
INSERT INTO `container_deploys` (`id`, `container_id`, `team_id`, `deploy_type`, `deploy_user`, `last_author_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 5, 1, 8, 1, '2018-08-14 21:18:40', '2018-08-14 21:18:40', NULL),
	(2, 2, 5, 1, 8, 1, '2018-08-15 14:38:56', '2018-08-15 14:38:56', NULL);
/*!40000 ALTER TABLE `container_deploys` ENABLE KEYS */;

-- 导出  表 apibuilder.container_params 结构
CREATE TABLE IF NOT EXISTS `container_params` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `container_id` int(10) NOT NULL DEFAULT '0',
  `key_string` varchar(50) NOT NULL,
  `value_string` varchar(255) NOT NULL,
  `last_author_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.container_params 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `container_params` DISABLE KEYS */;
INSERT INTO `container_params` (`id`, `container_id`, `key_string`, `value_string`, `last_author_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 'domain', 'localhost', 1, '2018-08-14 21:15:55', '2018-08-14 21:15:55', NULL);
/*!40000 ALTER TABLE `container_params` ENABLE KEYS */;

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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.modules 的数据：~7 rows (大约)
/*!40000 ALTER TABLE `modules` DISABLE KEYS */;
INSERT INTO `modules` (`id`, `title`, `spid`, `pid`, `author_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, '用户中心', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(2, '版本', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(3, '模块', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(4, '接口', '', 0, 1, '2018-07-29 10:53:53', '2018-07-29 10:53:53', NULL),
	(5, 'kai', '', 0, 1, '2018-07-29 16:57:27', '2018-07-29 16:57:27', NULL),
	(6, 'kai', NULL, 0, 1, '2018-07-31 13:39:13', '2018-07-31 13:39:13', NULL),
	(7, 'kai', '', 0, 1, '2018-08-15 17:27:04', '2018-08-15 17:27:04', NULL);
/*!40000 ALTER TABLE `modules` ENABLE KEYS */;

-- 导出  表 apibuilder.notifications 结构
CREATE TABLE IF NOT EXISTS `notifications` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL DEFAULT '0',
  `type` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL,
  `entity_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `entity_id` int(10) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `message` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未读 1已读',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.notifications 的数据：~13 rows (大约)
/*!40000 ALTER TABLE `notifications` DISABLE KEYS */;
INSERT INTO `notifications` (`id`, `user_id`, `type`, `entity_type`, `entity_id`, `title`, `message`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, '0', '', 0, '任务指定', '[yidao task222 yidao] 指定了任务（%!s(MISSING)）给%!s(MISSING)', 0, '2018-08-30 18:17:37', '2018-08-30 18:17:37', NULL),
	(2, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（%!!(string=yidao)!(string=task222)s(MISSING)）给%!!(MISSING)!(MISSING)s(MISSING)', 0, '2018-08-30 18:23:47', '2018-08-30 18:23:47', NULL),
	(3, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（%!!(string=yidao)!(string=task222)s(MISSING)）给%!!(MISSING)!(MISSING)s(MISSING)', 0, '2018-08-30 18:31:15', '2018-08-30 18:31:15', NULL),
	(4, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（task222）给yidao', 0, '2018-08-30 18:31:31', '2018-08-30 18:31:31', NULL),
	(5, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（task222）给yidao', 0, '2018-08-31 14:02:08', '2018-08-31 14:02:08', NULL),
	(6, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（task222）给yidao', 0, '2018-08-31 14:04:35', '2018-08-31 14:04:35', NULL),
	(7, 1, '0', '', 0, '任务指定', 'yidao 指定了任务（task222）给yidao', 0, '2018-08-31 14:04:58', '2018-08-31 14:04:58', NULL),
	(8, 1, 'task_dispatch', '', 0, '任务指定', 'yidao 指定了任务（task222）给yidao', 0, '2018-08-31 17:08:23', '2018-08-31 17:08:23', NULL),
	(9, 1, 'task_dispatch', '', 0, '分配团队任务', 'yidao 分配团队任务（task222）给', 0, '2018-08-31 17:10:44', '2018-08-31 17:10:44', NULL),
	(10, 9, 'task_separate', '', 0, '分解任务', ':yidao 分解团队任务（web） => 开发任务（子门店）， 指定给123', 0, '2018-08-31 19:07:44', '2018-08-31 19:07:44', NULL),
	(11, 2, 'task_develop', '', 0, '开发任务', 'yidao 分配了新开发任务（子门店）给你', 0, '2018-08-31 19:07:44', '2018-08-31 19:07:44', NULL),
	(12, 9, 'task_separate', '', 0, '分解任务', 'php:yidao 分解团队任务（web） => 开发任务（支付）， 指定给123', 0, '2018-08-31 19:09:40', '2018-08-31 19:09:40', NULL),
	(13, 2, 'task_develop', '', 0, '开发任务', 'yidao 分配了新开发任务（支付）给你', 0, '2018-08-31 19:09:40', '2018-08-31 19:09:40', NULL);
/*!40000 ALTER TABLE `notifications` ENABLE KEYS */;

-- 导出  表 apibuilder.proxys 结构
CREATE TABLE IF NOT EXISTS `proxys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `port` int(11) NOT NULL DEFAULT '0',
  `user_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1运行中 2锁定',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.proxys 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `proxys` DISABLE KEYS */;
/*!40000 ALTER TABLE `proxys` ENABLE KEYS */;

-- 导出  表 apibuilder.proxy_channels 结构
CREATE TABLE IF NOT EXISTS `proxy_channels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `proxy_id` int(11) NOT NULL DEFAULT '0',
  `user_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.proxy_channels 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `proxy_channels` DISABLE KEYS */;
/*!40000 ALTER TABLE `proxy_channels` ENABLE KEYS */;

-- 导出  表 apibuilder.proxy_reqs 结构
CREATE TABLE IF NOT EXISTS `proxy_reqs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `proxy_channel_id` int(11) NOT NULL DEFAULT '0',
  `remote_addr` varchar(50) NOT NULL DEFAULT '0',
  `user_agent` varchar(50) NOT NULL DEFAULT '0',
  `url` varchar(255) NOT NULL DEFAULT '0',
  `method` varchar(20) NOT NULL DEFAULT '0',
  `headers` json NOT NULL,
  `params` json NOT NULL,
  `response` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.proxy_reqs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `proxy_reqs` DISABLE KEYS */;
/*!40000 ALTER TABLE `proxy_reqs` ENABLE KEYS */;

-- 导出  表 apibuilder.scenes 结构
CREATE TABLE IF NOT EXISTS `scenes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='场景测试';

-- 正在导出表  apibuilder.scenes 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `scenes` DISABLE KEYS */;
INSERT INTO `scenes` (`id`, `author_id`, `title`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 't1', '2018-10-04 19:34:25', '2018-10-04 20:28:10', NULL),
	(2, 1, 't1', '2018-10-04 19:36:54', '2018-10-04 20:33:59', NULL);
/*!40000 ALTER TABLE `scenes` ENABLE KEYS */;

-- 导出  表 apibuilder.scene_asserts 结构
CREATE TABLE IF NOT EXISTS `scene_asserts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `target_id` int(11) NOT NULL,
  `target_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1对单条请求的断言 2对整个场景的断言',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1断言 2赋值',
  `src_expr` varchar(255) NOT NULL,
  `dest_expr` varchar(255) NOT NULL,
  `is_global` varchar(255) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.scene_asserts 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `scene_asserts` DISABLE KEYS */;
/*!40000 ALTER TABLE `scene_asserts` ENABLE KEYS */;

-- 导出  表 apibuilder.scene_historys 结构
CREATE TABLE IF NOT EXISTS `scene_historys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `scene_id` int(11) NOT NULL DEFAULT '0',
  `spend_time` int(10) NOT NULL DEFAULT '0',
  `err_num` mediumint(3) NOT NULL DEFAULT '0',
  `ok_num` mediumint(3) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1运行中 2已完成 3失败',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.scene_historys 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `scene_historys` DISABLE KEYS */;
/*!40000 ALTER TABLE `scene_historys` ENABLE KEYS */;

-- 导出  表 apibuilder.scene_history_asserts 结构
CREATE TABLE IF NOT EXISTS `scene_history_asserts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `scene_history_id` int(11) NOT NULL DEFAULT '0',
  `target_id` int(11) NOT NULL,
  `target_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1对单条请求的断言 2对整个场景的断言',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1断言 2赋值',
  `src_expr` varchar(255) NOT NULL,
  `dest_expr` varchar(255) NOT NULL,
  `is_global` varchar(255) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.scene_history_asserts 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `scene_history_asserts` DISABLE KEYS */;
/*!40000 ALTER TABLE `scene_history_asserts` ENABLE KEYS */;

-- 导出  表 apibuilder.scene_history_reqs 结构
CREATE TABLE IF NOT EXISTS `scene_history_reqs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `scene_history_id` int(11) NOT NULL DEFAULT '0',
  `scene_req_id` int(11) NOT NULL DEFAULT '0',
  `url` varchar(255) NOT NULL DEFAULT '0',
  `method` varchar(20) NOT NULL DEFAULT '0',
  `headers` json NOT NULL,
  `params` json NOT NULL,
  `response` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- 正在导出表  apibuilder.scene_history_reqs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `scene_history_reqs` DISABLE KEYS */;
/*!40000 ALTER TABLE `scene_history_reqs` ENABLE KEYS */;

-- 导出  表 apibuilder.scene_reqs 结构
CREATE TABLE IF NOT EXISTS `scene_reqs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `scene_id` int(11) NOT NULL DEFAULT '0',
  `url` varchar(255) NOT NULL DEFAULT '0',
  `method` varchar(20) NOT NULL DEFAULT '0',
  `headers` json NOT NULL,
  `params` json NOT NULL,
  `response` json NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.scene_reqs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `scene_reqs` DISABLE KEYS */;
/*!40000 ALTER TABLE `scene_reqs` ENABLE KEYS */;

-- 导出  表 apibuilder.students 结构
CREATE TABLE IF NOT EXISTS `students` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `age` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.students 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `students` DISABLE KEYS */;
/*!40000 ALTER TABLE `students` ENABLE KEYS */;

-- 导出  表 apibuilder.tasks 结构
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `author_id` int(10) NOT NULL DEFAULT '0',
  `appoint_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '指定人',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `priority` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1高 2中 3低',
  `deadline` datetime DEFAULT NULL,
  `version_id` int(10) NOT NULL DEFAULT '0',
  `has_prd` tinyint(1) NOT NULL DEFAULT '0',
  `is_check` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未校验 1申请校验 2打回开发 3校验ok',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0待分配 1已分配 2开发中 3测试中 4已上线 5终止，归档下线',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='新需求任务';

-- 正在导出表  apibuilder.tasks 的数据：~6 rows (大约)
/*!40000 ALTER TABLE `tasks` DISABLE KEYS */;
INSERT INTO `tasks` (`id`, `author_id`, `appoint_user_id`, `title`, `description`, `priority`, `deadline`, `version_id`, `has_prd`, `is_check`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 0, 2, 'task1', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 0, '2018-08-14 20:07:39', '2018-08-31 07:39:43', NULL),
	(2, 1, 22, 'ddd', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 0, '2018-08-14 20:18:25', '2018-10-06 19:15:20', NULL),
	(3, 1, 2, 'task1', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 0, '2018-08-14 21:09:23', '2018-08-31 07:39:46', NULL),
	(4, 1, 2, 'task1', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 0, '2018-08-15 14:40:47', '2018-08-31 07:39:47', NULL),
	(5, 1, 2, 'task1', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 0, '2018-08-16 18:33:07', '2018-08-31 07:39:49', NULL),
	(6, 1, 9, 'task222', 'test33', 1, '2018-08-30 23:04:05', 1, 0, 0, 2, '2018-08-30 17:48:34', '2018-08-31 18:20:07', NULL);
/*!40000 ALTER TABLE `tasks` ENABLE KEYS */;

-- 导出  表 apibuilder.task_containers 结构
CREATE TABLE IF NOT EXISTS `task_containers` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL DEFAULT '0',
  `author_id` int(10) NOT NULL DEFAULT '0',
  `container_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0未部署 1请求部署 2部署完成 3撤销这个环境部署',
  `reason` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `change_status_user_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='任务需要经过的测试环境';

-- 正在导出表  apibuilder.task_containers 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `task_containers` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_containers` ENABLE KEYS */;

-- 导出  表 apibuilder.task_logs 结构
CREATE TABLE IF NOT EXISTS `task_logs` (
  `id` int(11) DEFAULT NULL,
  `entity_id` int(11) DEFAULT NULL,
  `entity_type` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'clone分散给团队 assign_team 指派团队 assign_user 指派人 split创建子任务 develop开发  render对接 test测试',
  `from_user_id` int(11) DEFAULT NULL,
  `to_user_id` int(11) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.task_logs 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `task_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `task_logs` ENABLE KEYS */;

-- 导出  表 apibuilder.teams 结构
CREATE TABLE IF NOT EXISTS `teams` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `team_name` varchar(50) NOT NULL,
  `lead_id` int(10) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.teams 的数据：~4 rows (大约)
/*!40000 ALTER TABLE `teams` DISABLE KEYS */;
INSERT INTO `teams` (`id`, `team_name`, `lead_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'php', 1, '2018-08-31 17:03:29', '2018-08-31 11:08:46', NULL),
	(2, 'ios', 1, '2018-08-31 17:04:52', '2018-08-31 11:08:50', NULL),
	(3, 'android', 6, '2018-08-31 17:05:28', '2018-08-31 11:08:56', NULL),
	(4, 'web111', 6, '2018-08-31 17:07:10', '2018-08-31 17:07:10', NULL);
/*!40000 ALTER TABLE `teams` ENABLE KEYS */;

-- 导出  表 apibuilder.team_tasks 结构
CREATE TABLE IF NOT EXISTS `team_tasks` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `author_id` int(10) NOT NULL DEFAULT '0',
  `dispatch_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '分解人',
  `appoint_team_id` int(10) NOT NULL DEFAULT '0' COMMENT '指定团队',
  `task_id` int(10) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `deadline` datetime DEFAULT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='团队任务';

-- 正在导出表  apibuilder.team_tasks 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `team_tasks` DISABLE KEYS */;
INSERT INTO `team_tasks` (`id`, `author_id`, `dispatch_user_id`, `appoint_team_id`, `task_id`, `title`, `description`, `deadline`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 0, 1, 6, 'web', 'sdfasdf', '2018-08-30 23:04:05', 0, '2018-08-31 16:57:39', '2018-08-31 16:57:39', NULL),
	(2, 1, 0, 1, 6, 'web2123', 'sdfasdf', '2018-08-30 23:04:05', 0, '2018-08-31 17:07:26', '2018-08-31 17:13:11', NULL),
	(3, 1, 0, 1, 6, 'web', 'sdfasdf', '2018-08-30 23:04:05', 0, '2018-08-31 17:08:23', '2018-08-31 17:08:23', NULL),
	(4, 1, 0, 1, 6, 'web', 'sdfasdf', '2018-08-30 23:04:05', 0, '2018-08-31 17:10:43', '2018-08-31 17:10:43', NULL),
	(5, 1, 0, 0, 0, '子门店', 'sdfasdf', '2018-08-30 23:04:05', 0, '2018-08-31 18:17:52', '2018-08-31 18:17:52', NULL);
/*!40000 ALTER TABLE `team_tasks` ENABLE KEYS */;

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
  `team_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.users 的数据：~2 rows (大约)
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`id`, `username`, `password`, `avatar`, `email`, `phone`, `status`, `role_id`, `team_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 'yidao', '123123', 'sql.png', 'yi_dao@hunliji.com', '13567119103', 1, 0, 1, '2018-08-05 23:33:43', '2018-08-31 10:20:31', NULL),
	(2, '123', '123123', 'sql.png', 'yi_dao2@hunliji.com', '13567119102', 0, 0, 2, '2018-08-05 23:50:05', '2018-08-31 10:20:33', NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

-- 导出  表 apibuilder.user_roles 结构
CREATE TABLE IF NOT EXISTS `user_roles` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_name` int(11) NOT NULL DEFAULT '0',
  `permission` json NOT NULL,
  `last_author_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 正在导出表  apibuilder.user_roles 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;

-- 导出  表 apibuilder.user_tasks 结构
CREATE TABLE IF NOT EXISTS `user_tasks` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `author_id` int(10) NOT NULL DEFAULT '0',
  `appoint_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '指定人',
  `team_task_id` int(10) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0',
  `priority` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1高 2中 3低',
  `deadline` datetime DEFAULT NULL,
  `depend_id` int(10) NOT NULL DEFAULT '0' COMMENT '依赖的任务ID',
  `bind_api_id` int(10) NOT NULL DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0待开发 1开发中 2文档OK，可对接 3待测试 4待部署 5部署完成 6完成 7关闭',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.user_tasks 的数据：~5 rows (大约)
/*!40000 ALTER TABLE `user_tasks` DISABLE KEYS */;
INSERT INTO `user_tasks` (`id`, `author_id`, `appoint_user_id`, `team_task_id`, `title`, `description`, `priority`, `deadline`, `depend_id`, `bind_api_id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 2, 1, '子门店', 'sdfasdf', 0, '2018-08-30 23:04:05', 0, 0, 0, '2018-08-31 18:19:06', '2018-08-31 18:19:06', NULL),
	(2, 1, 2, 1, '子门店', 'sdfasdf', 0, '2018-08-30 23:04:05', 0, 0, 0, '2018-08-31 18:19:06', '2018-08-31 18:19:06', NULL),
	(3, 1, 2, 1, '子门店', 'sdfasdf', 0, '2018-08-30 23:04:05', 0, 0, 0, '2018-08-31 18:20:07', '2018-08-31 18:20:07', NULL),
	(4, 1, 2, 1, '子门店', 'sdfasdf', 0, '2018-08-30 23:04:05', 0, 0, 0, '2018-08-31 19:07:43', '2018-08-31 19:07:43', NULL),
	(5, 1, 2, 1, '支付', 'sdfasdf', 0, '2018-08-30 23:04:05', 0, 0, 0, '2018-08-31 19:09:40', '2018-08-31 19:09:40', NULL);
/*!40000 ALTER TABLE `user_tasks` ENABLE KEYS */;

-- 导出  表 apibuilder.user_task_apis 结构
CREATE TABLE IF NOT EXISTS `user_task_apis` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL DEFAULT '0',
  `api_id` int(10) NOT NULL DEFAULT '0',
  `user_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.user_task_apis 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `user_task_apis` DISABLE KEYS */;
INSERT INTO `user_task_apis` (`id`, `task_id`, `api_id`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 11, 1, '2018-08-31 20:10:07', '2018-08-31 20:10:07', NULL);
/*!40000 ALTER TABLE `user_task_apis` ENABLE KEYS */;

-- 导出  表 apibuilder.user_task_depends 结构
CREATE TABLE IF NOT EXISTS `user_task_depends` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `task_id` int(10) NOT NULL DEFAULT '0',
  `depend_id` int(10) NOT NULL DEFAULT '0',
  `user_id` int(10) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.user_task_depends 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `user_task_depends` DISABLE KEYS */;
INSERT INTO `user_task_depends` (`id`, `task_id`, `depend_id`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
	(1, 1, 2, 1, '2018-08-31 20:08:52', '2018-08-31 20:08:52', NULL);
/*!40000 ALTER TABLE `user_task_depends` ENABLE KEYS */;

-- 导出  表 apibuilder.versions 结构
CREATE TABLE IF NOT EXISTS `versions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag` int(11) DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 正在导出表  apibuilder.versions 的数据：~0 rows (大约)
/*!40000 ALTER TABLE `versions` DISABLE KEYS */;
/*!40000 ALTER TABLE `versions` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
