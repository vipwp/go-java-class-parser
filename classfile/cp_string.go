package classfile

type ConstantStringInfo struct {
	stringIndex uint16
}

func (this *ConstantStringInfo) ReadInfo(reader IClassReader) {
	this.stringIndex = reader.ReadUint16()
}

func (this *ConstantStringInfo) String(constantPool ConstantPool) string {
	if cp, ok := constantPool[this.stringIndex].(*ConstantUtf8Info); ok {
		return cp.String()
	}
	return ""
}
