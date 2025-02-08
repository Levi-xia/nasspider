CREATE TABLE `tv_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '链接',
  `name` varchar(1024) NOT NULL DEFAULT '' COMMENT '名称',
  `total_ep` int(11) NOT NULL DEFAULT 0 COMMENT '总集数',
  `current_ep` int(11) NOT NULL DEFAULT 0 COMMENT '当前集数',
  `status` int(11) NOT NULL DEFAULT 0 COMMENT '状态',
  `download_path` varchar(1024) NOT NULL DEFAULT '' COMMENT '下载路径',
  `xpath` varchar(2048) NOT NULL DEFAULT '' COMMENT 'xpath',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT '类型',
  `provider` varchar(32) NOT NULL DEFAULT '' COMMENT '提供商',
  `downloader` varchar(32) NOT NULL DEFAULT '' COMMENT '下载器',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='任务表';