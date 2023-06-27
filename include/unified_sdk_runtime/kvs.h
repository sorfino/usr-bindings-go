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

int client_create(const char *container);

int client_call(int fd, enum Op op, const uint8_t *buf, size_t len, int (*cb)(int,
                                                                              const uint8_t*,
                                                                              size_t));

int client_close(int fd);

#ifdef __cplusplus
} // extern "C"
#endif // __cplusplus
