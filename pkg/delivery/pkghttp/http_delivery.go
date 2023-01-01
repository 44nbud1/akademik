package pkghttp

import (
	"context"
	"net/http"
)

type Endpoint func(ctx context.Context, r *http.Request) (response interface{})

type Server interface {
	http.Handler
}

type server struct {
	e Endpoint
}

func NewServer(e Endpoint) *server {
	return &server{
		e,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response := s.e(ctx, r)
	err := EncodeResponse(ctx, w, response)
	if err != nil {
		return
	}

}
