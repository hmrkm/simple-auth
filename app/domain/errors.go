package domain

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrTokenWasExpired = errors.New("token was Expired")
	ErrInvalidVerify   = errors.New("invalid verify")
)
