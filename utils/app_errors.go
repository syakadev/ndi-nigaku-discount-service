package utils

import (
	"os"
)

type RequestError struct {
	StatusCode int                   `json:"status_code" xml:"status_code" example:"422"`
	Message    string                `json:"message" xml:"message" example:"Invalid email address"`
	Fields     []DataValidationError `json:"fields" xml:"fields"`
}

func (re RequestError) Error() string {
	return re.Message
}

type DataValidationError struct {
	Field   string `json:"field" xml:"field" example:"email"`
	Message string `json:"message" xml:"message" example:"Invalid email address"`
}

type GlobalError struct {
	Message string `json:"message" xml:"message" example:"invalid name"`
}

type NamaLevelResponse struct {
	NamaLevel string `json:"nama_level"`
}

type StandardError struct {
	Code    int    `json:"code" xml:"code" example:"422"`
	Message string `json:"message" xml:"message" example:"Invalid email address"`
}

func (se StandardError) Error() string {
	return se.Message
}

func GeneralErrorResponse(code int, err error) StandardError {
	env := os.Getenv("SERVER_LOGS_ENV")
	if env == "Production" {
		return StandardError{
			Code:    code,
			Message: "Terjadi kesalahan pada sistem. Silakan hubungi admin.",
		}
	}
	return StandardError{
		Code:    code,
		Message: err.Error(),
	}
}
