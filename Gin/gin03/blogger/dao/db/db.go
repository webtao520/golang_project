package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(addr string) (err error) {
	// dbParams
	// parseTime=true 将mysql中时间类型，自动解析为go结构体中的时间类型
	// 不加报错
	//addr := "root:root123@tcp(localhost:3306)/test?parseTime=true"
	DB, err = sqlx.Open("mysql", addr)
	if err != nil {
		panic(err.Error())
	}
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	// 默认情况下，池的增长是无限制的，当池中没有空闲连接时，就会创建连接。
	// 您可以使用DB。setmaxopenconn设置池的最大大小。未被使用的连接被标记为空闲，
	// 如果不需要，则关闭。为了避免创建和关闭大量的连接，
	// 使用DB设置最大空闲大小。将maxidleconns设置为适合查询加载的大小
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return
}
