package health

import (
	"context"
)

type Processor struct {
}

func DoHealth(ctx context.Context) (status int, message string, ret any, err error) {
	logger.Infof("Health info")
	return
}
