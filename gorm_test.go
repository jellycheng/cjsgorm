package cjsgorm

import (
	"fmt"
	"github.com/jellycheng/gosupport/dbutils"
	"github.com/jellycheng/gosupport/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"testing"
	"time"
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
// 返回表名
func (SystemModel)TableName() string {
	return "t_system"
}

// go test --run="TestNewMysqlGormV1"
func TestNewMysqlGormV1(t *testing.T) {
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


// go test --run="TestNewMysqlGormV2"
func TestNewMysqlGormV2(t *testing.T) {
	gormDb := MyMasterDb()

	//执行查询sql: SELECT * FROM `t_system`  WHERE (system_name='运营后台')
	var systemModel SystemModel
	gormDb.Where("system_name=?", "运营后台").Find(&systemModel)
	fmt.Println(fmt.Sprintf("%+v", systemModel))

}

func MyMasterDb() *gorm.DB {
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
	tmpLogObj := new(MyGormLogger)
	//根据db配置获取*gorm.DB对象
	myLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), //写日志接口
		tmpLogObj,
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, //gorm日志级别：Silent > Error > Warn > Info
			Colorful:      false,         // 禁用彩色打印
		},
	)
	//myLogger = myLogger.LogMode(logger.Error) //二次修改日志级别
	MyGormLogObj = tmpLogObj
	gormDb := NewMysqlGormV2(*mysqlDsnObj, &gorm.Config{Logger: myLogger,})
	//gorm设置 todo，如 log、debug、连接配置等

	return gormDb

}

// go test --run="TestNewMysqlGormGoroutine"
func TestNewMysqlGormGoroutine(t *testing.T) {
	wg := utils.WaitGroupWrapper{}
	for i:=0;i<100;i++ {
		wg.Wrap(func() {
			gormDb := MyMasterDb()
			if gormDb == nil {
				os.Exit(1)
				return
			}
			//执行查询sql: SELECT * FROM `t_system`  WHERE (system_name='运营后台')
			var systemModel SystemModel
			gormDb.Where("system_name=?", "运营后台").Find(&systemModel)
			fmt.Println(fmt.Sprintf("第1个goroutine： %+v",systemModel))
		})
	}

	for i:=0;i<100;i++ {
		wg.Wrap(func() {
			gormDb := MyMasterDb()
			if gormDb == nil {
				os.Exit(1)
				return
			}
			//执行查询sql: SELECT * FROM `t_system`  WHERE (system_name='运营后台')
			var systemModel SystemModel
			gormDb.Where("system_name=?", "运营后台").Find(&systemModel)
			fmt.Println(fmt.Sprintf("第2个goroutine： %+v", systemModel))
		})
	}

	wg.Wait()

}

type MyGormLogger struct{}
func (MyGormLogger) Printf(str string, values ...interface{}) {
	format := strings.Replace(str, "\n", " ", -1)
	//fmt.Println("str=", format)
	//fmt.Printf("%#v \n", values)
	fmt.Printf(format + "\n", values...)
	//logrus.Print(gorm.LogFormatter(values...))
}
