package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDB() (err error) {
	addr := "root:root123@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", addr)
	fmt.Println("===============>", db)
	if err != nil {
		return err
	}
	// 最大连接
	db.SetMaxOpenConns(100)
	// 最大空闲
	db.SetMaxIdleConns(16)
	return
}

// 查询图书
func QueryAllBook() (bookList []*Book, err error) {
	//bookList = []*Book{}
	sqlStr := "select id ,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	return
}

//保存图书
func InsertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	ret := db.MustExec(sqlStr, title, price)
	if _, err := ret.RowsAffected(); err != nil {
		panic(err.Error())
	}
	return
}

// 删除图书
func Delete(id int64) (err error) {
	sqlStr := "delete from book where id = ?"
	ret := db.MustExec(sqlStr, id)
	if _, err := ret.RowsAffected(); err != nil {
		panic(err.Error())
	}
	return
}
