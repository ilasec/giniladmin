package auth

import (
	"context"
	"github.com/gin-gonic/gin"
)

// Login
// @Tags Auth API
// @Summary auth login
// @Description auth login
// @Accept json
// @Produce json
// @Param user body LoginModel true "login info"
// @Success 200 {object} models.CommonResp "{"message":"success","status":200}"
// @Failure 400 {object} models.CommonResp "{"message":"user email [mritd@linux.com] already register","status":400}"
// @Failure 500 {object} models.CommonResp "{"message":"invalid connection","status":400}"
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) (status int, message string, ret any, err error) {
	param := LoginModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	key := c.ClientIP()

	return DoLogin(context.Background(), key, param)
}

// Captcha
// @Tags      Auth API
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /api/v1/auth/captcha [get]
func Captcha(c *gin.Context) (status int, message string, ret any, err error) {
	key := c.ClientIP()
	return DoCaptcha(context.Background(), key)
}
