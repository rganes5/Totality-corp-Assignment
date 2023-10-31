proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto
run:
	go run cmd/main.go
wire:
	cd pkg/wire/ && wire
swag:
	swag init -g cmd/main.go