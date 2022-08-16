# Library
Library is example service in microservices architecture

## Features
- transcoding api with grpc
- swagger api + swagger ui
- healthy grpc service `http://localhost/health`
- mongodb generic crud with transaction

# How to build and run Server?
- run make `make build`
- run mongodb instance
- rename `.env.example` to `.env` and set your environments
- run `./library`

# Environment variables

```shell
MONGO_LIBRARY_URI="mongodb://user:12345@localhost:27017/?directConnection=true"
SERVER_GRPC_ADDRESS="localhost"
SERVER_GRPC_PORT="9010"
SERVER_HTTP_ADDRESS="localhost"
SERVER_HTTP_PORT="8010"
```