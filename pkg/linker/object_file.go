package linker

type ObjectFile struct {
	InputFile

	SymtabSec *Shdr
}


func NewObjectFile(file *File) *ObjectFile {
	o := &ObjectFile{InputFile: NewInputFile(file)}
	return o
}


func (o *ObjectFile) Parse() {
	
}