package configure

import (
	"github.com/BurntSushi/toml"
)

// TOMLLoader loads config
type TOMLLoader struct {
}

// Load the configuration TOML file specified by path arg.
func (c TOMLLoader) Load(pathToToml string, conf any) error {
	// util.Log.Infof("Loading config: %s", pathToToml)
	if _, err := toml.DecodeFile(pathToToml, conf); err != nil {
		return err
	}

	ViperLoad(pathToToml, "toml", conf)

	return nil
}
