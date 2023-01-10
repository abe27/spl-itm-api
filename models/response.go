package models

type Response struct {
	StatusCode int         `json:"status_code" default:"200"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Error      string      `json:"error"`
}
