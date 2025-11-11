package subcmds

import (
	"context"
	"flag"
	"fmt"
	"giniladmin/cmd/options"
	"giniladmin/constants"
	"giniladmin/pkg/logging"
	"github.com/asaskevich/govalidator"
	"github.com/google/subcommands"
)

// DummyCmd is Subcommand of host discovery mode
type DummyCmd struct {
}

// Name return subcommand name
func (*DummyCmd) Name() string { return "dummy" }

// Synopsis return synopsis
func (*DummyCmd) Synopsis() string { return "dummy demo" }

// Usage return usage
func (*DummyCmd) Usage() string {
	return `dummy show`
}

// SetFlags set flag
func (p *DummyCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&options.Opt.Debug, "debug", false, "debug mode")
}

// Execute execute
func (p *DummyCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	logging.Log = logging.NewCustomLogger(false, false, false, "./logs", "", "")
	logging.Log.Infof("vuls-%s-%s", constants.Version, constants.Revision)
	// validate
	if len(f.Args()) == 0 {
		logging.Log.Errorf("Usage: " + p.Usage())
		return subcommands.ExitUsageError
	}

	if ok, _ := govalidator.IsFilePath("/var/logs"); !ok {
		logging.Log.Errorf("Cache DB path must be a *Absolute* file path. -cache-dbpath: %s",
			"/var/logs")
		return subcommands.ExitUsageError
	}

	for _, cidr := range f.Args() {
		fmt.Println(cidr)
	}
	return subcommands.ExitSuccess
}
