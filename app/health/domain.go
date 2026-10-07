package health

// Status es la entidad que representa el estado del servidor API
type Status struct {
	Version      string `json:"version"`
	OS           string `json:"operatingSystem"`
	Architecture string `json:"architecture"`
	Uptime       string `json:"uptime"`
}
