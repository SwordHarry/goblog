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
	UploadMDMaxSize       int
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	UploadImageSavePath   string
	UploadImageServerUrl  string
	UploadMDSavePath      string
	UploadMDServerUrl     string
	UploadSavePath        string
	UploadMDAllowExts     []string
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

type TracerSettingS struct {
	ServiceName string
	HostPort    string
}

var sections = make(map[string]interface{})

// 将配置文件中的部分 读取入结构体中
func (s *Setting) ReadSection(k string, v interface{}) error {
	// 读取 k 对应的值到 v 中
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
