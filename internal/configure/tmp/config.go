package tmp

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/mapstructure"
)

type AppConfig interface {
	Validate() error // 验证配置是否有效
	GetType() string // 返回配置类型，用于区分不同的应用配置
}

type Config struct {
	// 其他通用配置...
	Apps map[string]AppConfig // 使用 map 存储不同类型的应用配置
}

var Conf Config

func Init(cfgPath string) error {
	Conf.Apps = make(map[string]AppConfig)

	// 读取 TOML 文件
	var rawConfig map[string]interface{}
	if _, err := toml.DecodeFile(cfgPath, &rawConfig); err != nil {
		return fmt.Errorf("decode config file failed: %v", err)
	}

	// 检查 "apps" 部分是否存在
	appsConfig, ok := rawConfig["apps"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("no 'apps' section found in config file")
	}

	// 遍历 "apps" 部分，解析子配置
	for appName, appConfigData := range appsConfig {
		appConfigMap, ok := appConfigData.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid app config format for %s", appName)
		}
		switch appName {
		case "a":
			var aConfig AConfig
			if err := mapstructure.Decode(appConfigMap, &aConfig); err != nil {
				return fmt.Errorf("decode config for app 'a' failed: %v", err)
			}
			if err := aConfig.Validate(); err != nil {
				return fmt.Errorf("validate config for app 'a' failed: %v", err)
			}
			Conf.Apps["a"] = &aConfig
		case "b":
			var bConfig BConfig
			if err := mapstructure.Decode(appConfigMap, &bConfig); err != nil {
				return fmt.Errorf("decode config for app 'b' failed: %v", err)
			}
			if err := bConfig.Validate(); err != nil {
				return fmt.Errorf("validate config for app 'b' failed: %v", err)
			}
			Conf.Apps["b"] = &bConfig
		default:
			return fmt.Errorf("unknown app config: %s", appName)
		}
	}

	return nil
}

func GetAppConfig(appName string) (AppConfig, error) {
	if appConfig, ok := Conf.Apps[appName]; ok {
		return appConfig, nil
	}
	return nil, fmt.Errorf("app config for %s not found", appName)
}

/*
[apps]

[apps.a]
a_option1 = "value1"
a_option2 = 123

[apps.b]
b_option1 = true
b_option2 = ["item1", "item2"]
*/
