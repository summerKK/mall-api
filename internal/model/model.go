package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/summerKK/mall-api/global"
	"github.com/summerKK/mall-api/pkg/setting"
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

	formatted := fmt.Sprintf("\"%v\"", time.Time(t.Time).Format("2006-01-02 15:04:05"))
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
	db, err := gorm.Open(dbSetting.DBType, fmt.Sprintf(format,
		dbSetting.Username,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime,
	))

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunModel == "debug" {
		db.LogMode(true)
	}

	// 设置连接池最大连接数和空闲数
	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	return db, nil
}
