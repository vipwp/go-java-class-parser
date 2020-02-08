package classfile

type ConstantInterfaceMethodrefInfo struct {
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantInterfaceMethodrefInfo) ReadInfo(reader IClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}
