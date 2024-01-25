package apperrors

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

type AppError interface {
	Error() string
}
