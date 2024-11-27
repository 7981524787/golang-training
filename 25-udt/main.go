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

}
