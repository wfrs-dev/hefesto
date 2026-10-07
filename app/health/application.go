package health

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

type HealthApp struct {
	version string
	created int64
}

func NewHealthApp(version string, created int64) *HealthApp {
	return &HealthApp{
		version: version,
		created: created,
	}
}

func duration(created int64) string {
	d := time.Since(time.Unix(created, 0)).Round(time.Second)

	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}

func (h *HealthApp) Execute(ctx context.Context) Status {
	return Status{
		Version:      h.version,
		OS:           runtime.GOOS,
		Architecture: runtime.GOARCH,
		Uptime:       duration(h.created),
	}
}
