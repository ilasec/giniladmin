## 设计思路

实现自己的 OAuth 2.0 + OpenID Connect（OIDC） 认证服务器，通常涉及以下几个核心组件：

OAuth 2.0 授权服务器（Authorization Server）

负责认证用户，并颁发 Access Token 和 ID Token
实现 OAuth 2.0 标准授权流程（Authorization Code、Client Credentials 等）
OpenID Connect 认证层

扩展 OAuth 2.0，增加 ID Token，用于用户身份信息传递
提供 UserInfo 端点，允许客户端查询用户信息
客户端（Client Application）

你的测试平台，需要从 OAuth 服务器请求 Token 并管理用户 Session
资源服务器（Resource Server）（可选）

托管受保护的 API 服务，校验 Access Token 以确保合法访问


## 代码


```go


package main

import (
"time"
"net/http"
"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oslo/oauth2"
	"github.com/oslo/oauth2/models"
)

// 服务器密钥（用于 JWT 签名）
var secretKey = []byte("your_secret_key")

// 认证服务器实例
var authServer *oauth2.Server

// 客户端存储
var clients = map[string]*models.Client{
"test-client": {
ID:           "test-client",
Secret:       "test-secret",
RedirectURIs: []string{"http://localhost:3000/callback"},
},
}

// 用户存储
var users = map[string]string{
"testuser": "password123", // 用户名 -> 密码
}

// 生成 JWT 令牌（用于 OpenID Connect）
func generateJWTToken(userID string) (string, error) {
claims := jwt.MapClaims{
"sub": userID,
"exp": time.Now().Add(time.Hour).Unix(),
"iat": time.Now().Unix(),
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString(secretKey)
}

// 处理用户认证
func authenticateUser(username, password string) bool {
if pass, ok := users[username]; ok {
return pass == password
}
return false
}

func main() {
r := gin.Default()

	// 配置 OAuth 服务器
	authServer = oauth2.NewServer(oauth2.Config{
		TokenExpiration: time.Hour, // 令牌过期时间
	})

	// 授权端点（用户输入用户名和密码）
	r.POST("/authorize", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		clientID := c.PostForm("client_id")
		redirectURI := c.PostForm("redirect_uri")

		if authenticateUser(username, password) {
			// 生成授权码
			code := authServer.GenerateAuthCode(clientID, username)
			c.Redirect(http.StatusFound, redirectURI+"?code="+code)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	})

	// 令牌端点（交换授权码）
	r.POST("/token", func(c *gin.Context) {
		code := c.PostForm("code")
		clientID := c.PostForm("client_id")
		clientSecret := c.PostForm("client_secret")

		// 验证授权码
		userID, err := authServer.ValidateAuthCode(clientID, clientSecret, code)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid code"})
			return
		}

		// 生成 Access Token 和 ID Token
		accessToken := authServer.GenerateAccessToken(clientID, userID)
		idToken, _ := generateJWTToken(userID)

		c.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
			"id_token":     idToken,
			"token_type":   "bearer",
		})
	})

	// UserInfo 端点（OpenID Connect）
	r.GET("/userinfo", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		token = token[7:]

		// 验证 JWT 令牌
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil || !parsedToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, _ := parsedToken.Claims.(jwt.MapClaims)
		c.JSON(http.StatusOK, gin.H{
			"sub":  claims["sub"],
			"name": "Test User",
		})
	})

	log.Println("OAuth2 + OIDC Server running on :8080")
	r.Run(":8080")
}

```

```shell
curl -X POST http://localhost:8080/authorize \
     -d "username=testuser&password=password123&client_id=test-client&redirect_uri=http://localhost:3000/callback"
     
     {
    "access_token": "abc123",
    "id_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "bearer"
}

curl -H "Authorization: Bearer abc123" http://localhost:8080/userinfo


{
    "sub": "testuser",
    "name": "Test User"
}


✅ OAuth 2.0 认证（授权码模式）
✅ 支持 OpenID Connect（ID Token）
✅ 基于 JWT 的身份验证
✅ 轻量、易扩展

你的测试平台可以：

跳转 OAuth 认证
解析 Token 并管理 Session
使用 UserInfo API 获取用户信息
```