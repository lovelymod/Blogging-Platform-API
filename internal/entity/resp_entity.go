package entity

type PaginationMeta struct {
	Page       int   `json:"page,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	TotalRows  int64 `json:"totalRows,omitempty"`
	TotalPages int   `json:"totalPages,omitempty"`
}

type Resp struct {
	Message string         `json:"message,omitempty"`
	Data    any            `json:"data,omitempty"`
	Success bool           `json:"success"`
	Meta    PaginationMeta `json:"meta"`
}
