package dao

type Auth struct {
	*Common
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a *Auth) TableName() string {
	return "blog_auth"
}
