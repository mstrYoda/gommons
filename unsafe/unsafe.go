package unsafe

import (
	"reflect"
	"unsafe"
)

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func Byte(s string) []byte {
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))

	sliceHeader := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}
