package classfile

import "fmt"

type ConstantClassInfo struct {
	nameIndex uint16
}

func (this *ConstantClassInfo) ReadInfo(reader IClassReader) {
	this.nameIndex = reader.ReadUint16()
}

func (this ConstantClassInfo) String(constantPool ConstantPool) string {
	return fmt.Sprint(constantPool[this.nameIndex])
}
