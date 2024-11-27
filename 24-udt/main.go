package main

import "fmt"

func main() {

	var C struct {
		R         float32
		Area      func() float32
		Perimeter func() float32
	}

	C.R = 123.123

	C.Area = func() float32 {
		return 3.14 * C.R * C.R
	}

	C.Perimeter = func() float32 {
		return 2 * 3.14 * C.R
	}

	if C.Area != nil {
		a1 := C.Area()
		fmt.Println("Area of C a1:", a1)
	}
	if C.Perimeter != nil {
		p1 := C.Perimeter()
		fmt.Println("Perimeter of C p1:", p1)
	}

}

// type Circle struct {
// 	R         float32
// 	Area      func() float32
// 	Perimeter func() float32
// }
