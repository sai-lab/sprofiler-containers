// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	if err := os.Mkdir("hoge", 0o660); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
