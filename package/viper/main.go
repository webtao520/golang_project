package main 

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	 "github.com/fsnotify/fsnotify" // 监听和重新读取配置文件
)


func init(){
	viper.SetConfigName("config") //　获取json配置文件
	viper.AddConfigPath(".") // 设置配置文件和可执行二进制文件在用一个目录
	if err:=viper.ReadInConfig();err!=nil {  // viper.ReadInConfig () 加载配置文件内容
		if _,ok:=err.(viper.ConfigFileNotFoundError);ok {
			log.Println("no such config file")
		}else {
			log.Panicln("read config error")
		}
		log.Fatal(err)//读取配置文件失败致命错误
	}
}

func main (){
	// viper.Get*** 获取配置文件中配置项的信息
    fmt.Println("获取配置文件的port", viper.GetInt("port"))
    fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.url`))
    fmt.Println("获取配置文件的mysql.username", viper.GetString(`mysql.username`))
    fmt.Println("获取配置文件的mysql.password", viper.GetString(`mysql.password`))
    fmt.Println("获取配置文件的redis", viper.GetStringSlice("redis"))
	fmt.Println("获取配置文件的smtp", viper.GetStringMap("smtp"))

	//viper 设置配置项的默认值
	viper.SetDefault("ContentDir","content") 
	viper.SetDefault("Taxonomies",map[string]string{"tag":"tags","category":"cateforyes"}) 

	fmt.Println(viper.GetBool("ContentDir"))
	fmt.Println(viper.GetString("ContentDir"))
	fmt.Println(viper.GetStringMapString("Taxonomies"))
	
}