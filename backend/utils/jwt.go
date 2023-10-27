package utils

import (
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 签名密钥
var jwtSecretKey = []byte("This is sparks")

// GenerateToken 创建token， 过期时间设置已注释
func GenerateToken(userId int64, username string) string {

	nowTime := time.Now()

	claims := Claims{
		ID:       userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: nowTime.Unix(),
			Issuer:   "Sparks",
		},
	}
	// 使用用于签名的算法和令牌
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建JWT字符串
	if token, err := tokenClaims.SignedString(jwtSecretKey); err != nil {
		zap.L().Error("generate token fail!", zap.Error(err))
		return "fail"
	} else {
		zap.L().Info("generate token success!")
		return token
	}
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
