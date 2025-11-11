package jwtutil

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JWTManager 结构体
type JWTManager struct {
	SecretKey     string        // 签名密钥
	Issuer        string        // 发行者
	TokenExpiry   time.Duration // Token 过期时间
	RefreshExpiry time.Duration // Refresh Token 过期时间
}

// NewJWTManager 创建 JWT 管理器
func NewJWTManager(secretKey, issuer string, tokenExpiry, refreshExpiry time.Duration) *JWTManager {
	return &JWTManager{
		SecretKey:     secretKey,
		Issuer:        issuer,
		TokenExpiry:   tokenExpiry,
		RefreshExpiry: refreshExpiry,
	}
}

// CreateToken 生成 JWT Token
func (jm *JWTManager) CreateToken(claimsData map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}

	// 复制用户传递的 claims 数据
	for key, value := range claimsData {
		claims[key] = value
	}

	// 设置 JWT 标准字段
	now := time.Now()
	claims["iss"] = jm.Issuer
	claims["iat"] = now.Unix()                     // 签发时间
	claims["exp"] = now.Add(jm.TokenExpiry).Unix() // 过期时间

	// 生成 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jm.SecretKey))
}

// ParseToken 解析 JWT Token
func (jm *JWTManager) ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jm.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 确保 Token 有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result := make(map[string]interface{})
		for key, value := range claims {
			result[key] = value
		}
		return result, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新 JWT 令牌
func (jm *JWTManager) RefreshToken(tokenString string) (string, error) {
	claims, err := jm.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// 更新 `exp` 过期时间
	claims["exp"] = time.Now().Add(jm.RefreshExpiry).Unix()

	// 重新生成 Token
	return jm.CreateToken(claims)
}
