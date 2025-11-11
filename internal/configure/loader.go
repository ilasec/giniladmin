package configure

// Load loads configuration
func Load(path string, conf any) error {
	loader := TOMLLoader{}
	return loader.Load(path, conf)
}

// Loader is interface of concrete loader
type Loader interface {
	Load(string, string) error
}
