package model

import "time"

// CREATE TABLE `leave` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '留言id',
//   `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
//   `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
//   `content` text COLLATE utf8mb4_unicode_ci NOT NULL,
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (`id`)
// ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

// 留言表实体
type Leave struct {
	 Id  int64 `db:"id"`
	 Content string `db:"content"` 
	 Username string `db:"username"`
	 CreateTime time.Time `db:"create_time"`
	 Email string   `db:"email"`
}