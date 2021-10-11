package bolt

import (
	"syscall"
	"unsafe"
)

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0x40000000 // 1GB

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0xFFFFFFF

// brokenUnaligned Are unaligned load/stores broken on this arch?
var brokenUnaligned = false

// NOTE: This function is copied from stdlib because it is not available on darwin.
func madvise(b []byte, advice int) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}
