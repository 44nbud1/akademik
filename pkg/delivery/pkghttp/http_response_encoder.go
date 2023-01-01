package pkghttp

import (
	"context"
	"encoding/json"
	"net/http"
)

var (
	applicationJson = "application/json; charset=utf-8"
	contentType     = "Content-Type"
)

type errorer interface {
	GetErrorResponse() error
	StatusCode() int
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set(contentType, applicationJson)
	if headerEr, ok := response.(Headerer); ok {
		for k, values := range headerEr.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	code := http.StatusOK
	if sc, ok := response.(StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	if code == http.StatusNoContent {
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

type StatusCoder interface {
	StatusCode() int
}

type Headerer interface {
	Headers() http.Header
}
