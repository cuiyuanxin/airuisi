package setting

import "time"

// 服务地址和端口
type Server struct {
	GinMode      string        `mapstructure:"GinMode"`
	HttpAddr     string        `mapstructure:"HttpAddr"`
	HttpPort     string        `mapstructure:"HttpPort"`
	ReadTimeout  time.Duration `mapstructure:"ReadTimeout"`
	WriteTimeout time.Duration `mapstructure:"WriteTimeout"`
}

// 数据库
//type Database struct {
//	DriverName   string   `mapstructure:"DriverName"`
//	Conns        []string `mapstructure:"Conns"`
//	TZLocation   string   `mapstructure:"TZLocation"`
//	MaxIdleConns int      `mapstructure:"MaxIdleConns"`
//	MaxOpenConns int      `mapstructure:"MaxOpenConns"`
//}

// 日志配置
type Logger struct {
	FilePath   string `mapstructure:"FilePath"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxBackups int    `mapstructure:"MaxBackups"`
	MaxAge     int    `mapstructure:"MaxAge"`
	LocalTime  bool   `mapstructure:"LocalTime"`
	Compress   bool   `mapstructure:"Compress"`
}

// App通用配置
type App struct {
	Limit                 int           `mapstructure:"Limit"`
	IsDel                 bool          `mapstructure:"IsDel"`
	UploadSavePath        string        `mapstructure:"UploadSavePath"`
	UploadServerUrl       string        `mapstructure:"UploadServerUrl"`
	UploadImageMaxSize    int           `mapstructure:"UploadImageMaxSize"`
	UploadImageAllowExts  []string      `mapstructure:"UploadImageAllowExts"`
	UploadFileMaxSize     int           `mapstructure:"UploadFileMaxSize"`
	UploadFileAllowExts   []string      `mapstructure:"UploadFileAllowExts"`
	DefaultContextTimeout time.Duration `mapstructure:"DefaultContextTimeout"`
	CryptKey              string        `mapstructure:"CryptKey"`
	InfoValidTime         int           `mapstructure:"InfoValidTime"`
	QrCryptLength         int           `mapstructure:"QrCryptLength"`
}

// jwt
//type JWT struct {
//	AppKey    string        `mapstructure:"AppKey"`
//	Secret    string        `mapstructure:"Secret"`
//	Issuer    string        `mapstructure:"Issuer"`
//	ExpiresAt time.Duration `mapstructure:"ExpiresAt"`
//}

// Tracer
type Tracer struct {
	AgentHostPort string `mapstructure:"AgentHostPort"`
}

// Email
type Email struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
