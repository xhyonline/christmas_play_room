package component

import (
	"sync"

	"gorm.io/gorm"

)

// Server 组件服务
type Server struct {
	MySQL *gorm.DB
}

var (
	Instance *Server
	once     sync.Once
)

type Option func()

// Init 初始化组建服务
func Init(options ...Option) {
	once.Do(func() {
		Instance = new(Server)
		for _, f := range options {
			f()
		}
	})
}
