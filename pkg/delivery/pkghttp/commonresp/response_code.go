package commonresp

import (
	"github.com/44nbud1/akademik/pkg/pkgservice"
	"net/http"
)

type Response interface {
	GetHttp() int
	GetCaseCode() string
	Error() string
}

func NewResponseCode(errx *pkgservice.ErrorService) Response {

	if errx == nil {
		return ResponseCode{
			HttpCode: http.StatusOK,
			CaseCode: "00",
			Message:  "Success",
		}
	}

	switch errx.GetCode() {

	case pkgservice.ErrBadRequest:
		return ResponseCode{
			HttpCode: http.StatusBadRequest,
			CaseCode: "99",
			Message:  "Bad Request",
		}

	case pkgservice.ErrInvalid:
		return ResponseCode{
			HttpCode: http.StatusBadRequest,
			CaseCode: "99",
			Message:  "Invalid Request",
		}

	case pkgservice.ErrNotFound:
		return ResponseCode{
			HttpCode: http.StatusNotFound,
			CaseCode: "99",
			Message:  "Not found",
		}

	default:
		return ResponseCode{
			HttpCode: http.StatusInternalServerError,
			CaseCode: "99",
			Message:  "General error",
		}
	}
}
