//go:build tools
// +build tools

// Package tools tracks dependencies for tools that are required to generate the protobuf code.
// See https://github.com/golang/go/issues/25922
package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
)
