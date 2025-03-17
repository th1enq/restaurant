package helper

import "errors"

var (
	ErrInvalidFood     = errors.New("invalid food name")
	ErrInvalidDrinking = errors.New("invalid drinking name")
	ErrInvalidEmployee = errors.New("invalid employee type")
	ErrInvalidStatus   = errors.New("invalid status")
)
