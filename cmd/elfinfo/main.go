package main

import (
	"fmt"
	"os"

	"elfinfo/internal/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: elfinfo <elf-file>")
		os.Exit(1)
	}

	err := parser.PrintInfo(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}