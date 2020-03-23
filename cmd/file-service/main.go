package main

import (
	"file-service/cmd/file-service/app"
	"file-service/pkg/services/files"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	hostF        = flag.String("host", "", "Server host")
	portF        = flag.String("port", "", "Server port")
	storagePathF = flag.String("filepath", "files", "Files store directory")
)

const (
	envHost        = "HOST"
	envPort        = "PORT"
	envStoragePath = "STORAGE_PATH"
)

func fLagOrEnv(flag *string, envName string) (server string, ok bool) {
	if *flag != "" {
		return *flag, true
	}
	return os.LookupEnv(envName)
}

func main() {
	flag.Parse()
	host, ok := fLagOrEnv(hostF, envHost)
	if !ok {
		log.Panic("can't get host")
	}
	port, ok := fLagOrEnv(portF, envPort)
	if !ok {
		log.Panic("can't get port")
	}
	storagePath, ok := fLagOrEnv(storagePathF, envStoragePath)
	if !ok {
		log.Panic("can't get storage path")
	}
	addr := net.JoinHostPort(host, port)
	start(addr, storagePath)
}

func start(addr string, path string) {
	createStorageFolder(path)
	mux := http.NewServeMux()
	fileSvc := files.NewFileService(path)
	server := app.NewServer(
		mux,
		fileSvc,
		path,
	)
	server.InitRoutes()
	panic(http.ListenAndServe(addr, server))
}

func createStorageFolder(path string) {
	err := os.Mkdir(path, 0666)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatalf("can't create directory: %s", err)
		}
	}
}
