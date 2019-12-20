package crypt

import (
	"unsafe"
)

// #include <unistd.h>
// #include <stdlib.h>
import "C"

// Crypt is wrapper for crypt(3)
func Crypt(key, salt string) string {
	//data := C.struct_crypt_data{}
	cKey := C.CString(key)
	cSalt := C.CString(salt)
	out := C.GoString(C.crypt(cKey, cSalt)) //, &data))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cSalt))
	return out
}
