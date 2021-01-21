// 分类数据层
package db

import (
	"blogger/model"

	"github.com/jmoiron/sqlx"
)

func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	//fmt.Println("+++++++++++++", &categoryList)
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

func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{} // &{0  0}
	//fmt.Println("==========>", category) // nil 空的结构体
	sqlstr := "select id, category_name, category_no from category where id=?"
	err = DB.Get(category, sqlstr, id)
	return
}
