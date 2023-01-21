proto:
	protoc ./pkg/pb/*.proto -I. --go_out=:.  --go-grpc_opt=require_unimplemented_servers=false  --go-grpc_out=:. --experimental_allow_proto3_optional

server:
	go run cmd/main.go