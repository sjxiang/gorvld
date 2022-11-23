package linker

import "bytes"

// 验证 magic number
func CheckMagic(contents []byte) bool {
	return bytes.HasPrefix(contents, []byte("\177ELF"))
}