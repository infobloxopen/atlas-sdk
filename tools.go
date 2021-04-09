//+build tools
package tools

import (
	_ "github.com/infobloxopen/atlas-app-toolkit/query"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
