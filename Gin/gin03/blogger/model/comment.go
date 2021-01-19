package model

import "time"

// CREATE TABLE `comment` (
// 	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '评论id',
// 	`content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论内容',
// 	`username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论作者',
// 	`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
// 	`status` int(255) NOT NULL DEFAULT '1' COMMENT '评论状态:0 删除 1',
// 	`article_id` bigint(20) NOT NULL,
// 	PRIMARY KEY (`id`)
//   ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

// 评论实体
type Comment struct {
	 Id   int64  `db:"id"`
	 Content string `db:"content"`
	 Username string `db:"username"`
	 CreateTime time.Time `db:"create_time"`
	 Status    uint  `db:"status"`
	 ArticleId int64 `db:"article_id"`
}