package app

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

const multipartMaxBytes = 15 * 1024 * 1024

func (s *server) handleMultipartUpload(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		return
	}

	err := request.ParseMultipartForm(multipartMaxBytes)
	if err != nil {
		log.Print(err)
		http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fileHeaders := request.MultipartForm.File["file"]

	type FileURL struct {
		Id  string
		URL string
	}
	fileURLs := make([]FileURL, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		name, err := s.saveFile(fileHeader)
		if err != nil {
			http.Error(responseWriter, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Print(err)
			return
		}

		fileURLs = append(fileURLs, FileURL{
			Id:  name[:len(name)-len(filepath.Ext(name))],
			URL: s.storagePath + "/" + name,
		})
	}

	urlsJSON, err := json.Marshal(fileURLs)
	if err != nil {
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	_, err = responseWriter.Write(urlsJSON)
	if err != nil {
		log.Print(err)
		return
	}

	return
}

func (s *server) saveFile(fileHeader *multipart.FileHeader) (name string, err error) {
	file, err := fileHeader.Open()
	if err != nil {
		return
	}
	defer func() {
		err = file.Close()
	}()

	contentType := fileHeader.Header.Get("Content-Type")
	name, err = s.fileSvc.Save(file, contentType)
	if err != nil {
		return
	}

	return //nil
}
