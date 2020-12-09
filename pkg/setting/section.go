package setting

import "time"

type ServerSettingS struct {
	RunModel     string
	HttpPort     string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	ProjectName  string
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type AppSettingS struct {
	MaxPageSize     int
	DefaultPageSize int
	AllowDownload   bool
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}
