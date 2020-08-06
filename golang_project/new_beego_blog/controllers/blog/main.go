package blog

import (
	"fmt"
	"new_beego_blog/models"
)

type MainController struct {
	baseController
}

// 首页
func (this *MainController) Index(){
  //fmt.Println("首页文件...")
  var list []*models.Post
  query:=new(models.Post).Query().Filter("status",0).Filter("urltype",0)
  count,_:=query.Count()
  fmt.Println("======",count,list)

  this.Data["pagebar"]=""
  this.display("index")
}
