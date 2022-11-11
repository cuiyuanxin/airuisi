package setting

import (
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	env := os.Getenv("Env")
	if env == "" {
		env = "prod"
	}
	vp := viper.New()
	vp.AddConfigPath("configs/")
	switch env {
	case "dev":
		vp.SetConfigName("dev")
	case "prod":
		vp.SetConfigName("prod")
	case "test":
		vp.SetConfigName("test")
	default:
		vp.SetConfigName("prod")
	}
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
