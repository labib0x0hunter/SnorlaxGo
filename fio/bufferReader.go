package fio

import (
	"io"
	"os"
)

type Reader struct {
	buf []byte
	r io.Reader
	n int // how much read
}

func BufferReader(r io.Reader) *Reader {
	return &Reader{
		buf: make([]byte, bufSize),
		r: os.Stdin,
	}
}

func (r *Reader) Avalable() int {
	return bufSize - r.n
}

func (r *Reader) Read() {

}

func (r *Reader) Write(p []byte) (n int, err error) {

}