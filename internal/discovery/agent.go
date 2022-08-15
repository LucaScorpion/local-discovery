package discovery

import "time"

type Agent struct {
	Name         string         `json:"name"`
	LocalAddress string         `json:"localAddress"`
	Info         map[string]any `json:"info"`
	registered   time.Time
}
