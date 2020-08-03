package setting

import "time"

// 配置的结构体表示
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	UploadImageMaxSize    int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageAllowExts  []string
	DefaultContextTimeout time.Duration
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type EmailSettingS struct {
	Host     string
	UserName string
	Password string
	From     string
	Port     int
	IsSSL    bool
	To       []string
}

// 将配置文件中的部分 读取入结构体中
func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vp.UnmarshalKey(k, v)
}
