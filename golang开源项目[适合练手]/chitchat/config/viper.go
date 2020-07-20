package config

import (
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

var ViperConfig Configuration

func init() {
	runtimeViper := viper.New()
	runtimeViper.AddConfigPath(".")
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	err := runtimeViper.ReadInConfig() // // viper.ReadInConfig () 加载配置文件内容
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 再通过 runtimeViper.Unmarshal(&ViperConfig) 将其映射到之前定义的位于 config.go 中的 Configuration 结构体变量 ViperConfig
	runtimeViper.Unmarshal(&ViperConfig)

	// 本地化初始设置
	bundle := i18n.NewBundle(language.English)           // 创建一个用于应用程序生命周期的包。
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal) // 在初始化期间将翻译加载到包中
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active.en.json")
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	ViperConfig.LocaleBundle = bundle

	// 监听配置文件变更
	runtimeViper.WatchConfig()
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		runtimeViper.Unmarshal(&ViperConfig)
		ViperConfig.LocaleBundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	})
}
