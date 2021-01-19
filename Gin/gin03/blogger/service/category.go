package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
