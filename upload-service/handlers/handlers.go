package handlers

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"time"

	pb "github.com/benjih/grpc-services-test/grpc"
	"google.golang.org/grpc"
)

func UploadCustomerContactsCSVHandler(consumerContactServiceUrl string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := grpc.Dial(consumerContactServiceUrl, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewCustomerContactServiceClient(conn)

		fscanner := bufio.NewScanner(r.Body)
		for fscanner.Scan() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err = client.AddOrUpdateCustomerContact(ctx, &pb.CustomerContactRequest{Data: fscanner.Text()})
			if err != nil {
				log.Fatalf("Failed to update customer-contact-service: %v", err)
			}
		}
		w.WriteHeader(200)
	})
}
