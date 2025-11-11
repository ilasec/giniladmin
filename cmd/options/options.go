package options

var Opt Options

type Options struct {
	Debug      bool
	ConfigPath string
	LogDir     string
	LogToFile  bool
}
