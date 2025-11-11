package core

import (
	"context"
	"fmt"
	"giniladmin/internal/apps/admin/config"
	_ "giniladmin/internal/apps/admin/service"
	"giniladmin/pkg/logging"
	"giniladmin/pkg/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	ctx    *context.Context
	conf   *config.Config
	engine *Engine
	log    *logging.Logger
}

func New(ctx context.Context) (s *Server, err error) {
	v := ctx.Value("value").(map[string]any)
	e := Engine{}
	e.Setup(ctx)

	s = &Server{
		ctx:    &ctx,
		conf:   v["conf"].(*config.Config),
		engine: &e,
		log:    v["log"].(*logging.Logger),
	}
	return
}

func (s *Server) Run() (err error) {
	// run gin http server
	addr := fmt.Sprint(":", s.conf.Server.Port)
	s.log.Infof("server listen at %s", addr)
	srv := InitServer(addr, s.engine.E)
	go func() {
		// service connections
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			utils.CheckAndExit(err)
		}
	}()

	// graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs

	// stop gin engine
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return
}

// Run runs the application and launch servers
func Run(ctx context.Context) error {
	server, err := New(ctx)
	if err != nil {
		return err
	}

	return server.Run()
}
