run-customer-contact-service:
	go run customer-contact-service/main.go

build-customer-contact-service:
	cd customer-contact-service && go build 

docker-build-customer-contact-service:
	cd upload-service && \
	env GOOS=linux GOARCH=amd64 go build && \
	docker build -t benjih-customer-contact-service .

run-upload-service:
	go run upload-service/main.go

build-upload-service:
	cd upload-service && go build

docker-build-upload-service:
	cd upload-service && \
	env GOOS=linux GOARCH=amd64 go build && \
	docker build -t benjih-upload-service .

pb:
	protoc -I grpc grpc/grpc.proto --go_out=plugins=grpc:grpc

upload:
	curl -X POST http://localhost:3000/upload --data-binary "@data.csv"