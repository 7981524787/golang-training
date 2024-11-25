package main

import "fmt"

func main() {
	var num1 uint8 = 123
	var num2 int64 = int64(num1)
	println(num2)
	var num3 = 12312321312
	var num4 uint8 = uint8(num3)
	var num5 uint16 = uint16(num3)
	fmt.Printf("%b %d %d", num3, num4, num5)

	//var ok1 bool = true

	//var num5 int8 = int8(ok1)
}
