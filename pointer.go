package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 08/05/2022;
 * Time : 22:44;
**/

func main() {
	var namesA string = "rian"
	var namesB *string = &namesA

	println("names : ", namesA)
	println("names : ", &namesA)
	println("names : ", *namesB)
	println("names : ", namesB)

	fmt.Println("")

	namesA = "Tiara"

	println("names : ", namesA)
	println("names : ", &namesA)
	println("names : ", *namesB)
	println("names : ", namesB)
}
