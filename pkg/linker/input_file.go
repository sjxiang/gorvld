package linker

import (
	"bytes"
	"encoding/binary"

	"github.com/sjxiang/gorvld/pkg/utils"
)

type InputFile struct {
	File *File
	ElfSections []Shdr
}


func NewInputFile(file *File) InputFile {
	f := InputFile{ File: file }

	if len(file.Contents) < EhdrSize {
		utils.Fatal("file too small")
	}

	if !CheckMagic(file.Contents) {
		utils.Fatal("not an ELF file")
	}

	ehdr := Ehdr{}
	reader := bytes.NewReader(file.Contents)
	err := binary.Read(reader, binary.LittleEndian, &ehdr)  // 小端模式，从 reader 中读取数据，写入到 ehdr 中
	utils.MustNo(err)

	contents := file.Contents[ehdr.ShOff:]

	shdr := utilRead(contents)  // 辅助函数

	numSections := int64(ehdr.ShNum)
	if numSections == 0 {
		numSections = int64(shdr.Size)
	}

	f.ElfSections = []Shdr{ shdr }
	for numSections > 1 {
		contents = contents[ShdrSize:]
		f.ElfSections = append(f.ElfSections, utilRead(contents))
		numSections--
	}
	

	return f
}


func utilRead(data []byte) Shdr {
	shdr := Shdr{}

	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &shdr)
	utils.MustNo(err)

	return shdr
}
