package linker

import (
	"bytes"
	"unsafe"
)


const EhdrSize = int(unsafe.Sizeof(Ehdr{}))
const ShdrSize = int(unsafe.Sizeof(Shdr{}))


// ELF header
type Ehdr struct {
	Ident    [16]uint8 // Magic：   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00 
	Type      uint16  // 类别:                              ELF64
	Machine   uint16  // 系统架构:                          Advanced Micro Devices X86-64
	Version   uint32  // 版本:                              0x1
	Entry     uint64  // 入口点地址：                        0x0
	PhOff     uint64  // 程序头起点：                        0 (bytes into file)
	ShOff     uint64  // Start of section headers:          360 (bytes into file)
	Flags     uint32  // 标志：             0x0
	EhSize    uint16  // Size of this header:               64 (bytes)
	PhEntSize uint16  // Size of program headers:           0 (bytes)
	PhNum     uint16  // Number of program headers:         0
	ShEntSize uint16  // Size of section headers:           64 (bytes)
	ShNum     uint16  // Number of section headers:         8
	ShStrndx  uint16  // Section header string table index: 7
}


// Section header
type Shdr struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	AddrAlign uint64
	EntSize   uint64
}


func ELFGetName(strTab []byte, offset uint32) string {
	length := uint32(bytes.Index(strTab[offset:], []byte{0}))
	return string(strTab[offset : offset+length])
}
