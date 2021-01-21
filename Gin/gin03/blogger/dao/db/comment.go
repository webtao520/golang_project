package db

import (
	"blogger/model"
	"fmt"
)


//  获取所有评论
func GetCommentList(articleId int64, pageNum,pageSize int) (commentList []*model.Comment,err error){
    if pageNum <0 || pageSize <0 {
		err=fmt.Errorf("")
	}
}



