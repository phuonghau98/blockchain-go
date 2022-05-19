package main

import (
	"hvphuong98/tensor-blc/cli"
	"os"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
