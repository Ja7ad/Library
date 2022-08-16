.PHONY: statik proto build

statik:
	statik -f -src=api/swagger/swaggerui/ -dest=server/

proto:
	protoc --proto_path=proto --go-grpc_out=proto/protoModel --go_out=proto/protoModel --grpc-gateway_out=logtostderr=true:proto/protoModel --openapiv2_out api/swagger --openapiv2_opt logtostderr=true --openapiv2_opt generate_unbound_methods=true proto/library.proto

build:
	go fmt ./...
	go mod tidy
	go mod vendor
	go build -o ./library server/main.go