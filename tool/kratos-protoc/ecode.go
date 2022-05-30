package main

import (
	"os/exec"
)

const (
	_ecodeProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --ecode_out=" +
		"Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types," +
		"Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:."
)

func installEcodeGen() error {
	if _, err := exec.LookPath("protoc-gen-ecode"); err != nil {
		panic("protoc-gen-ecode not exist")
	}
	return nil
}

func genEcode(files []string) error {
	return generate(_ecodeProtoc, files)
}
