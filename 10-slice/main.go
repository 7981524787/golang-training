package main

import (
	"fmt"

	"math/rand"
)

func main() {

	var slice1 []int // this is slice declaration but not instantiation
	//var arr1 [4]int
	if slice1 == nil {
		println("yes it is ")
	}

	slice1 = make([]int, 5) // zero value
	fmt.Println(slice1)

	fmt.Printf("address of slice1:%p\n", &slice1)
	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1), "pointer:", &slice1[0])

	slice1 = append(slice1, 123)
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1), "pointer:", &slice1[0])
	slice1 = append(slice1, 1234)
	fmt.Printf("slice1 address:%p", &slice1)
	fmt.Println(" slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1), "pointer:", &slice1[0])
	slice2 := slice1 // shallow copy
	fmt.Printf("slice1 address:%p", &slice2)
	fmt.Println(" slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2), "pointer:", &slice2[0])

	slice2[0] = 999999
	fmt.Println(" slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1), "pointer:", &slice1[0])

	slice2 = append(slice2, 6666, 7777, 8888, 9999)
	slice2[1] = 1111
	slice1 = slice2
	fmt.Println(" slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1), "pointer:", &slice1[0])

}
