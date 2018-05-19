package webhook

import "time"

type Config struct {
	ContentType string `json:"content_type"`
	InsecureSsl string `json:"insecure_ssl"`
	Secret      string `json:"secret"`
	URL         string `json:"url"`
}

type LastResponse struct {
	Code    *string `json:"code"`
	Status  string  `json:"status"`
	Message *string `json:"message"`
}

type Hook struct {
	Type         string       `json:"type"`
	ID           int64        `json:"id"`
	Name         string       `json:"name"`
	Active       bool         `json:"active"`
	Events       []string     `json:"events"`
	Config       Config       `json:"config"`
	UpdatedAt    time.Time    `json:"updated_at"`
	CreatedAt    time.Time    `json:"created_at"`
	URL          string       `json:"url"`
	TestURL      string       `json:"test_url"`
	PingURL      string       `json:"ping_url"`
	LastResponse LastResponse `json:"last_response"`
}

type Payload struct {
	Hook Hook `json:"hook"`
}
