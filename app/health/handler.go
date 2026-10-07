package health

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type StatusHandler struct {
	health  *HealthApp
	created int64
	version string
}

func NewStatusHandler(version string, created int64) *StatusHandler {
	return &StatusHandler{
		health:  NewHealthApp(version, created),
		created: created,
		version: version,
	}
}

func (h *StatusHandler) Health(c *echo.Context) error {
	resp := h.health.Execute(c.Request().Context())

	return c.JSON(http.StatusOK, resp)
}

func (h *StatusHandler) Routes(g *echo.Group) {
	g.GET("/health", h.Health)
}
