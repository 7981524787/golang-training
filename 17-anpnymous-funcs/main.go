package main

import "fmt"

func main() {
	map1 := make(map[string]any, 10)
	map1["F1-Add"] = add
	map1["F2-Sub"] = sub
	map1["F3-Greet"] = greet
	map1["F4-Square"] = sq

	m := func(a, b int) int {
		return a * b
	}(100, 34)

	f := func(a, b int) int {
		return a * b
	}
	map1["F5-Mult"] = f
	map1["V1"] = m

	map1["F6-Div"] = func(a, b int) int {
		return a / b
	}
	map1["Arr1"] = [...]int{123, 123, 232}
	map1["Slice1"] = []int{123, 23, 34, 5, 665}
	map1["Map1-Internal"] = map[string]string{"a": "1", "b": "2"}

	for key, val := range map1 {
		switch val := val.(type) {
		case func():
			fmt.Println("key:", key)
			//val.(func())()
			val()
		case func(int, int) int:
			r := val(100, 200)
			fmt.Println("Key:", key, "result of ", key, ":", r)
		case func(int) int:
			r := val(100)
			fmt.Println("Key:", key, "result of ", key, ":", r)
		case [3]int:
			fmt.Println("key:", key, "array len:", len(val), "values:", val)
		case []int:
			fmt.Println("key:", key, "slice len:", len(val), "cap:", cap(val), "values:", val)
		case int:
			fmt.Println("key:", key, "value:", val)
		case map[string]string:
			fmt.Println("key:", key, "value:", val)
			fmt.Println("Internal map")
			for k1, v1 := range val {
				fmt.Println("Key:", k1, "Value:", v1)
			}
		default:
			fmt.Println("key:", key, "Value:", "Undefined value, not implemented")
		}
	}

	// num := 100
	// func(a int, b int) {
	// 	//a, b := 100, 200
	// 	c := a + b + num
	// 	println("c:", c)
	// }(100, 200)

	// v1 := calc(100, 200, func(x1, x2 int) int {
	// 	return x1 + x2
	// })

	// v2 := calc(100, 200, func(x1, x2 int) int {
	// 	return x1 - x2
	// })

	// v3 := calc(100, 200, func(x1, x2 int) int {
	// 	return x1 * x2
	// })

	f1 := calc2(100, 2300)
	r1 := f1()
	fmt.Println("Result of a function:", r1)

	// var f2 func()
	// var f3 func(int) int
	// var f4 func(int, int) (int, int, int, error)
}

func greet() {
	println("hello world")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func sq(n int) int {
	return n * n
}

func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func calc2(a, b int) func() int {
	return func() int {
		return a + b
	}
}
