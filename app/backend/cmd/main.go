package main

import (
	"fmt"
	"os"
	"sandbox/cmd/serve"
)

func main() {
	if os.Args[1] == "serve" {
		fmt.Println("Serving...")
		serve.Run()
	}
}