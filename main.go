package main

import (
	"fmt"
	"os"
)

var version = "1.0.0-dev"

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		fmt.Println(version)
		return
	}
	fmt.Println("This is app.")
}
