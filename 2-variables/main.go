package main

import "fmt"

type integer = int

var global float32 = 12312.23

const PI = 3.14

const (
	MAX    = 9999
	MIN    = 1
	Minute = 60

	HOUR = Minute * 60
)

const (
	iota = 4
	EAST = iota
	WEST
	SOUTH
	_
	NORTH
)

func main() {

	var num1 uint8 = 255
	var ok1 bool
	var str1 string = "Hello World"
	var char1 int32 = 'A'
	var char2 int32 = 'B'

	var num2 rune = '好'
	//你 好
	var char4 = '你'

	var num3 = 100

	var float1 = 12.34

	num1 = 123

	float2 := 123.123

	ok2 := true

	str2 := "Hello ICICI"

	var char3 int64 = 'C'
	println(char4)
	println(num1, ok1, char1, char2, char3, str1, num2, num3, float1, float2, ok2, str2)

	var (
		num4, num5        = 123, 123
		str3       string = "Hello"
		ok3               = true
	)

	println(num4, num5, str3, ok3, WEST)

	//var hours uint8 = 24

	const Day = Minute * 60 * 24

	const sample = (Minute + Day) + 100*MIN + MAX/PI

	c1 := complex(123.34, 1.23)

	var r1, im1 float32 = 123.11, 1.23
	c2 := complex(r1, im1)

	c3 := 123.123 + 2.3i

	c4 := c1 + c3
	c5 := c1 - c3
	c6 := c1 * c3
	c7 := c1 / c3

	{
		a, b := 10, 20
		println(a, b)
	}

	fmt.Println(c1, c2, c3, c4, c5, c6, c7)
	// --------- web server

}

/*
numbers
int ,uint, int8,int16,int32,int64, uint8,uint16,uint32,uint64,float32,float64,
rune, byte ,uinptr

boolen:
bool

string:
string
*/
