package classfile

import "fmt"

type ConstantFieldrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantFieldrefInfo) ReadInfo(reader IClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}

func (this ConstantFieldrefInfo) String(constantPool ConstantPool) string {
	class, _ := constantPool[this.classIndex].(*ConstantClassInfo)
	nameAndType, _ := constantPool[this.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
	return fmt.Sprintf("%s.%s", class.String(constantPool), nameAndType.String(constantPool))
}
