# Hasher service
## Description
Stateful application containing generated uuid hash in its memory.
Hash is recreated every HASH_TTL interval.
Application contains two api servers: gRPC and http.
Each api implements single endpoint to get actual hash string and hash generation datetime.
Covered with unit tests where it’s needed.
This app demonstrates coding quality, app design skills, golang best practices, etc.

## Prerequisites

### Protobuf compiler
Install [protoc](https://grpc.io/docs/languages/go/quickstart/)

### Open API code generation
Install [oapi-codegen](https://github.com/deepmap/oapi-codegen)

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Generate boilerplate code
go generate ./...