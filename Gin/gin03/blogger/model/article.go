//   文章表实体
package model

import "time"

// CREATE TABLE `article` (
// 	`id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章id',
// 	`category_id` bigint(20) NOT NULL COMMENT '分类id',
// 	`content` longtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章内容',
// 	`title` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章标题',
// 	`view_count` int(255) NOT NULL COMMENT '阅读次数',
// 	`comment_count` int(255) NOT NULL COMMENT '评论次数',
// 	`username` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '作者',
// 	`status` int(10) NOT NULL DEFAULT '1' COMMENT '状态，正常为1',
// 	`summary` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章摘要',
// 	`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
// 	`update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
// 	PRIMARY KEY (`id`)
//   ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
// id, summary, title, view_count,
// create_time, comment_count, username, category_id
// 文章的结构体  可以不写文章内容，因为是大文本
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CreateTime   time.Time `db:"create_time"`
	CommentCount uint32    `db:"comment_count"`
	Username     string    `db:"username"`
}

// 用户文章详情页的实体
// 为了提升效率
type ArticleDetail struct {
	ArticleInfo
	// 文章内容
	Content  string `db:"content"`
	Category        // 匿名结构体嵌套 Category：Category
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
