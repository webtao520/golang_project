package db

// 数据层
import (
	"blogger/model"
	"fmt"
)

//  获取文章列表
func GetArticleList(pageNum,pageSize int)(articleList []*model.ArticleInfo,err error){
	// fmt.Println("======",articleList,err)//  ====== [] <nil>
		if pageNum < 0 || pageSize < 0 {
			err=fmt.Errorf("invalid parameter ,page_num:%d, page_size:%d",pageNum,pageSize)
			return
		}
		sqlstr := `select 
						id, summary, title, view_count,
						 create_time, comment_count, username, category_id
					from 
						article 
					where 
						status = 1
					order by create_time desc
					limit ?, ?`
					err = DB.Select(&articleList, sqlstr, pageNum, pageSize)	
		return
}