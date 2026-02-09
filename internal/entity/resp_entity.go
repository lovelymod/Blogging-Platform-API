package entity

type Resp struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Success bool   `json:"success"`
}
