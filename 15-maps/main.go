package main

import "fmt"

func main() {

	var map1 map[string]any

	map1 = make(map[string]any)
	map1["560011"] = "Some place-1"
	map1["560012"] = "Some place-2"
	map1["1231"] = "some place-3"
	map1["560013"] = "Some place-4"
	map1["560014"] = "Some place-5"
	map1["12311"] = "some place-6"

	for key, value := range map1 {
		fmt.Println("key", key, "value:", value)
	}

	//var any1 any

	// v1, ok1 := any1.(bool)

	// if !ok1 {
	// 	println(v1)
	// }

	v2, ok2 := map1["5600343"]
	if ok2 {
		fmt.Println(v2)
	} else {
		println("key-value pair is not available")
	}

	delete(map1, "560011")
	fmt.Println()
	delete(map1, "okay")

	for key, value := range map1 {
		fmt.Println("key", key, "value:", value)
	}

}

// 0a4d55a8d778e5022fab701977c5d840bbc486d0
// 0a4d55a8d778e5022fab701977c5d840bbc486d0
// 0a4d55a8d778e5022fab701977c5d840bbc486d0
// 053bebe714723f6bd76ba08f869ad8113a408d3e
