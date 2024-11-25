package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var any1 any // interface{}

	var str1 string = "Hello world"
	var b1 byte = str1[0]
	fmt.Println("address of str1:", &b1, &str1)
	var str2 string = "Welcome to ICICI Securities Golang Training"
	var str3 = "erwerwerewwerwe we we rwerwer we wer w er wer we rwe rw erwerwer w r werwe"

	str1 = "hello Universe"
	var b2 byte = str1[0]
	fmt.Println("address of str1:", &b2, &str1)

	fmt.Println("Size of str1 and str2:", unsafe.Sizeof(str1), unsafe.Sizeof(str2), unsafe.Sizeof(str3))

	fmt.Println("any1:", unsafe.Sizeof(any1))
	any1 = "hello world"

	var str5 string = any1.(string) // type assertion
	str5, isit1 := any1.(string)
	fmt.Println(str5, isit1)
	var ok4 bool
	ok4, isit2 := any1.(bool)
	if isit2 {
		fmt.Println(ok4, isit2)
	}

	//var ok3 bool = any1.(bool) // type assertion

	any1 = true
	var ok5 bool = any1.(bool)
	any1 = 12312.123
	var float2 float64 = any1.(float64)
	any1 = 234234
	any1 = 'A'
	fmt.Println(any1, str5, ok5, float2)

	var (
		n1   uint8   = 123
		n2   uint16  = 1231
		n3   float32 = 1231.123
		n4   float64 = 12312.123
		any2 any     = 123.1231 //float64
		any3 any     = 123123   // int
	)

	var sum float64 = float64(n1) + float64(n2) + float64(n3) + n4 + any2.(float64) + float64(any3.(int))
	fmt.Println(sum)

	a1 := add(int64(n1), int64(n2))
	fmt.Println(a1)

}

func add(a, b int64) int64 {
	return a + b
}

// func adda(a, b any) any {
// 	return a + b
// }
