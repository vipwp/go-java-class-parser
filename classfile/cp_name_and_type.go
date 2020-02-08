package classfile

import "fmt"

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) ReadInfo(reader IClassReader) {
	this.nameIndex = reader.ReadUint16()
	this.descriptorIndex = reader.ReadUint16()
}

func (this ConstantNameAndTypeInfo) String(constantPool ConstantPool) string {
	return fmt.Sprintf("%s:%s", constantPool[this.nameIndex], constantPool[this.descriptorIndex])
}
