package kvsbinding

/*
#cgo CFLAGS: -I../../build/include
#cgo LDFLAGS: -L../../lib -lusr_kvs -Wl,-rpath,../../lib
#include "unified_sdk_runtime/kvs.h"
*/
import "C"
import (
	"github/mercadolibre/go-bindings/pkg/kvsbinding/internal/pointer"
	"unsafe"
)

//export go_callback
func go_callback(ptr *C.uchar, length C.ulong, userData *C.void, errno C.int) {
	r := C.GoBytes(unsafe.Pointer(ptr), C.int(length))
	ch := pointer.Restore(unsafe.Pointer(userData)).(chan []byte)
	// send the response to the channel.
	ch <- r
}
