package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str string = "hello I am Jitendranadh Palaparthi"

	ptr1 := uintptr(unsafe.Pointer(&str))
	ptr1 = ptr1 + 8
	//v1 := (*[2]int)(unsafe.Pointer(ptr1))

	ptr2 := (*int)(unsafe.Pointer(ptr1))
	*ptr2 = 100
	//println(v1[1])
	//fmt.Println(v1)
	fmt.Println("length", len(str))

	for i, v := range str {
		println(i, "-->", string(v))
	}
}
