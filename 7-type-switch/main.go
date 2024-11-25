package main

import "fmt"

func main() {

	var any1, any2 any = 123.123, 12312.123
	sum1, err := add(any1, any2)
	if err != nil {
		println(err.Error())
	} else {
		println("sum1:", sum1)
	}
	var a, b uint8 = 12, 123
	sum1, err = add(a, b)
	if err != nil {
		println(err.Error())
	} else {
		println("sum1:", sum1)
	}

	var a1, b1 float64 = 12.123, 123.123
	sum1, err = add(a1, b1)
	if err != nil {
		println(err.Error())
	} else {
		println("sum1:", sum1)
	}

}

func isNumber(a any) bool {
	switch a.(type) {
	case int, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	}
	return false
}

func add(a, b any) (float64, error) {
	ok1, ok2 := isNumber(a), isNumber(b)
	sum := 0.0
	if ok1 && ok2 {
		switch a.(type) { // type switch
		case int:
			sum = float64(a.(int) + b.(int))

		case float32:
			sum = float64(a.(float32) + b.(float32))
		case float64:
			sum = a.(float64) + b.(float64)
		case uint:
			sum = float64(a.(uint) + b.(uint))
		case int8:
			sum = float64(a.(int8) + b.(int8))
		case int16:
			sum = float64(a.(int16) + b.(int16))
		case int32:
			sum = float64(a.(int32) + b.(int32))
		case int64:
			sum = float64(a.(int64) + b.(int64))
		case uint8:
			sum = float64(a.(uint8) + b.(uint8))
		case uint16:
			sum = float64(a.(uint16) + b.(uint16))
		case uint32:
			sum = float64(a.(uint32) + b.(uint32))
		case uint64:
			sum = float64(a.(uint64) + b.(uint64))
		default:
			sum = 0

		}
		return sum, nil
	}

	//return nil, errors.New("invalid number")

	return nil, fmt.Errorf("invalid number")

}
