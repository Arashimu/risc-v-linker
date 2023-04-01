package main

import (
	"fmt"
	"os"
	"riscv_link/pkg/linker"
	"riscv_link/pkg/utils"
)

func main() {
	if len(os.Args) < 2 {
		utils.Fatal("wrong args")
	}
	file := linker.MustNewFile(os.Args[1])

	objFile := linker.NewObjectFile(file)

	//for _, shdr := range inputFile.ElfSections {
	//	fmt.Println(linker.ElfGetName(inputFile.ShStrtab, shdr.Name))
	//}
	objFile.Parse()

	for _, sym := range objFile.ElfSyms {
		fmt.Println(linker.ElfGetName(objFile.SymbolStrtab, sym.Name))
	}
}
