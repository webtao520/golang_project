package service

import (
	"blogger/dao/db"
	"blogger/model"
	"fmt"
	"math"
)

// 业务逻辑
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Printf("get article list failed, err:%v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	// 获取分类ids
	categoryIds := getCategoryIds(articleInfoList)
	// 依据分类id 获取分类信息
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	for _, article := range articleInfoList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId //分类id
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}

func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		for _, id := range ids { // 去重
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// 根据分类id获取文章信息 (用于文章上下页)
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleInfoList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Printf("get artcle list failed,err %v\n", err)
		return
	}
	if len(articleInfoList) == 0 {
		return
	}
	// fmt.Printf("%p\n", articleRecordList)
	// fmt.Println("+++++++++++++++++++", articleRecordList)
	categoryIds := getCategoryIds(articleInfoList)       // 过滤到重复的分类id
	categoryList, err := db.GetCategoryList(categoryIds) //  根据分类id 获取分类信息
	if err != nil {
		fmt.Printf("2 get category list failed, err:%v\n", err)
		return
	}

	for _, article := range articleInfoList {
		//fmt.Println("article==========>", *article) // article==========> &{2 1 好啊后 php 1 2021-01-19 15:21:07 +0000 UTC 1 张涛} 取值
		//fmt.Printf("content:%s\n", article.Summary) // content:好啊后
		// 初始化分页结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}

	return
}

func InsertArticle(content, author, title string, categoryId int64) (err error) {
	// 函数数据量比较大传参数  用指针 性能提升
	articleDetail := &model.ArticleDetail{} // 结构体初始化
	articleDetail.ArticleInfo.CategoryId = categoryId
	articleDetail.Content = content
	articleDetail.ArticleInfo.Username = author
	articleDetail.ArticleInfo.Title = title
	//fmt.Println("content==", content)
	// 字符转换 utf-8
	contentUtf8 := []rune(content) // 字节表示长度
	minLength := int(math.Min(float64(len(contentUtf8)), 128.0))
	//fmt.Println("=======", len(contentUtf8)) //  <p><a href="http://localhost:8008/">个人博客</a></p> 个人博客 48
	//获取文章摘要
	articleDetail.Summary = string([]rune(content)[:minLength])
	id, err := db.InsertArticle(articleDetail)
	fmt.Printf("insert article succ,id:%d,err:%v\n", id, err)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	//获取分类信息
	category, err := db.GetCategoryById(articleDetail.ArticleInfo.CategoryId)
	if err != nil {
		return
	}
	// fmt.Println("====================>", category) // ====================> &{1 技术 1}
	articleDetail.Category = *category // * 取值
	return
}

func GetRelativeAricleList(articleId int64) (articleList []*model.RelativeArticle, err error) {
	articleList, err = db.GetRelativeArticle(articleId)
	return
}

//   获取上一篇/下一篇文章 结构体
func GetPrevAndNextArticleInfo(articleId int64) (prevArticle *model.RelativeArticle, nextArticle *model.RelativeArticle, err error) {
	prevArticle, err = db.GetPrevArticleById(articleId)
	if err != nil {
		return
	}
	nextArticle, err = db.GetNextArticleById(articleId)
	if err != nil {
		return
	}
	return
}
