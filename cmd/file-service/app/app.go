package app

import (
	"errors"
	"file-service/pkg/services/files"
	"net/http"
)

type server struct {
	router        http.Handler
	fileSvc       *files.FileService
	storagePath   string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(router http.Handler, fileSvc *files.FileService, storagePath string) *server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if fileSvc == nil {
		panic(errors.New("fileSvc can't be nil"))
	}
	if storagePath == "" {
		panic(errors.New("storagePath can't be nil"))
	}

	return &server{fileSvc: fileSvc, storagePath: storagePath, router: router}
}
