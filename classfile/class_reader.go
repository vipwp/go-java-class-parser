package classfile

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"strings"
)

var bigEndian = binary.BigEndian

type IClassReader interface {
	ReadUint32() uint32
	ReadUint16() uint16
	ReadUint8() uint8
	ReadBytes(len int) []byte
	Length() int
	Position() int
}

type ClassReader struct {
	bytecode []byte
	position int
}

func NewClassReader(bytecode []byte) *ClassReader {
	return &ClassReader{bytecode: bytecode, position: 0}
}

func (cr *ClassReader) ReadUint32() uint32 {
	value := bigEndian.Uint32(cr.bytecode[cr.position : cr.position+4])
	cr.position += 4
	return value
}

func (cr *ClassReader) ReadUint16() uint16 {
	value := bigEndian.Uint16(cr.bytecode[cr.position : cr.position+2])
	cr.position += 2
	return value
}

func (cr *ClassReader) ReadUint8() uint8 {
	return uint8(cr.ReadBytes(1)[0])
}

func (cr *ClassReader) ReadBytes(len int) []byte {
	bytes := cr.bytecode[cr.position : cr.position+len]
	cr.position += len
	return bytes
}

func (cr *ClassReader) Length() int {
	return len(cr.bytecode)
}

func (cr *ClassReader) Position() int {
	return cr.position
}

type debugClassReader struct {
	*ClassReader
}

func trace() {
	programCounters := make([]uintptr, 4)
	n := runtime.Callers(2, programCounters)
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		traceTarget, _ := frames.Next()

		fmt.Printf("%s called\n", getFuncName(traceTarget.Func.Name()))
		count := 0
		for caller, more := frames.Next(); more; caller, more = frames.Next() {
			count++
			if strings.Contains(caller.Func.Name(), "(*debugClassLoader)") {
				continue
			}
			file, len := caller.Func.FileLine(caller.PC)
			fmt.Printf("%s%s:%d %s \n", strings.Repeat(" ", count), file, len, getFuncName(caller.Func.Name()))
		}

	}
}

func getFuncName(name string) string {
	n := strings.LastIndex(name, ".")
	return name[n+1:]
}

func newDebugClassReader(cr *ClassReader) *debugClassReader {
	return &debugClassReader{cr}
}

func (cr *debugClassReader) ReadUint32() uint32 {
	defer trace()
	return cr.ClassReader.ReadUint32()
}

func (cr *debugClassReader) ReadUint16() uint16 {
	defer trace()
	return cr.ClassReader.ReadUint16()
}

func (cr *debugClassReader) ReadUint8() uint8 {
	defer trace()
	return cr.ClassReader.ReadUint8()
}

func (cr *debugClassReader) ReadBytes(len int) []byte {
	defer trace()
	return cr.ClassReader.ReadBytes(len)
}

func (cr *debugClassReader) Length() int {
	defer trace()
	return cr.ClassReader.Length()
}

func (cr *debugClassReader) Position() int {
	defer trace()
	return cr.ClassReader.Position()
}
