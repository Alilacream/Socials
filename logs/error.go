package logs

import (
	"net/http"
)

type HTTPError struct {
	Writer  http.ResponseWriter
	ErrStat string
	Status  int
}

// the display function
func DisplayErr(w http.ResponseWriter, err string) {
	httpErr := Errors(w, err)
	if httpErr != nil {
		http.Error(httpErr.Writer, httpErr.ErrStat, int(httpErr.Status))
	}
}

// switch case for possible errors
func Errors(w http.ResponseWriter, err string) *HTTPError {
	switch err {
	case "BadRequest":
		return new(HTTPError{
			Writer:  w,
			ErrStat: err,
			Status:  http.StatusBadRequest,
		})
	case "Internal":
		return new(HTTPError{
			Writer:  w,
			ErrStat: err,
			Status:  http.StatusInternalServerError,
		})
	case "UnAuthorized":
		return new(HTTPError{
			Writer:  w,
			ErrStat: err,
			Status:  http.StatusUnauthorized,
		})

	case "Method":
		return new(HTTPError{
			Writer:  w,
			ErrStat: "UnAuthorized Method",
			Status:  http.StatusMethodNotAllowed,
		})
	case "Parse":
		return new(HTTPError{
			Writer:  w,
			ErrStat: "Unknown Request Format",
			Status:  http.StatusBadRequest,
		})
	case "Password":
		return new(HTTPError{
			Writer:  w,
			ErrStat: "Incorect Password",
			Status:  http.StatusConflict,
		})
	case "Userregister":
		return new(HTTPError{
			Writer:  w,
			ErrStat: "User Already Exists",
			Status:  http.StatusConflict,
		})
	case "Userlogin":
		return new(HTTPError{
			Writer:  w,
			ErrStat: "Used Does not exist",
			Status:  http.StatusConflict,
		})
	}
	return nil
}
