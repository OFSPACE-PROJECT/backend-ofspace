package errors

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrMovieResource = errors.New("movie id not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrFillData = errors.New("please fill in the required data")

	ErrUsernamePasswordNotFound = errors.New("check Your email and password")
)

func ErrorCode(err error) (code int) {
	if err == ErrDuplicateData {
		code = http.StatusBadRequest
	} else if err == ErrFillData {
		code = http.StatusBadRequest
	} else if err == ErrInternalServer {
		code = http.StatusInternalServerError
	} else if err == ErrMovieResource {
		code = http.StatusNoContent
	} else if err == ErrNotFound {
		code = http.StatusNoContent
	} else if err == ErrUsernamePasswordNotFound {
		code = http.StatusBadRequest
	}
	return code
}
