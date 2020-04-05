// +build tools

package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"                       // protoc-gen-go
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"            // linters aggregator
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway" // protoc-gen-grpc-gateway
	_ "github.com/uber/prototool/cmd/prototool"                        // swiss army knife for protocol buffers
	_ "golang.org/x/tools/cmd/goimports"                               // updates imports and formats code
)
