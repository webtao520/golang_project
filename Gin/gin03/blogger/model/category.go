package model

// CREATE TABLE `category` (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类id',
//   `category_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '分类名字',
//   `category_no` int(10) NOT NULL COMMENT '分类排序',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (`id`)
// ) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

// 定义分类实体
type Category struct {
	CategoryId   int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   uint   `db:"category_no"`
}
