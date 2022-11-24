package linker

import (
	"fmt"
	"github.com/sjxiang/gorvld/pkg/utils"
)

type InputFile struct {
	File        *File
	ElfSections []Shdr
	ShStrtab    []byte
}


func NewInputFile(file *File) InputFile {
	f := InputFile{ File: file }

	// 确保不小于 ELF header 长度
	if len(file.Contents) < EhdrSize {
		utils.Fatal("file too small")
	}

	// 校验 ELF header magic number
	if !CheckMagic(file.Contents) {
		utils.Fatal("not an ELF file")
	}

	// 填充 ELF header
	ehdr := FillIntoEhdr(file.Contents)
	
	contents := file.Contents[ehdr.ShOff:]  // 剩余都是 Section header table 


	// 填充 Section header table
	numSections := int64(ehdr.ShNum)  
	for numSections >= 1 {

		// 填充 Section header 
		shdr := FillIntoShdr(contents)
		f.ElfSections = append(f.ElfSections, shdr)

		contents = contents[ShdrSize:]
		numSections--
	}


	// section header table 索引，指向 shstrtab 
	shstrndx := int64(ehdr.ShStrndx)  
	// 保存 section name 
	f.ShStrtab = f.GetBytesFromIdx(shstrndx)
	
	return f
}


func (f *InputFile) GetBytesFromShdr(s *Shdr) []byte {
	end := s.Offset + s.Size
	if uint64(len(f.File.Contents)) < end {
		utils.Fatal(fmt.Sprintf("sectio header is out of range: %d", s.Offset))
	}

	return f.File.Contents[s.Offset : s.Offset+s.Size]
}

// 根据下标，找 section header，再根据 header 中偏移字段读取数据
func (f InputFile) GetBytesFromIdx(idx int64) []byte {
	return f.GetBytesFromShdr(&f.ElfSections[idx])
}


func (f *InputFile) FindSection(typeField uint32) *Shdr {
	for i := 0; i < len(f.ElfSections); i++ {
		shdr := &f.ElfSections[i]
		if shdr.Type == typeField {
			return shdr
		}
	} 

	return nil
}