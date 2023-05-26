package util

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTConfig struct {
	Secret []byte //秘钥
}

type CustomClaims struct {
	Id   uint
	Name string
	jwt.RegisteredClaims
}

var (
	config JWTConfig
)

func init() {
	config = JWTConfig{
		Secret: []byte("study-group"),
	}
}

// GenerateTokenWithHS256 使用HS256算法生成token
func GenerateTokenWithHS256(name string, id uint) (string, error) {
	claims := CustomClaims{
		Id:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "ERIC",
			ExpiresAt: &jwt.NumericDate{
				time.Now(),
			},
			NotBefore: &jwt.NumericDate{
				time.Now(),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.Secret)
	return tokenString, err
}

// ParseToken 解析token，返回用户的id
func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Secret, nil
	})
	//如果token解析成功
	if token.Valid {
		return token.Claims.(*CustomClaims).Id, nil
	} else {
		//解析出现错误, 直接返回错误，在上层函数处理相关的错误
		return 0, err
	}
}

// RefreshToken  刷新token，token一旦签发是不可变的，所以需要重新定义结构体进行生成
func RefreshToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.Secret, nil
	})
	if err != nil {
		return "", nil
	}
	claims := token.Claims.(*CustomClaims)
	return GenerateTokenWithHS256(claims.Name, claims.Id)
}

// IsValidTokenString 验证前端传过来的token是否是有效的
func IsValidTokenString(tokenString string) bool {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false
	} else {
		return true
	}
}
