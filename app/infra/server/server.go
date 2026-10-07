package server

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/wfrs-dev/hefesto/app/config"
)

type Modular interface {
	Register(e *echo.Echo)
}

type Server struct {
	instance *echo.Echo
	config   *config.Config
	modules  []Modular
}

func New(cfg *config.Config) (*Server, error) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: string2LogLevel(cfg.Server.LogLevel),
	}))
	slog.SetDefault(log)

	e := echo.New()
	e.Use(requestLoggerMiddleware())
	return &Server{
		instance: e,
		config:   cfg,
		modules:  []Modular{},
	}, nil
}

func (s *Server) AddModule(m Modular) {
	s.modules = append(s.modules, m)
}

func (s *Server) Start() error {
	for _, m := range s.modules {
		m.Register(s.instance)
	}

	return s.instance.Start(fmt.Sprintf(":%d", s.config.Server.Port))
}

func (s *Server) EchoInstance() *echo.Echo {
	return s.instance
}
