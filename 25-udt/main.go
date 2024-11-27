package main

import "fmt"

type myint int

func (mi myint) ToString() string {
	return fmt.Sprint(mi)
}
func (mi myint) Sq() int {
	return int(mi * mi)
}

type myanotherint myint

func (mai myanotherint) Cube() int {
	return int(mai * mai * mai)
}

type mybool bool

func (mb mybool) Touint8() uint8 {
	if mb {
		return 1
	} else {
		return 0
	}
}

func (mb mybool) ToString() string {
	return fmt.Sprint(mb)
}

type integer = int // in all aspects both integer and int are same

func main() {

	var num1 myint = 100
	str1 := num1.ToString()
	s1 := num1.Sq()
	c1 := myanotherint(num1).Cube()

	fmt.Println(str1)
	fmt.Println(s1)
	fmt.Println(c1)

	var num2 int = 25
	str2 := myint(num2).ToString()
	s2 := myint(num2).Sq()
	c2 := myanotherint(num2).Cube()
	fmt.Println(str2)
	fmt.Println(s2)
	fmt.Println(c2)

	var num3 myanotherint = 30
	str3 := myint(num3).ToString()
	s3 := myint(num3).Sq()
	c3 := num3.Cube()

	fmt.Println(str3)
	fmt.Println(s3)
	fmt.Println(c3)

	var float1 float32 = 123.123

	str4 := myint(float1).ToString()
	s4 := myint(float1).Sq()
	c4 := myanotherint(float1).Cube()

	fmt.Println(str4)
	fmt.Println(s4)
	fmt.Println(c4)

	var b1 mybool = true

	str5 := myint(b1.Touint8()).ToString()
	str6 := b1.ToString()
	fmt.Println(str5)
	fmt.Println(str6)
}
