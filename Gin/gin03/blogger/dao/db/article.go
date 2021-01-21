package db

// 数据层
import (
	"blogger/model"
	"fmt"
)

// 保存文章
func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//fmt.Println(article) // &{{0 1 <p>This is some sample content.</p> 1 0 0001-01-01 00:00:00 +0000 UTC 0 1} <p>This is some sample content.</p> {0  0}}
	if article == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlstr := `insert into 
	article(content, summary, title, username, 
		category_id, view_count, comment_count)
		values(?, ?, ?, ?, ?, ?, ?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

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

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		err = fmt.Errorf("invalid parameter,article_id:%d", articleId)
		return
	}

	//articleDetail = &model.ArticleDetail{} // TODO
	articleDetail = &model.ArticleDetail{}
	//var a []*model.ArticleInfo
	// fmt.Println("=====================", a) //nil
	sqlstr := `select 
							id, summary, title, view_count, content,
							 create_time, comment_count, username, category_id
						from 
							article 
						where 
							id = ?
						and
							status = 1
						`
	err = DB.Get(articleDetail, sqlstr, articleId)
	return
}
