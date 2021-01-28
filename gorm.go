package cjsgorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/jellycheng/gosupport/dbutils"
	"sync"
)

type MysqlGormInstance struct {
	mysql  map[string]*gorm.DB
	lock sync.RWMutex
}
//公有
func (mysqlInstance *MysqlGormInstance) GetMysql(dsnKey string) *gorm.DB {
	mysqlInstance.lock.Lock()
	defer mysqlInstance.lock.Unlock()
	if d, ok := mysqlInstance.mysql[dsnKey]; ok {
		return d
	}
	return nil

}

//私有
func (mysqlInstance *MysqlGormInstance) registerMysql(dsn string, db *gorm.DB) *gorm.DB {
	mysqlInstance.lock.Lock()
	defer mysqlInstance.lock.Unlock()
	mysqlInstance.mysql[dsn] = db
	return db
}


var mysqlGormObj = MysqlGormInstance{}
//实例化
func NewMysqlGorm(mysqlDsn dbutils.MysqlDsn) *gorm.DB {
	if mysqlGormObj.mysql == nil{
		mysqlGormObj.mysql = make(map[string]*gorm.DB)
	}
	if d := mysqlGormObj.GetMysql(mysqlDsn.Key()); d != nil {
		return d
	}
	//实例化
	if d, err := gorm.Open(mysql.Open(mysqlDsn.ToDsn()), &gorm.Config{}); err != nil {
		fmt.Println("connect mysql error: " + err.Error())
		return nil
	} else {//注册
		return mysqlGormObj.registerMysql(mysqlDsn.Key(), d)
	}

}

