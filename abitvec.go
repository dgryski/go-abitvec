// Package abitvec is bit-vector with atomic access
package abitvec

import "sync/atomic"

// Vec is a bitvector
type Vec []uint64

// New returns a new bitvector with the given size
func New(size uint32) Vec {
	return make(Vec, uint(size+63)/64)
}

// Get returns the given bit
func (b Vec) Get(bit uint32) uint {
	shift := bit % 64
	bb := b[bit/64]
	bb &= (1 << shift)
	return uint(bb >> shift)
}

// Set sets the given bit
func (b Vec) Set(bit uint32) {
	b[bit/64] |= (1 << (bit % 64))
}

// AGet atomically returns the given bit
func (b Vec) AGet(bit uint32) uint {
	shift := bit % 64
	bb := atomic.LoadUint64(&b[bit/64])
	bb &= (1 << shift)
	return uint(bb >> shift)
}

// ASet atomically sets the given bit
func (b Vec) ASet(bit uint32) {
	set := uint64(1) << (bit % 64)
	addr := &b[bit/64]
	var old uint64
	for {
		old = atomic.LoadUint64(addr)
		if atomic.CompareAndSwapUint64(addr, old, old|set) {
			break
		}
	}
}
