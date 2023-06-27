package kvsbinding

/*
#cgo CFLAGS: -I../../include
#cgo LDFLAGS: -L../../lib -lusr_kvs -Wl,-rpath,../../lib
#include "unified_sdk_runtime/kvs.h"
*/
import "C"
import "unsafe"

var _cb map[int]func(int, []byte) int

//export go_callback
func go_callback(fd C.int, ptr *C.uchar, length C.ulong) C.int {
	r := _cb[int(fd)](int(fd), C.GoBytes(unsafe.Pointer(ptr), C.int(length)))

	delete(_cb, int(fd))
	return C.int(r)
}

func registerCallback(fd int, cb func(int, []byte) int) {
	if _cb == nil {
		_cb = make(map[int]func(int, []byte) int)
	}

	_cb[fd] = cb
}
