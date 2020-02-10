package classfile

import "fmt"

type ConstantPackageInfo struct {
	nameIndex uint16
}

func (this *ConstantPackageInfo) ReadInfo(reader IClassReader) {
	this.nameIndex = reader.ReadUint16()
}

func (this ConstantPackageInfo) String(constantPool ConstantPool) string {
	return fmt.Sprint(constantPool[this.nameIndex])
}
