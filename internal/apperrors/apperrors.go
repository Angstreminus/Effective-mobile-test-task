package apperrors

import "net/http"

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

type InvalidDataErr struct {
	Message string
}

func (error *InvalidDataErr) Error() string {
	return error.Message
}

type DBoperationErr struct {
	Message string
}

func (error *DBoperationErr) Error() string {
	return error.Message
}

type GatewayOperationErr struct {
	Message string
}

func (error *GatewayOperationErr) Error() string {
	return error.Message
}

type AppError interface {
	Error() string
}

func MatchError(appErr AppError) *ResponseError {
	switch ae := appErr.(type) {
	case *DBoperationErr, *GatewayOperationErr:
		return &ResponseError{
			Message: ae.Error(),
			Status:  http.StatusInternalServerError,
		}
	case *InvalidDataErr:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
