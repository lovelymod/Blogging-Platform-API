package utils

import (
	"blogging-platform-api/internal/entity"
	"net/http"
)

func GetHttpStatus(err error) int {
	switch err {
	case entity.ErrGlobalServerErr:
		return http.StatusInternalServerError

	case entity.ErrGlobalNotFound:
		return http.StatusNotFound
	}

	return http.StatusBadRequest
}
