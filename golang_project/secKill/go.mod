module secKill

go 1.14

require (
	github.com/astaxie/beego v1.12.2
	github.com/beego/admin v0.0.0-20171002082758-83609ddd6f2d
	github.com/coreos/etcd v3.3.22+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1 // indirect
	go.uber.org/zap v1.15.0 // indirect
	google.golang.org/genproto v0.0.0-20200701001935-0939c5918c31 // indirect
	google.golang.org/grpc v1.30.0 // indirect
)

replace (
    google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.26.0
)