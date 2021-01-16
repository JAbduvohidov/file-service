package main

import (
	"file-service/cmd/file-service/app"
	"file-service/pkg/services/files"
	"flag"
	"github.com/JAbduvohidov/jwt"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	hostF        = flag.String("host", "", "Server host")
	portF        = flag.String("port", "", "Server port")
	secretF      = flag.String("secret", "", "Secret key")
	storagePathF = flag.String("filepath", "files", "Files store directory")
)

const (
	envHost        = "HOST"
	envPort        = "PORT"
	envSecret      = "SECRET"
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
	secret, ok := fLagOrEnv(secretF, envSecret)
	if !ok {
		log.Panic("can't get storage path")
	}
	addr := net.JoinHostPort(host, port)
	start(addr, storagePath, jwt.Secret(secret))
}

func start(addr string, path string, secret jwt.Secret) {
	createStorageFolder(path)

	mux := http.NewServeMux()
	fileSvc := files.NewFileService(path)
	server := app.NewServer(
		mux,
		&secret,
		fileSvc,
		path,
	)
	server.Start()
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
