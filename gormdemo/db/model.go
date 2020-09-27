package db

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	DB string
}

func (m *Model) Master() *gorm.DB {
	//读取配置
	cfg := GormConfig{
		Dns:          "root:123456@tcp(192.168.56.106:3306)/test?charset=utf8&parseTime=True&loc=Local",
		MaxIdleConns: 10,
		MaxOpenConns: 20,
	}
	// fmt.Print(MasterConn(cfg))
	return MasterConn(cfg)
}

func (m *Model) CheckError(err error) bool {
	if err != nil && err != gorm.ErrRecordNotFound {
		return true
	}
	return false
}
