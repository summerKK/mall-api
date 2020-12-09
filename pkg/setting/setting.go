package setting

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

// 配置文件分类
var sections = make(map[string]interface{})

func NewSetting(paths ...string) (*Setting, error) {
	vp := viper.New()
	// 设置配置名
	vp.SetConfigName("config")
	// 设置文件格式
	vp.SetConfigType("yaml")
	// 配置文件目录
	for _, path := range paths {
		if path != "" {
			vp.AddConfigPath(path)
		}
	}

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp: vp}
	// 检测文件修改
	s.WatchSettingChange()

	return s, nil
}

// 配置文件修改刷新配置项
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			// 重新读取所有配置
			err := s.ReadAllSection()
			if err != nil {
				log.Errorf("s.ReadAllSection error:%v", err)
				return
			}
			log.Println("配置文件已更新")
		})
	}()
}

func (s *Setting) ReadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	// 根据 key ,把数据读到 v 中
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; ok {
		sections[k] = v
	}

	return nil
}
