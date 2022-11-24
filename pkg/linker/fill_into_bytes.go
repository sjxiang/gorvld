package linker

import (
	"bytes"
	"encoding/binary"

	"github.com/sjxiang/gorvld/pkg/utils"
)

// 填充 ELF header
func FillIntoEhdr(data []byte) Ehdr {
	ehdr := Ehdr{}
	
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &ehdr)  // 小端模式，从 reader 中读取数据，写入到 ehdr 中
	utils.MustNo(err)

	return ehdr
} 


// 填充 Sections header 
func FillIntoShdr(data []byte) Shdr {
	shdr := Shdr{}

	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &shdr)
	utils.MustNo(err)

	return shdr
}