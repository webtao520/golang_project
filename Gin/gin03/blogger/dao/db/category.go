// 分类数据层
package db

import (
	"blogger/model"

	"github.com/jmoiron/sqlx"
)

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {

	sqlstr, args, err := sqlx.In("select id, category_name, "+
		"category_no from category where id in(?)", categoryIds)
	if err != nil {
		return
	}

	err = DB.Select(&categoryList, sqlstr, args...)
	return

}

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	sqlstr := "select id, category_name, category_no from category order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return
}
