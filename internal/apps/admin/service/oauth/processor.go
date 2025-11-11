package oauth

import (
	"context"
	"errors"
	"giniladmin/internal/apps/admin/service/system/application"
	"giniladmin/internal/apps/admin/service/system/users"
	jwtutil "giniladmin/pkg/utils/jwtutils"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

var (
	store = base64Captcha.DefaultMemStore
)

type Processor struct {
}

func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

func DoCaptcha(ctx context.Context, appId, key string) (status int, message string, ret any, err error) {
	_, err = application.CheckAppId(context.Background(), appId)
	if err != nil {
		logger.Errorf("无效的appId", zap.Error(err))
		// 验证码次数+1
		blackCache.Increment(key, 1)
		return 0, "", nil, err
	}

	openCaptcha := conf.Captcha.OpenCaptcha               // 判断验证码是否开启
	openCaptchaTimeOut := conf.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := blackCache.Get(key)
	if !ok {
		blackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(conf.Captcha.ImgHeight, conf.Captcha.ImgWidth, conf.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logger.Errorf("验证码获取失败! %v", err)
		return
	}
	ret = map[string]any{
		"CaptchaId":     id,
		"PicPath":       b64s,
		"CaptchaLength": conf.Captcha.KeyLong,
		"OpenCaptcha":   oc,
	}
	return
}

func DoLogin(ctx context.Context, appId, key string, param LoginModel) (status int, message string, ret any, err error) {

	openCaptcha := conf.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := conf.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := blackCache.Get(key)
	if !ok {
		blackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if oc {
		err = errors.New("账号锁定")
		return
	}

	if !oc && (param.CaptchaId != "" && param.Captcha != "" && store.Verify(param.CaptchaId, param.Captcha, true)) {
		u := &users.UserModel{Username: param.Username, Password: param.Password}
		securityKey, err := application.CheckAppId(context.Background(), appId)
		if err != nil {
			logger.Errorf("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			blackCache.Increment(key, 1)
			return 0, "", nil, err
		}

		user, err := users.Login(u)
		if err != nil {
			logger.Errorf("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			blackCache.Increment(key, 1)
			return 0, "", nil, err
		}

		if user.Enable != 1 {
			err = errors.New("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			blackCache.Increment(key, 1)
			return 0, "", nil, err
		}
		ret, err = tokenNext(user, securityKey)
		return 0, "success", ret, nil
	}
	// 验证码次数+1
	blackCache.Increment(key, 1)
	err = errors.New("验证码错误")
	return
}

// TokenNext 登录以后签发jwt
func tokenNext(user users.UserModel, key string) (ret any, err error) {

	jwtManager := jwtutil.NewJWTManager(key, "oauth", 15*time.Minute, 24*time.Hour)
	// 创建 Claims 数据
	claimsData := map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
	}

	// 生成 JWT Token
	token, err := jwtManager.CreateToken(claimsData)
	if err != nil {
		logger.Printf("生成 Token 失败:", err)
		return
	}


	ret = map[string]any{
		"user": map[string]any{
			"username": user.Username,
			"id":       user.ID,
		},
		"token": token,
	}
	return
}
