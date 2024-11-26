package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := make([]int, 2, 6)
	slice1[0], slice1[1] = 10, 11
	//slice1[2] = 4

	slice2 := slice1
	Printheader(slice1, "slice1")
	Printheader(slice2, "slice2")

	slice2 = append(slice2, 12)
	println("\nafter appending slice2")
	Printheader(slice1, "slice1")
	Printheader(slice2, "slice2")

	println("\nafter appending slice1")

	slice1 = append(slice1, 13)
	Printheader(slice1, "slice1")
	Printheader(slice2, "slice2")

	println("\again appending slice1")
	slice4 := append(slice1, 14)
	Printheader(slice1, "slice1")
	Printheader(slice2, "slice2")

	//
	Printheader(slice4, "slice4")

	slice5 := make([]int, 5)
	copy(slice5, slice4) // deep

	slice6 := make([]int, 2)
	//var slice7 []int
	err := copys(slice6, slice1)
	if err != nil {
		println(err.Error())
	} else {
		Printheader(slice6, "slice6")
	}
}

func copys(dest, source []int) error {
	if dest == nil || source == nil {
		return errors.New("source or destination slice is nil")
	}
	minlen := min(len(dest), len(source))

	for i := 0; i < minlen; i++ {
		dest[i] = source[i]
	}
	return nil
}

func Printheader(slice []int, name string) error {
	if slice == nil {
		return errors.New("nil slice")
	}
	fmt.Printf("Slice:%s Address:%p Values: %v ptr:%p Len: %d Cap: %d\n", name, &slice, slice, &slice[0], len(slice), cap(slice))
	return nil
}
