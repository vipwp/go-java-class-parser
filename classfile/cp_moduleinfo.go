package classfile

import "fmt"

type ConstantModuleInfo struct {
	nameIndex uint16
}

func (this *ConstantModuleInfo) ReadInfo(reader IClassReader) {
	this.nameIndex = reader.ReadUint16()
}

func (this ConstantModuleInfo) String(constantPool ConstantPool) string {
	return fmt.Sprint(constantPool[this.nameIndex])
}
