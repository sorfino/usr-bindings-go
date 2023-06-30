package kvsbinding

/*
#include "unified_sdk_runtime/kvs.h"

extern int go_callback(int, const uint8_t*, size_t);

typedef int (*gateway_function_t)(int, const uint8_t*, size_t);
int gateway_function(int fd, const uint8_t* ptr, size_t length)
{
    return go_callback(fd, ptr, length);
}
*/
import "C"
import (
	"context"
	"errors"
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

// DefaultTimeout is the timeout for the USR call, not for the call to the underlying service.
const DefaultTimeout = time.Second * 10

type Client struct {
	// fd is the file descriptor of this client.
	fd C.int
}

func NewClient(container string) (*Client, error) {
	name := C.CString(container)
	defer C.free(unsafe.Pointer(name))

	fd := C.client_create(name)
	if fd < 0 {
		return nil, errors.New("failure to create client")
	}

	return &Client{fd: fd}, nil
}

// Call will call the underlying USR client and wait for the response.
// The response contains the raw bytes of the flatbuffer in question, depending on the operation type.
// ptr cannot be nil or empty, if that is the case, the function will panic.
//
// This implementation is rudementary and should be improved for a first release.
func (c *Client) Call(ctx context.Context, operation uint32, buffer []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()

	response := make(chan []byte, 1)
	defer close(response)

	registerCallback(int(c.fd), func(fd int, buffer []byte) int {
		response <- buffer
		return 0
	})

	// get a pointer to the first item of the slice.
	ptr := (*C.uchar)(unsafe.SliceData(buffer))
	length := C.size_t(len(buffer))

	// convert the callback function to a C function pointer.
	cb := C.gateway_function_t(C.gateway_function)

	// call the USR.
	errno := C.client_call(c.fd, operation, ptr, length, cb)
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
	errno := C.client_close(c.fd)
	if errno > 0 {
		return fmt.Errorf("failure to close client: %w", syscall.Errno(errno))
	}

	return nil
}
