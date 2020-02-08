package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (this *ConstantMethodHandleInfo) ReadInfo(reader IClassReader) {
	this.referenceKind = reader.ReadBytes(1)[0]
	this.referenceIndex = reader.ReadUint16()
}

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (this *ConstantMethodTypeInfo) ReadInfo(reader IClassReader) {
	this.descriptorIndex = reader.ReadUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantInvokeDynamicInfo) ReadInfo(reader IClassReader) {
	this.bootstrapMethodAttrIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}
