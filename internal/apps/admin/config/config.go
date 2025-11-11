package config

import (
	"giniladmin/internal/configure"
)

type Config struct {
	Server   Server             `toml:"server"`
	Log      Log                `toml:"log,omitempty" json:"log,omitempty"`
	Backend  Backend            `toml:"backend,omitempty" json:"backend,omitempty"`
	Jwt      Jwt                `toml:"jwt,omitempty" json:"jwt,omitempty"`
	Captcha  Captcha            `toml:"capcha,omitempty" json:"capcha,omitempty"`
	DataBase configure.DataBase `toml:"db,omitempty" json:"db,omitempty"`
}

// Server .
type Server struct {
	Debug     bool   `toml:"debug,omitempty" json:"debug,omitempty"`
	Port      string `toml:"port,omitempty" json:"port,omitempty"`
	SessionDb string `toml:"sessiondb,omitempty" json:"sessiondb,omitempty"`
}

type Log struct {
	Dir string `toml:"dir,omitempty" json:"dir,omitempty"`
}

type Backend struct {
	ApiKey string `toml:"apikey,omitempty" json:"apikey,omitempty"`
	Url    string `toml:"url,omitempty" json:"url,omitempty"`
}

type Jwt struct {
	SigningKey  string `toml:"signing-key,omitempty" json:"signing-key,omitempty"`
	ExpiresTime string `toml:"expires-time,omitempty" json:"expires-time,omitempty"`
	BufferTime  string `toml:"buffer-time,omitempty" json:"buffer-time,omitempty"`
	Issuer      string `toml:"issuer,omitempty" json:"issuer,omitempty"`
}

type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // 验证码长度
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // 验证码宽度
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // 验证码高度
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}
