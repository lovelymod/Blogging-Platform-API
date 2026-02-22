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
	if errors.Is(err, entity.ErrAuthTokenExpired) ||
		errors.Is(err, entity.ErrAuthTokenInvalid) ||
		errors.Is(err, entity.ErrAuthTokenNotProvided) {
		return http.StatusUnauthorized
	}

	return http.StatusBadRequest
}
