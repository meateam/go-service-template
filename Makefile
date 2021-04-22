# Basic go commands
PROTOC=protoc

# Binary names
BINARY_NAME=go-template-service

build-proto:
		rm -f proto/*.go
		go clean -modcache
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/template.proto
show-problems:
		go vet .