package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type mysqlConfigPool struct {
	Master *gorm.DB
	Slaves []*gorm.DB
}

var (
	defaultConnectPool map[string]mysqlConfigPool
)

type GormConfig struct {
	Dns          string
	DbName       string
	MaxIdleConns int
	MaxOpenConns int
}

// type MysqlConfig

func newConn(cfg GormConfig) *gorm.DB {
	var err error
	db, err = gorm.Open("mysql", cfg.Dns)
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	db.DB().SetMaxOpenConns(cfg.MaxOpenConns)

	db.LogMode(true)
	// db.DB().SetConnMaxLifetime(20)
	return db
}

// func CheckConnInit() {

// }

func checkInit(cfgs []GormConfig) {
	// masterFlag := md5(fmt.Sprintf("%s-%s", cfgs.Master.DbName, "Master"))
	// slaveFlag := fmt.Sprintf("%s-%s", cfgs.Master, "Slave")

}

//主从配置 待优化
func MasterConn(cfg GormConfig) *gorm.DB {
	return newConn(cfg)
}

// func SlaveConn(cfg []GormConfig) *gorm.DB {

// }
