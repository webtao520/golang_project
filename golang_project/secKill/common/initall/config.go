package initall

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

type MysqlConfig struct {
	UserName string
	PassWd   string
	Port     int
	DbName   string
	Host     string
}

type EtcdConfig struct {
	Address     string
	PrefixKey   string
	ProductKey  string
	DailTimeOut int
	PutTimeOut  int
	GetTimeOut  int
}

type LogConfig struct {
	LogPath string
}

type ConfigAll struct {
	MysqlConfig
	EtcdConfig
	LogConfig
	AccessRedisConfig
	DisposeRedisConfig
	BlacklistRedisConfig
}

type AccessRedisConfig struct {
	Address           string
	MaxIdle           int
	MaxActive         int
	IdleTimeout       int
	WriteGoroutineNum int
	ReadGoroutineNum  int
	ListName          string
}
type DisposeRedisConfig struct {
	Address           string
	MaxIdle           int
	MaxActive         int
	IdleTimeout       int
	WriteGoroutineNum int
	ReadGoroutineNum  int
	ListName          string
}

type BlacklistRedisConfig struct {
	Address     string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

var conf config.Configer

// 初始化config
func InitConfig() (ConfigAll ConfigAll, err error) {
	conf, err = config.NewConfig("ini", "../common/conf/common.conf")
	if err != nil {
		err = errors.New(fmt.Sprintf("new config failed, err:", err))
		return
	}
	LogConfig, err := GetLogConfig()
	if err != nil {
		return
	}
	ConfigAll.LogConfig = LogConfig

	MysqlConfig, err := GetMysqlConfig()
	if err != nil {
		return
	}
	ConfigAll.MysqlConfig = MysqlConfig

	EtcdConfig, err := GetEtcdConfig()
	if err != nil {
		return
	}

	ConfigAll.EtcdConfig = EtcdConfig

	AccessRedisConfig, err := GetAccessRedisConfig()
	if err != nil {
		return
	}
	ConfigAll.AccessRedisConfig = AccessRedisConfig

	DisposeRedisConfig, err := GetDisposeRedisConfig()
	if err != nil {
		return
	}
	ConfigAll.DisposeRedisConfig = DisposeRedisConfig

	BlacklistRedisConfig, err := GetBlacklistRedisConfig()
	if err != nil {
		return
	}
	ConfigAll.BlacklistRedisConfig = BlacklistRedisConfig

	SecKillConf = ConfigAll
	logs.Debug("init config success")
	return
}

func GetLogConfig() (LogConfig LogConfig, err error) {
	LogPath := conf.String("logs::admin_log_path")
	if len(LogPath) == 0 {
		logs.Error("load config of admin_log_path failed, is null")
		err = errors.New("load config of admin_log_path failed, is null")
		return
	}
	LogConfig.LogPath = LogPath
	return
}

func GetMysqlConfig() (MysqlConfig MysqlConfig, err error) {
	UserName := conf.String("mysql::mysql_user_name")
	if len(UserName) == 0 {
		logs.Error("load config of mysql_user_name failed, is null")
		err = errors.New("load config of mysql_user_name failed, is null")
		return
	}
	MysqlConfig.UserName = UserName

	PassWd := conf.String("mysql::mysql_pass")
	if len(PassWd) == 0 {
		logs.Error("load config of mysql_pass failed, is null")
		err = errors.New("load config of mysql_pass failed, is null")
		return
	}
	MysqlConfig.PassWd = PassWd

	Port, err := conf.Int("mysql::mysql_port")
	if err != nil {
		logs.Error("load config of mysql_port failed, is err : ", err)
		err = errors.New(fmt.Sprintf("load config of mysql_port failed, is err : ", err))
		return
	}
	MysqlConfig.Port = Port

	Host := conf.String("mysql::mysql_host")
	if len(Host) == 0 {
		logs.Error("load config of mysql_host failed , is null")
		err = errors.New("load config of mysql_host failed , is null")
		return
	}
	MysqlConfig.Host = Host

	DbName := conf.String("mysql::mysql_db_name")
	if len(DbName) == 0 {
		logs.Error("load config of mysql_db_name failed , is null")
		err = errors.New("load config of mysql_db_name failed , is null")
		return
	}
	MysqlConfig.DbName = DbName

	return
}

func GetEtcdConfig() (EtcdConfig EtcdConfig, err error) {
	Address := conf.String("etcd::etcd_addr")
	if len(Address) == 0 {
		err = errors.New("load config of etcd_addr failed , is null")
		logs.Error(err)
		return
	}
	EtcdConfig.Address = Address

	PrefixKey := conf.String("etcd::etcd_prefix_key")
	if len(PrefixKey) == 0 {
		logs.Error("load config of etcd_prefix_key , is null")
		err = errors.New("load config of etcd_prefix_key , is null")
		return
	}
	EtcdConfig.PrefixKey = PrefixKey

	ProductKey := conf.String("etcd::etcd_product_key")
	if len(ProductKey) == 0 {
		logs.Error("load config of etcd_product_key , is null")
		err = errors.New("load config of etcd_product_key , is null")
		return
	}
	EtcdConfig.ProductKey = ProductKey

	DailTimeOut, err := conf.Int("etcd::etcd_dail_timeout")
	if err != nil {
		logs.Error("load config of etcd_dail_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_dail_timeout err : ", err))
		return
	}
	EtcdConfig.DailTimeOut = DailTimeOut

	PutTimeOut, err := conf.Int("etcd::etcd_put_timeout")
	if err != nil {
		logs.Error("load config of etcd_put_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_put_timeout err : ", err))
		return
	}
	EtcdConfig.PutTimeOut = PutTimeOut

	GetTimeOut, err := conf.Int("etcd::etcd_get_timeout")
	if err != nil {
		logs.Error("load config of etcd_get_timeout err : ", err)
		err = errors.New(fmt.Sprintf("load config of etcd_get_timeout err : ", err))
		return
	}
	EtcdConfig.GetTimeOut = GetTimeOut
	return
}

func GetAccessRedisConfig() (AccessRedisConfig AccessRedisConfig, err error) {
	Address := conf.String("redis::access_to_dispose_addr")
	if len(Address) == 0 {
		err = errors.New("load config of access_to_dispose_addr , is null")
		logs.Error(err)
		return
	}
	AccessRedisConfig.Address = Address
	MaxIdle, err := conf.Int("redis::access_to_dispose_MaxIdle")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of access_to_dispose_MaxIdle err : %v", err))
		logs.Error(err)
		return
	}
	AccessRedisConfig.MaxIdle = MaxIdle
	MaxActive, err := conf.Int("redis::access_to_dispose_MaxActive")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of access_to_dispose_MaxActive err : %v", err))
		logs.Error(err)
		return
	}
	AccessRedisConfig.MaxActive = MaxActive
	IdleTimeout, err := conf.Int("redis::access_to_dispose_IdleTimeout")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of access_to_dispose_IdleTimeout err : %v", err))
		logs.Error(err)
		return
	}
	AccessRedisConfig.IdleTimeout = IdleTimeout
	WriteGoroutineNum, err := conf.Int("redis::write_access_to_dispose_goroutine_num")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of write_access_to_dispose_goroutine_num err : %v", err))
		logs.Error(err)
		return
	}
	AccessRedisConfig.WriteGoroutineNum = WriteGoroutineNum
	ReadGoroutineNum, err := conf.Int("redis::read_dispose_to_access_goroutine_num")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of read_dispose_to_access_goroutine_num err : %v", err))
		logs.Error(err)
		return
	}
	AccessRedisConfig.ReadGoroutineNum = ReadGoroutineNum
	ListName := conf.String("redis::ccess_list_name")
	if len(ListName) == 0 {
		err = errors.New("load config of ccess_list_name failed , is null")
		logs.Error(err)
		return
	}
	AccessRedisConfig.ListName = ListName
	return
}

func GetDisposeRedisConfig() (DisposeRedisConfig DisposeRedisConfig, err error) {
	Address := conf.String("redis::redis_dispose_to_access_addr")
	if len(Address) == 0 {
		err = errors.New("load config of redis_dispose_to_access_addr , is null")
		logs.Error(err)
		return
	}
	DisposeRedisConfig.Address = Address
	MaxIdle, err := conf.Int("redis::redis_dispose_to_access_MaxIdle")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_dispose_to_access_MaxIdle err : %v", err))
		logs.Error(err)
		return
	}
	DisposeRedisConfig.MaxIdle = MaxIdle
	MaxActive, err := conf.Int("redis::redis_dispose_to_access_MaxActive")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_dispose_to_access_MaxActive err : %v", err))
		logs.Error(err)
		return
	}
	DisposeRedisConfig.MaxActive = MaxActive
	IdleTimeout, err := conf.Int("redis::redis_dispose_to_access_IdleTimeout")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_dispose_to_access_IdleTimeout err : %v", err))
		logs.Error(err)
		return
	}
	DisposeRedisConfig.IdleTimeout = IdleTimeout
	WriteGoroutineNum, err := conf.Int("redis::write_dispose_to_access_goroutine_num")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of write_dispose_to_access_goroutine_num err : %v", err))
		logs.Error(err)
		return
	}
	DisposeRedisConfig.WriteGoroutineNum = WriteGoroutineNum
	ReadGoroutineNum, err := conf.Int("redis::read_access_to_dispose_goroutine_num")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of read_access_to_dispose_goroutine_num err : %v", err))
		logs.Error(err)
		return
	}
	DisposeRedisConfig.ReadGoroutineNum = ReadGoroutineNum
	ListName := conf.String("redis::dispose_list_name")
	if len(ListName) == 0 {
		err = errors.New("load config of dispose_list_name failed , is null")
		logs.Error(err)
		return
	}
	DisposeRedisConfig.ListName = ListName
	return
}

func GetBlacklistRedisConfig() (BlacklistRedisConfig BlacklistRedisConfig, err error) {
	// Address     string
	// MaxIdle     int
	// MaxActive   int
	// IdleTimeout int
	Address := conf.String("redis::redis_blacklist_addr")
	if len(Address) == 0 {
		err = errors.New("load config of redis_blacklist_addr , is null")
		logs.Error(err)
		return
	}
	BlacklistRedisConfig.Address = Address
	MaxIdle, err := conf.Int("redis::redis_blacklist_MaxIdle")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_blacklist_MaxIdle err : %v", err))
		logs.Error(err)
		return
	}
	BlacklistRedisConfig.MaxIdle = MaxIdle
	MaxActive, err := conf.Int("redis::redis_blacklist_MaxActive")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_blacklist_MaxActive err : %v", err))
		logs.Error(err)
		return
	}
	BlacklistRedisConfig.MaxActive = MaxActive
	IdleTimeout, err := conf.Int("redis::redis_blacklist_IdleTimeout")
	if err != nil {
		err = errors.New(fmt.Sprintf("load config of redis_blacklist_IdleTimeout err : %v", err))
		logs.Error(err)
		return
	}
	BlacklistRedisConfig.IdleTimeout = IdleTimeout
	return
}
