package main

import (
	"os"
	"sandbox/cmd/serve"
)

func main() {
	if os.Args[1] == "serve" {
		serve.Run()
	}
}