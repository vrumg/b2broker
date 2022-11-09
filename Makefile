generate:
	protoc --proto_path=api/b2broker \
			--go_out=pkg/b2brokerpb --go_opt=paths=source_relative \
  			--go-grpc_out=pkg/b2brokerpb --go-grpc_opt=paths=source_relative \
   			b2broker.proto