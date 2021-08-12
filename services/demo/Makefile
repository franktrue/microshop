build:
	protoc -I. --micro_out=. --go_out=. proto/demo/demo.proto  
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-service .

run:
	docker run -d -p 9999:9999 -e MICRO_SERVER_ADDRESS=:9999 -e MICRO_REGISTRY=mdns laracom-demo-service