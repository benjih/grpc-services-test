package main

import (
	"log"
	"net/http"

	"github.com/benjih/grpc-services-test/upload-service/handlers"
)

const (
	port                      = ":3000"
	consumerContactServiceUrl = "localhost:3001"
)

func main() {
	log.Print(" starting upload-service on " + port)

	http.HandleFunc("/upload", handlers.UploadCustomerContactsCSVHandler(consumerContactServiceUrl))

	http.ListenAndServe(port, nil)
}
