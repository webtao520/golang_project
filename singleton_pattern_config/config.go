/*
     初始化配置文件代码
*/
package  singleton_pattern_config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	Address string
	Static string 
	Log string 
}

type Database struct {
	 Driver string
	 Address string
	 Database string
	 User    string
	 Password string
}

type  Configuration struct {
	 App  App
	 Db  Database
 }

var config *Configuration
var once  sync.Once  //   执行一次动作的对象

// 通过单例模式初始化全局配置
func LoadConfig () *Configuration{
	 once.Do(func (){
		  file,err:=os.Open("config.json")
		  
	 })
}



