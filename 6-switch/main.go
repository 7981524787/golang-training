package main

func main() {

	var d uint8 = 4

	switch d {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	case 3:
		println("tuesday")
	case 4:
		println("wednesday")
	case 5:
		println("thursday")
	case 6:
		println("friday")
	case 7:
		println("saturday")
	default:
		println("no day")
	}

	var num = -50

	switch {
	case num >= 0 && num < 50:
		println(num, "is betwwen 0 to 50")
	case num >= 50 && num < 100:
		println(num, "is 50 or less than 100")
	case num >= 100:
		{
			println(num, "is 100 or more")
		}
	default:
		println(num, "is less than 0")

	}

	num = 24
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	char1 := 'a'

	switch char1 {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		println(string(char1), " is vovel")
	default:
		println(string(char1), " is a consonent or a special char")
	}

}
