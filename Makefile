generate:
	protoc --proto_path=api/b2broker \
			--go_out=pkg/b2brokerpb --go_opt=paths=source_relative \
  			--go-grpc_out=pkg/b2brokerpb --go-grpc_opt=paths=source_relative \
   			b2broker.proto

run-server:
	go run ./cmd/server/main.go

run-client:
	go run ./cmd/client/main.go

run:
	go run ./cmd/server/main.go
	go run ./cmd/client/main.go