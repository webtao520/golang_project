package db

// 定义book 数据结构体
type Book struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
	Price int64  `db:"price"`
}
