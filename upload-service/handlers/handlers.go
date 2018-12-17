package handlers

import (
	"bufio"
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"strings"
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
			text := fscanner.Text()
			csvReader := csv.NewReader(strings.NewReader(text))
			record, err := csvReader.Read()
			if err != nil {
				log.Print(err.Error())
				w.WriteHeader(500)
				return
			}

			_, err = client.AddOrUpdateCustomerContact(ctx, &pb.CustomerContactRequest{
				Id:              record[0],
				Name:            record[1],
				EmailAddress:    record[2],
				TelephoneNumber: record[3],
			})
			if err != nil {
				log.Fatalf("Failed to update customer-contact-service: %v", err)
			}
		}
		w.WriteHeader(200)
	})
}
