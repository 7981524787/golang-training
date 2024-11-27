package main

import (
	"fmt"
)

func main() {

	//var num int = 100
	//num1 := 100

	var arr [5]int
	arr[0] = 123
	var slice []int
	slice = make([]int, 1)
	slice[0] = 123

	// var f1 func() = greet
	// f1()
	// var f2 func(int, int) int = add
	// //s1 := f2(100, 200)
	// f2 = sub
	// s2 := f2(100, 200)

	map1 := make(map[string]any)

	map1["F1"] = greet
	map1["F2"] = add
	map1["F3"] = sub(100, 200)
	map1["V1"] = 123

	for k, v := range map1 {
		//fmt.Println(k, v)
		fmt.Printf("key:%s type:%T\n", k, v)
		//tpe := fmt.Sprintf("%T", v) // formatter
	}

	// num := 100
	// ok := true
	// float := 123.123

	// str := fmt.Sprintf("Num:%d Ok:%v Float:%.2f", num, ok, float)
	// fmt.Println(str)
	//let str1 = format!("{} {} {}",num,ok,float)

	//str1 := fmt.Sprint(num)
	//str1 := strconv.Itoa(num)

	//str2 := fmt.Sprintf("%s", num)

}

func greet() {
	println("hello world")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a + b
}
