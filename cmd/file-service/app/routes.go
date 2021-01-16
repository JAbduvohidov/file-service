package app

import (
	"net/http"
)

func (s *Server) InitRoutes() {
	mux := s.router.(*http.ServeMux)

	mux.HandleFunc("/files", s.handleMultipartUpload)
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(s.storagePath))))
}
