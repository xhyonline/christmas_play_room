module github.com/xhyonline/christmas_play_room

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/gin-gonic/gin v1.7.4
	github.com/gogo/protobuf v1.3.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/xhyonline/xutil v0.1.2021111118
	go.etcd.io/etcd v3.3.27+incompatible
	google.golang.org/grpc v1.38.0
	gorm.io/gorm v1.21.16
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
