package kvsbinding

/*
#cgo CFLAGS: -I../../include
#cgo LDFLAGS: -L../../lib -lusr_kvs -Wl,-rpath,../../lib
#include "unified_sdk_runtime/kvs.h"

extern int go_callback(const uint8_t*, size_t, void*, int);

typedef int (*gateway_function_t)(const uint8_t*, size_t, void*, int);
int gateway_function(const uint8_t* ptr, size_t length, void *user_data, int fd)
{
    return go_callback(ptr, length, user_data, fd);
}
*/
import "C"
import (
	"context"
	"fmt"
	"runtime/cgo"
	"syscall"
	"time"
	"unsafe"
)

//go:generate flatc --go --gen-onefile --go-namespace protocol -o protocol --gen-object-api ../../build/flatbuffers/kvs.fbs

//go:generate cp -R ../../build/include/ ../../include

// DefaultTimeout is the timeout for the USR call, not for the call to the underlying service.
const DefaultTimeout = time.Second * 10

type Client struct {
	// h is the handler to the client.
	h *C.ClientHandle
}

func NewClient(container string) (*Client, error) {
	name := C.CString(container)
	defer C.free(unsafe.Pointer(name))

	h, err := C.client_create(name)
	if h == nil {
		return nil, fmt.Errorf("failure to create client: %w", err)
	}

	return &Client{h: h}, nil
}

// Call will call the underlying USR client and wait for the response.
// The response contains the raw bytes of the flatbuffer in question, depending on the operation type.
// buffer cannot be nil or empty, if that is the case, the function will panic.
//
// This implementation is rudementary and should be improved for a first release.
func (c *Client) Call(ctx context.Context, operation uint32, buffer []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()

	response := make(chan []byte, 1)

	// get a pointer to the first item of the slice.
	ptr := (*C.uchar)(unsafe.SliceData(buffer))
	length := C.size_t(len(buffer))

	// convert the callback function to a C function pointer.
	cb := C.gateway_function_t(C.gateway_function)

	userData := cgo.NewHandle(response)
	defer userData.Delete()

	// call the runtime.
	errno := C.client_call(c.h, operation, ptr, length, unsafe.Pointer(&userData), cb)
	if errno > 0 {
		return nil, fmt.Errorf("failure to call client: %w", syscall.Errno(errno))
	}

	// wait until callback execution.
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-response:
		return r, nil
	}
}

func (c *Client) Close() error {
	errno := C.client_close(c.h)
	if errno > 0 {
		return fmt.Errorf("failure to close client: %w", syscall.Errno(errno))
	}

	return nil
}
