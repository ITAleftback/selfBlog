/**
 * @Author: Anpw
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2021/5/28 19:41
 */

package app

import (
	"github.com/dgrijalva/jwt-go"
	"selfblog/global"
	"selfblog/pkg/util"
	"time"
)

type Claims struct {
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetJWTSecret() string {
	return global.JWTSetting.Secret
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:         util.EncodeMD5(appKey),
		AppSecret:      util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:     expireTime.Unix(),
			Issuer: global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(GetJWTSecret()))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetJWTSecret()), nil
	})

	if tokenClaims !=nil {
		if Claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return Claims, nil
		}
	}

	return nil, err
}
