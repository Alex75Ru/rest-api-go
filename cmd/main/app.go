package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api-go/internal/user"
	"rest-api-go/pkg/logging"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	log.Println("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		// need to write timeout
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening the port 127.0.0.1:1234")
	logger.Fatal(server.Serve(listener))
}
