package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	bytes uint32
}

func (this *ConstantIntegerInfo) ReadInfo(reader IClassReader) {
	this.bytes = reader.ReadUint32()
}

type ConstantFloatInfo struct {
	bytes uint32
}

func (this *ConstantFloatInfo) ReadInfo(reader IClassReader) {
	this.bytes = reader.ReadUint32()
}

func (this *ConstantFloatInfo) Value() float32 {
	return math.Float32frombits(this.bytes)
}

type ConstantLongInfo struct {
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantLongInfo) ReadInfo(reader IClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
}

type ConstantDoubleInfo struct {
	highBytes uint32
	lowBytes  uint32
}

func (this *ConstantDoubleInfo) ReadInfo(reader IClassReader) {
	this.highBytes = reader.ReadUint32()
	this.lowBytes = reader.ReadUint32()
}
