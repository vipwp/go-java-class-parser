package classfile

import (
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []string{
		"../testdata/Lambda.class",
		"../testdata/Sample.class",
		"../testdata/org.springframework.boot.loader.jar.JarFileEntries.class",
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
			for _, method := range v.BootstrapMethods {
				if method.ClassName() != "java/lang/invoke/LambdaMetafactory" {
					t.Errorf("MethodHandle.MethodRef.className is unexpected. got=%s, want=%s", method.ClassName(), "java/lang/invoke/LambdaMetafactory")
				}
				methodName, _ := method.NameAndDescriptor()
				if methodName != "metafactory" {
					t.Errorf("MethodHandle.MethodRef.methodName is unexpected. got=%s, want=%s", methodName, "metafactory")
				}
			}
		}
	}
}
