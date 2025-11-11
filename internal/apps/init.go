package apps

import (
	"giniladmin/internal/apps/admin/cmd"
	"github.com/google/subcommands"
)

func init() {
	subcommands.Register(&cmd.Cmd{}, "admin")
}
