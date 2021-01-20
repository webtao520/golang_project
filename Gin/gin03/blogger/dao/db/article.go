package db

// 数据层
import (
	"blogger/model"
	"fmt"
)

//  获取文章列表
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	// fmt.Println("======",articleList,err)//  ====== [] <nil>
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid parameter ,page_num:%d, page_size:%d", pageNum, pageSize)
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

func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize < 0 {
		err = fmt.Errorf("invalid paramter,page_num:%d, page_size:%d", pageNum, pageSize)
		return
	}

	// type ArticleInfo struct {
	// 	Id           int64     `db:"id"`
	// 	CategoryId   int64     `db:"category_id"`
	// 	Summary      string    `db:"summary"`
	// 	Title        string    `db:"title"`
	// 	ViewCount    uint32    `db:"view_count"`
	// 	CreateTime   time.Time `db:"create_time"`
	// 	CommentCount uint32    `db:"comment_count"`
	// 	Username     string    `db:"username"`
	// }

	sqlstr := `select 
		 id,summary,title,view_count,create_time,comment_count,username,category_id
		from  article 
		where 
		 status = 1
		 and 
		 category_id=?
		 order by create_time desc
		 limit ?,?
	`
	err = DB.Select(&articleList, sqlstr, categoryId, pageNum, pageSize)
	return
}
