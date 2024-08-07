package main

import "fmt"

type BufferPool struct {
	pool chan []byte
}

func NewBufferPool(maxSize int) *BufferPool {
	return &BufferPool{
		pool: make(chan []byte, maxSize),
	}
}

func (bp *BufferPool) Get() []byte {
	select {
	case buf := <-bp.pool:
		return buf
	default:
		return make([]byte, 1024)
	}
}

func (bp *BufferPool) Put(buf []byte) {
	select {
	case bp.pool <- buf:
	default:

	}
}

func main() {
	bp := NewBufferPool(10)

	buf := bp.Get()
	fmt.Println("Buffer length: %d\n", len(buf))

	bp.Put(buf)

	buf2 := bp.Get()
	fmt.Println("Buffer length: %d\n", len(buf2))

}
