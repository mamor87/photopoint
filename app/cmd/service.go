package main

import (
	"os"

	"github.com/marmor87/photopoint/app/cmd/parser"
)

func setupServer() {}

func startClient() {}

func main() {
	args := parser.ParseArguments(os.Args[1:])
	setupServer()
	if !args.RunInBackground {
		startClient()
	}
}
