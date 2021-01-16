package app

import (
	"errors"
	"file-service/pkg/services/files"
	"github.com/JAbduvohidov/jwt"
	"net/http"
)

type Server struct {
	router      http.Handler
	fileSvc     *files.FileService
	secret      *jwt.Secret
	storagePath string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) Start() {
	s.InitRoutes()
}

func NewServer(router http.Handler, secret *jwt.Secret, fileSvc *files.FileService, storagePath string) *Server {
	if router == nil {
		panic(errors.New("router can't be nil"))
	}
	if fileSvc == nil {
		panic(errors.New("fileSvc can't be nil"))
	}
	if storagePath == "" {
		panic(errors.New("storagePath can't be nil"))
	}

	return &Server{fileSvc: fileSvc, secret: secret, storagePath: storagePath, router: router}
}
