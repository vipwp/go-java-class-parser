package classfile

import "fmt"

type ConstantUtf8Info struct {
	bytes []byte //u2 length
}

func (this *ConstantUtf8Info) ReadInfo(reader IClassReader) {
	this.bytes = reader.ReadBytes(int(reader.ReadUint16()))
}

func (this ConstantUtf8Info) String() string {
	return fmt.Sprintf("%s", this.bytes)
}
