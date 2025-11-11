package middleware

import (
	"context"
	"errors"
	"fmt"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

type middleware struct {
	HandlerFunc func() gin.HandlerFunc
	Weight      int
}

type middlewareSlice []middleware

var mws middlewareSlice

func (m middlewareSlice) Len() int { return len(m) }

func (m middlewareSlice) Less(i, j int) bool { return m[i].Weight > m[j].Weight }

func (m middlewareSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// registering new middleware
func register(handlerFunc func() gin.HandlerFunc) {
	mw := middleware{
		HandlerFunc: handlerFunc,
		Weight:      50,
	}
	mws = append(mws, mw)
}

// registering new middleware with weight
func registerWithWeight(weight int, handlerFunc func() gin.HandlerFunc) {

	if weight > 100 || weight < 0 {
		utils.CheckAndExit(errors.New(fmt.Sprintf("middleware weight must be >= 0 and <=100")))
	}

	mw := middleware{
		HandlerFunc: handlerFunc,
		Weight:      weight,
	}
	mws = append(mws, mw)
}

// framework init func
func Setup(ctx context.Context, e *gin.Engine) {
	sort.Sort(mws)
	for _, mw := range mws {
		e.Use(mw.HandlerFunc())
	}
}
