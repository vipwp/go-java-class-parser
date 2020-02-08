package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/

type BootstrapMethodsAttribute struct {
	cp               ConstantPool
	BootstrapMethods []BootstrapMethod
}

func (b *BootstrapMethodsAttribute) ReadInfo(reader *ClassReader) {
	num := reader.ReadUint16()
	for i := uint16(0); i < num; i++ {
		m := BootstrapMethod{cp: b.cp, BootstrapMethodArguments: []uint16{}}
		m.BootstrapMethodRef = reader.ReadUint16()
		numArguments := reader.ReadUint16()
		for j := uint16(0); j < numArguments; j++ {
			m.BootstrapMethodArguments = append(m.BootstrapMethodArguments, reader.ReadUint16())
		}
		b.BootstrapMethods = append(b.BootstrapMethods, m)
	}
}

type BootstrapMethod struct {
	cp                       ConstantPool
	BootstrapMethodRef       uint16
	BootstrapMethodArguments []uint16
}

func (b *BootstrapMethod) ClassName() string {
	h := b.cp.GetConstantInfo(b.BootstrapMethodRef).(*ConstantMethodHandleInfo)
	ref := b.cp.GetConstantInfo(h.referenceIndex).(*ConstantMethodrefInfo)
	return ref.ClassName()
}

func (b *BootstrapMethod) NameAndDescriptor() (string, string) {
	h := b.cp.GetConstantInfo(b.BootstrapMethodRef).(*ConstantMethodHandleInfo)
	ref := b.cp.GetConstantInfo(h.referenceIndex).(*ConstantMethodrefInfo)
	return ref.NameAndDescriptor()
}
