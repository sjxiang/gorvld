package main

import (
	"os"

	"github.com/sjxiang/gorvld/pkg/linker"
	"github.com/sjxiang/gorvld/pkg/utils"
)


func main() {
	
	if len(os.Args) < 2 {
		utils.Fatal("wrong args")
	}
	
	file := linker.MustNewFile(os.Args[1])
	
	inputFile := linker.NewInputFile(file)
	utils.Assert(len(inputFile.ElfSections) == 12)
}