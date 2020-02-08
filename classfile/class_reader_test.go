package classfile

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []string{
		"../testdata/Lambda.class",
		"../testdata/Sample.class",
	}
	for _, tt := range tests {
		bs, err := ioutil.ReadFile(tt)
		if err != nil {
			t.Fatal(err)
		}
		Parse(bs)
	}
}

func TestParseBootstrapMethods(t *testing.T) {
	tt := "../testdata/Lambda.class"
	bs, err := ioutil.ReadFile(tt)
	if err != nil {
		t.Fatal(err)
	}
	cf := Parse(bs)
	for _, a := range cf.attributes {
		if v, ok := a.(*BootstrapMethodsAttribute); ok {
			for _, x := range v.BootstrapMethods {
				h := v.cp.GetConstantInfo(x.BootstrapMethodRef).(*ConstantMethodHandleInfo)
				ref := v.cp.GetConstantInfo(h.referenceIndex).(*ConstantMethodrefInfo)
				method, _ := ref.NameAndDescriptor()
				if ref.ClassName() != "java/lang/invoke/LambdaMetafactory" {
					t.Errorf("MethodHandle.MethodRef.className is unexpected. got=%s, want=%s", ref.ClassName(), "java/lang/invoke/LambdaMetafactory")
				}
				if method != "metafactory" {
					t.Errorf("MethodHandle.MethodRef.methodName is unexpected. got=%s, want=%s", ref.ClassName(), "metafactory")
				}
			}
		}
	}
}
