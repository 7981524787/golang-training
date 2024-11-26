package main

import (
	_ "time"
)

func main() {
	println("even numbers")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "is even")
		}
	}

	num := 1
	println("odd numbers")
	for num <= 20 { // like while loop
		if num%2 == 1 {
			println(num, "is odd number")
		}
		num++
	}

	num = 1
	for {
		if num >= 10 {
			break
		}
		print(num, " ")
		num++
	}

	println()
	i := 1
	for ; ; i++ {
		if i%2 == 1 {
			continue
		}
		if i >= 20 {
			break
		}
		print(i, " ")
	}
	println()

	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println(i, "-->", j)
	}

	//done := false
out:
	for i := 1; i <= 10; i++ {
		// if done {
		// 	break
		// }
		for j := 5; j <= 10; j++ {
			println("i:", i, "j:", j)
			if i == j {
				//done = true
				break out
			}
		}
	}

	str1 := "Hello, world!" // < 255

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ") // byte
	}
	println()

	str2 := "Hello, 世界!" // unicode chars
	println("len:", len(str2))
	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ") // byte
	}
	println()

	for _, v := range str2 {
		println(":->", string(v))
	}
	// _

	a, _ := 10, 20
	println(a)

	a1, p1 := AreaAndPerimeter(12.34, 13.23)
	a2, _ := AreaAndPerimeter(12.34, 13.23)
	_, p3 := AreaAndPerimeter(12.34, 13.23)

	println("area a1:", a1, "perimeter p1:", p1)
	println("area a2:", a2, "perimeter p3:", p3)

	// n, err := fmt.Println("hello world")
	// if err != nil {

	// }
	num1 := 1
loop:
	println(num1)
	num1++
	if num1 == 20 {
		goto end
	}
	if num1 <= 20 {
		if num1%2 == 0 {
			goto even
		} else {
			goto odd
		}
	}
even:
	println("even:", num1)
	goto loop
odd:
	println("odd:", num1)
	goto loop
end:
}

func AreaAndPerimeter(l, b float32) (a float32, p float32) {
	a = l * b
	p = 2 * (l + b)
	//return l * b, 2 * (l + b)
	return a, p
}
