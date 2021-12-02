package cjsgorm

import (
	"github.com/jellycheng/gosupport/dbutils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type MysqlGormInstance struct {
	mysql  map[string]*gorm.DB
	lock sync.RWMutex
}

// 获取*gorm.DB对象
func (mysqlInstance *MysqlGormInstance) GetMysql(dsnKey string) *gorm.DB {
	mysqlInstance.lock.RLock()
	defer mysqlInstance.lock.RUnlock()
	if d, ok := mysqlInstance.mysql[dsnKey]; ok {
		return d
	}
	return nil

}

func (mysqlInstance *MysqlGormInstance) registerMysql(dsn string, db *gorm.DB) *gorm.DB {
	mysqlInstance.lock.Lock()
	defer mysqlInstance.lock.Unlock()
	mysqlInstance.mysql[dsn] = db
	return db
}


var mysqlGormObj = MysqlGormInstance{
	mysql: make(map[string]*gorm.DB),
}

// 实例化,使用默认配置
func NewMysqlGorm(mysqlDsn dbutils.MysqlDsn) *gorm.DB {
	if d := mysqlGormObj.GetMysql(mysqlDsn.Key()); d != nil {
		return d
	}
	//实例化
	if d, err := gorm.Open(mysql.Open(mysqlDsn.ToDsn()), &gorm.Config{}); err != nil {
		MyGormLogObj.Printf("connect mysql error: " + err.Error())
		return nil
	} else {//注册
		return mysqlGormObj.registerMysql(mysqlDsn.Key(), d)
	}

}

// 实例化,通过外部传入配置
func NewMysqlGormV2(mysqlDsn dbutils.MysqlDsn, gormCfg *gorm.Config) *gorm.DB {
	if d := mysqlGormObj.GetMysql(mysqlDsn.Key()); d != nil {
		return d
	}
	if d, err := gorm.Open(mysql.Open(mysqlDsn.ToDsn()), gormCfg); err != nil {
		MyGormLogObj.Printf("connect mysql error: " + err.Error())
		return nil
	} else {
		return mysqlGormObj.registerMysql(mysqlDsn.Key(), d)
	}

}
