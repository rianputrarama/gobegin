package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 07/12/2022;
 * Time : 20:22;
**/

func factiorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		// fmt.Println(result, " * ", i)
		result *= i
	}
	return result
}

func recursiveFunc(value int) int {
	if value == 1 {
		return value
	} else {
		return value * recursiveFunc(value-1)
	}
}

func main() {
	loop := factiorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)

	recusive := recursiveFunc(5)
	fmt.Println(recusive)
}
