package app

import (
	"github.com/dgrijalva/jwt-go"
	"goblog/global"
	"goblog/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return global.JWTSetting.Secret
}

// 生成最终的 JWT token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// jwt.NewWithClaims 根据结构体生成 token 实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名字符串
	// key 需要是 []byte 类型
	return tokenClaims.SignedString([]byte(GetJWTSecret()))
}

// 解析和校验 token
func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 用于解析鉴权的声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetJWTSecret()), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
