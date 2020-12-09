package init

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/summerKK/go-code-snippet-library/koel-api/global"
	"github.com/summerKK/go-code-snippet-library/koel-api/internal/model"
	"github.com/summerKK/go-code-snippet-library/koel-api/pkg/setting"
)

var (
	cfgPath = []string{"configs/"}
)

func SetConfig(configPath []string) {
	if configPath != nil {
		cfgPath = configPath
	}
}

func Boot() {
	var err error
	// 读取配置文件
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.Setting error:%v", err)
	}

	// 初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.DBEngine error:%v", err)
	}

	// 设置运行模式
	gin.SetMode(global.ServerSetting.RunModel)
}

func setupSetting() error {
	var err error
	SettingS, err := setting.NewSetting(cfgPath...)
	if err != nil {
		return err
	}

	err = SettingS.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = SettingS.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = SettingS.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = SettingS.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDbEngine(global.DatabaseSetting)

	return err
}
