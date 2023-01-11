package models

import "time"

type Response struct {
	StatusCode int         `json:"status_code,omitempty" default:"200"`
	Success    bool        `json:"success,omitempty" default:"false"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	At         time.Time   `json:"at,omitempty"`
}
