package classfile

import "fmt"

type ConstantMethodrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMethodrefInfo) ReadInfo(reader IClassReader) {
	this.classIndex = reader.ReadUint16()
	this.nameAndTypeIndex = reader.ReadUint16()
}

func (this *ConstantMethodrefInfo) String() string {
	class, _ := this.cp[this.classIndex].(*ConstantClassInfo)
	nameAndType, _ := this.cp[this.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
	return fmt.Sprintf("%s.%s", class.String(this.cp), nameAndType.String(this.cp))
}

func (self *ConstantMethodrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMethodrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

func (self *ConstantMethodrefInfo) ConstantPool() ConstantPool {
	return self.cp
}
