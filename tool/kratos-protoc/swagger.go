package main

import (
	"os/exec"
)

const (
	_swaggerProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bswagger_out=" +
		"Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:."
)

func installSwaggerGen() error {
	if _, err := exec.LookPath("protoc-gen-bswagger"); err != nil {
		panic("protoc-gen-bswagger not exist")

	}
	return nil
}

func genSwagger(files []string) error {
	return generate(_swaggerProtoc, files)
}
