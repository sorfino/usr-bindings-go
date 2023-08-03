/*
 DO NOT MODIFY THIS MANUALLY! This file was generated using cbindgen. 

 This file contains the C bindings for interacting with
 the KVS Service using the Unified SDK Runtime.
*/

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

#define ENUM_MIN_METHOD 0

#define ENUM_MAX_METHOD 2

#define ENUM_MIN_ERROR_CODE 0

#define ENUM_MAX_ERROR_CODE 101

typedef enum Op {
  Get,
  Set,
  Del,
} Op;

typedef struct ClientHandle ClientHandle;

#define Item_VT_KEY 4

#define Item_VT_VALUE 6

#define Item_VT_VERSION 8

#define Item_VT_TIMESTAMP 10

#define Error_VT_CODE 4

#define Error_VT_MESSAGE 6

#define Request_VT_KEYS 4

#define Response_VT_ITEMS 4

#define Response_VT_ERROR 6

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

/**
 * Creates a new `Client` handle.
 * Use this handle to call the other functions. The caller is resposible for calling `client_close` when the handler is no longer needed.
 * Returns a null pointer if any error occurs and sets `errno` with the appropriate error code.
 */
struct ClientHandle *client_create(const char *container);

/**
 * Calls the client with the given operation.
 * callback function will be called when a response comes available.
 */
int client_call(struct ClientHandle *handle,
                enum Op op,
                const uint8_t *root,
                size_t len,
                void *user_data,
                void (*f)(const uint8_t*, size_t, void*, int));

/**
 * closes the client. Calling this function with an already closed handle will cause undefined behavior, most proabably will crash your app.
 */
int client_close(struct ClientHandle *c);

#ifdef __cplusplus
} // extern "C"
#endif // __cplusplus
