package parser

import (
	"debug/elf"
	"fmt"
)

func PrintInfo(path string) error {
	f, err := elf.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Println("=== ELF INFO ===")
	fmt.Printf("Class: %s\n", f.Class)
	fmt.Printf("Data: %s\n", f.Data)
	fmt.Printf("Machine: %s\n", f.Machine)
	fmt.Printf("Entry: 0x%x\n", f.Entry)

	fmt.Println("\nSections:")
	for _, sec := range f.Sections {
		fmt.Printf(
			"%-20s %-10s %d bytes\n",
			sec.Name,
			sec.Type,
			sec.Size,
		)
	}

	return nil
}