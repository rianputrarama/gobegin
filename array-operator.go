package main

import "fmt"

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 03/11/2022;
 * Time : 11:02;
**/
func main() {
	var names [3]string
	names[0] = "Febriansyah"
	names[1] = "Putra"
	names[2] = "Ramadhan"

	fmt.Println(names[0])

	var age = []int{
		80, 91, 70,
	}

	fmt.Println(age)
}
