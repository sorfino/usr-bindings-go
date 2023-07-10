package kvsbinding

/*
#cgo CFLAGS: -I../../build/include
#cgo LDFLAGS: -L../../lib -lusr_kvs -Wl,-rpath,../../lib
#include "unified_sdk_runtime/kvs.h"
*/
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export go_callback
func go_callback(ptr *C.uchar, length C.ulong, userData *C.void, errno C.int) {
	r := C.GoBytes(unsafe.Pointer(ptr), C.int(length))
	h := *(*cgo.Handle)(unsafe.Pointer(userData))
	ch := h.Value().(chan []byte)

	// send the response to the channel.
	ch <- r
}
