package configure

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ViperLoad is monitor configuration
func ViperLoad(p string, t string, c any) error {
	v := viper.New()
	v.SetConfigFile(p)
	v.SetConfigType(t)
	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(c); err != nil {
		}
	})
	if err = v.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
