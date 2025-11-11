package main

import (
	"context"
	"flag"
	"fmt"
	"giniladmin/constants"
	_ "giniladmin/internal/apps"
	"github.com/google/subcommands"
	"os"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	var v = flag.Bool("v", false, "Show version")

	flag.Parse()

	if *v {
		fmt.Printf("giniladmin-%s\n", constants.Version)
		os.Exit(int(subcommands.ExitSuccess))
	}

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}

// 初次安装，日志里输出临时账号密码， 启动初始化页面。
