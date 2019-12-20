package crypt

import (
	"unsafe"
)

// #cgo LDFLAGS: -lcrypt
// #define _GNU_SOURCE
// #include <crypt.h>
// #include <stdlib.h>
import "C"

// Crypt is wrapper for crypt_r(3)
func Crypt(key, salt string) string {
	data := C.struct_crypt_data{}
	cKey := C.CString(key)
	cSalt := C.CString(salt)
	out := C.GoString(C.crypt_r(cKey, cSalt, &data))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cSalt))
	return out
}
