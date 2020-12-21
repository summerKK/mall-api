package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/pkg/setting"
	"github.com/summerKK/mall-api/pkg/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var zeroTime = time.Time{}

type ID struct {
	Id uint `json:"id" gorm:"primaryKey"`
}

// 数据库时间格式化问题
type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	if t.Time == zeroTime {
		return []byte(`""`), nil
	}

	formatted := fmt.Sprintf("\"%v\"", time.Time(t.Time).Format(util.TimeLayout))
	return []byte(formatted), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 初始化数据库
func NewDbEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	format := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	dns := fmt.Sprintf(format,
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// 开发者模式,打印SQL
	if global.ServerSetting.RunModel == "debug" {
		db = db.Debug()
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池最大连接数和空闲数
	sqlDb.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	sqlDb.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	return db, nil
}
