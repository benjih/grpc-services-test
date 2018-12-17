run-customer-contact-service:
	go run customer-contact-service/main.go

run-upload-service:
	go run upload-service/main.go

pb:
	protoc -I grpc grpc/grpc.proto --go_out=plugins=grpc:grpc

upload:
	curl -X POST http://localhost:3000/upload --data-binary "@data.csv"