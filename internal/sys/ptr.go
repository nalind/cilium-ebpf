package sys

import (
	"unsafe"

	"github.com/cilium/ebpf/internal/unix"
)

// NewPointer creates a 64-bit pointer from an unsafe Pointer.
func NewPointer(ptr unsafe.Pointer) Pointer {
	return Pointer{ptr: ptr}
}

// NewSlicePointer creates a 64-bit pointer from a slice.
func NewSlicePointer[T comparable](buf []T) Pointer {
	if len(buf) == 0 {
		return Pointer{}
	}

	return Pointer{ptr: unsafe.Pointer(unsafe.SliceData(buf))}
}

// NewSlicePointerLen creates a 64-bit pointer from a byte slice.
//
// Useful to assign both the pointer and the length in one go.
func NewSlicePointerLen(buf []byte) (Pointer, uint32) {
	return NewSlicePointer(buf), uint32(len(buf))
}

// NewStringPointer creates a 64-bit pointer from a string.
func NewStringPointer(str string) Pointer {
	slice, err := unix.ByteSliceFromString(str)
	if err != nil {
		return Pointer{}
	}

	return Pointer{ptr: unsafe.Pointer(&slice[0])}
}

// NewStringSlicePointer allocates an array of Pointers to each string in the
// given slice of strings and returns a 64-bit pointer to the start of the
// resulting array.
//
// Use this function to pass arrays of strings as syscall arguments.
func NewStringSlicePointer(strings []string) Pointer {
	sp := make([]Pointer, 0, len(strings))
	for _, s := range strings {
		sp = append(sp, NewStringPointer(s))
	}

	return Pointer{ptr: unsafe.Pointer(&sp[0])}
}
