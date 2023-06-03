package entity

import (
	"errors"
	"source-base-go/config"
)

var ErrInternalServerError = errors.New(config.INTERNAL_SERVER_ERROR)
var ErrUnauthorized = errors.New(config.UNAUTHORIZED)
var ErrLock = errors.New(config.LOCK)
var ErrBadRequest = errors.New(config.BAD_REQUEST)
var ErrForbidden = errors.New(config.FORBIDDEN)
