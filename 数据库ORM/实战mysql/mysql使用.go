


新建test数据库，person、place 表

------------------------------------------------------------------------------------------

CREATE TABLE `person` (
	`user_id` int(11) NOT NULL AUTO_INCREMENT,
	`username` varchar(260) DEFAULT NULL,
	`sex` varchar(260) DEFAULT NULL,
	`email` varchar(260) DEFAULT NULL,
	PRIMARY KEY (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

------------------------------------------------------------------------------------------


mysql使用

	使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql数据库驱动）
	github.com/jmoiron/sqlx （基于mysql驱动的封装）

	驱动包文档
		https://gowalker.org/github.com/jmoiron/sqlx

命令行输入 ：

	go get github.com/go-sql-driver/mysql 
	go get github.com/jmoiron/sqlx  


ORM 特性：

	支持 Go 的所有类型存储
	轻松上手，采用简单的 CRUD 风格
	自动 Join 关联表
	跨数据库兼容查询
	允许直接使用 SQL 查询／映射
	严格完整的测试保证 ORM 的稳定与健壮
	更多特性请在文档中自行品读。

链接mysql

database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")