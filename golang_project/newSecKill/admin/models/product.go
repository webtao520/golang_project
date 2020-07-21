package models

/*
默认的表名规则，使用驼峰转蛇形：
AuthUser -> auth_user
Auth_User -> auth__user
DB_AuthUser -> d_b__auth_user
FG_RE  -> f_g__r_e
SecKillProduct -> sec_kill_product
除了开头的大写字母以外，遇到大写会增加 _，原名称中的下划线保留。
*/

// 设置参数
type SecKillProduct struct {
	ProductId   int64  `orm:"pk;auto"` // 当 Field 类型为 int, int32, int64, uint, uint32, uint64 时，可以设置字段为自增健
	ProductName string `orm:unique;size(60);default('')" form:"ProductName" valid:"Required"`
	//Total       int    `orm:"size(10);default(0) form:"Total" valid:"Min(0)"`
	Total  int `orm:"size(10);default(0)" form:"Total" valid:"Min(0)"`
	Status int `orm:"default(2)" form:"status" valid:"Range(1,2)"`
}

/*
// 自定义表名 如果前缀设置为 prefix_ 那么表名为：prefix_auth_user
func (s *SecKillProduct) TableName() string {
	return "kill_product"
}
*/
