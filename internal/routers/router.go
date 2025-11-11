package routers

import (
	"errors"
	"fmt"
	"giniladmin/pkg/logging"
	"giniladmin/pkg/utils"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"sync"
)

type routerFunc struct {
	Name   string
	Weight int
	Func   func(router *gin.Engine)
}

type routerSlice []routerFunc

func (r routerSlice) Len() int { return len(r) }

func (r routerSlice) Less(i, j int) bool { return r[i].Weight > r[j].Weight }

func (r routerSlice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

var userRouterOnce sync.Once
var routers routerSlice

// register new router with key name
// key name is used to eliminate duplicate routes
// key name not case sensitive
func register(name string, f func(router *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

func Register(name string, f func(router *gin.Engine)) {
	registerWithWeight(name, 50, f)
}

// register new router with weight
func registerWithWeight(name string, weight int, f func(router *gin.Engine)) {
	if weight > 100 || weight < 0 {
		utils.CheckAndExit(errors.New(fmt.Sprintf("router weight must be >= 0 and <=100")))
	}

	for _, r := range routers {
		if strings.ToLower(name) == strings.ToLower(r.Name) {
			utils.CheckAndExit(errors.New(fmt.Sprintf("router [%s] already register", r.Name)))
		}
	}

	routers = append(routers, routerFunc{
		Name:   name,
		Weight: weight,
		Func:   f,
	})
}

// framework init
func Setup(e *gin.Engine, log *logging.Logger) {
	userRouterOnce.Do(func() {
		sort.Sort(routers)
		for _, r := range routers {
			r.Func(e)
			log.Infof("load router [%s] success...", r.Name)
		}
	})
}
