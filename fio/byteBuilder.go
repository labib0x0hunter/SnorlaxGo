package fio

type ByteBuilder struct {
	buf    []byte
	curPos int
}

func (b *ByteBuilder) grow(n int) (int) {

}

func (b *ByteBuilder) tryGrowByReslice(n int) (int, bool) {

}

func (b *ByteBuilder) WriteString(s string) int {
	m, ok := b.tryGrowByReslice(len(s));
	if !ok {
		m = b.grow(len(s))
	}
	return copy(b.buf[m:], s)
}

// func main() {

// }