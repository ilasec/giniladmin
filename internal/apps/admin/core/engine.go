package core

import (
	"context"
	"giniladmin/internal/apps/admin/config"
	"giniladmin/internal/apps/admin/service"
	"giniladmin/internal/middleware"
	"giniladmin/internal/repository"
	"giniladmin/internal/routers"
	"giniladmin/pkg/logging"
	"sync"

	"github.com/gin-gonic/gin"
)

var once sync.Once

var EngineInstance *gin.Engine

type Engine struct {
	E *gin.Engine
	R repository.Repository
	L *logging.Logger
}

// init gin router engine
func (p *Engine) Init(debug bool) {
	once.Do(func() {
		if debug {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
		p.E = gin.New()
		EngineInstance = p.E

	})
}

func (p *Engine) Setup(ctx context.Context) (err error) {
	// Extract values from context, handling potential panics.
	v, ok := ctx.Value("value").(map[string]any)
	if !ok {
		return errMissingContextValue("value") // Create a custom error type
	}

	c, ok := v["conf"].(*config.Config)
	if !ok {
		return errMissingContextValue("conf")
	}

	p.L, ok = v["log"].(*logging.Logger)
	if !ok {
		return errMissingContextValue("log")
	}

	p.Init(c.Server.Debug) // Initialize Gin Engine

	// repository init
	if err := p.R.Setup(context.WithValue(ctx, "value", c.DataBase)); err != nil {
		return err // Return the error from repository setup
	}

	// Service setup.  Create the value map directly.
	serviceValue := map[string]any{
		"conf": c,
		"log":  p.L,
		"repo": &p.R,
	}
	service.Setup(context.WithValue(ctx, "value", serviceValue))
	middleware.Setup(ctx, p.E)
	routers.Setup(p.E, p.L)
	return
}

// Custom error type for missing context values.  Improves error handling.
type missingContextValueError string

func (e missingContextValueError) Error() string {
	return "missing context value: " + string(e)
}

func errMissingContextValue(key string) error {
	return missingContextValueError(key)
}
