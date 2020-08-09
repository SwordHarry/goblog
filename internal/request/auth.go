package request

type AuthRequest struct {
	AppKey    string `form:"app_key" json:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" json:"app_secret" binding:"required"`
}
