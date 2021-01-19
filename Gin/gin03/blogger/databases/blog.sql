CREATE TABLE `article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `category_id` bigint(20) NOT NULL COMMENT '分类id',
  `content` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
  `title` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
  `view_count` int(255) NOT NULL COMMENT '阅读次数',
  `comment_count` int(255) NOT NULL COMMENT '评论次数',
  `username` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '作者',
  `status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
  `summary` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章摘要',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci



CREATE TABLE `category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `category_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名字',
  `category_no` int(10) NOT NULL COMMENT '分类排序',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci




CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
  `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论作者',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态:0 删除 1',
  `article_id` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci




CREATE TABLE `leave` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '留言id',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci