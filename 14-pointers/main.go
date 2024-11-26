package main

import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	var num1 int = 123
	Incr(num1)
	println(num1)
	var ptr1 *int = &num1
	IncrP(ptr1)
	IncrP(&num1)

	println(num1)
	println(*ptr1)

	// var ptr2 *int // zero value is nil

	// IncrP(ptr2)

	var ptr3 *int = new(int) // allocates 8 bytes and gives that address to ptr3

	println(ptr3)
	println(*ptr3) // 0
	IncrP(ptr3)
	println(*ptr3) // 0

	var ptr4 **int = &ptr3

	IncrP(*ptr4)

	**ptr4 += 1

	//var ptr5 ***int = &ptr4

	var ok1 = true
	var ptr6 *bool = &ok1

	fmt.Println("Size of ptr6", unsafe.Sizeof(ptr6))

	arr1 := [...]int{10, 11, 12, 13, 14, 15}

	//var ptr7 *int = &arr1[0]

	//3434343453
	//3434343461
	//3434343469

	//var uintptr1 int = int(uintptr(unsafe.Pointer(&arr1[0])))
	var uintptr1 uintptr = uintptr(unsafe.Pointer(&arr1[0]))
	for i := 1; i < len(arr1); i++ {
		uintptr1 += unsafe.Sizeof(&arr1[0])
		//v := (*int)(unsafe.Pointer(uintptr(uintptr1)))
		v := (*int)(unsafe.Pointer(uintptr1))
		println(*v)
	}

	slice1 := []int{1, 2, 3, 4, 5}
	Printheader(slice1, "slice1")
	SqAndAddElems(&slice1, 6, 7, 8)
	//fmt.Println(slice1)
	Printheader(slice1, "slice1")

}

func CheckNil1(slice []int) error {
	if slice == nil { // header ptr is checked whether nil or not
		return errors.New("nil slice")
	}
	return nil
}

func CheckNil2(slice *[]int) error {
	if slice == nil { // it does not check the header pointer, it checks the slice itself
		return errors.New("nil slice")
	}

	if *slice == nil { // header pointer
		return errors.New("nil slice")
	}
	return nil
}

func SqAndAddElems(slice *[]int, nums ...int) error {
	if *slice == nil {
		return errors.New("nil slice")
	}

	*slice = append(*slice, nums...)
	for i, v := range *slice {
		(*slice)[i] = v * v
		//(*slice)[i] = v * v
	}
	return nil
}

func Incr(n int) {
	n++
}

func IncrP(n *int) {
	if n != nil {
		*n++
	}
}

func Printheader(slice []int, name string) error {
	if slice == nil {
		return errors.New("nil slice")
	}
	fmt.Printf("Slice:%s Address:%p Values: %v ptr:%p Len: %d Cap: %d\n", name, &slice, slice, &slice[0], len(slice), cap(slice))
	return nil
}
