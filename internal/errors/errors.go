package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppError struct {
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(message string, statusCode int) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func BadRequest(msg string) *AppError {
	return New(msg, fiber.ErrBadRequest.Code)
}

func NotFound(msg string) *AppError {
	return New(msg, fiber.ErrNotFound.Code)
}

func Internal(msg string) *AppError {
	return New(msg, fiber.ErrInternalServerError.Code)
}

func DBError(err error) *AppError {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return New("duplicated key not allowed", fiber.StatusBadRequest)
	}

	if errors.Is(err, gorm.ErrInvalidData) {
		return New("invalid or missing data", fiber.StatusBadRequest)
	}

	return Internal(err.Error())
}
