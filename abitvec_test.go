package abitvec

import (
	"sync"
	"testing"
)

func writeBits(b Vec, every, size uint32, wg *sync.WaitGroup) {
	for i := uint32(0); i < size; i += every {
		b.ASet(i)
	}
	wg.Done()
}

func TestAtomic(t *testing.T) {

	const size = 1024

	b := New(size)
	var wg sync.WaitGroup

	wg.Add(2)

	go writeBits(b, 3, size, &wg)
	go writeBits(b, 5, size, &wg)

	wg.Wait()

	for i := uint32(0); i < size; i++ {
		want := i%3 == 0 || i%5 == 0
		if b.Get(i) != b.AGet(i) || b.Get(i) != want {
			t.Errorf("failed: i=%v want=%v get=%v aget=%v", i, want, b.Get(i), b.AGet(i))
		}
	}
}
