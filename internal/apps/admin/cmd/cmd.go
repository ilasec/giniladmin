package cmd

import (
	"context"
	"flag"
	"fmt"
	"giniladmin/cmd/options"
	"giniladmin/constants"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/apps/admin/core"
	"giniladmin/internal/configure"
	"giniladmin/pkg/logging"
	"giniladmin/pkg/utils"
	"github.com/google/subcommands"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Cmd is Subcommand of host discovery mode
type Cmd struct {
}

// Name return subcommand name
func (*Cmd) Name() string { return "admin" }

// Synopsis return synopsis
func (*Cmd) Synopsis() string { return "admin..." }

// Usage return usage
func (*Cmd) Usage() string {
	return `server:
	server
		[-config=/path/to/config.toml]
		[-log-to-file=/path/to/log]
`
}

// SetFlags set flag
func (p *Cmd) SetFlags(f *flag.FlagSet) {
	wd, _ := os.Getwd()

	defaultConfPath := filepath.Join(wd, "config.toml")
	f.BoolVar(&options.Opt.Debug, "debug", false, "debug mode")
	f.StringVar(&options.Opt.ConfigPath, "config", defaultConfPath, "/path/to/toml")

	defaultLogDir := filepath.Join(wd, "logs")
	f.StringVar(&options.Opt.LogDir, "log-dir", defaultLogDir, "/path/to/log")
	f.BoolVar(&options.Opt.LogToFile, "log-to-file", true, "Output log to file")

}

// Execute execute
func (p *Cmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	logging.Log = logging.NewCustomLogger(false, false, options.Opt.LogToFile, options.Opt.LogDir, "", constants.Name)
	logging.Log.Infof("%s-%s-%s", constants.Name, constants.Version, constants.Revision)

	var conf config.Config

	if !utils.CheckFileExit(options.Opt.ConfigPath) {
		options.Opt.ConfigPath = "config/config.sample.toml"
	}

	if err := configure.Load(options.Opt.ConfigPath, &conf); err != nil {
		msg := []string{
			fmt.Sprintf("Error loading %s", options.Opt.ConfigPath),
			"If you update server and get this error, there may be incompatible changes in config.toml",
			"Please check server.toml template : ",
		}
		logging.Log.Errorf("%s\n%+v", strings.Join(msg, "\n"), err)
		return subcommands.ExitUsageError
	}

	if options.Opt.Debug {
		conf.Server.Debug = true
	}

	logging.Log.Info("Start running")
	logging.Log.Infof("config: %s", options.Opt.ConfigPath)
	logging.Log.Infof("log: %s", options.Opt.LogDir)

	value := map[string]any{
		"log":  &logging.Log,
		"conf": &conf,
	}
	ctx = context.WithValue(ctx, "value", value)

	log.Fatal(core.Run(ctx))
	return subcommands.ExitSuccess
}
