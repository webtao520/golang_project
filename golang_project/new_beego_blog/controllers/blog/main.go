package blog

import (
	"new_beego_blog/models"
	//"fmt"
)

type MainController struct {
	baseController
}

// 首页
func (this *MainController) Index(){
  var list []*models.Post
  query:=new(models.Post).Query().Filter("status",0).Filter("urltype",0)
  count,_:=query.Count()
   if count > 0 {
	   query.OrderBy("-istop", "-posttime").Limit(this.pagesize, (this.page-1)*this.pagesize).RelatedSel().All(&list)
   }
  this.Data["list"]=list
  /*
	for k,v:=range list {
		fmt.Println(k,"======>",v.Comments)
		fmt.Printf("%T",v) // *models.Post
	}
 */
  this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize), "/index%d.html").ToString()
  this.setHeadMetas()  // 设置seo
  this.display("index")
}
