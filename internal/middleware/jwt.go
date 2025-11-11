package middleware

import (
	"errors"
	"giniladmin/pkg/utils"
	"giniladmin/pkg/utils/cache"
	"giniladmin/pkg/utils/jtoken"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	JwtBlackCache *cache.Cache
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := jtoken.GetToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]any{})
			c.Abort()
			return
		}
		if JwtBlackCache != nil {
			_, ok := JwtBlackCache.Get(token)
			if ok {
				c.JSON(http.StatusUnauthorized, map[string]any{})
				jtoken.ClearToken(c)
				c.Abort()
				return
			}

		}
		j := jtoken.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if errors.Is(err, jtoken.TokenExpired) {
				c.JSON(http.StatusUnauthorized, map[string]any{})
				jtoken.ClearToken(c)
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, map[string]any{})
			jtoken.ClearToken(c)
			c.Abort()
			return
		}

		// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
		// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

		//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
		//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
		//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		//	c.Abort()
		//}
		c.Set("claims", claims)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(jtoken.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			jtoken.SetToken(c, newToken, int(dr.Seconds()))
		}
		c.Next()

		if newToken, exists := c.Get("new-token"); exists {
			c.Header("new-token", newToken.(string))
		}
		if newExpiresAt, exists := c.Get("new-expires-at"); exists {
			c.Header("new-expires-at", newExpiresAt.(string))
		}
	}
}
