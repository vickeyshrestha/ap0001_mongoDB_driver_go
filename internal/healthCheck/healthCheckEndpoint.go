package healthCheck

type HealthEndpoint struct {
	Application  string `json:"Application"`
	Version      string `json:"Version"`
	HealthStatus string `json:"Health Status"`
	Message      string `json:"Message"`
}
