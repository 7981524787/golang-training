package main

import (
	"fmt"

	"math/rand"
)

func main() {
	var arr1 [5]int // zero values , type of an array
	var arr2 [5]string

	for i := range arr1 {
		arr1[i] = rand.Intn(100)
	}
	sum1 := SumOfArr(arr1)
	println("sum1 of arr1:", sum1)

	arr3 := [...]int{123, 1, 54, 6, 7, 8, 4, 45, 4, 5, 7, 45} // length of array must be evalued at compile time
	arr4 := [12]int{123, 1, 54, 6, 7, 8, 4, 45, 4, 5, 7, 45}  // length of array must be evalued at compile time

	//SumOfArr(arr3)
	println("len of arr3 and arr4:", len(arr3), len(arr4))
	fmt.Println(arr1, arr2)

	arr2d := [2][2]int{{1, 2}, {3, 4}}

	println("2d array")
	for _, v1 := range arr2d {
		for _, v2 := range v1 {
			println(v2)
		}
	}
	arr3d := [...][][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	println("3d array")

	for _, v1 := range arr3d {
		for _, v2 := range v1 {
			for _, v3 := range v2 {
				println(v3)
			}
		}
	}

}

func SumOfArr(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v

	}
	return sum
}

func SumOfArr2(arr [12]int) int {
	sum := 0
	for _, v := range arr {
		sum += v

	}
	return sum
}

func SumOfArr3(arr [10]int) int {
	sum := 0
	for _, v := range arr {
		sum += v

	}
	return sum
}
