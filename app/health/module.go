package health

import (
	"github.com/labstack/echo/v5"
)

type Module struct {
	handler *StatusHandler
	version string
	created int64
}

func New(version string, created int64) *Module {
	return &Module{
		handler: NewStatusHandler(version, created),
		version: version,
		created: created,
	}
}

func (m *Module) Register(e *echo.Echo) {
	m.handler.Routes(e.Group("/status"))
}
