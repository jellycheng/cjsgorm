package cjsgorm

import (
	"fmt"
	"github.com/jellycheng/gosupport/dbutils"
	"testing"
)
//一个表一个结构体
type SystemModel struct {
	ID        uint64 `gorm:"primary_key;Column:id"`
	IsDelete  uint8  `gorm:"Column:is_delete;DEFAULT:0"`
	CreateTime uint64 `gorm:"Column:create_time"`
	UpdateTime uint64 `gorm:"Column:update_time"`
	DeleteTime uint64 `gorm:"Column:delete_time"`

	SystemCode string `gorm:"Column:system_code"`
	SystemName string `gorm:"Column:system_name"`
	Appid string `gorm:"Column:app_id"`
	Secret string `gorm:"Column:secret"`
}
//返回表名
func (SystemModel)TableName() string {
	return "t_system"
}
//go test --run="TestNewMysqlGorm222"
func TestNewMysqlGorm222(t *testing.T) {
	//数据库配置
	mysqlDsnObj := dbutils.NewMysqlDsn(map[string]interface{}{
		"host":"localhost",
		"username":"root",
		"password":"88888888",
		"port":3306,
		"dbname":"db_common",
		"charset":"utf8mb4",
		"extparam":"parseTime=True&loc=Local",
	})
	//打印dsn串
	fmt.Println(mysqlDsnObj.ToDsn())
	//根据db配置获取*gorm.DB对象
	gormDb := NewMysqlGorm(*mysqlDsnObj)
	//gorm设置 todo，如 log、debug、连接配置等

	//执行查询sql: SELECT * FROM `t_system`  WHERE (system_name='运营后台')
	var systemModel SystemModel
	gormDb.Debug().Where("system_name=?", "运营后台").Find(&systemModel)
	fmt.Println(fmt.Sprintf("%+v", systemModel))

}
