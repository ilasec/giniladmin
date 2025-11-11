package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func CheckAndExit(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}

func CheckFileExit(path string) (b bool) {
	info, err := os.Stat(path)
	if err == nil {
		b = true != info.IsDir()
	}
	return
}
