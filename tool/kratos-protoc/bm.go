package main

import (
	"os/exec"
)

const (
	_bmProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --bm_out=:."
)

func installBMGen() error {
	if _, err := exec.LookPath("protoc-gen-bm"); err != nil {
		panic("protoc-gen-bm not exist")
	}
	return nil
}

func genBM(files []string) error {
	return generate(_bmProtoc, files)
}
