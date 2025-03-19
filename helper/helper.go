package helper

import "errors"

var (
	ErrInvalidFood     = errors.New("invalid food name")
	ErrInvalidDrinking = errors.New("invalid drinking name")
	ErrInvalidItem     = errors.New("invalid item")
	ErrInvalidEmployee = errors.New("invalid employee type")
	ErrInvalidStatus   = errors.New("invalid status")
)
