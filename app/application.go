package app

import (
	"log"
	"time"

	"github.com/wfrs-dev/hefesto/app/config"
	"github.com/wfrs-dev/hefesto/app/health"
	"github.com/wfrs-dev/hefesto/app/infra/server"
)

type EchoApplication struct {
	srv *server.Server
}

func New(cfg *config.Config, version string, created int64) *EchoApplication {
	srv, err := server.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	srv.AddModule(health.New(version, time.Now().Unix()))

	return &EchoApplication{
		srv: srv,
	}
}
