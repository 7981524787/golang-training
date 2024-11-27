package main

import "fmt"

func main() {

	cc1 := ColourCode{int: 101, string: "come colour"}
	fmt.Println(cc1)
	fmt.Println(cc1.ToString())

}

type integer = int

type ColourCode struct {
	int // anonymous fileds
	integer
	string
}

func (cc *ColourCode) ToString() string {
	return fmt.Sprintln("ID:", cc.int, "Another ID", cc.integer, "Code:", cc.string)
}
