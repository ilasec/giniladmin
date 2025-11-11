package oauth

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login
// @Tags OAuth API
// @Summary auth login
// @Description auth login
// @Accept json
// @Produce json
// @Param x-appid header string true "token for authentication"
// @Param user body LoginModel true "login info"
// @Success 200 {object} models.CommonResp "{"message":"success","status":200}"
// @Failure 400 {object} models.CommonResp "{"message":"user email [mritd@linux.com] already register","status":400}"
// @Failure 500 {object} models.CommonResp "{"message":"invalid connection","status":400}"
// @Router /oauth/v1/login [post]
func Login(c *gin.Context) (status int, message string, ret any, err error) {
	param := LoginModel{}
	if err = c.ShouldBindJSON(&param); err != nil {
		return
	}

	appID := c.GetHeader("x-appid")
	if appID == "" {
		status = http.StatusBadRequest
		message = "APPID不能为空"
		return
	}

	key := c.ClientIP()

	return DoLogin(context.Background(), appID, key, param)
}

// Captcha
// @Tags      OAuth API
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param x-appid header string true "token for authentication"
// @Success   200  {object} models.CommonResp "{"message":"success","status":200}"  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /oauth/v1/captcha [get]
func Captcha(c *gin.Context) (status int, message string, ret any, err error) {
	key := c.ClientIP()
	appID := c.GetHeader("x-appid")
	if appID == "" {
		status = http.StatusBadRequest
		message = "APPID不能为空"
		return
	}

	return DoCaptcha(context.Background(), appID, key)
}
