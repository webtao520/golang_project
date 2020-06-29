package controllers

import (
	"crawl_movie/models"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/httplib"
  "time"
)

type CrawlMovieController struct {
	beego.Controller
}

/**
 目前这个爬虫只能爬取静态数据 对于像京东的部分动态数据 无法爬取
 对于动态数据 可以采用 一个组件 phantomjs
*/
func (c *CrawlMovieController) CrawlMovie() {
  var movieInfo models.MovieInfo
  //连接到redis
  models.ConnectRedis("127.0.0.1:6379")
    
  //爬虫入口url
  sUrl := "https://movie.douban.com/subject/25827935/"
  models.PutinQueue(sUrl)

  for{
        length := models.GetQueueLength()
        if length == 0{
            break //如果url队列为空 则退出当前循环
        }

        sUrl = models.PopfromQueue()
        //我们应当判断sUrl是否应该被访问过
        if models.IsVisit(sUrl){
           continue
        }

        rsp := httplib.Get(sUrl)
        //设置User-agent以及cookie是为了防止  豆瓣网的 403
        rsp.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0")
        rsp.Header("Cookie", `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`)
        sMovieHtml,err := rsp.String()
        if err != nil{
            panic(err)
        }

        movieInfo.Movie_name            = models.GetMovieName(sMovieHtml)
        //记录电影信息
        if movieInfo.Movie_name != ""{
            movieInfo.Movie_director        = models.GetMovieDirector(sMovieHtml)
            movieInfo.Movie_main_character  = models.GetMovieMainCharacters(sMovieHtml) 
            movieInfo.Movie_type            = models.GetMovieGenre(sMovieHtml)
            movieInfo.Movie_on_time         = models.GetMovieOnTime(sMovieHtml)
            movieInfo.Movie_grade           = models.GetMovieGrade(sMovieHtml)
            movieInfo.Movie_span            = models.GetMovieRunningTime(sMovieHtml)
            
            models.AddMovie(&movieInfo)
        }

        //提取该页面的所有连接
        urls := models.GetMovieUrls(sMovieHtml)

        for _,url := range urls{
            models.PutinQueue(url)
            c.Ctx.WriteString("<br>" + url + "</br>")
        }

        //sUrl 应当记录到 访问set中
        models.AddToSet(sUrl)

        time.Sleep(time.Second)
    }

    c.Ctx.WriteString("end of crawl!")
}
