package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var age uint8 = 26

	var ok = age >= 18 && true && (true || false)
	if ok {
		println("eligible for vote bcz age is ", age)
	} else {
		println("not ligible for vote bcz age is ", age)
	}

	gender := 109 //m

	if age >= 18 && gender == 'F' || gender == 102 {
		println("she is eligible for marriage becase of ", age)
	} else if age >= 21 && gender == 77 || gender == 'm' {
		println("he is eligible for marriage becase of ", age)
	} else {
		println("not eligible for marriage")
	}

	char := 20909
	char1 := '冭'
	println(string(char), string(char1))

	str1 := "109123o"

	n, err := strconv.Atoi(str1)

	if err != nil {
		println("----<", err.Error())
	} else {
		fmt.Println("type of n:", reflect.TypeOf(n), n)
		fmt.Printf("type of n %T value of n:%v\n", n, n)
	}
	println(age)

	str2 := "123123.123123"

	float1, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		println("---->", err.Error())
	} else {
		fmt.Println(float1)
	}

	str3 := "true123"
	//str4 := "123.123+12.2i"
	ok1, err := strconv.ParseBool(str3)
	if err != nil {
		println("---->", err.Error(), ok1)
	} else {
		fmt.Println(ok1)
	}

	a, b, c := 10, 20, 30
	// t := a
	// a = b
	// b = t
	//a, b = b, a

	a, b, c = c, a, b

	str4 := "123.123+45.5i"

	c1, err := strconv.ParseComplex(str4, 64)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(c1)
	}

}
