package models

type Response struct {
	StatusCode int         `json:"status_code" default:"200"`
	Success    bool        `json:"success,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}
