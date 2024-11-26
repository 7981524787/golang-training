package main

import "fmt"

func main() {

	//fmt.Println()

	//var nums ...int

	s1 := SumOf()
	s2 := SumOf(10, 20)
	s3 := SumOf(12, 123, 23, 23, 232, 23, 23, 23, 232, 323, 232, 3232)
	println(s1, s2, s3)
	var slice1 []int                            // slice1 is nil,only declare
	slice1 = append(slice1, 12, 12, 13, 14, 15) // instantiate
	println("cap:", cap(slice1))

	slice2 := []int{123, 123, 34, 23, 232, 23} // shorthqnd declaration of slice with values
	// arr1 := [3]int{1, 2, 3}
	// arr2 := [...]int{1, 2, 3}
	if slice2 == nil {
		println("slice 2 is nil")
	}

	// pass the slice as a variadic arguemnt
	SumOf(slice1...) // variadic argument

	arr1 := [3]int{12, 23, 23} //array

	slice4 := arr1[:] // convert an array to a slice
	slice4[0] = 9999
	fmt.Println(arr1)
	slice4 = append(slice4, 8888) // divergent
	slice4[0] = 1111              // changing the slice do not impact the arrrauy
	fmt.Println(arr1)
	s5 := SumOf(arr1[:]...)
	println(s5)
}

func SumOf(nums ...int) int {
	println(len(nums))
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func SumMultipleOf(nums ...int, m int) int {
func SumMultipleOf(m int, nums ...int) int {
	println(len(nums))
	sum := 0
	for _, v := range nums {
		sum += v * m
	}
	return sum
}

// variadic must be the last parameter
// variaic cannpt be the return type
// cannot create variadic as a variable outside of the function
