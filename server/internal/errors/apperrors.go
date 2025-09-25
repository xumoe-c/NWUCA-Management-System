package apperrors

import "errors"

var (
	ErrUsernameExists     = errors.New("username already exists")
	ErrEmailExists        = errors.New("email already exists")
	ErrInvalidCredits     = errors.New("invalid email or password")
	ErrPositionNameExists = errors.New("position name already exists")
	ErrNotFound           = errors.New("not found")
)
