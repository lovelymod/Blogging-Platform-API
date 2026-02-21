package utils

import (
	"blogging-platform-api/internal/entity"
	"errors"
	"net/http"
)

func GetHttpErrStatus(err error) int {
	if errors.Is(err, entity.ErrGlobalServerErr) {
		return http.StatusInternalServerError
	}
	if errors.Is(err, entity.ErrGlobalNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, entity.ErrAuthTokenExpired) {
		return http.StatusUnauthorized
	}

	return http.StatusBadRequest
}
