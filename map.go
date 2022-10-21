package main

import "fmt"

func main() {
	// map declaration example
	var names map[string]string
	// by default the value of the map variable is nil, it needs to be initialized by adding {}
	names = map[string]string{}

	names["january"] = "java"
	names["february"] = "golang"

	value, isExist := names["c++"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	for key, element := range names {
		fmt.Println("Key :", key, "=>", "Element :", element)
	}

	fmt.Printf("%s", names["january"])

}
