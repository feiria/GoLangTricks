package zero_copy_between_string_and_bytes

// 以下代码来自《Go程序员面试笔试宝典》 P106

import (
	"reflect"
	"unsafe"
)

func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func string2bytesError(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	// stringHeader.Data本身是uintptr类型，由于goroutine的栈空间可能发生移动，因此不能将其作为中间态的值复制到bh，再转换为[]byte

	return *(*[]byte)(unsafe.Pointer(&bh))
}
