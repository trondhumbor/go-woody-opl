package opl

/*
#include "./woody-opl/opl.c"
*/
import "C"

type Opl struct {
}

func NewOpl() *Opl {
	return &Opl{}
}

func (o Opl) Adlib_init(sampleRate int) {
	C.adlib_init(C.uint(sampleRate))
}

func (o Opl) Adlib_write(idx byte, val byte) {
	C.adlib_write(C.ulonglong(idx), C.uchar(val))
}

func (o Opl) Adlib_getsample() int16 {
	var sample C.short
	C.adlib_getsample(&sample, 1)
	return int16(sample)
}

func (o Opl) Adlib_reg_read(port uint64) uint64 {
	return uint64(C.adlib_reg_read(C.ulonglong(port)))
}

func (o Opl) Adlib_write_index(port uint64, val uint8) {
	C.adlib_write_index(C.ulonglong(port), C.uchar(val))
}
