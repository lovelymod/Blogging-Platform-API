package entity

import "errors"

// Global Error
var (
	ErrGlobalNotFound  = errors.New("not_found")
	ErrGlobalServerErr = errors.New("internal_server_error")
)

// Blog Error
var (
	ErrBlogTagMustBeUnique = errors.New("tag_must_be_unique")
)

// Blog User
var (
	ErrUserThisEmailIsAlreadyUsed = errors.New("this_email_is_already_used")
)
