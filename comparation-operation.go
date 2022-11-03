package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 31/10/2022;
 * Time : 19:01;
**/

func main() {
	var name1 = "Rian"
	var name2 = "Rian"

	var result bool = name1 == name2
	fmt.Println("Result :", result)

	var value1 = 100
	var value2 = 200

	fmt.Printf("%T, %v\n", value1, value1)
	fmt.Printf("%T, %v\n", value2, value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
