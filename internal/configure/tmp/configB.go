package tmp

import "fmt"

type BConfig struct {
	BOption1 bool     `toml:"b_option1" mapstructure:"b_option1"`
	BOption2 []string `toml:"b_option2" mapstructure:"b_option2"`
}

func (c *BConfig) Validate() error {
	if len(c.BOption2) == 0 {
		return fmt.Errorf("BOption2 is required")
	}
	return nil
}

func (c *BConfig) GetType() string {
	return "b"
}
