.PHONY: statik proto build submodule

submodule:
	git submodule update --init --recursive

statik:
	statik -f -src=api/swagger/swaggerui/ -dest=server/

proto:
	protoc --proto_path=proto --proto_path=proto/include/googleapis --proto_path=proto/include/grpc-gateway --go-grpc_out=proto/protoModel --go_out=proto/protoModel --grpc-gateway_out=logtostderr=true:proto/protoModel --openapiv2_out=logtostderr=true:api/swagger proto/server/*.proto
	cat api/swagger/server/*.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' > api/swagger.json

build:
	go fmt ./...
	go mod tidy
	go mod vendor
	go build -o ./library server/main.go
