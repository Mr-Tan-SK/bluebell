package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Hour * 2

var secret = []byte("WorryFree")

type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 获取生成 access token
func GenToken(userID int64, username string) (aToken, rToken string, err error) {
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                 // 签发人
		},
	}).SignedString(secret)
	return
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	mc := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	// token有效，返回 MyClaim信息
	if token.Valid {
		return mc, nil
	}
	// 否则返回无效的token错误
	return nil, errors.New("无效的token")
}

// RefreshToken 刷新AccessToken
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, func(jt *jwt.Token) (interface{}, error) {
		return secret, nil
	}); err != nil {
		return
	}
	// 从旧access token中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(jt *jwt.Token) (interface{}, error) {

		return secret, nil
	})
	v, _ := err.(*jwt.ValidationError)
	// 当access token是过期错误 并且 refresh token没有过期时就创建⼀个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID, claims.Username)
	}
	return
}
