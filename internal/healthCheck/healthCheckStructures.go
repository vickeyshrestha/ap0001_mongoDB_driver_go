package healthCheck

type HealthEndpoint struct {
	Application string `json:"Application"`
	HealthStatus string `json:"Health Status"`
	Message     string `json:"Message"`
}

