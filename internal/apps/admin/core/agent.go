package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type agent interface {
	ListenAndServe() error
}

func InitServer(address string, router *gin.Engine) agent {
	p := endless.NewServer(address, router)
	p.ReadHeaderTimeout = 20 * time.Second
	p.WriteTimeout = 20 * time.Second
	p.MaxHeaderBytes = 1 << 20
	return p
}
