package tmp

import "fmt"

type AConfig struct {
	AOption1 string `toml:"a_option1" mapstructure:"a_option1"`
	AOption2 int    `toml:"a_option2" mapstructure:"a_option2"`
}

func (c *AConfig) Validate() error {
	if c.AOption1 == "" {
		return fmt.Errorf("AOption1 is required")
	}
	return nil
}

func (c *AConfig) GetType() string {
	return "a"
}
