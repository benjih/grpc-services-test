run-customer-contact-service:
	go run customer-contact-service/main.go

build-customer-contact-service:
	cd customer-contact-service && go build 

run-upload-service:
	go run upload-service/main.go

build-upload-service:
	cd upload-service && go build 

pb:
	protoc -I grpc grpc/grpc.proto --go_out=plugins=grpc:grpc

upload:
	curl -X POST http://localhost:3000/upload --data-binary "@data.csv"